package main

//Given a string s and a string t, check if s is subsequence of t.
//
//You may assume that there is only lower case English letters in both s and t. t is potentially a very long (length ~= 500,000) string, and s is a short string (<=100).
//
//A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ace" is a subsequence of "abcde" while "aec" is not).
//
//Example 1:
//s = "abc", t = "ahbgdc"
//
//Return true.
//
//Example 2:
//s = "axc", t = "ahbgdc"
//
//Return false.
//
//Follow up:
//If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you want to check one by one to see if T has its subsequence. In this scenario, how would you change your code?
//
//Credits:
//Special thanks to @pbrother for adding this problem and creating all test cases.

// for multiple queries, use a map to store each character appear index
// compare those index with s
func isSubsequence(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)

	if lenS == 0 {
		return true
	}

	if lenS > lenT {
		return false
	}

	// store index of each character
	mapping := make(map[byte][]int)

	for i := 0; i < lenT; i++ {
		c := t[i]
		if _, ok := mapping[c]; !ok {
			mapping[c] = []int{i}
		} else {
			mapping[c] = append(mapping[c], i)
		}
	}

	tIndex := 0

	for i := 0; i < lenS; i++ {
		if pos, ok := mapping[s[i]]; !ok {
			return false
		} else {
			// check index is valid
			found := false
			for _, j := range pos {
				// make sure char in order
				if j > tIndex {
					tIndex = j
					found = true
					break
				}
			}

			if !found {
				return false
			}
		}
	}

	return true
}

func isSubsequence2(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)

	// remove all character form t
	if lenS == 0 {
		return true
	}

	// substring length must be less
	if lenS > lenT {
		return false
	}

	if lenS == lenT {
		return s == t
	}

	var i, j int

	for i, j = 0, 0; i < lenS && j < lenT; j++ {
		if s[i] == t[j] {
			i++
		}
	}

	return i == lenS
}

// problems
// 1. t start index should be -1, otherwise if first occurrence char at index 0 will be counted
