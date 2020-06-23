package main

import "strings"

// To some string S, we will perform some replacement operations that replace groups of letters with new ones (not necessarily the same size).
//
// Each replacement operation has 3 parameters: a starting index i, a source word x and a target word y.  The rule is that if x starts at position i in the original string S, then we will replace that occurrence of x with y.  If not, we do nothing.
//
// For example, if we have S = "abcd" and we have some replacement operation i = 2, x = "cd", y = "ffff", then because "cd" starts at position 2 in the original string S, we will replace it with "ffff".
//
// Using another example on S = "abcd", if we have both the replacement operation i = 0, x = "ab", y = "eee", as well as another replacement operation i = 2, x = "ec", y = "ffff", this second operation does nothing because in the original string S[2] = 'c', which doesn't match x[0] = 'e'.
//
// All these operations occur simultaneously.  It's guaranteed that there won't be any overlap in replacement: for example, S = "abc", indexes = [0, 1], sources = ["ab","bc"] is not a valid test case.
//
// Example 1:
//
// Input: S = "abcd", indexes = [0,2], sources = ["a","cd"], targets = ["eee","ffff"]
// Output: "eeebffff"
// Explanation: "a" starts at index 0 in S, so it's replaced by "eee".
// "cd" starts at index 2 in S, so it's replaced by "ffff".
//
// Example 2:
//
// Input: S = "abcd", indexes = [0,2], sources = ["ab","ec"], targets = ["eee","ffff"]
// Output: "eeecd"
// Explanation: "ab" starts at index 0 in S, so it's replaced by "eee".
// "ec" doesn't starts at index 2 in the original S, so we do nothing.
//
// Notes:
//
//     0 <= indexes.length = sources.length = targets.length <= 100
//     0 < indexes[i] < S.length <= 1000
//     All characters in given inputs are lowercase letters.

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	replacements := make(map[int]int)
	for i := range indexes {
		replacements[indexes[i]] = i
	}

	var strBuilder strings.Builder
	for i := 0; i < len(S); i++ {
		if idx, ok := replacements[i]; ok {
			src, dst := sources[idx], targets[idx]
			if i+len(src) <= len(S) && S[i:i+len(src)] == src {
				strBuilder.WriteString(dst)
				i += len(src) - 1
				continue
			}
		}

		strBuilder.WriteByte(S[i])
	}

	return strBuilder.String()
}

//	problems
//	1.	from sample code, no need to store whole string, can store index to
//		save space

//	2.	from another sample code, sort indexes by its value, and traverse
//		through indexes and create new string

//	3.	inspired from https://leetcode.com/problems/find-and-replace-in-string/discuss/130587/C%2B%2BJavaPython-Replace-S-from-right-to-left

//		could also do replacement from end of string, which may avoid some
//		operation when replace S in place

//		since S is immutable in go, I think it's still better to use string
//		builder

//	4.	inspired from https://leetcode.com/problems/find-and-replace-in-string/discuss/134758/Java-O(n)-solution

//		there's also a technique to preserve sources that matches original
//		string
