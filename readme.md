# GO-VS expression tree V2

 
[![Status](https://github.com/go-vs/exp-tree/actions/workflows/go.yml/badge.svg?branch=v2](https://github.com/go-vs/exp-tree/actions/workflows/go.yml)

Exp-tree is go library for parsing expression tree. V2 is a major update with new features and improvements.
- Variable struct/map type supported
- Use `gEnErIc`

TODO:
- Custom operator
- Shunting Yard algorithm parser
- Array index

## Installation

```sh
go get -u github.com/go-vs/exp-tree/v2
```

## Quick start

### Format

Expression tree is in format:

```json
{
  "<operator>": [
    "<arg1>",
    "<arg2>",
    {
      "<opetator>": [
        "<arg>",
        "<arg>",
        "<@variable>"
      ]
    }
  ]
}
```

example

```json
{"and":["@a",{"lt":[1,2]}]}
```

is equivalent to `@a and (1 < 2)` with `a` is a variable

### Variable

Variables use format `@<string>`, and will be replaced as `Value` from `Variables` when
call `Tree.Calculate(v Variables)`

Support for `String`, `Number` ( as float64), `Bool`, `Array` type

With `Bool` type, we already define `True` and `False`
You could use `et.Var(value)` to auto convert value into corresponding type

### Data type
[Bool](doc/bool.md)

[Number](doc/number.md)

[Array](doc/arr.md)

### Operator
[Operator](doc/operator.md)

### Parse tree

```go
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
	treeJSON, err := tree.JSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(treeJSON)
}
```

### Calculate

```go
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
```
