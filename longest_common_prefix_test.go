package main

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{"one", []string{"flower", "flow", "flight"}, "fl"},
		{"two", []string{"dog", "racecar", "car"}, ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LongestCommonPrefix(tc.strs)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tc.want)
			}
		})
	}
}
