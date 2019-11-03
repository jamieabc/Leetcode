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

// reverse string, find common string
// if common string is found, check if index of that character match
// then record that number if matches reflect
func longestPalindrome2(s string) string {
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

func longestPalindrome(s string) string {
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

func main() {
	fmt.Printf("babad longest: %s\n", longestPalindrome("babad"))
	fmt.Printf("abcda longest: %s\n", longestPalindrome("abcda"))
	fmt.Printf("cbbd longest: %s\n", longestPalindrome("cbbd"))
	fmt.Printf("ccc longest: %s\n", longestPalindrome("ccc"))
}
