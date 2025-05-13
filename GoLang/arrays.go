package main

import "fmt"

func main() {
	// int -> 0, string -> "", bool -> false
	var nums [4]int

	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums)
	// array length
	fmt.Println(len(nums))

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	var name [3]string
	name[0] = "golang"
	fmt.Println(name)

	numsArray := [3]int{1, 2, 3}
	fmt.Println(numsArray)

	numsArrayNested := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(numsArrayNested)

}
