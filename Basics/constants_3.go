package main

import (
	"fmt"
)

// iota is not a function or keyword — it’s a built-in constant generator that works only inside const blocks.
func main() {

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Println("KB ->", KB)
	fmt.Println("MB ->", MB)
	fmt.Println("GB ->", GB)
	fmt.Println("TB ->", TB)

}
