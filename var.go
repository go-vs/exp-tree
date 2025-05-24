package exp_tree

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type VariableMap map[Variable]any

func (v VariableMap) Set(key Variable, value any) {
	v[key] = value
}
func (v VariableMap) Get(path Variable) (Value, error) {
	paths := strings.Split(string(path), ".")
	if len(paths) == 1 {
		if val, ok := v[path]; ok {
			return Var(val), nil
		}
		return nil, ErrVarNotFound(string(path))
	}
	return v.extract(paths[0], paths[1:])
}

func (v VariableMap) extract(base string, paths []string) (Value, error) {
	data, ok := v[Variable(base)]
	if !ok {
		return nil, ErrVarNotFound(base)
	}
	val := reflect.ValueOf(data)
	for _, path := range paths {
		val = reflect.Indirect(val)
		base += "." + path
		switch val.Kind() {
		case reflect.Struct:
			val = val.FieldByName(path)
			if !val.IsValid() {
				return nil, ErrVarNotFound(base)
			}
		case reflect.Map:
			val = val.MapIndex(reflect.ValueOf(path))
			if !val.IsValid() {
				return nil, ErrVarNotFound(base)
			}
			val = reflect.ValueOf(val.Interface())
		default:
			return nil, ErrVarNotFound(base)
		}
	}
	res := Var(val.Interface())
	if res == nil {
		return nil, ErrVarNotFound(base)
	}
	return res, nil
}

func Var(value ...interface{}) Value {
	if len(value) == 1 {
		return varOne(value[0])
	}
	return varOne(value)
}

func varOne(value interface{}) Value {
	t := reflect.TypeOf(value).Kind()
	val := reflect.ValueOf(value)
	switch t {
	case reflect.String:
		return String(val.String())
	case reflect.Bool:
		return Bool(val.Bool())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Float32, reflect.Float64:
		v, _ := strconv.ParseFloat(fmt.Sprint(value), 64)
		return Number(v)
	case reflect.Slice, reflect.Array:
		arr := make(Array, 0, val.Len())
		for i := 0; i < val.Len(); i++ {
			arr = append(arr, Var(val.Index(i).Interface()))
		}
		return arr
	default:
		return nil
	}
}
