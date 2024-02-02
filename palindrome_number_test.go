package main

import (
	"reflect"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"one", 1, true},
		{"two", 10, false},
		{"three", 101, true},
		{"four", -101, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome(tc.num)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("twoSum() = %v, want %v", got, tc.want)
			}
		})
	}
}
