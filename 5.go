package main

import (
	"fmt"
)

//Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
//
//Example 1:
//
//Input: "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
//
//Example 2:
//
//Input: "cbbd"
//Output: "bb"

func longestPalindrome(s string) string {
	size := len(s)
	if size <= 1 {
		return s
	}

	longest := 1
	start, end := 0, 0

	// dp[i][j]: is palindrome from i ~ j
	dp := make([][]bool, size)
	for i := range dp {
		dp[i] = make([]bool, size)
		dp[i][i] = true

		if i+1 < size && s[i] == s[i+1] {
			dp[i][i+1] = true
			longest = 2
			start, end = i, i+1
		}
	}

	for i := 2; i < size; i++ {
		for j := 0; i+j < size; j++ {
			if s[j] == s[i+j] && dp[j+1][i+j-1] {
				dp[j][i+j] = true

				if i+1 > longest {
					longest = i + 1
					start, end = j, i+j
				}
			}
		}
	}

	return s[start : end+1]
}

func longestPalindrome4(s string) string {
	size := len(s)
	if size <= 1 {
		return s
	}

	// dp[i][j]: is palindrome from i ~ j
	dp := make([][]bool, size)
	for i := range dp {
		dp[i] = make([]bool, size)
		dp[i][i] = true
	}

	var longest string

	for i := 1; i < size; i++ {
		for j := 0; i+j < size; j++ {
			if dp[j][j+i] && s[j] == s[j+i] {
				dp[j][j+i] = true

				if j-i+1 > len(longest) {
					longest = s[i : j+1]
				}
			}
		}
	}

	return longest
}

// tc: O(n^2)
func longestPalindrome3(s string) string {
	counter := make(map[byte][]int)
	for i := range s {
		counter[s[i]] = append(counter[s[i]], i)
	}

	longest := s[:1]

	for i := range s {
		for j := len(counter[s[i]]) - 1; counter[s[i]][j] > i; j-- {
			idx := counter[s[i]][j]

			if idx-i+1 <= len(longest) {
				break
			}

			if isPalindrome(s[i:idx+1]) && len(longest) < idx-i+1 {
				longest = s[i : idx+1]

			}
		}
	}

	return longest
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func longestPalindrome2(s string) string {
	length := len(s)
	if 0 == length {
		return ""
	} else if 1 == length {
		return s
	}

	var i, j, d, start, end int

	array := make([][]int, length)
	for i = range array {
		array[i] = make([]int, length)
	}

	// setup initial
	for i = 0; i < length; i++ {
		if i < length-1 && s[i] == s[i+1] {
			array[i][i+1] = 2
			start = i
			end = i + 1
		}
		array[i][i] = 1
	}

	for d = 2; d < length; d++ {
		for i = 0; i < length-d; i++ {
			j = i + d
			if s[i] == s[j] {
				if array[i+1][j-1] == d-1 {
					array[i][j] = array[i+1][j-1] + 2
					if array[i][j] > array[start][end] {
						start = i
						end = j
					}
				} else {
					array[i][j] = array[i+1][j-1]
				}
			} else {
				if array[i][j-1] >= array[i+1][j] {
					array[i][j] = array[i][j-1]
				} else {
					array[i][j] = array[i+1][j]
				}
			}
		}
	}

	return s[start : end+1]
}

// reverse string, find common string
// if common string is found, check if index of that character match
// then record that number if matches reflect
func longestPalindrome1(s string) string {
	length := len(s)

	if 0 == length {
		return ""
	} else if 1 == length {
		return s
	}

	mapping := make(map[uint8]int)

	start := 0
	end := 0
	max := 1
	for i := 0; i < length; i++ {
		if _, ok := mapping[s[i]]; !ok {
			mapping[s[i]] = i
		}
		for j := i + max; j < length; j++ {
			if _, ok := mapping[s[j]]; !ok {
				continue
			}
			if palindrone(s, i, j) && j-i+1 > max {
				end = j
				start = i
				max = j - i + 1
			}
		}
	}

	return s[start : end+1]
}

func palindrone(s string, i, j int) bool {
	for i <= j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("babad longest: %s\n", longestPalindrome("babad"))
	fmt.Printf("abcda longest: %s\n", longestPalindrome("abcda"))
	fmt.Printf("cbbd longest: %s\n", longestPalindrome("cbbd"))
	fmt.Printf("ccc longest: %s\n", longestPalindrome("ccc"))
}

//	Notes
//	1.	for using dp to calculate palindrome, there are 2 kind of ways: even &
//		odd, both of conditions required (current length - 2) also palindrome,
//		and new added char matches start of palindrome

//	2.	inspired from https://leetcode.com/problems/longest-palindromic-substring/discuss/151144/Bottom-up-DP-Logical-Thinking

//		author provides a very good explanation

//	3.	inspired from https://havincy.github.io/blog/post/ManacherAlgorithm/
//		animation http://manacher-viz.s3-website-us-east-1.amazonaws.com/#/

//		manacher's algorithm, not implement it
