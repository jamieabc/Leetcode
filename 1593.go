package main

import "math"

// Given a string s, return the maximum number of unique substrings that the given string can be split into.
//
// You can split string s into any list of non-empty substrings, where the concatenation of the substrings forms the original string. However, you must split the substrings such that all of them are unique.
//
// A substring is a contiguous sequence of characters within a string.
//
//
//
// Example 1:
//
// Input: s = "ababccc"
// Output: 5
// Explanation: One way to split maximally is ['a', 'b', 'ab', 'c', 'cc']. Splitting like ['a', 'b', 'a', 'b', 'c', 'cc'] is not valid as you have 'a' and 'b' multiple times.
//
// Example 2:
//
// Input: s = "aba"
// Output: 2
// Explanation: One way to split maximally is ['a', 'ba'].
//
// Example 3:
//
// Input: s = "aa"
// Output: 1
// Explanation: It is impossible to split the string any further.
//
//
//
// Constraints:
//
//    1 <= s.length <= 16
//
//    s contains only lower case English letters.

func maxUniqueSplit(s string) int {
	set := make(map[string]bool)

	return recursive(s, 0, set)
}

func recursive(str string, start int, set map[string]bool) int {
	// terminate condition
	if start == len(str) {
		return 0
	}

	maxPossible := math.MinInt32

	for i := start; i < len(str); i++ {
		if !set[str[start:i+1]] {
			set[str[start:i+1]] = true

			after := recursive(str, i+1, set)
			set[str[start:i+1]] = false

			if after == -1 {
				continue
			}

			maxPossible = max(maxPossible, after+1)
		}
	}

	if maxPossible == math.MinInt32 {
		return -1
	}

	return maxPossible
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
