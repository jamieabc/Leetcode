package main

// Given a list of unique words, return all the pairs of the distinct indices (i, j) in the given list, so that the concatenation of the two words words[i] + words[j] is a palindrome.

// Example 1:

// Input: words = ["abcd","dcba","lls","s","sssll"]
// Output: [[0,1],[1,0],[3,2],[2,4]]
// Explanation: The palindromes are ["dcbaabcd","abcddcba","slls","llssssll"]

// Example 2:

// Input: words = ["bat","tab","cat"]
// Output: [[0,1],[1,0]]
// Explanation: The palindromes are ["battab","tabbat"]

// Example 3:

// Input: words = ["a",""]
// Output: [[0,1],[1,0]]

// Constraints:

//     1 <= words.length <= 5000
//     0 <= words[i].length <= 300
//     words[i] consists of lower-case English letters.

type Trie struct {
	IsWord bool
	Children [26]*Trie
}

// reverse str to build prefix tree
func (t *Trie) Insert(str string, idx int) {

}

func (t *Trie) IsPalindrome() (bool, []int) {

}

func palindromePairs(words []string) [][]int {
	root := &Trie{}

	for i, word := range words {
		root.Insert(word, i)
	}

	pairs := make([][]int, 0)

	for _, word := range words {

	}

	return pairs
}

//	Notes
//	1.	naive way: combination n pick 2, combine string & scan, O(n^2 * n) = O(n^3)
//
//	2.	to reduce combination, which is a waste of time, use trie (prefix tree) to find directly expected
//		string that has same reverse order as target string, this part takes O(n)
//
//		after one string reaches end, start dfs to check remain candidates exist palindrome, this part
//		takes O(n)
//
//		overall tc: O(n^2)
