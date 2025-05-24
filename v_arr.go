package exp_tree

// Array store array of value
type Array []Value

func (a Array) Type() NodeType {
	return NValue
}

func (Array) Variables() VariableMap {
	return nil
}

func (a Array) F(op Operator) CalcFn {

	switch op {
	case In:
		_, err := as[Array](a[0])
		if err != nil {
			return a[0].F(op)
		}
		fallthrough
	default:
		m := arrMap[op]
		if m == nil {
			return a[0].F(op)
		}
		return m
	}
}

var arrMap = map[Operator]CalcFn{
	In: chainValue(chain(as[Array], asArr[Array]), func(arrs []Array) (Bool, error) {
		mp := arrs[0].toMap()
		for _, arr := range arrs[1:] {
			vMp := arr.toMap()
			for k := range mp {
				if vMp[k] == 0 {
					return False, nil
				}
			}
		}
		return True, nil
	}),
	OneIn: chainValue(chain(as[Array], asArr[Array]), func(arrs []Array) (Bool, error) {
		mp := arrs[0].toMap()
		for _, arr := range arrs[1:] {
			ok := False
			vMp := arr.toMap()
			for k := range mp {
				if vMp[k] != 0 {
					ok = True
					break
				}
			}
			if !ok {
				return False, nil
			}
		}
		return True, nil
	}),
}

// toMap convert Array to with key is Value and value is number of Value in Array
func (a Array) toMap() map[Value]int {
	mp := make(map[Value]int)
	for _, v := range a {
		mp[v]++
	}
	return mp
}
