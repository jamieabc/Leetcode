package main

import "sort"

//Let's define a function f(s) over a non-empty string s, which calculates the frequency of the smallest character in s. For example, if s = "dcce" then f(s) = 2 because the smallest character is "c" and its frequency is 2.
//
//Now, given string arrays queries and words, return an integer array answer, where each answer[i] is the number of words such that f(queries[i]) < f(W), where W is a word in words.
//
//
//
//Example 1:
//
//Input: queries = ["cbd"], words = ["zaaaz"]
//Output: [1]
//Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so f("cbd") < f("zaaaz").
//
//Example 2:
//
//Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
//Output: [1,2]
//Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").
//
//
//
//Constraints:
//
//    1 <= queries.length <= 2000
//    1 <= words.length <= 2000
//    1 <= queries[i].length, words[i].length <= 10
//    queries[i][j], words[i][j] are English lowercase letters.

func numSmallerByFrequency(queries []string, words []string) []int {
	wNums := make([]int, len(words))

	for i, w := range words {
		wNums[i] = smallestFreq(w)
	}

	sort.Ints(wNums)

	result := make([]int, len(queries))
	length := len(words)

	var left, right int
	for i, str := range queries {
		q := smallestFreq(str)
		left = 0
		right = length

		for left < right {
			mid := left + (right-left)/2
			if wNums[mid] <= q {
				left = mid + 1
			} else {
				right = mid
			}
		}

		result[i] = length - left
	}

	return result
}

func smallestFreq(str string) int {
	if str == "" {
		return 0
	}

	mapping := make([]int, 26)

	for _, s := range str {
		mapping[s-'a']++
	}

	for i := 0; i < 26; i++ {
		if mapping[i] != 0 {
			return mapping[i]
		}
	}
	return 0
}

// problems
// 1. mixed up with byte & rune
// 2. to use this way, array needs to be sorted
// 3. wrong comparison of rune and int
// 4. refactor, use binary search
// 5. use array for instead of map, better performance
// 6. wrong index, it should be bases of 'a'
// 7. wrong index of array, it should be 26
// 8. refactor, no need to store 2 slices, only 1 in necessary
