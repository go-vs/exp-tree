package exp_tree

// NodeType indicate what type of note
type NodeType int

const (
	NOperation NodeType = iota
	NGroup
	NValue
	NVariable
)

// Node interface
type Node interface {
	Type() NodeType
	Variables() VariableMap
}

// Group composite node
type Group []Node

func (Group) Type() NodeType {
	return NGroup
}

func (g Group) Variables() VariableMap {
	res := make(VariableMap)
	for _, node := range g {
		for k := range node.Variables() {
			res[k] = nil
		}
	}
	return res
}
