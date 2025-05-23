package exp_tree

import (
	"errors"
	"fmt"
)

var ErrVarNotFound = func(key string) error {
	return fmt.Errorf("ErrVarNotFound: variable %v not found", key)
}

var ErrParseTree = errors.New("ErrParseTree: error when parse tree")

var ErrOpMustBeUnique = errors.New("ErrOpMustBeUnique: operation must be map with single element")

var ErrCalcTree = errors.New("ErrCalcTree: error when calculate")

var ErrOperatorNotSupported = func(op Operator, value Value) error {
	return fmt.Errorf(`ErrOperatorNotSupported: operator "%v" not supported for value %v`, op, value)
}

var ErrInvalidConversion = func(v Value, dstType string) error {
	return fmt.Errorf(`invalid conversion "%v" to %s`, v, dstType)
}
