package exp_tree

// Bool store bool value
type Bool bool

func (b Bool) F(op Operator) MathFunc {
	return bMp[op]
}

func (b Bool) Type() NodeType {
	return NValue
}

func (b Bool) Variables() Variables {
	return nil
}

const True = Bool(true)
const False = Bool(false)

var bAnd = chainValue(chain(as[Array], asArr[Bool]), typedReduce(func(acc Bool, v Bool, idx int) (Bool, error) {
	if v == False {
		return False, Break
	}
	return acc, nil
}, True))

var bOr = chainValue(chain(as[Array], asArr[Bool]), typedReduce(func(acc Bool, v Bool, idx int) (Bool, error) {
	if v == True {
		return True, Break
	}
	return acc, nil
}, False))

var bNot = chainValue(as[Bool], func(v Bool) (Bool, error) {
	return !v, nil
})

var bMp = map[Operator]MathFunc{
	None: NoneFn,
	And:  bAnd,
	Or:   bOr,
	Not:  bNot,
}
