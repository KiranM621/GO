package main

import "fmt"

func main() {

	var s []string

	s = make([]string, 3)
	// print's empty list
	fmt.Println("Printing slices ", s)

	s[0] = "Jessi"
	s[1] = "Jey"
	s[2] = "Meowth"

	s = append(s, "Team Rocket")
	fmt.Println("Printing slices ", s)

	var cp_s []string

	cp_s = make([]string, len(s))

	// empty list
	fmt.Println("Copy of list : ", cp_s)

	//copy from one array to other

	copy(cp_s, s)
	fmt.Println("Copy of list : ", cp_s)

	//scling

	fmt.Println("s[:]", s[:])
	fmt.Println("s[:1]", s[:1])
	fmt.Println("s[1:]", s[1:])
	fmt.Println("s[0:1]", s[0:2])

}
