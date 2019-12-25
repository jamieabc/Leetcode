package main

//Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together. You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.
//
//
//
//Example 1:
//
//Input: "abab"
//Output: True
//Explanation: It's the substring "ab" twice.
//Example 2:
//
//Input: "aba"
//Output: False
//Example 3:
//
//Input: "abcabcabcabc"
//Output: True
//Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)

// a a a
// ab ab
// abad abad
// abadac abadac

func repeatedSubstringPattern(s string) bool {
	// the length of substring should be able to divide total length
	// this problem is actually to find divisible number

	divisable := []int{1}
	length := len(s)

	if length == 1 {
		return false
	}

	// find all numbers that is able to divide length, e.g. 14 => 1, 2, 7
	for i := 2; i <= length/2; i++ {
		if length%i == 0 {
			divisable = append(divisable, i)
		}
	}

	var i int
	for _, n := range divisable {
		subString := s[:n]
		for i = n; i < length; i += n {
			if subString != s[i:i+n] {
				break
			}
		}

		// repeated, that's how i goes to length
		if i == length {
			return true
		}
	}

	return false
}

// problems
// 1. length 1 is always false, but in the algorithm it's true
