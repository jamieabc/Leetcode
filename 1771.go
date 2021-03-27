package main

import "fmt"

// You are given two strings, word1 and word2. You want to construct a string in the following manner:
//
// Choose some non-empty subsequence subsequence1 from word1.
// Choose some non-empty subsequence subsequence2 from word2.
// Concatenate the subsequences: subsequence1 + subsequence2, to make the string.
//
// Return the length of the longest palindrome that can be constructed in the described manner. If no palindromes can be constructed, return 0.
//
// A subsequence of a string s is a string that can be made by deleting some (possibly none) characters from s without changing the order of the remaining characters.
//
// A palindrome is a string that reads the same forward as well as backward.
//
//
//
// Example 1:
//
// Input: word1 = "cacb", word2 = "cbba"
// Output: 5
// Explanation: Choose "ab" from word1 and "cba" from word2 to make "abcba", which is a palindrome.
//
// Example 2:
//
// Input: word1 = "ab", word2 = "ab"
// Output: 3
// Explanation: Choose "ab" from word1 and "a" from word2 to make "aba", which is a palindrome.
//
// Example 3:
//
// Input: word1 = "aa", word2 = "bb"
// Output: 0
// Explanation: You cannot construct a palindrome from the described method, so return 0.
//
//
//
// Constraints:
//
// 1 <= word1.length, word2.length <= 1000
// word1 and word2 consist of lowercase English letters.

func longestPalindrome(word1 string, word2 string) int {
	dp := lps(fmt.Sprintf("%s%s", word1, word2))

	var ans int

	for i := range word1 {
		for j := range word2 {
			if word1[i] == word2[j] {
				ans = max(ans, dp[i+1][len(word1)+j-1]+2)
			}
		}
	}

	return ans
}

func lps(str string) [][]int {
	size := len(str)

	// dp[i][j]: longest palindrome subsequence from i ~ j
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = 1
		if i < size-1 && str[i] == str[i+1] {
			dp[i][i+1] = 2
		}
	}

	for d := 1; d < size; d++ {
		for i := 0; i+d < size; i++ {
			if str[i] == str[i+d] {
				dp[i][i+d] = max(dp[i][i+d], dp[i+1][i+d-1]+2)
			} else {
				dp[i][i+d] = max(dp[i+1][i+d], dp[i][i+d-1])
			}
		}
	}

	return dp
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	longest palindrome subsequence take O(n^2)

//	2.	inspired from https://leetcode.com/problems/maximize-palindrome-length-from-subsequences/discuss/1075504/Java-Detailed-Explanation-DP-Enumerate-Every-Char-Pair

//		need to take one char from word1 and one char from word2, need to check
//		word1[i] == word2[j], and find values from dp
