package main

import "sort"

// Given a list of strings words representing an English Dictionary, find the longest word in words that can be built one character at a time by other words in words. If there is more than one possible answer, return the longest word with the smallest lexicographical order.
//
// If there is no answer, return the empty string.
// Example 1:
//
// Input:
// words = ["w","wo","wor","worl", "world"]
// Output: "world"
// Explanation:
// The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
// Example 2:
//
// Input:
// words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
// Output: "apple"
// Explanation:
// Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
// Note:
//
// All the strings in the input will only contain lowercase letters.
// The length of words will be in the range [1, 1000].
// The length of words[i] will be in the range [1, 30].

type Trie struct {
	IsWord   bool
	Children [26]*Trie
}

func (t *Trie) Build(str string) int {
	node := t

	var idx int
	for ; idx < len(str)-1; idx++ {
		if !node.IsWord || node.Children[str[idx]-'a'] == nil {
			return idx
		}

		node = node.Children[str[idx]-'a']
	}

	if node.Children[str[idx]-'a'] == nil {
		node.Children[str[idx]-'a'] = &Trie{
			IsWord: true,
		}
	}

	return len(str)
}

// tc: O(n log(n))
func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	root := &Trie{IsWord: true}

	var longest int
	var str string

	for _, word := range words {
		cur := root.Build(word)
		if cur > longest {
			longest = cur
			str = word
		}
	}

	return str
}

// tc: O(n log n)
func longestWord3(words []string) string {
	sort.Strings(words)

	table := make(map[string]bool)
	table[""] = true
	maxSize, result := 0, ""

	for _, word := range words {
		if _, ok := table[word[:len(word)-1]]; ok {
			table[word] = true

			if len(word) > maxSize || (len(word) == maxSize && word < result) {
				result = word
				maxSize = len(word)
			}
		}
	}

	return result
}

type Trie1 struct {
	Val      byte
	Children map[byte]*Trie1
	IsWord   bool
}

func (t *Trie1) Insert(str string) {
	node := t
	for i := range str {
		if _, ok := node.Children[str[i]]; !ok {
			node.Children[str[i]] = &Trie1{
				Val:      str[i],
				Children: make(map[byte]*Trie1),
			}
		}
		node = node.Children[str[i]]
	}
	node.IsWord = true
}

func (t *Trie1) Dfs(prev string) string {
	var current string
	if t.Val != 0 {
		current = prev + string(t.Val)
	}

	str := current

	for _, n := range t.Children {
		if !n.IsWord {
			continue
		}

		tmp := n.Dfs(current)
		if len(tmp) > len(str) || (len(tmp) == len(str) && tmp < str) {
			str = tmp
		}
	}

	return str
}

func longestWord2(words []string) string {
	t := &Trie1{
		Children: make(map[byte]*Trie1),
	}

	// create trie
	for _, word := range words {
		t.Insert(word)
	}

	// find longest
	return t.Dfs("")
}

func longestWord1(words []string) string {
	var maxSize int
	counter := make(map[int]map[string]bool)

	for _, word := range words {
		if _, ok := counter[len(word)]; !ok {
			counter[len(word)] = make(map[string]bool)
		}
		counter[len(word)][word] = true

		if len(word) > maxSize {
			maxSize = len(word)
		}
	}

	var result string
	var j int

	for i := maxSize; i >= 1; i-- {
		for str := range counter[i] {
			// check if this word can be come up from shorter words
			for j = len(str) - 1; j > 0; j-- {
				if !counter[j][str[:j]] {
					break
				}
			}

			// valid word
			if j == 0 {
				if len(result) == 0 {
					result = str
				} else {
					// choose smaller lexical
					for k := 0; k < len(str); k++ {
						if str[k] > result[k] {
							break
						} else if str[k] < result[k] {
							result = str
						}
					}
				}
			}
		}

		if len(result) > 0 {
			return string(result)
		}
	}

	return ""
}

//	Notes
//	1.	wrong logic about choosing smaller lexical, because it should top when
//		char is larger than original, or update when char is smaller

//	2.	inspired from sample code

//		I was trying to build a hashmap for size to words, then start from
//		longest word, traverse down to check if it can be build from scratch

//		but solution provides a way to sort words, then build from smallest
//		I think this is similar to do searching, choose one after search all
//		afterwards; if not found, jump to next one and only search from that one
//		afterwards

//	3.	when building trie, this problem is different from previous one, it
//		needs to check if all previous words exist

//	4.	for trie to build word, check all previous conditions
