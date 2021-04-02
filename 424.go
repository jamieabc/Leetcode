package main

// Given a string s that consists of only uppercase English letters, you can perform at most k operations on that string.
//
// In one operation, you can choose any character of the string and change it to any other uppercase English character.
//
// Find the length of the longest sub-string containing all repeating letters you can get after performing the above operations.
//
// Note:
// Both the string's length and k will not exceed 104.
//
// Example 1:
//
// Input:
// s = "ABAB", k = 2
//
// Output:
// 4
//
// Explanation:
// Replace the two 'A's with two 'B's or vice versa.
//
//
// Example 2:
//
// Input:
// s = "AABABBA", k = 1
//
// Output:
// 4
//
// Explanation:
// Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.

func characterReplacement(s string, k int) int {
	counter := make([]int, 26)
	var ans, majority int
	size := len(s)

	for i, j := 0, 0; i < size && j < size; {
		if i == j || j-i-majority <= k {
			idx := int(s[j] - 'A')
			counter[idx]++

			majority = max(majority, counter[idx])

			if (j - i + 1 - majority) <= k {
				ans = max(ans, j-i+1)
			}

			j++
		} else {
			counter[s[i]-'A']--
			i++

			majority = 0
			for _, c := range counter {
				majority = max(majority, c)
			}
		}
	}

	return ans
}

func characterReplacement1(s string, k int) int {
	counter := make([]int, 26)
	var ans, majority int
	size := len(s)

	for i, j := 0, 0; i < size && j < size; {
		if i == j || j-i-majority <= k {
			idx := int(s[j] - 'A')
			counter[idx]++

			majority = max(majority, counter[idx])

			if (j - i + 1 - majority) <= k {
				ans = max(ans, j-i+1)
			}

			j++
		} else {
			counter[s[i]-'A']--
			i++

			majority = 0
			for _, c := range counter {
				majority = max(majority, c)
			}
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
//	1.	only need to know majority, majority + diff == total length, diff <= k
//		update majority everytime

//	2.	when sliding window is expanded, no need to recheck, just update largest

//	3.	inspired from https://leetcode.com/problems/longest-repeating-character-replacement/discuss/181382/Java-Sliding-Window-with-Explanation

//		don't need to shrink window, just keep original window size and try to see if
//		different range of word can expand

//		this is the most brilliant solution to this problem
