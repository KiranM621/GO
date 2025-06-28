package main

import "fmt"

func main() {
	//	# command-line-arguments
	//./arraydecalartion.go:6:10: invalid use of [...] array (outside a composite literal)
	var num [...]int
	fmt.Println("Array ", num)
}
