package main

import (
	parent "packages/family/father"
	child "packages/family/father/son"

	"fmt"
)

func main() {
	f := new(parent.Father)
	fmt.Println(f.Data("abc"))

	c := new(child.Son)
	fmt.Println(c.Data("def"))
}
