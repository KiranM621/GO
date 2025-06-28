package main

import "fmt"

func main() {

	// intialize arry with lenght 0
	var num []int

	fmt.Println("Empty Array : ", num)

	// array index out of range error
	num[2] = 10
}
