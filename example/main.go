package main

import (
	"fmt"

	et "github.com/go-vs/exp-tree/v2"
)

func main() {
	tree, err := et.ParseTree(`{"and":["@a",{"lt":[1,2]}]}`)
	if err != nil {
		panic(err)
	}
	res, err := tree.Calculate(et.VariableMap{
		"a": et.True, // or et.Var(true)
	})
	fmt.Println(res) // true
}
