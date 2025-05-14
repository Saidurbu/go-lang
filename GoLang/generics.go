package main

import "fmt"

func printSlices[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}

}

type stack[T any] struct {
	elements []T
}

func main() {
	myStack := stack[string]{
		elements: []string{"golang"},
	}

	fmt.Println(myStack)

	//nums := []int{1, 2, 3}
	names := []string{"golang", "typescript"}
	//	vals := []bool{true, false, true}
	printSlices(names)
	//printSlice(myStack)
}
