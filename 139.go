package main

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.
//
// Note:
//
//    The same word in the dictionary may be reused multiple times in the segmentation.
//    You may assume the dictionary does not contain duplicate words.
//
// Example 1:
//
// Input: s = "leetcode", wordDict = ["leet", "code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".
//
// Example 2:
//
// Input: s = "applepenapple", wordDict = ["apple", "pen"]
// Output: true
// Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//             Note that you are allowed to reuse a dictionary word.
//
// Example 3:
//
// Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// Output: false

// tc: O(mnk), m: string length, n: dictionary length, k: average dictionary
// string length
func wordBreak(s string, wordDict []string) bool {
	size := len(s)

	// dp[i]: ok to be composed start from i
	dp := make([]bool, size+1)
	dp[0] = true

	for i := range dp {
		if dp[i] {
			for _, word := range wordDict {
				if i+len(word) <= size && s[i:i+len(word)] == word {
					dp[i+len(word)] = true
				}
			}
		}
	}

	return dp[size]
}

type Trie struct {
	IsWord   bool
	Children map[byte]*Trie
	Val      byte
}

func (this *Trie) Add(str string, idx int) {
	if idx == len(str) {
		this.IsWord = true
		return
	}

	if _, ok := this.Children[str[idx]]; !ok {
		this.Children[str[idx]] = &Trie{
			Children: make(map[byte]*Trie),
			Val:      str[idx],
		}
	}

	this.Children[str[idx]].Add(str, idx+1)
}

func wordBreak1(s string, wordDict []string) bool {
	trie := &Trie{
		Children: make(map[byte]*Trie),
	}

	for _, word := range wordDict {
		trie.Add(word, 0)
	}

	return traverse(trie, trie.Children[s[0]], s, 0)
}

func traverse(root, node *Trie, str string, idx int) bool {
	// not exist
	if node == nil {
		return false
	}

	// reaches end
	if idx == len(str)-1 {
		return node.IsWord
	}

	var separate bool
	if node.IsWord {
		separate = traverse(root, root.Children[str[idx+1]], str, idx+1)
	}

	return separate || traverse(root, node.Children[str[idx+1]], str, idx+1)
}

//	Notes
//	1.	when reaches end of target string, need to check current position is a
//		word end

//	2.	recurring problem: from range 0~i, can it be composed by dictionary?

//		use dp to solve, but becareful about time complexity, string compare
//		takes linear time
