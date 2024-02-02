package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	nums := []int{2, 7, 11, 15}
	target := 9
	// Use the TwoSum function from the twosum package
	fmt.Println(TwoSum(nums, target))
	// Add more calls to twosum.TwoSum as needed
}
