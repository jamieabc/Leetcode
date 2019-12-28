package main

import "strings"

//Write a function that takes a string as input and reverse only the vowels of a string.
//
//Example 1:
//
//Input: "hello"
//Output: "holle"
//
//Example 2:
//
//Input: "leetcode"
//Output: "leotcede"
//
//Note:
//The vowels does not include the letter "y".

func reverseVowels(s string) string {
	length := len(s)

	if length <= 1 {
		return s
	}

	var sb strings.Builder
	j := length - 1

	for i := range s {
		if isVowel(s[i]) {
			for j >= 0 {
				if isVowel(s[j]) {
					sb.WriteByte(s[j])
					j--
					break
				}
				j--
			}
		} else {
			sb.WriteByte(s[i])
		}
	}

	return sb.String()
}

func isVowel(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U'
}

// problems
// 1. forget about capital
