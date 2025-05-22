package exp_tree

type Number float64

func (n Number) F(op Operator) MathFunc {
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

func (Number) Variables() Variables {
	return nil
}

var numberMap = map[Operator]MathFunc{
	Sum: chainValue(chain(as[Array], asArr[Number]), numberSum),
	Mul: numberMul,
	Gt:  numberGt,
	Gte: numberGte,
	Lt:  numberLt,
	Lte: numberLte,
	Div: numberDiv,
	In:  numberIn,
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

func NumberIn(t []Number) (Bool, error) {
	for i := 1; i < len(t); i++ {
		if t[i-1] == t[i] {
			return True, nil
		}
	}
	return False, nil
}

var numberIn = &Math{
	v: func(value Value) error {
		if err := isArr(value); err != nil {
			return err
		}
		values := value.(Array)
		if err := isNumber(values[0]); err != nil {
			return err
		}
		for _, v := range values[1:] {
			if err := isNumberArr(v); err != nil {
				return err
			}
		}
		return nil
	},
	f: func(value Value) Value {
		values := value.(Array)
		now := values[0].(Number)
		for _, arr := range values[1:] {
			if arr.(Array).toMap()[now] == 0 {
				return False
			}
		}
		return True
	},
}
