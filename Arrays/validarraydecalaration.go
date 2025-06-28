package main

import (
	"fmt"
	"strings"
)

func main() {

	var num [3]int
	var text [3]string
	var result [2]bool
	runtime := [...]int{1, 2, 3}

	// printing runtime

	fmt.Println("Runtime : ", runtime)

	// [0,0,0]
	fmt.Println("Num : ", num)
	//[ ]
	fmt.Println("Text ", text)
	//[false,false]
	fmt.Println("Result ", result)

	// setting value
	num[2] = 10
	text[1] = "Go"
	result[1] = true

	// [0,0,0]
	fmt.Println("Num : ", num)
	//[	Go]
	fmt.Println("Text ", text)
	//[false,false]
	fmt.Println("Result ", result)

	text[2] = "Java 8"
	//[ Go,Java]
	fmt.Println("Text ", text)

	// [Go, ,Java]
	fmt.Println("Text ", strings.Join(text[:], ","))
}
