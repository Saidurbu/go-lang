package main

import "fmt"

// by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// by reference

func changeNumer(num *int) {
	*num = 10
	fmt.Println("After change", *num)

}

func main() {
	num := 1
	// changeNum(num)
	changeNumer(&num)

	// fmt.Println("Memory address", &num)
	fmt.Println("After changeNum in main", num)

}
