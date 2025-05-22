package exp_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_In(t *testing.T) {
	tree := &Operation{
		op: In,
		args: Group{
			Var("hello"),
			Group{
				Var("hello"),
				Var("world"),
			},
		},
	}
	res, err := calc(None, tree, nil)
	assert.Nil(t, err)
	assert.Equal(t, True, res)
}
