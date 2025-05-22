package exp_tree

// Array store array of value
type Array []Value

func (a Array) Type() NodeType {
	return NValue
}

func (Array) Variables() Variables {
	return nil
}

func (a Array) F(op Operator) MathFunc {

	switch op {
	case In:
		if err := isArr(a[0]); err != nil {
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

var arrMap = map[Operator]MathFunc{
	In: chainValue(as[Array], reduce(func(acc Bool, v Value, idx int) (Bool, error) {
		return acc, nil
	}, False)),
	OneIn: chainValue(as[Array], func(arr Array) (Bool, error) {
		firstElem, err := as[Array](arr[0])
		if err != nil {
			return False, err
		}
		mp := firstElem.toMap()
		return reduce(func(acc Bool, v Array, idx int) (Bool, error) {
			e, err := as[Array](v)
			if err != nil {
				return False, err
			}
			vMp := e.toMap()
			for k := range mp {
				if vMp[k] != 0 {
					return True, nil
				}
				return False, Break
			}
			return acc, nil
		}, True)(arr[1:])

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
