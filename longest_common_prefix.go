package main

import (
	"sort"
)

func LongestCommonPrefix(strs []string) string {
	sort.Strings(strs)

	var min string
	var res string

	first := strs[0]
	last := strs[len(strs)-1]

	if len(first) < len(last) {
		min = first
	} else {
		min = last
	}

	for i, c := range min {
		if first[i] == last[i] {
			res += string(c)
		} else {
			return res
		}
	}

	return res
}

/*
14. Longest Common Prefix
Solved
Easy
Topics
Companies
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
*/
