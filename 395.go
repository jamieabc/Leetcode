package main

import "math"

// Find the length of the longest substring T of a given string (consists of lowercase letters only) such that every character in T appears no less than k times.
//
// Example 1:
//
// Input:
// s = "aaabb", k = 3
//
// Output:
// 3
//
// The longest substring is "aaa", as 'a' is repeated 3 times.
//
// Example 2:
//
// Input:
// s = "ababbc", k = 2
//
// Output:
// 5
//
// The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.

func longestSubstring(s string, k int) int {
	size := len(s)
	if size == 0 || size < k {
		return 0
	}

	chars := make([]bool, 26)
	var maxUniq int
	for i := range s {
		if !chars[s[i]-'a'] {
			maxUniq++
			chars[s[i]-'a'] = true
		}
	}

	var longest, start, end int

	for i := 1; i <= maxUniq; i++ {
		counter := make([]int, 26)
		var uniq, meetCriteria int

		for start, end = 0, 0; end < size; {
			if start == end || uniq <= i {
				idx := s[end] - 'a'
				counter[idx]++
				end++

				if counter[idx] == 1 {
					uniq++
				}

				if counter[idx] == k {
					meetCriteria++
				}

				if uniq <= i && meetCriteria > 0 {
					longest = max(longest, end-start)
				}
			} else {
				idx := s[start] - 'a'
				counter[idx]--
				start++

				if counter[idx] == 0 {
					uniq--
				}

				if counter[idx] == k-1 {
					meetCriteria--
				}
			}
		}
	}

	return longest
}

// tc: average O(n), worst case O(n^2), if every time last character is removed,
// overall will be n + (n-1) + (n-2) + ... + 1 = O(n^2)
func longestSubstring3(s string, k int) int {
	var longest int

	recursive(s, 0, len(s)-1, k, &longest)

	return longest
}

func recursive(s string, start, end, target int, longest *int) {
	if end-start+1 < target {
		return
	}

	counter := make([]int, 26)
	for i := start; i <= end; i++ {
		counter[s[i]-'a']++
	}

	minCount := math.MaxInt32
	var char int
	for i := range counter {
		if counter[i] > 0 && counter[i] < target && counter[i] < minCount {
			minCount = counter[i]
			char = i
		}
	}

	// whole string meets criteria
	if minCount == math.MaxInt32 {
		*longest = max(*longest, end-start+1)
		return
	}

	for prev, i := start, start; i <= end; i++ {
		if int(s[i]-'a') == char {
			recursive(s, prev, i-1, target, longest)
			i++
			prev = i
		} else if i == end {
			recursive(s, prev, i, target, longest)
		}
	}
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func longestSubstring2(s string, k int) int {
	var maxSize int

	divide(s, k, &maxSize)

	return maxSize
}

func divide(s string, k int, maxSize *int) {
	if len(s) < k {
		return
	}

	counter := make([][]int, 26)

	for i := range s {
		counter[s[i]-'a'] = append(counter[s[i]-'a'], i)
	}

	idx, count := -1, 0

	for i, c := range counter {
		if len(c) > 0 && len(c) < k {
			if len(c) > count {
				idx, count = i, len(c)
			}
		}
	}

	if idx == -1 {
		*maxSize = max(*maxSize, len(s))
		return
	}

	var start int
	for _, c := range counter[idx] {
		divide(s[start:c], k, maxSize)
		start = c + 1
	}
	divide(s[start:], k, maxSize)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func longestSubstring1(s string, k int) int {
	var maxLength int

	for i := range s {
		maxLength = max(maxLength, check1(s, i, k))
	}

	return maxLength
}

// tc: O(n^2)
func check1(s string, idx, k int) int {
	counter := make([]int, 26)
	size := len(s)
	var maxLength int

	for i := idx; i < size; i++ {
		counter[s[i]-'a']++

		found := true
		for _, count := range counter {
			if count > 0 && count < k {
				found = false
				break
			}
		}

		if found {
			maxLength = max(maxLength, i-idx+1)
		}
	}

	return maxLength
}

//	Notes
//	1.	to find substring with all chars within appear at least k times, the goal
//		is to remove all chars appear time < k, can min-heap help to solve?

//		or use a counter to store frequency, each time scan for 26 chars and find
//		those chars appear time < k, reduce range

//		the problem is about choosing interval, say char c appear time < k
//		e.g. . . . . c . . . . c. . . . . c. . . . c. . . c. . . . . c. . .
//					 ^         ^          ^        ^      ^          ^
//    		  intr 1    intr 2     intr 3   intr 4   intr 5    intr 6   intr 7

// 		which interval to select? all intr 1 ~ 7 might be valid

//		because char appear time is always increasing, so it's safe to cut whole
//		array into smaller array apart at intr1, intr2, intr3, etc.

//		of course, it might be all c, it could be any character total appear time
//		< k

//	2.	above thinking is wrong, because it should be a recursive process.
//		e.g. bbaaacbd, k = 3
//                ^ ^

//		algorithm assumes 0 ~ c is valid, but it found bbaaa invalid, and stop
//		process

//	3.	inspired from https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/discuss/87768/4-lines-Python

//		it's basically same idea, but split string by smallest count. It's
//		divide & conquer algorithm

//	4.	another way is too select first character not meet criteria, it's by
//		seeing this process as recursion, but it takes more time than select
//		smallest appear time char

//	5.	inspired from solution, sliding window of tc: O(n) is possible

//		for a string s with maximum unique character count c. Based one c,
//		longest substring with each character inside can at most be c.
//		Iterate from 1 ~ c, using sliding window to find such range has only
//		specific unique characters with each appear time >= k

//		sliding window expands when this range unique char count not meet target
//		count, and shrinks if unique char count > target count

//		e.g. s = abababbdabcabc, k = 2
//		maximum uniq character count = 4 (a, b, c, d)

//		for uniq char count 1, found bb
//		for uniq char count 2, found abababb
//		for uniq char count 3, found abcabc
//		for uniq char count 4, found nothing

//		this is a really smart solution, while I was thinking about the problem
//		and don't know how to decide start/end of an interval, and once unique
//		count is determined, there's a way to decide start/end of an interval

//		big problem is separated into smaller subset of uniq char count, and
//		longest substring must in those conditions
