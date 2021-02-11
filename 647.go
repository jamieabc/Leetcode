package main

// Given a string, your task is to count how many palindromic substrings in this string.
//
// The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.
// Example 1:
//
// Input: "abc"
// Output: 3
// Explanation: Three palindromic strings: "a", "b", "c".
//
//
//
// Example 2:
//
// Input: "aaa"
// Output: 6
// Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
//
//
//
// Note:
//
//     The input string length won't exceed 1000.

// tc: O(n^2)
func countSubstrings(s string) int {
	size := len(s)
	var count int

	for i := range s {
		count++

		// center at i
		for j := 1; i-j >= 0 && i+j < size; j++ {
			if s[i-j] == s[i+j] {
				count++
			} else {
				break
			}
		}

		// center at i & i+1
		if i+1 < size && s[i] == s[i+1] {
			count++
			for j := 2; i-j+1 >= 0 && i+j < size; j++ {
				if s[i-j+1] == s[i+j] {
					count++
				} else {
					break
				}
			}
		}
	}

	return count
}

// tc:　O(n^2)
func countSubstrings2(s string) int {
	var count int
	size := len(s)

	// dp[i][j]: # of palindrome from i ~ j
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = 1

		if i+1 < size && s[i] == s[i+1] {
			dp[i][i+1] = 1
		}
	}

	for d := 2; d < size; d++ {
		for i := range s {
			if i+d < size && s[i] == s[i+d] && dp[i+1][i+d-1] > 0 {
				dp[i][i+d] += dp[i+1][i+d-1]
			}
		}
	}

	for i := range dp {
		for j := range dp[0] {
			count += dp[i][j]
		}
	}

	return count
}

// tc: O(n^3), for every character, check from i+1 ~ size-1, and iterate through range of chars
func countSubstrings1(s string) int {
	size := len(s)
	var count int

	for i := range s {
		count++
		for j := i+1; j < size; j++ {
			if isPalindrome(s[i:j+1]) {
				count++
			}
		}
	}

	return count
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

//	Notes
//	1.	reference from https://leetcode.com/problems/palindromic-substrings/discuss/105707/Java-Python-DP-solution-based-on-longest-palindromic-substring

//		author uses dp, I think dp didn't benefit much in this problem cause
//		for every i, j calculation needs to happen once
