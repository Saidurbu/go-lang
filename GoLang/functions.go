package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

func processIt(fn func(a int) int) int {
	return fn(1)
}

func main() {
	fn := func(a int) int {
		return 2
	}

	newValue := processIt(fn)

	fmt.Println("processIt returned:", newValue)

	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)
	result := add(3, 5)
	fmt.Println(result)
}
