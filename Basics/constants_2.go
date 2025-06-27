package main

import (
	"fmt"
)

// iota is not a function or keyword — it’s a built-in constant generator that works only inside const blocks.
func main() {
	const a = 1
	fmt.Println("a ->", a)

	const (
		x = iota
		b
		c
	)
	fmt.Println("x ->", x)
	fmt.Println("b ->", b)
	fmt.Println("c ->", c)

	// type
	const y int = 10

	// untyped

	const z = "Z"
	fmt.Println("y ->", y)
	fmt.Println("z ->", z)

}
