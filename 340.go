package main

// Given a string s and an integer k, return the length of the longest substring of s that contains at most k distinct characters.
//
//
//
// Example 1:
//
// Input: s = "eceba", k = 2
// Output: 3
// Explanation: The substring is "ece" with length 3.
//
// Example 2:
//
// Input: s = "aa", k = 1
// Output: 2
// Explanation: The substring is "aa" with length 2.
//
//
//
// Constraints:
//
//     1 <= s.length <= 5 * 104
//     0 <= k <= 50

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	counter := make(map[int]int)
	size := len(s)

	if size == 0 || k == 0 {
		return 0
	}

	var longest, distinct int

	for low, high := 0, 0; high < size; {
		if low == high || distinct <= k {
			counter[int(s[high])]++

			if counter[int(s[high])] == 1 {
				distinct++
			}

			if distinct <= k {
				longest = max(longest, high-low+1)
			}

			high++
		} else {
			counter[int(s[low])]--

			if counter[int(s[low])] == 0 {
				distinct--
			}

			low++
		}
	}

	return longest
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/discuss/80044/Java-O(nlogk)-using-TreeMap-to-keep-last-occurrence-Interview-%22follow-up%22-question!

//		author uses another way to track window: last occurrence of a character,
//		not implement
