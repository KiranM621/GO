package main

import "fmt"

func main() {

	var s []string

	// print's empty list
	fmt.Println("Printing slices ", s)

	// creates a slice with length 3
	s = make([]string, 3)
	s[0] = "Jessi"
	s[1] = "Jey"
	s[2] = "Meowth"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	//panic: runtime error: index out of range [4] with length 3
	//panic: runtime error: index out of range [4] with length 3
	s[4] = "Team Rocket"

	fmt.Println("set:", s)
	fmt.Println("get:", s[4])
}
