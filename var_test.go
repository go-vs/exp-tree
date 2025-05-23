package exp_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Var(t *testing.T) {
	// string
	assert.Equal(t, String("123"), Var("123"))
	// bool
	assert.Equal(t, True, Var(true))
	// uint
	assert.Equal(t, Number(1), Var(uint(1)))
	//int
	assert.Equal(t, Number(1), Var(1))
	// float
	assert.Equal(t, Number(1.5), Var(1.5))
	// slice
	assert.Equal(t, Array{
		Var(1),
		Var(true),
		Var(1.5),
	}, Var(1, true, 1.5))
	assert.Equal(t, Array{
		Var(1),
		Var(true),
		Var(1.5),
	}, Var([]interface{}{1, true, 1.5}))
}

func Test_Extract(t *testing.T) {
	type testStruct struct {
		Bool   bool
		Int    int
		String string
		Struct *testStruct
		Map    map[string]any
	}
	data := testStruct{
		Bool:   true,
		Int:    1,
		String: "123",
		Struct: &testStruct{
			Bool:   false,
			Int:    2,
			String: "456",
			Struct: nil,
			Map:    nil,
		},
		Map: map[string]any{
			"key":  "value",
			"key2": 1,
			"key3": map[string]any{
				"key4": "value4",
				"key5": 2,
			},
		},
	}
	v := make(VariableMap)
	v.Set("data", data)
	b, err := v.Get("data.Bool")
	assert.NoError(t, err)
	assert.Equal(t, True, b)
	i, err := v.Get("data.Int")
	assert.NoError(t, err)
	assert.Equal(t, Number(1), i)
	s, err := v.Get("data.String")
	assert.NoError(t, err)
	assert.Equal(t, String("123"), s)
	b2, err := v.Get("data.Struct.Bool")
	assert.NoError(t, err)
	assert.Equal(t, False, b2)
	i2, err := v.Get("data.Struct.Int")
	assert.NoError(t, err)
	assert.Equal(t, Number(2), i2)
	s2, err := v.Get("data.Struct.String")
	assert.NoError(t, err)
	assert.Equal(t, String("456"), s2)
	m, err := v.Get("data.Map.key")
	assert.NoError(t, err)
	assert.Equal(t, String("value"), m)
	m2, err := v.Get("data.Map.key2")
	assert.NoError(t, err)
	assert.Equal(t, Number(1), m2)
	m3, err := v.Get("data.Map.key3.key4")
	assert.NoError(t, err)
	assert.Equal(t, String("value4"), m3)
	_, err = v.Get("data.Map.key3.key6")
	assert.Error(t, err)
}
