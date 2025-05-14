package main

import "fmt"

func sum1(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}
func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := sum(nums...)
	fmt.Println(result)
}
