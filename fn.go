package exp_tree

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrInvalidConversion = func(v Value, dstType string) error {
	return fmt.Errorf(`invalid conversion "%v" to %s`, v, dstType)
}

func as[T Value](v Value) (T, error) {
	var zero T
	dst, ok := v.(T)
	if !ok {
		tName := reflect.TypeOf(zero).Name()
		return zero, ErrInvalidConversion(v, tName)
	}
	return dst, nil
}

func asArr[T Value](v Array) ([]T, error) {
	res := make([]T, len(v))
	for i, item := range v {
		elem, err := as[T](item)
		if err != nil {
			return nil, err
		}
		res[i] = elem
	}
	return res, nil
}

func asValue[T Value](v T) (Value, error) {
	return v, nil
}

var Break = errors.New("break")

func reduce[T Value, V Value](fn func(acc T, v V, idx int) (T, error), initValue T) func(arr Array) (T, error) {
	return func(arr Array) (T, error) {
		var res = initValue
		var err error
		for i, v := range arr {
			res, err = chain(as[V], func(v V) (T, error) {
				return fn(res, v, i)
			})(v)
			if err != nil {
				if errors.Is(err, Break) {
					return res, nil
				}
				return res, err
			}
		}
		return res, nil
	}
}

func chain[F, T, O any](a func(F) (T, error), b func(T) (O, error)) func(data F) (O, error) {
	return func(data F) (O, error) {
		var zero O
		f, err := a(data)
		if err != nil {
			return zero, err
		}
		return b(f)
	}
}

func chainValue[F, T any, O Value](a func(F) (T, error), b func(T) (O, error)) func(data F) (Value, error) {
	return func(data F) (Value, error) {
		var zero O
		f, err := a(data)
		if err != nil {
			return zero, err
		}
		return b(f)
	}
}
