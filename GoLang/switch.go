package main

import (
	"fmt"
	"time"
)

func main() {

	i := 3
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}
	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Friday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other")
		}
	}

	whoAmI(55)

}
