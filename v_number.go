package exp_tree

type Number float64

func (n Number) F(op Operator) CalcFn {
	switch op {
	case None:
		return NoneFn
	default:
		return numberMap[op]
	}
}

func (n Number) Type() NodeType {
	return NValue
}

func (Number) Variables() VariableMap {
	return nil
}

var numberMap = map[Operator]CalcFn{
	Sum: chainValue(chain(as[Array], asArr[Number]), numberSum),
	Mul: chainValue(chain(as[Array], asArr[Number]), numberMul),
	Gt:  chainValue(chain(as[Array], asArr[Number]), numberGT),
	Gte: chainValue(chain(as[Array], asArr[Number]), numberGTE),
	Lt:  chainValue(chain(as[Array], asArr[Number]), numberLT),
	Lte: chainValue(chain(as[Array], asArr[Number]), numberLTE),
	Div: chainValue(chain(as[Array], asArr[Number]), numberDiv),
	In:  chainValue(as[Array], arrIn[Number]),
	Eq:  chainValue(chain(as[Array], asArr[Number]), numberEq),
}

func numberEq(t []Number) (Bool, error) {
	for _, v := range t[1:] {
		if v != t[0] {
			return False, nil
		}
	}
	return True, nil
}

func numberSum(t []Number) (Number, error) {
	res := Number(0)
	for _, v := range t {
		res += v
	}
	return res, nil
}

func numberMul(t []Number) (Number, error) {
	res := Number(1)
	for _, v := range t {
		res *= v
	}
	return res, nil
}

func numberLT(t []Number) (Bool, error) {
	for i := 1; i < len(t); i++ {
		if t[i-1] >= t[i] {
			return False, nil
		}
	}
	return True, nil
}

func numberLTE(t []Number) (Bool, error) {
	for i := 1; i < len(t); i++ {
		if t[i-1] > t[i] {
			return False, nil
		}
	}
	return True, nil
}

func numberGT(t []Number) (Bool, error) {
	for i := 1; i < len(t); i++ {
		if t[i-1] <= t[i] {
			return False, nil
		}
	}
	return True, nil
}

func numberGTE(t []Number) (Bool, error) {
	for i := 1; i < len(t); i++ {
		if t[i-1] < t[i] {
			return False, nil
		}
	}
	return True, nil
}

func numberDiv(t []Number) (Number, error) {
	res := t[0]
	for _, v := range t[1:] {
		res /= v
	}
	return res, nil
}
