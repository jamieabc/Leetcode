package main

import "sort"

//Given a list of words, each word consists of English lowercase letters.
//
//Let's say word1 is a predecessor of word2 if and only if we can add exactly one letter anywhere in word1 to make it equal to word2.  For example, "abc" is a predecessor of "abac".
//
//A word chain is a sequence of words [word_1, word_2, ..., word_k] with k >= 1, where word_1 is a predecessor of word_2, word_2 is a predecessor of word_3, and so on.
//
//Return the longest possible length of a word chain with words chosen from the given list of words.
//
//
//
//Example 1:
//
//Input: ["a","b","ba","bca","bda","bdca"]
//Output: 4
//Explanation: one of the longest word chain is "a","ba","bda","bdca".
//
//
//
//Note:
//
//    1 <= words.length <= 1000
//    1 <= words[i].length <= 16
//    words[i] only consists of English lowercase letters.

func longestStrChain(words []string) int {
	length := len(words)
	dp := make([]int, length)

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	var m int
	for i := range dp {
		for j := i - 1; j >= 0; j-- {
			if len(words[j]) < len(words[i])-1 {
				break
			}

			if chainable(words[i], words[j]) {
				dp[i] = max(dp[i], dp[j]+1)
				m = max(m, dp[i])
			}
		}
	}

	return m + 1
}

func longestStrChain1(words []string) int {
	length := len(words)
	if length <= 1 {
		return 0
	}

	var min, max int
	mapping := make(map[int][]int)
	for i := range words {
		length := len(words[i])
		if _, ok := mapping[length]; ok {
			mapping[length] = append(mapping[length], i)
		} else {
			mapping[length] = []int{i}
		}

		if min == 0 {
			min = length
		} else if length < min {
			min = length
		}

		if max == 0 {
			max = length
		} else if length > max {
			max = length
		}
	}

	dp := make([]int, length)
	var i, j, k, tmp, result int

	for i = min + 1; i <= max; i++ {
		dsts := mapping[i]
		srcs := mapping[i-1]
		if len(srcs) == 0 || len(dsts) == 0 {
			continue
		}

		for j = 0; j < len(dsts); j++ {
			tmp = 0
			for k = 0; k < len(srcs); k++ {
				if chainable(words[dsts[j]], words[srcs[k]]) {
					if dp[srcs[k]]+1 > tmp {
						tmp = dp[srcs[k]] + 1
					}
				}
			}
			dp[dsts[j]] = tmp
			if tmp > result {
				result = tmp
			}
		}
	}

	return result + 1
}

func chainable(dst, src string) bool {
	var i, j int
	var added bool
	for i, j = 0, 0; i < len(dst) && j < len(src); i++ {
		if dst[i] == src[j] {
			j++
		} else {
			if added {
				return false
			} else {
				added = true
			}
		}
	}

	// first condition is that all src are same w/ one different char
	// second condition is last char different
	return (added && j == len(src)) || (!added && i == len(dst)-1)
}

func chainable2(dst, src string) bool {
	length := len(dst)
	if length-1 != len(src) {
		return false
	}

	for i := 0; i < length; i++ {
		bytes := []byte(dst)
		tmp := append(bytes[:i], bytes[i+1:]...)
		if string(tmp) == src {
			return true
		}
	}
	return false
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

// problems
// 1. can random pick any string as predecessor
// 2. string order could be different, I am wring program that has specific order
//     e.g. a b ab abc or abc a ab b
// 3. the problem is to find maximum sequence of words, not only the last one. The
// 	  mistake I made is to assume maximum appears at last, but this isn't true
//	  because longest sequence might not at end
// 4. optimize, since for ever previous words, it always to find max. e.g. if a3 ->
//    a4, a2 -> a4, then it needs to find max(a2, a3).
//    this means no need to store in 2D array, 1D is enough
// 5. optimize, use sort.Slice to reduce memory copy and hash allocation
// 6. optimize, use map to faster operation
// 7. use same name of variable, I accidently replace the other one
// 8. seems like I don't need to sort strings
// 9. optimize, the result can be stored when traversing dp

//	10.	inspired from sample code, use map to store length n words, reduce
//		time by space

//	11.	inspired from https://leetcode.com/problems/longest-string-chain/discuss/294890/C%2B%2BJavaPython-DP-Solution

//		chainable is to loop all possible previous word with 1 letter
//		missing

//		also, can use map to store all words with same length (in index)

//		I think tc: O(n log n) + O(n * 2L * k), n: input length, L: average
//		word length, k: # of words w/ same length
