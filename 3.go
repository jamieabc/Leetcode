package main

// Given a string, find the length of the longest substring without repeating characters.
//
// Example 1:
//
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:
//
// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:
//
// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
//              Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	mapping := make(map[byte]int)

	var prev, count int

	// idea is to use a map to store all char occur index, and use another
	// variable to store start of current non-repeated range
	for i := range s {
		if d, ok := mapping[s[i]]; ok && d >= prev {
			// non-repeated means within range of i to prev
			// prev is where unique char start, it doesn't care about ith one
			count = max(count, i-prev)
			prev = d + 1
			mapping[s[i]] = i
		} else {
			mapping[s[i]] = i
		}
	}

	return max(count, len(s)-prev)
}

func lengthOfLongestSubstring1(s string) int {
	pos := make(map[byte][]int)

	for i := range s {
		pos[s[i]] = append(pos[s[i]], i)
	}

	var dist, end int

	for i := range s {
		if len(pos[s[i]]) == 1 {
			end = len(s)
		} else {
			end = search(pos[s[i]], i, len(s))
		}

		for j := i + 1; j < end; j++ {
			if tmp := search(pos[s[j]], j, end); tmp < end {
				end = tmp
			}
		}
		dist = max(dist, end-i)
	}

	return dist
}

func search(data []int, target, orig int) int {
	var j, k int
	for j, k = 0, len(data)-1; j < k; {
		mid := j + (k-j)/2
		if data[mid] > target {
			k = mid
		} else {
			j = mid + 1
		}
	}

	if data[j] <= target {
		return orig
	}
	return data[j]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	when next duplicate index not found, return original value

//	2.	could be any char, including 1, 2, 3, a, b, c, ' ', etc.

//	3.	tc: O(n log m + n)

//	4.	every time use map to store char exist

//	5.	only use map to store char existence isn't working, it a duplicate
//		char encountered and remove map, but distance could still be valid

//		e.g. dvdf

//	6.	add reference https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/1730/Shortest-O(n)-DP-solution-with-explanations
