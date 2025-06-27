package main

import "fmt"

func main() {
	const a = 1

	fmt.Println("a ->", a)

	a = a + 1

	//The Go error:
	//"cannot assign to a (neither addressable nor a map index expression)"
	//means you're trying to assign a value to something that cannot be modified — i.e., it's not addressable (can’t be stored to).

	fmt.Println(a)
}
