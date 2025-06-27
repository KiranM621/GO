package main

import "fmt"

func main() {
	if num := 10; num < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Positive number")
	}
}
