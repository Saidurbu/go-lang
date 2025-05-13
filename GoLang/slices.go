package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var nums1 = make([]int, 0, 5)

	fmt.Println(cap(nums1))
	fmt.Println(nums1 == nil)

	nums3 := []int{}

	nums3 = append(nums3, 1)
	nums3 = append(nums3, 2)

	fmt.Println(nums3)
	fmt.Println(cap(nums3))
	fmt.Println(len(nums3))

	var nums4 = make([]int, 0, 5)
	nums4 = append(nums4, 2)
	var nums5 = make([]int, len(nums4))

	// copy function
	copy(nums5, nums4)
	fmt.Println(nums4, nums5)

	// slice operator

	var nums7 = []int{1, 2, 3, 4, 5}
	fmt.Println(nums7[0:1])
	fmt.Println(nums7[:2])
	fmt.Println(nums7[1:])

	// slices
	var nums8 = []int{1, 2, 3}
	var nums9 = []int{1, 2, 3}

	fmt.Println(slices.Equal(nums8, nums9))

}
