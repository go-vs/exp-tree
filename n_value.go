package exp_tree

// Value node that store value
type Value interface {
	Type() NodeType
	F(op Operator) MathFunc
	Variables() Variables
}
