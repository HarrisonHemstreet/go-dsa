package main

import (
	"reflect"
	"testing"
)

func TestMergeStringsAlternately(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{"one", "cat", "dog", "cdaotg"},
		{"two", "abc", "pqr", "apbqcr"},
		{"three", "ab", "pqrs", "apbqrs"},
		{"four", "abcd", "pq", "apbqcd"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MergeStringsAlternately(tc.word1, tc.word2)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ran MergeStringsAlternately() = %v, want = %v", got, tc.want)
			}
		})
	}
}
