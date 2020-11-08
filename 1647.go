package main

import "sort"

// A string s is called good if there are no two different characters in s that have the same frequency.
//
// Given a string s, return the minimum number of characters you need to delete to make s good.
//
// The frequency of a character in a string is the number of times it appears in the string. For example, in the string "aab", the frequency of 'a' is 2, while the frequency of 'b' is 1.
//
//
//
// Example 1:
//
// Input: s = "aab"
// Output: 0
// Explanation: s is already good.
//
// Example 2:
//
// Input: s = "aaabbbcc"
// Output: 2
// Explanation: You can delete two 'b's resulting in the good string "aaabcc".
// Another way it to delete one 'b' and one 'c' resulting in the good string "aaabbc".
//
// Example 3:
//
// Input: s = "ceabaacb"
// Output: 2
// Explanation: You can delete both 'c's resulting in the good string "eabaab".
// Note that we only care about characters that are still in the string at the end (i.e. frequency of 0 is ignored).
//
//
//
// Constraints:
//
//     1 <= s.length <= 105
//     s contains only lowercase English letters.

func minDeletions(s string) int {
	counter := make([]int, 26)

	for i := range s {
		counter[s[i]-'a']++
	}

	// sort string in descending order
	sort.Slice(counter, func(i, j int) bool {
		return counter[i] > counter[j]
	})

	var moves int
	for i := 1; i < len(counter); i++ {
		if counter[i] == 0 {
			break
		}

		if counter[i] >= counter[i-1] {
			// at most remove all characters
			count := max(0, counter[i-1]-1)
			moves += counter[i] - count
			counter[i] = count
		}
	}

	return moves
}

// tc: O(n)
func minDeletions1(s string) int {
	counter := make([]int, 26)

	for i := range s {
		counter[s[i]-'a']++
	}

	table := make(map[int]bool)
	var removed int

	for _, c := range counter {
		if c == 0 {
			continue
		}

		for _, ok := table[c]; c > 0 && ok; {
			c--
			removed++
			_, ok = table[c]
		}

		if c > 0 {
			table[c] = true
		}
	}

	return removed
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	it's not always to find smallest, but smaller one without any frequency

//	2.	inspired from alex, for char frequency, this is actually turn frequency
//		array into descending

//	3.	to use greedy, important thing is to find smaller count that has not
//		used, so the point is to find a set stores occurrence of integer
