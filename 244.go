package main

import "math"

// Design a class which receives a list of words in the constructor, and implements a method that takes two words word1 and word2 and return the shortest distance between these two words in the list. Your method will be called repeatedly many times with different parameters.
//
// Example:
// Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
//
// Input: word1 = “coding”, word2 = “practice”
// Output: 3
//
// Input: word1 = "makes", word2 = "coding"
// Output: 1
//
// Note:
// You may assume that word1 does not equal to word2, and word1 and word2 are both in the list.

type WordDistance struct {
	mapping map[string][]int
}

func Constructor(words []string) WordDistance {
	mapping := make(map[string][]int)

	for i, w := range words {
		mapping[w] = append(mapping[w], i)
	}

	return WordDistance{
		mapping: mapping,
	}
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	loc1 := this.mapping[word1]
	loc2 := this.mapping[word2]

	dist := math.MaxInt32
	for i, j := 0, 0; i < len(loc1) && j < len(loc2); {
		if loc1[i] < loc2[j] {
			dist = min(dist, loc2[j]-loc1[i])
			i++
		} else {
			dist = min(dist, loc1[i]-loc2[j])
			j++
		}
	}

	return dist
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Shortest(word1,word2);
 */

//	problems
//	1.	from sample code, author uses sort.SearchInts, return index that is
//		>= target value

//	2.	tc O(m+n), m n are position length of target strings
