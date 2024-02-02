package main

import (
	"reflect"
	"testing"
)

func TestValidParens(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"one", "()", true},
		{"two", "()[]{}", true},
		{"three", "(]", false},
		{"four", "[", false},
		{"five", "", false},
		{"six", "([]){", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValidParens(tc.s)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("IsValidParens = %v, want %v", got, tc.want)
			}
		})
	}
}
