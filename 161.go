package main

// Given two strings s and t, return true if they are both one edit distance apart, otherwise return false.
//
// A string s is said to be one distance apart from a string t if you can:
//
//     Insert exactly one character into s to get t.
//     Delete exactly one character from s to get t.
//     Replace exactly one character of s with a different character to get t.
//
//
//
// Example 1:
//
// Input: s = "ab", t = "acb"
// Output: true
// Explanation: We can insert 'c' into s to get t.
//
// Example 2:
//
// Input: s = "", t = ""
// Output: false
// Explanation: We cannot get t from s by only one step.
//
// Example 3:
//
// Input: s = "a", t = ""
// Output: true
//
// Example 4:
//
// Input: s = "", t = "A"
// Output: true
//
//
//
// Constraints:
//
//     0 <= s.length <= 104
//     0 <= t.length <= 104
//     s and t consist of lower-case letters, upper-case letters and/or digits.

func isOneEditDistance(s string, t string) bool {
	sizeS, sizeT := len(s), len(t)

	// make sure length s >= length t
	if sizeS < sizeT {
		return isOneEditDistance(t, s)
	}

	// only one different char
	if sizeS == sizeT {
		for i := range s {
			if s[i] != t[i] {
				return s[i+1:] == t[i+1:]
			}
		}

		// no difference found
		return false
	}

	// different length, can only by 1 difference (insert or delete)
	if sizeS != sizeT+1 {
		return false
	}

	for i := 0; i < sizeS-1; i++ {
		if s[i] != t[i] {
			// found different char, all afterward chars should all be same
			return s[i+1:] == t[i:]
		}
	}

	// different char is last one
	return true
}

func isOneEditDistance1(s string, t string) bool {
	sizeS, sizeT := len(s), len(t)

	if abs(sizeS-sizeT) > 1 {
		return false
	}

	counter := make([]int, 256)

	for i := range s {
		counter[s[i]]++
	}

	for i := range t {
		counter[t[i]]--
	}

	var plus, minus int

	for i := range counter {
		if counter[i] > 0 {
			plus += counter[i]
		} else if counter[i] < 0 {
			minus -= counter[i]
		}
	}

	var i, j, diff int

	if plus == 1 && minus == 1 {
		// replace, make sure other chars are same at same position

		if sizeS != sizeT {
			return false
		}

		for i := range s {
			if s[i] != t[i] {
				diff++
			}
		}

		return diff == 1

	} else if plus == 1 && minus == 0 {
		// delete one char from s

		if sizeS != sizeT+1 {
			return false
		}

		for i, j = 0, 0; i < sizeS && j < sizeT; i++ {
			if s[i] != t[j] {
				diff++
			} else {
				j++
			}
		}

		return diff == 1 || i == sizeS-1

	} else if plus == 0 && minus == 1 {
		// delete one char from t

		if sizeT != sizeS+1 {
			return false
		}

		for ; i < sizeS && j < sizeT; j++ {
			if s[i] != t[j] {
				diff++
			} else {
				i++
			}
		}

		return diff == 1 || j == sizeT-1
	}

	return false
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	Notes
//	1.	initially I thought it was about checking char frequency, but this is wrong because
//		rules means: only one difference exist, two strings can be with 1 different count
//		and multiple different, e.g. abcde, edcbb, only one char b is different, but almost
//		every char at same position are different

//	2.	not only lower case, but also upper case & digits

//	3.	there are 3 conditions: delete, increase, replace, so I first categorize
//		condition, then check again under specific condition

//	4.	boundary conditions are messy, if not deal with it carefully, it might
//		access out of boundary memory

//		the problem I think of is assuming only one char is different, but there
//		could also be multiple different chars, which cause one pointer to go
//		beyond boundary

//	5.	inspired from solution, there's more elegant way to solve this:
//		- same length, once different char is found, all afterward chars
//		are same (not including self)

//		- length difference 1, once different char is found, all afterward
//		chars are same (including self)

//		actually, insert & delete is the same case, differs only by which view
//		point, longer or shorter string, so, one smarter way is to fix view
//		point by shorter/longer string, then condition count can be reduced by 1

//	6.	the mose important insight is that: delete/insert in same condition,
//		once different char found, same condition applies: afterwards char are same
