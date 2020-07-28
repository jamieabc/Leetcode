package main

// Given two strings text1 and text2, return the length of their longest common subsequence.
//
// A subsequence of a string is a new string generated from the original string with some characters(can be none) deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.
//
//
//
// If there is no common subsequence, return 0.
//
//
//
// Example 1:
//
// Input: text1 = "abcde", text2 = "ace"
// Output: 3
// Explanation: The longest common subsequence is "ace" and its length is 3.
//
// Example 2:
//
// Input: text1 = "abc", text2 = "abc"
// Output: 3
// Explanation: The longest common subsequence is "abc" and its length is 3.
//
// Example 3:
//
// Input: text1 = "abc", text2 = "def"
// Output: 0
// Explanation: There is no such common subsequence, so the result is 0.
//
//
//
// Constraints:
//
//     1 <= text1.length <= 1000
//     1 <= text2.length <= 1000
//     The input strings consist of lowercase English characters only.

func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] means longest common subsequence for text1 from 0~i & text2 form
	// 0~j
	// dp[0] or dp[][0] are empty set
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				// if same character, add 1 from previous
				dp[i+1][j+1] = 1 + dp[i][j]
			} else {
				// main same from previous
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func longestCommonSubsequence1(text1 string, text2 string) int {
	// memo[i][j] means longest common subsequence size for text1 from i & text2
	// from j
	memo := make([][]int, len(text1))
	for i := range memo {
		memo[i] = make([]int, len(text2))
		for j := range memo[0] {
			memo[i][j] = -1
		}
	}

	return dfs(text1, text2, 0, 0, memo)
}

func dfs(str1, str2 string, ptr1, ptr2 int, memo [][]int) int {
	if ptr1 == len(str1) || ptr2 == len(str2) {
		return 0
	}

	if memo[ptr1][ptr2] != -1 {
		return memo[ptr1][ptr2]
	}

	if str1[ptr1] == str2[ptr2] {
		memo[ptr1][ptr2] = 1 + dfs(str1, str2, ptr1+1, ptr2+1, memo)
	} else {
		memo[ptr1][ptr2] = max(dfs(str1, str2, ptr1+1, ptr2, memo), dfs(str1, str2, ptr1, ptr2+1, memo))
	}

	return memo[ptr1][ptr2]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	cannot find solution

//	2.	inspired from https://leetcode.com/problems/longest-common-subsequence/discuss/436719/Python-very-detailed-solution-with-explanation-and-walkthrough-step-by-step.

//	3.	the other video about LCS https://www.youtube.com/watch?v=V5hZoJ6uK-s

//	4.	add another inspiration https://leetcode.com/problems/longest-common-subsequence/discuss/436719/Python-very-detailed-solution-with-explanation-and-walkthrough-step-by-step.
