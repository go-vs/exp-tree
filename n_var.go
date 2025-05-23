package exp_tree

import "fmt"

// Variable node that take external value
type Variable string

func (Variable) Type() NodeType {
	return NVariable
}
func (v Variable) Variables() VariableMap {
	return VariableMap{v: nil}
}

func (v Variable) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%c%s"`, VariableIndicator, v)), nil
}
