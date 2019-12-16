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
	mapping1 := make(map[string]string)
	mapping2 := make(map[string]string)
	length := len(pattern)

	strs := strings.Split(str, " ")

	if length != len(strs) {
		return false
	}

	for i, p := range pattern {
		char := string(p)
		if _, ok := mapping1[char]; !ok {
			if _, ok := mapping2[strs[i]]; ok {
				return false
			}
			mapping1[char] = strs[i]
			mapping2[strs[i]] = char
		} else {
			if mapping1[char] != strs[i] || mapping2[strs[i]] != char {
				return false
			}
		}
	}
	return true
}
