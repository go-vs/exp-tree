package exp_tree

import "encoding/json"

// Operation node
type Operation struct {
	op     Operator
	args   Node
	result Value // for debugging
}

func (*Operation) Type() NodeType {
	return NOperation
}

func (o Operation) Variables() Variables {
	return o.args.Variables()
}

// MarshalJSON custom JSON marshal
func (o Operation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[Operator]Node{o.op: o.args})
}

// Op create pointer to Operation
func Op(op Operator, args ...Node) *Operation {
	var arg Node
	if len(args) > 1 {
		arg = Group(args)
	} else {
		arg = args[0]
	}
	return &Operation{
		op:     op,
		args:   arg,
		result: nil,
	}
}
