package main

// Given a string s, return true if it is possible to split the string s into three non-empty palindromic substrings. Otherwise, return false.​​​​​
//
// A string is said to be palindrome if it the same string when reversed.
//
//
//
// Example 1:
//
// Input: s = "abcbdd"
// Output: true
// Explanation: "abcbdd" = "a" + "bcb" + "dd", and all three substrings are palindromes.
// Example 2:
//
// Input: s = "bcbddxy"
// Output: false
// Explanation: s cannot be split into 3 palindromes.
//
//
// Constraints:
//
// 3 <= s.length <= 2000
// s consists only of lowercase English letters.

// tc: O(n^2)
func checkPartitioning(s string) bool {
	size := len(s)

	// dp[i][j]: i~j is palindrome
	dp := make([][]bool, size)
	for i := range dp {
		dp[i] = make([]bool, size)
		dp[i][i] = true
		if i+1 < size && s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}

	for d := 2; d < size; d++ {
		for i := range s {
			if i+d < size && dp[i+1][i+d-1] && s[i] == s[i+d] {
				dp[i][i+d] = true
			}
		}
	}

	for i := range s {
		for j := i + 1; j < size-1; j++ {
			if dp[0][i] && dp[i+1][j] && dp[j+1][size-1] {
				return true
			}
		}
	}

	return false
}
