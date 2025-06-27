package main

import "fmt"

func main() {

	// basic variable decalaration
	// Go can infer the type from the assigned value
	var a = "Hello"
	var b = 1
	var c = true
	var d = 1.2

	fmt.Println("a -> ", a)
	fmt.Println("b -> ", b)
	fmt.Println("c -> ", c)
	fmt.Println("d -> ", d)

	// decalaration with type
	var e int
	var f string
	var g float64
	var h bool

	fmt.Println("e -> ", e)
	fmt.Println("f -> ", f)
	fmt.Println("g -> ", g)
	fmt.Println("h -> ", h)

	// multiple variable decalaration

	// type will be interfered at compile time
	var x, y, z = 1.2, "abc", 1

	fmt.Println("x -> ", x)
	fmt.Println("y -> ", y)
	fmt.Println("z -> ", z)

}
