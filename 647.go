package main

// Given a string, your task is to count how many palindromic substrings in this string.
//
// The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.
//
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

func countSubstrings(s string) int {
	length := len(s)
	var count int

	for i := range s {
		// odd
		count++ // self
		for j := 1; i-j >= 0 && i+j < length; j++ {
			if s[i-j] == s[i+j] {
				count++
			} else {
				break
			}
		}

		// even
		if i < length-1 && s[i] == s[i+1] {
			count++
			for j := 1; i-j >= 0 && i+1+j < length; j++ {
				if s[i-j] == s[i+1+j] {
					count++
				} else {
					break
				}
			}
		}
	}

	return count
}

//	problems
//	1.	reference from https://leetcode.com/problems/palindromic-substrings/discuss/105707/Java-Python-DP-solution-based-on-longest-palindromic-substring

//		author uses dp, I think dp didn't benefit much in this problem cause
//		for every i, j calculation needs to happen once
