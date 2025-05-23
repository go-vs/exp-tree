package exp_tree

import "encoding/json"

type Tree struct {
	head Node
}

func (t *Tree) Calculate(v VariableMap) (Value, error) {
	if v == nil {
		v = make(VariableMap)
	}
	return calc(None, t.head, v)
}

func (t *Tree) JSON() (string, error) {
	data, err := json.Marshal(t.head)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (t *Tree) Variables() VariableMap {
	return t.head.Variables()
}
