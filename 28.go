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
	needleLength := len(needle)

	if 0 == needleLength {
		return 0
	}
	haystackLength := len(haystack)

	if haystackLength < needleLength {
		return -1
	}

	for i := range haystack {
		if haystack[i] == needle[0] && haystackLength-i >= needleLength && haystack[i+1:i+needleLength] == needle[1:] {
			return i
		}
	}

	return -1
}

func main() {
	source := "hello"
	target := "ll"

	result := strStr(source, target)
	fmt.Printf("result: %d\n", result)
}
