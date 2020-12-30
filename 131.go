package main

import "strings"

// Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.
//
// A palindrome string is a string that reads the same backward as forward.
//
//
//
// Example 1:
//
// Input: s = "aab"
// Output: [["a","a","b"],["aa","b"]]
// Example 2:
//
// Input: s = "a"
// Output: [["a"]]
//
//
// Constraints:
//
// 1 <= s.length <= 16
// s contains only lowercase English letters.

func partition(s string) [][]string {
	size := len(s)

	// dp[i][j]: is palindrome from index i ~ j
	dp := make([][]bool, size)
	for i := range dp {
		dp[i] = make([]bool, size)
	}

	ans := make([][]string, 0)
	dfs(s, dp, 0, []string{}, &ans)

	return ans
}

func dfs(s string, dp [][]bool, start int, current []string, ans *[][]string) {
	if start == len(s) {
		*ans = append(*ans, current)
		return
	}

	for i := start; i < len(s); i++ {
		// bottom-up dp
		if s[start] == s[i] && (i-start <= 2 || dp[start+1][i-1]) {
			dp[start][i] = true

			tmp := append([]string{}, current...)
			tmp = append(tmp, s[start:i+1])
			dfs(s, dp, i+1, tmp, ans)
		}
	}
}

//	Notes
//	1.	everytime a palindrome string is found, keep looking to next one

//		worst case scenario is string is composed by one character, thus it could
//		have 2^n ways to separate string, and takes n to check palindrome, O(2^n * n)

//	2.	inspired from solution, a way to check palindrome can be further optimized

//		checking palindrome s[0:2], s[0:3], s[0:4], ..., s[1:2], s[1:3], s[1:4], ...
//		there's a repeated process to be optimized

//		if build from bottom-up, for s[i] == s[j] and j - i <= 2, it always a valid
//		palindrome

//		e.g. i == j, valid
//			 i+1 == j, valid (aa)
//			 i+2 == j, valid (aba)

//		if sor s[i] == s[j], j - i > 2, but i+1 ~ j-1 is palindrome, then it's
//		also a palindrome
