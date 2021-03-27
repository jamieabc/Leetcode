package main

// Given a string s, find the longest palindromic subsequence's length in s.
//
// A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.
//
//
//
// Example 1:
//
// Input: s = "bbbab"
// Output: 4
// Explanation: One possible longest palindromic subsequence is "bbbb".
//
// Example 2:
//
// Input: s = "cbbd"
// Output: 2
// Explanation: One possible longest palindromic subsequence is "bb".
//
//
//
// Constraints:
//
// 1 <= s.length <= 1000
// s consists only of lowercase English letters.

// tc: O(n^2)
func longestPalindromeSubseq(s string) int {
	size := len(s)

	// dp[i][j]: longest paindrome from i ~ j
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = 1
		if i+1 < size && s[i] == s[i+1] {
			dp[i][i+1] = 2
		}
	}

	for d := 1; d < size; d++ {
		for i := 0; i+d < size; i++ {
			if s[i] == s[i+d] {
				dp[i][i+d] = max(dp[i][i+d], dp[i+1][i+d-1]+2)
			} else {
				dp[i][i+d] = max(dp[i+1][i+d], dp[i][i+d-1])
			}
		}
	}

	return dp[0][size-1]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/longest-palindromic-subsequence/discuss/99101/Straight-forward-Java-DP-solution

//		if left-most & right-most characters are different, use from left-most
//		or right-most

//		dp[i][j] = max(dp[i+1][j], dp[i][j-1])
