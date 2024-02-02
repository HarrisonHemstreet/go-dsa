package main

func TwoSum(nums []int, target int) []int {
	// Define a function named twoSum that takes a slice of integers and a target integer, returning a slice of integers.
	m := make(map[int]int)
	// Initialize a map to keep track of numbers and their indices encountered so far.
	for i, num := range nums {
		// Loop over the slice, getting both the index (i) and value (num) of each element.
		if j, ok := m[target-num]; ok {
			// Check if the complement of the current number (target-num) exists in the map. If so, 'ok' is true, and 'j' is its index.
			return []int{j, i}
			// If a complement is found, return a slice with the indices of the complement and the current number.
		}
		m[num] = i
		// If no complement is found, add the current number and its index to the map for future reference.
	}
	return nil
	// If no two numbers add up to the target, return nil.
}
