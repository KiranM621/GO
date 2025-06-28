package main

import "fmt"

func main() {

	var s []string

	// print's empty list
	fmt.Println("Printing slices ", s)

	// panic: runtime error: index out of range [0] with length 0
	// let needs to specified using builtin make
	s[0] = "Jessi"
	s[1] = "Jey"
	s[2] = "Meowth"

}
