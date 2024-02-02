package main

//import ( // Begin import block to include necessary packages.
//	"reflect" // Import the reflect package for deep equality checks.
//	"testing" // Import the testing package, providing support for automated testing.
//)

// import twosum from current project
import (
	"reflect"
	"testing"
)

// Define a test function for the twoSum function.
func TestTwoSum(t *testing.T) {
	// Initialize a slice of anonymous structs for test cases.
	tests := []struct {
		name   string // Name of the test case for identification.
		nums   []int  // Input slice of integers for the twoSum function.
		target int    // Target sum for which two numbers in `nums` must add up to.
		want   []int  // Expected result: indices of the two numbers in `nums` that add up to `target`.
	}{
		// Define individual test cases with name, nums, target, and expected output.
		{"test1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"test2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"test3", []int{3, 3}, 6, []int{0, 1}},
		{"test4", []int{3, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 23, []int{10, 11}},
	}

	// Iterate through each test case.
	for _, tt := range tests {
		// Use t.Run to execute a subtest for each case. `tt.name` gives the subtest its name.
		t.Run(tt.name, func(t *testing.T) {
			// Call the twoSum function with the current test case's inputs.
			got := TwoSum(tt.nums, tt.target)
			// Check if the function's output matches the expected output.
			if !reflect.DeepEqual(got, tt.want) {
				// If they don't match, report an error for this test case.
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
