package exp_tree

type CalcFn func(value Value) (Value, error)

var NoneFn CalcFn = func(value Value) (Value, error) {
	return value, nil
}

func arrIn[T Value](arrs Array) (Value, error) {
	now, err := as[T](arrs[0])
	if err != nil {
		return False, err
	}
	return chain(
		asArr[Array], func(arrs []Array) (Bool, error) {
			for _, arr := range arrs {
				if arr.toMap()[now] != 0 {
					return True, nil
				}
			}
			return False, nil
		},
	)(arrs[1:])
}
