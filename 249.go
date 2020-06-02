package main

import "strings"

// Given a string, we can "shift" each of its letter to its successive letter, for example: "abc" -> "bcd". We can keep "shifting" which forms the sequence:
//
// "abc" -> "bcd" -> ... -> "xyz"
//
// Given a list of strings which contains only lowercase alphabets, group all strings that belong to the same shifting sequence.
//
// Example:
//
// Input: ["abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"],
// Output:
// [
//   ["abc","bcd","xyz"],
//   ["az","ba"],
//   ["acef"],
//   ["a","z"]
// ]

func groupStrings(strings []string) [][]string {
	mapping := make(map[string][]string)

	for i := range strings {
		k := key(strings[i])
		mapping[k] = append(mapping[k], strings[i])
	}

	result := make([][]string, 0)

	for _, v := range mapping {
		result = append(result, v)
	}

	return result
}

func key(str string) string {
	if len(str) == 0 {
		return ""
	}

	if len(str) == 1 {
		return "1"
	}

	var sb strings.Builder
	for i := 1; i < len(str); i++ {
		diff := int(str[i]) - int(str[i-1])
		if diff < 0 {
			diff += 26
		}

		sb.WriteByte(byte('a' + diff))
	}

	return sb.String()
}

//	problems
//	1.	add reference https://leetcode.com/problems/group-shifted-strings/discuss/67459/1-4-lines-in-Java

//		author uses (c + 26 ) % 26 to avoid diff < 0 statement
