package exp_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAs(t *testing.T) {
	a := Var(1)
	b, err := as[Number](a)
	assert.NoError(t, err)
	assert.Equal(t, Number(1), b)
	_, err = as[Bool](a)
	fmt.Println(err)
	assert.Error(t, err)
}
