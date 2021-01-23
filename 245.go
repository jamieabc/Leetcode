package main

import (
	"math"
)

// Given a list of words and two words word1 and word2, return the shortest distance between these two words in the list.
//
// word1 and word2 may be the same and they represent two individual words in the list.
//
// Example:
// Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
//
// Input: word1 = “makes”, word2 = “coding”
// Output: 1
//
// Input: word1 = "makes", word2 = "makes"
// Output: 3
//
// Note:
// You may assume word1 and word2 are both in the list.

func shortestWordDistance(words []string, word1 string, word2 string) int {
	minDist := math.MaxInt32

	prev := -1
	same := word1 == word2
	for i := range words {
		if words[i] == word1 || words[i] == word2 {
			if prev != -1 && (same || words[prev] != words[i]) {
				minDist = min(minDist, i-prev)
			}
			prev = i
		}
	}

	return minDist
}

func shortestWordDistance1(words []string, word1 string, word2 string) int {
	shortest := len(words)

	if word1 == word2 {
		prev := -1

		for i := range words {
			if words[i] == word1 {
				if prev != -1 {
					shortest = min(shortest, i-prev)
				}
				prev = i
			}
		}

		return shortest
	}

	idx1, idx2 := -1, -1

	for i := range words {
		if words[i] == word1 {
			if idx2 != -1 {
				shortest = min(shortest, i-idx2)
			}
			idx1 = i
		} else if words[i] == word2 {
			if idx1 != -1 {
				shortest = min(shortest, i-idx1)
			}
			idx2 = i
		}
	}

	return shortest
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	same word should exist different position

//	2.	too slow, use binary search

//	3.	still too slow, iterate through array once to find

//	4. 	wrong condition, as long as larger than 0 are valid

//	5.	forget to check when same word, ith word need to match

//	6.	inspired from https://leetcode.com/problems/shortest-word-distance-iii/discuss/67095/Short-Java-solution-10-lines-O(n)-modified-from-Shortest-Word-Distance-I

//		only one extra variable is needed, not 2, and can combine same /
//		different cases together

//		also, in discussion, string comparison is expensive, no need to do in
//		loop, can make it outside
