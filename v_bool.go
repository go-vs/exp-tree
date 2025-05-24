package exp_tree

// Bool store bool value
type Bool bool

func (b Bool) F(op Operator) CalcFn {
	return bMp[op]
}

func (b Bool) Type() NodeType {
	return NValue
}

func (b Bool) Variables() VariableMap {
	return nil
}

const True = Bool(true)
const False = Bool(false)

var bAnd = chainValue(chain(as[Array], asArr[Bool]), func(arr []Bool) (Bool, error) {
	for _, v := range arr {
		if v == False {
			return False, nil
		}
	}
	return True, nil
})

var bOr = chainValue(chain(as[Array], asArr[Bool]), func(arr []Bool) (Bool, error) {
	for _, v := range arr {
		if v == True {
			return True, nil
		}
	}
	return False, nil
})

var bNot = chainValue(as[Bool], func(v Bool) (Bool, error) {
	return !v, nil
})

var bMp = map[Operator]CalcFn{
	None: NoneFn,
	And:  bAnd,
	Or:   bOr,
	Not:  bNot,
}
