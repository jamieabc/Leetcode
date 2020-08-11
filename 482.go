package main

import "strings"

//You are given a license key represented as a string S which consists only alphanumeric character and dashes. The string is separated into N+1 groups by N dashes.
//
//Given a number K, we would want to reformat the strings such that each group contains exactly K characters, except for the first group which could be shorter than K, but still must contain at least one character. Furthermore, there must be a dash inserted between two groups and all lowercase letters should be converted to uppercase.
//
//Given a non-empty string S and a number K, format the string according to the rules described above.
//
//Example 1:
//
//Input: S = "5F3Z-2e-9-w", K = 4
//
//Output: "5F3Z-2E9W"
//
//Explanation: The string S has been split into two parts, each part has 4 characters.
//Note that the two extra dashes are not needed and can be removed.
//
//Example 2:
//
//Input: S = "2-5g-3-J", K = 2
//
//Output: "2-5G-3J"
//
//Explanation: The string S has been split into three parts, each part has 2 characters except the first part as it could be shorter as mentioned above.
//
//Note:
//
//    The length of string S will not exceed 12,000, and K is a positive integer.
//    String S consists only of alphanumerical characters (a-z and/or A-Z and/or 0-9) and dashes(-).
//    String S is non-empty.

// 123-456 K = 4
// 12-34-56 K = 4
// 1234-5678 K = 4

func licenseKeyFormatting(S string, K int) string {
	// count chars in S
	var charCount int
	for i := range S {
		if isChar(S[i]) {
			charCount++
		}
	}

	// allocate enough space for new string
	remain := charCount % K
	var newStr []byte
	if remain == 0 && charCount > 0 {
		newStr = make([]byte, (charCount/K)*(K+1)-1)
	} else {
		newStr = make([]byte, (charCount/K)*(K+1)+remain)
	}

	// put first remain chars
	var i, j int
	for ; i < remain; i++ {
		for ; !isChar(S[j]); j++ {
		}

		newStr[i] = capitalize(S[j])
		j++
	}

	if remain > 0 && i < len(newStr) {
		newStr[i] = '-'
		i++
	}

	for count := 1; i < len(newStr); i, j, count = i+1, j+1, count+1 {
		for !isChar(S[j]) {
			j++
		}

		newStr[i] = capitalize(S[j])
		if count%K == 0 && i < len(newStr)-1 {
			i++
			newStr[i] = '-'
		}
	}

	return string(newStr)
}

func capitalize(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return 'A' + b - 'a'
	}
	return b
}

func isChar(b byte) bool {
	return b != '-'
}

func licenseKeyFormatting(S string, K int) string {
	length := len(S)

	var reversed strings.Builder
	count := 0
	for i := length - 1; i >= 0; i-- {
		// skip any -
		if S[i] == '-' {
			continue
		}

		// write capitalized
		if count < K {
			count++
			reversed.WriteString(capitalize(string(S[i])))
		} else {
			// write separator
			// prevent last - without any character left
			reversed.WriteString("-")
			reversed.WriteString(capitalize(string(S[i])))
			count = 1
		}
	}

	var result strings.Builder
	str := reversed.String()
	for i := len(str) - 1; i >= 0; i-- {
		result.WriteByte(str[i])
	}

	return result.String()
}

func capitalize(char string) string {
	return strings.ToUpper(char)
}

//	Notes
//	1.	be careful about string boundaries

//	2.	inspired from https://leetcode.com/problems/license-key-formatting/discuss/96506/Concise-C%2B%2B-solution-(scan-string-backward)

//		every K+1 from tail must be '-', and since I already calculated
//		remain, this one could be used, which reduce a variable space

//		as it's clear, start from beginning requires many checking for
//		boundary conditions (empty string, space calcualtion, etc.). If
//		interviewer allows additional space for reversed string, it would
//		be better to write from back
