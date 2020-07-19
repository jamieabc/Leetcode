package main

import (
	"fmt"
	"math"
	"sort"
)

// Given a string s of lowercase letters, you need to find the maximum number of non-empty substrings of s that meet the following conditions:
//
// The substrings do not overlap, that is for any two substrings s[i..j] and s[k..l], either j < k or i > l is true.
// A substring that contains a certain character c must also contain all occurrences of c.
//
// Find the maximum number of substrings that meet the above conditions. If there are multiple solutions with the same number of substrings, return the one with minimum total length. It can be shown that there exists a unique solution of minimum total length.
//
// Notice that you can return the substrings in any order.
//
//
//
// Example 1:
//
// Input: s = "adefaddaccc"
// Output: ["e","f","ccc"]
// Explanation: The following are all the possible substrings that meet the conditions:
// [
// "adefaddaccc"
// "adefadda",
// "ef",
// "e",
// "f",
// "ccc",
// ]
// If we choose the first string, we cannot choose anything else and we'd get only 1. If we choose "adefadda", we are left with "ccc" which is the only one that doesn't overlap, thus obtaining 2 substrings. Notice also, that it's not optimal to choose "ef" since it can be split into two. Therefore, the optimal way is to choose ["e","f","ccc"] which gives us 3 substrings. No other solution of the same number of substrings exist.
//
// Example 2:
//
// Input: s = "abbaccd"
// Output: ["d","bb","cc"]
// Explanation: Notice that while the set of substrings ["d","abba","cc"] also has length 3, it's considered incorrect since it has larger total length.
//
//
//
// Constraints:
//
// 1 <= s.length <= 10^5
// s contains only lowercase English letters.

func maxNumOfSubstrings(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	// find intervals
	counters := buildCounters(s)
	intervals := buildIntervals(counters)

	end := -1
	result := make([]string, 0)

	for _, intr := range intervals {
		if intr[0] > end {
			result = append(result, s[intr[0]:intr[1]+1])
			end = intr[1]
		}
	}

	return result
}

func buildIntervals(counters [][]int) [][]int {
	intervals := make([][]int, 0)
	for i := range counters {
		if len(counters[i]) > 0 {
			intervals = append(intervals, counters[i])
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	return intervals
}

func buildCounters(s string) [][]int {
	chars := make([][]int, 26)
	for i := range chars {
		chars[i] = make([]int, 0)
	}

	// find max range for single char
	for i := range s {
		idx := s[i] - 'a'
		if len(chars[idx]) == 0 {
			chars[idx] = append(chars[idx], i, i)
		} else {
			chars[idx][1] = i
		}
	}

	// make sure every words are included in it
	// abab, a: 0-3, b: 1-4
	for i := range chars {
		if len(chars[i]) == 0 {
			continue
		}

		oldStart, oldEnd := chars[i][0], chars[i][1]

		for j := chars[i][0] + 1; j <= chars[i][1]; j++ {
			if s[j] != s[chars[i][0]] {
				chars[i][1] = max(chars[i][1], chars[s[j]-'a'][1])
				chars[i][0] = min(chars[i][0], chars[s[j]-'a'][0])
			}
		}

		for oldStart > chars[i][0] || oldEnd < chars[i][1] {
			tmpStart, tmpEnd := chars[i][0], chars[i][1]

			for j := chars[i][0]; j < oldStart; j++ {
				chars[i][1] = max(chars[i][1], chars[s[j]-'a'][1])
				chars[i][0] = min(chars[i][0], chars[s[j]-'a'][0])
			}

			for j := oldEnd; j < chars[i][1]; j++ {
				chars[i][1] = max(chars[i][1], chars[s[j]-'a'][1])
				chars[i][0] = min(chars[i][0], chars[s[j]-'a'][0])
			}

			oldStart, oldEnd = tmpStart, tmpEnd
		}
	}

	return chars
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	problems
//	1.	inspired from https://leetcode.com/problems/maximum-number-of-non-overlapping-substrings/discuss/743248/Java-O(n)-solution-with-comment

//		most important intuition: after find intervals, with only one meeting
//		room available, how many meetings can hold?

//	2.	be careful about find start/end, it shoud be a recursive process because
//		each time checking, if could find additional chars

//		e.g cabbcba for b: 2~5
//		1st round: c is added and range extend to 0~5
//		2nd round: a is added and range extend to 0~6
