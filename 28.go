package main

import "fmt"

//Implement strStr().
//
//Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//Example 1:
//
//Input: haystack = "hello", needle = "ll"
//Output: 2
//Example 2:
//
//Input: haystack = "aaaaa", needle = "bba"
//Output: -1
//Clarification:
//
//What should we return when needle is an empty string? This is a great question to ask during an interview.
//
//For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

func strStr(haystack string, needle string) int {
	length := len(needle)

	if 0 == length {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	haystackLength := len(haystack)

	for i := range haystack {
		if haystack[i] == needle[0] && haystackLength-i >= length && same(haystack[i+1:i+length], needle[1:]) {
			return i
		}
	}

	return -1
}

func same(source, target string) bool {
	if len(source) != len(target) {
		return false
	}

	for i := range target {
		if target[i] != source[i] {
			return false
		}
	}
	return true
}

func main() {
	source := "hello"
	target := "ll"

	result := strStr(source, target)
	fmt.Printf("result: %d\n", result)
}
