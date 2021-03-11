package main

// Given a string s, return the length of the longest substring that contains at most two distinct characters.
//
//
//
// Example 1:
//
// Input: s = "eceba"
// Output: 3
// Explanation: The substring is "ece" which its length is 3.
//
// Example 2:
//
// Input: s = "ccaabbb"
// Output: 5
// Explanation: The substring is "aabbb" which its length is 5.
//
//
//
// Constraints:
//
// 1 <= s.length <= 104
// s consists of English letters.

func lengthOfLongestSubstringTwoDistinct(s string) int {
	var char1, char2 byte
	var count1, count2, ans int

	for left, right := 0, 0; right < len(s); {
		if char1 == byte(0) || char1 == s[right] || char2 == s[right] || char2 == byte(0) {
			if char1 == byte(0) || char1 == s[right] {
				char1 = s[right]
				count1++
			} else {
				char2 = s[right]
				count2++
			}

			ans = max(ans, count1+count2)
			right++
		} else {
			if char1 == s[left] {
				count1--
				if count1 == 0 {
					char1 = byte(0)
				}
			} else {
				count2--
				if count2 == 0 {
					char2 = byte(0)
				}
			}
			left++
		}
	}

	return ans
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	becareful about boundary condition: initially not matching any char

//	2.	inspired from solution, can use hash table size to check if there are
//		more than 2 different chars
