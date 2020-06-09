package main

import (
	"strconv"
	"strings"
)

// Given an encoded string, return its decoded string.
//
// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.
//
// You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.
//
// Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].
//
//
//
// Example 1:
//
// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"
//
// Example 2:
//
// Input: s = "3[a2[c]]"
// Output: "accaccacc"
//
// Example 3:
//
// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"
//
// Example 4:
//
// Input: s = "abc3[cd]xyz"
// Output: "abccdcdcdxyz"

func decodeString(s string) string {
	chars := []byte{byte(0)}
	nums := make([]int, 0)
	var num int

	for i := range s {
		if s[i] == '[' {
			nums = append(nums, num)
			num = 0
			chars = append(chars, byte(0))
		} else if s[i] == ']' {
			var tmp []byte
			for chars[len(chars)-1] != byte(0) {
				tmp = append(tmp, chars[len(chars)-1])
				chars = chars[:len(chars)-1]
			}

			for j, k := 0, len(tmp)-1; j < k; j, k = j+1, k-1 {
				tmp[j], tmp[k] = tmp[k], tmp[j]
			}

			// remove separator byte(0)
			chars = chars[:len(chars)-1]

			for j := 0; j < nums[len(nums)-1]; j++ {
				chars = append(chars, tmp...)
			}
			nums = nums[:len(nums)-1]
		} else {
			if isNum(s[i]) {
				num *= 10
				num += int(s[i] - '0')
			} else {
				chars = append(chars, s[i])
			}
		}
	}

	var sb strings.Builder
	for i := range chars {
		if chars[i] == byte(0) {
			continue
		} else {
			sb.WriteByte(chars[i])
		}
	}

	return sb.String()
}

func decodeString1(s string) string {
	return traverse(s)
}

func traverse(s string) string {
	var sb strings.Builder
	length := len(s)

	var j, k int
	for i := 0; i < length; i++ {
		if isChar(s[i]) {
			sb.WriteByte(s[i])
		} else {
			for j = i + 1; j < len(s); j++ {
				if !isNum(s[j]) {
					break
				}
			}

			n, _ := strconv.Atoi(s[i:j])
			count := 1
			for k = j + 1; k < len(s); k++ {
				if s[k] == '[' {
					count++
				} else if s[k] == ']' {
					count--
				}

				if count == 0 {
					break
				}
			}
			tmp := traverse(s[j+1 : k])

			for ; n > 0; n-- {
				sb.WriteString(tmp)
			}

			i = k
		}
	}

	return sb.String()
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func isChar(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

//	problems
//	1.	char could be capital

//	2.	inspired from https://leetcode.com/problems/decode-string/discuss/87662/Python-solution-using-stack

//		using stack to do recursion, when [ is encountered, add a
//		number and separator, when ] is encountered, pop stack until
//		separator is encountered. expand encoded into decoded string

//		this is really beautiful solution...extract recursion into basic
//		steps, when [ is meet, repeat all words to ] number of times
