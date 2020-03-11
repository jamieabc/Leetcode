package main

import "strings"

//Given a pattern and a string str, find if str follows the same pattern.
//
//Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.
//
//Example 1:
//
//Input: pattern = "abba", str = "dog cat cat dog"
//Output: true
//
//Example 2:
//
//Input:pattern = "abba", str = "dog cat cat fish"
//Output: false
//
//Example 3:
//
//Input: pattern = "aaaa", str = "dog cat cat dog"
//Output: false
//
//Example 4:
//
//Input: pattern = "abba", str = "dog dog dog dog"
//Output: false
//
//Notes:
//You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.

func wordPattern(pattern string, str string) bool {
	mapping1 := make(map[byte]string)
	mapping2 := make(map[string]byte)

	length := len(pattern)
	strs := strings.Split(str, " ")
	if length != len(strs) {
		return false
	}

	for i := 0; i < length; i++ {
		if mapped, ok := mapping1[pattern[i]]; !ok {
			mapping1[pattern[i]] = strs[i]
		} else {
			if mapped != strs[i] {
				return false
			}
		}

		if mapped, ok := mapping2[strs[i]]; !ok {
			mapping2[strs[i]] = pattern[i]
		} else {
			if mapped != pattern[i] {
				return false
			}
		}
	}

	return true
}

// problems
//	1.	map is one way, so it can find duplicates when a -> b but not b -> a
