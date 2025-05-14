package main

import "fmt"

// iterating over data structures
func main() {
	// nums := []int{6, 7, 8}
	nums := []int{1, 2, 3}

	for i, num := range nums {
		fmt.Println(num, i)
	}
	m := map[string]string{"fname": "Md", "lname": "Saidur"}
	for v, k := range m {
		fmt.Println(v, k)
	}

	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte , 2 byte
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}
