package exp_tree

// String store string value
type String string

func (s String) Type() NodeType {
	return NValue
}

func (s String) F(op Operator) MathFunc {
	return stringMap[op]
}

func (String) Variables() Variables {
	return nil
}

var stringMap = map[Operator]MathFunc{
	None: NoneFn,
	In:   chainValue(as[Array], arrIn[String]),
	Eq:   chainValue(chain(as[Array], asArr[String]), stringEq),
	Lt:   chainValue(chain(as[Array], asArr[String]), stringLt),
	Lte:  chainValue(chain(as[Array], asArr[String]), stringLte),
	Gt:   chainValue(chain(as[Array], asArr[String]), stringGt),
	Gte:  chainValue(chain(as[Array], asArr[String]), stringGte),
}

func stringLt(t []String) (Bool, error) {
	for i := 1; i < len(t); i++ {
		if t[i] <= t[i-1] {
			return False, nil
		}
	}
	return True, nil
}

func stringLte(t []String) (Bool, error) {
	for i := 1; i < len(t); i++ {
		if t[i] < t[i-1] {
			return False, nil
		}
	}
	return True, nil
}

func stringGt(t []String) (Bool, error) {
	for i := 1; i < len(t); i++ {
		if t[i] >= t[i-1] {
			return False, nil
		}
	}
	return True, nil
}

func stringGte(t []String) (Bool, error) {
	for i := 1; i < len(t); i++ {
		if t[i] > t[i-1] {
			return False, nil
		}
	}
	return True, nil
}

func stringEq(t []String) (Bool, error) {
	for i := 1; i < len(t); i++ {
		if t[i] != t[i-1] {
			return False, nil
		}
	}
	return True, nil
}
