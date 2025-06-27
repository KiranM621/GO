package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Weekday : ", time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend!")
	default:
		fmt.Println("Weekdays!")
	}

}
