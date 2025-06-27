package main

import "fmt"

func main() {
	fmt.Println("1+1", 1+1)
	fmt.Println("A+A", "A"+"A")
	fmt.Println("7.0/3.0 = ", (7.0 / 3.0))
	fmt.Println("7.1+3.0 = ", (7.1 + 3))
	//fmt.Println("A+1", "A"+1) - type error - invalid conversion
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
