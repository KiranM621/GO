package main

import "fmt"

func main() {
	fmt.Println("For Loop!")

	i := 0

	for i <= 3 {
		fmt.Println("i ->", i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println("j ->", j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 3 {
		fmt.Println("n ->", n)
	}
}
