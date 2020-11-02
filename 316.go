package main

// Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is the smallest in lexicographical order among all possible results.
//
// Note: This question is the same as 1081: https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
//
//
//
// Example 1:
//
// Input: s = "bcabc"
// Output: "abc"
// Example 2:
//
// Input: s = "cbacdcbc"
// Output: "acdb"
//
//
// Constraints:
//
// 1 <= s.length <= 104
// s consists of lowercase English letters.

func removeDuplicateLetters(s string) string {
	counter := make([]int, 26)
	for i := range s {
		counter[s[i]-'a']++
	}

	stack := make([]byte, 0)
	used := make([]bool, 26)

	for i := range s {
		counter[s[i]-'a']--

		if used[s[i]-'a'] {
			continue
		}

		// remove characters that is larger than current and appears later on
		for len(stack) > 0 && counter[stack[len(stack)-1]-'a'] > 0 && stack[len(stack)-1] > s[i] {
			idx := len(stack) - 1
			used[stack[idx]-'a'] = false
			stack = stack[:idx]
		}

		stack = append(stack, s[i])
		used[s[i]-'a'] = true
	}

	return string(stack)
}

func removeDuplicateLetters1(s string) string {
	counter := make([]int, 26)
	for i := range s {
		counter[s[i]-'a']++
	}

	pos := -1
	ans := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		// already used character, skip
		if counter[s[i]-'a'] < 0 {
			continue
		}

		// select smallest when other characters after still exist
		if pos == -1 || s[i] < s[pos] {
			pos = i
		}

		counter[s[i]-'a']--

		if counter[s[i]-'a'] == 0 {
			// after i's position, every character with at least one
			// it's safe to select smallest character
			ans = append(ans, s[pos])
			counter[s[pos]-'a'] = -1

			// recover counter, check again after pos
			if pos != i {
				for j := pos + 1; j <= i; j++ {
					if counter[s[j]-'a'] != -1 {
						counter[s[j]-'a']++
					}
				}
			}

			i = pos
			pos = -1
		}
	}

	return string(ans)
}

//	Notes
//	1.	inspired from solution
//
//		lexicographical order means smaller character is placed earlier. Since
//		every character from original needs to appear at least once, there are
//		two conditions:

//		- if a character has no duplicates after that position, then this character
//	      needs to be selected
//		- if both characters with duplicates after, then select smaller one

//		based on this observation, iterate from start of string, store position
//		that has smallest during iteration. If any character has no duplicates
//		after some position, then smallest character in this iteration will be
//		selected. Repeat this process, start over from position smallest+1

//	2.	the other way of solving this problem is to find rule: for a character c
//		not unique, and there are other character smaller, this character c can
//		be safely removed.

//		order of string matters, removal of character depends on last selected
//		character with duplicates, this acts as a stack
