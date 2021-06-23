package main

// Given a string s and an array of strings words, return the number of words[i] that is a subsequence of s.

// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

//     For example, "ace" is a subsequence of "abcde".



// Example 1:

// Input: s = "abcde", words = ["a","bb","acd","ace"]
// Output: 3
// Explanation: There are three strings in words that are a subsequence of s: "a", "acd", "ace".

// Example 2:

// Input: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
// Output: 2



// Constraints:

//     1 <= s.length <= 5 * 104
//     1 <= words.length <= 5000
//     1 <= words[i].length <= 50
//     s and words[i] consist of only lowercase English letters.

func numMatchingSubseq(s string, words []string) int {
	var count int

	table := make(map[rune][]string)
	for _, word := range words {
		c := rune(word[0])
		table[c] = append(table[c], word)
	}

	for _, c := range s {
		array := table[c]
		table[c] = []string{}

		for _, word := range array {
			if len(word) == 1 {
				count++
				continue
			}

			c := rune(word[1])
			table[c] = append(table[c], word[1:])
		}
	}

	return count
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/number-of-matching-subsequences/discuss/117634/Efficient-and-simple-go-through-words-in-parallel-with-explanation
