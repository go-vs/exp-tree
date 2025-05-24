package exp_tree

import (
	"encoding/json"
	"reflect"
)

const VariableIndicator = '@'

func parseNode(v any) (Node, error) {
	rV := reflect.ValueOf(v)
	switch rV.Kind() {
	case reflect.Bool:
		return Bool(rV.Bool()), nil
	case reflect.Float64, reflect.Float32:
		return Number(rV.Float()), nil
	case reflect.String:
		vString := rV.String()
		if len(vString) > 0 && vString[0] == VariableIndicator {
			return Variable(vString[1:]), nil
		}
		return String(vString), nil
	case reflect.Slice:
		arr := reflect.ValueOf(v)
		group := make(Group, 0, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			e := arr.Index(i).Interface()
			eVal, err := parseNode(e)
			if err != nil {
				return nil, err
			}
			group = append(group, eVal)
		}
		return group, nil
	case reflect.Map:
		mp, _ := v.(map[string]any)
		if len(mp) != 1 {
			return nil, ErrOpMustBeUnique
		}
		for op, node := range mp {
			args, err := parseNode(node)
			if err != nil {
				return nil, err
			}
			return &Operation{
				op:   Operator(op),
				args: args,
			}, nil
		}
		return nil, ErrParseTree
	default:
		return nil, ErrParseTree
	}
}

func ParseTree(s string) (*Tree, error) {
	res := make(map[string]any)
	if err := json.Unmarshal([]byte(s), &res); err != nil {
		return nil, err
	}
	head, err := parseNode(res)
	if err != nil {
		return nil, err
	}
	return &Tree{
		head: head,
	}, nil
}
