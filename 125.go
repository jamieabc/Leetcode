package main

import (
	"fmt"
	"strings"
	"unicode"
)

//Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
//
//Note: For the purpose of this problem, we define empty string as valid palindrome.
//
//Example 1:
//
//Input: "A man, a plan, a canal: Panama"
//Output: true
//
//Example 2:
//
//Input: "race a car"
//Output: false

func isPalindrome(s string) bool {
	if 0 == len(s) {
		return true
	}

	trimmed := trimNonChar(s)
	length := len(trimmed)

	i := 0
	j := length - 1

	for i < j {
		if string(trimmed[i]) == string(trimmed[j]) {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}

func trimNonChar(source string) string {
	bytes := make([]byte, 0)
	for _, s := range source {
		if unicode.IsLetter(s) || unicode.IsDigit(s) {
			bytes = append(bytes, byte(s))
		}
	}

	return strings.ToLower(string(bytes))
}

func isPalindrome2(s string) bool {
	if 0 == len(s) {
		return true
	}

	i, j := 0, len(s)-1
	r := []rune(s)

	for i < j {
		if !isAlphaNumerical(r[i]) {
			i++
			continue
		}

		if !isAlphaNumerical(r[j]) {
			j--
			continue
		}

		if unicode.ToLower(r[i]) == unicode.ToLower(r[j]) {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}

func isAlphaNumerical(r rune) bool {
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return true
	}
	return false
}

func main() {
	s := "A man, a plan, a canal: Panama"

	result := isPalindrome(s)
	fmt.Printf("result: %t\n", result)
}
