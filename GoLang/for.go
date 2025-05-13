package main

import "fmt"

// for -> only construct in go for looping
func main() {
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// infinite loop
	// for {
	// 	println("1")
	// }

	for i := 0; i <= 5; i++ {
		if i == 4 {
			continue
		}
		fmt.Println(i)
	}

	// 1.22 range
	for i := range 9 {
		fmt.Println(i)
	}
}
