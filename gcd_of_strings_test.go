package main

import (
	"reflect"
	"testing"
)

func TestGCDOfStrings(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{"one", "ABCABC", "ABC", "ABC"},
		{"two", "ABABAB", "ABAB", "AB"},
		{"three", "LEET", "CODE", ""},
		{"four", "TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXX"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GCDOfStrings(tc.str1, tc.str2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GCDOfSTrings() = %v, want = %v", got, tc.want)
			}
		})
	}
}
