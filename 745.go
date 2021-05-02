package main

import "fmt"

//Design a special dictionary which has some words and allows you to search the words in it by a prefix and a suffix.
//
//Implement the WordFilter class:
//
//WordFilter(string[] words) Initializes the object with the words in the dictionary.
//f(string prefix, string suffix) Returns the index of the word in the dictionary which has the prefix prefix and the suffix suffix. If there is more than one valid index, return the largest of them. If there is no such word in the dictionary, return -1.
//
//
//
//Example 1:
//
//Input
//["WordFilter", "f"]
//[[["apple"]], ["a", "e"]]
//Output
//[null, 0]
//
//Explanation
//WordFilter wordFilter = new WordFilter(["apple"]);
//wordFilter.f("a", "e"); // return 0, because the word at index 0 has prefix = "a" and suffix = 'e".
//
//
//
//Constraints:
//
//1 <= words.length <= 15000
//1 <= words[i].length <= 10
//1 <= prefix.length, suffix.length <= 10
//words[i], prefix and suffix consist of lower-case English letters only.
//At most 15000 calls will be made to the function f.

type Trie struct {
	Index    int
	Children [27]*Trie
}

func (root *Trie) Insert(str string, idx int) {
	node := root
	var afterSeparator bool

	for i := range str {
		if node.Children[str[i]-'a'] == nil {
			node.Children[str[i]-'a'] = &Trie{}
		}

		node = node.Children[str[i]-'a']

		if afterSeparator {
			node.Index = idx
		} else {
			node.Index = -1
		}

		if str[i] == '{' {
			afterSeparator = true
		}
	}
}

func (root *Trie) Search(str string) int {
	node := root
	var i int
	var afterSeparator bool

	for ; i < len(str); i++ {
		node = node.Children[str[i]-'a']

		if node == nil {
			break
		}

		if str[i] == '{' {
			afterSeparator = true
		}
	}

	if afterSeparator && i == len(str) {
		return node.Index
	}

	return -1
}

type WordFilter struct {
	Root *Trie
}

func Constructor(words []string) WordFilter {
	root := &Trie{
		Index: -1,
	}

	for i, word := range words {
		for j := len(word) - 1; j >= 0; j-- {
			str := fmt.Sprintf("%s{%s", word[j:], word)
			root.Insert(str, i)
		}
	}

	return WordFilter{
		Root: root,
	}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	str := fmt.Sprintf("%s{%s", suffix, prefix)
	return this.Root.Search(str)
}

type WordFilter1 struct {
	Table map[string]int
}

func Constructor1(words []string) WordFilter1 {
	table := make(map[string]int)

	for i, word := range words {
		size := len(word)

		for j := range word {
			for k := size - 1; k >= 0; k-- {
				str := fmt.Sprintf("%s-%s", word[:j+1], word[k:])
				table[str] = i
			}
		}
	}

	return WordFilter1{
		Table: table,
	}
}

func (this *WordFilter1) F(prefix string, suffix string) int {
	str := fmt.Sprintf("%s-%s", prefix, suffix)

	if idx, ok := this.Table[str]; !ok {
		return -1
	} else {
		return idx
	}
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */

//	Notes
//	1.	inspired from https://zxi.mytechroad.com/blog/tree/leetcode-745-prefix-and-suffix-search/

//		brute force way is to generate all possible prefix-suffix combinations and check
//		tc will be O(NL^3 + QL), N: size of word array, L: average string length, Q: query size

//		better way is to use trie (prefix tree), but there's also suffix tree, so it's
//		a bit tricky to apply suffix to prefix tree: reverse the suffix, so that search
//		meets original prefix tree way

//		tc: O(NL^2 + QL), sc: O(NL^2)

//		e.g	apple
//		prefix			suffix			insert
//		a				e				e_apple
//		a				le				le_apple
//		a				ple				ple_apple
//		a				pple			pple_apple
//		a				apple			apple_apple

//		the smart point is that prefix is always full word, so any break point will
//		meets

//		to make sure both prefix & suffix are matched, before separator (_), trie
//		stores index -1 to denote not match, and after separator (_), trie stores
//		actual word index

//		the other small tricky part is that char '{' - 'a' = 26, which can be matched
//		to array

//	2.	inspired from https://leetcode.com/problems/prefix-and-suffix-search/discuss/320712/Different-Python-solutions-with-thinking-process

//		could also uses two prefix trees to store, and compare all matched to find largest
//		index
