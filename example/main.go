package main

import (
	"fmt"

	et "github.com/go-vs/exp-tree/v2"
)

func main() {
	tree, err := et.ParseTree(`{"and":["@data.a", "@b",{"lt":[1,2]}]}`)
	if err != nil {
		panic(err)
	}
	vMp := make(et.VariableMap)
	vMp.Set("data", map[string]any{"a": true})
	vMp.Set("b", true)
	res, err := tree.Calculate(vMp)
	if err != nil {
		panic(err)
	}
	fmt.Println(res) // true
}
