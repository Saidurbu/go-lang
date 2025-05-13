package main

import "fmt"

func main() {
	age := 13
	if age > 18 {
		fmt.Println("Person is adult")
	} else if age >= 12 {
		fmt.Println("person is an teenager")
	} else {
		fmt.Println("Person is not adult")

	}

	role := "admin"
	hasPermissions := true

	if role == "admin" && hasPermissions {
		fmt.Println("Allow to access")
	}

	// we can declare a variable inside if construct
	if age := 20; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	// go does not have ternary, you will have to use normal if else

}
