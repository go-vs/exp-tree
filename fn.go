package exp_tree

import (
	"reflect"
)

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
