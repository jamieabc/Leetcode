package main

import "sort"

//Given a text string and words (a list of strings), return all index pairs [i, j] so that the substring text[i]...text[j] is in the list of words.
//
//
//
//Example 1:
//
//Input: text = "thestoryofleetcodeandme", words = ["story","fleet","leetcode"]
//Output: [[3,7],[9,13],[10,17]]
//Example 2:
//
//Input: text = "ababa", words = ["aba","ab"]
//Output: [[0,1],[0,2],[2,3],[2,4]]
//Explanation:
//Notice that matches can overlap, see "aba" is found in [0,2] and [2,4].
//
//
//Note:
//
//All strings contains only lowercase English letters.
//It's guaranteed that all strings in words are different.
//1 <= text.length <= 100
//1 <= words.length <= 20
//1 <= words[i].length <= 50
//Return the pairs [i,j] in sorted order (i.e. sort them by their first coordinate in case of ties sort them by their second coordinate).

type Trie struct {
	IsWord   bool
	Children [26]*Trie
}

func (t *Trie) Build(str string) {
	node := t
	size := len(str)

	var i int
	for i = 0; i < size; i++ {
		idx := int(str[i] - 'a')

		if node.Children[idx] == nil {
			node.Children[idx] = &Trie{}
		}

		node = node.Children[idx]
	}

	node.IsWord = true
}

func (t *Trie) Search(str string, start int) [][]int {
	ans := make([][]int, 0)

	for node, i := t, start; i < len(str); i++ {
		idx := int(str[i] - 'a')

		if node.Children[idx] == nil {
			break
		}

		node = node.Children[idx]

		// becareful abour order, need to switch to next char and check,
		// because initial root is empty
		if node.IsWord {
			ans = append(ans, []int{start, i})
		}
	}

	return ans
}

func indexPairs(text string, words []string) [][]int {
	root := &Trie{}

	for _, w := range words {
		root.Build(w)
	}

	ans := make([][]int, 0)

	for i := range text {
		ans = append(ans, root.Search(text, i)...)
	}

	sort.Slice(ans, func(i, j int) bool {
		if ans[i][0] != ans[j][0] {
			return ans[i][0] < ans[j][0]
		}
		return ans[i][1] < ans[j][1]
	})

	return ans
}

type Trie1 struct {
	IsWord   bool
	Val      byte
	Children []*Trie1
}

func (t *Trie1) Add(bytes []byte) {
	var ptr *Trie1
	var idx int

	for ptr, idx = t, 0; idx < len(bytes); idx++ {
		pos := bytes[idx] - 'a'
		if ptr.Children[pos] == nil {
			ptr.Children[bytes[idx]-'a'] = &Trie1{
				IsWord:   idx == len(bytes)-1,
				Children: make([]*Trie1, 26),
			}
		}
		ptr = ptr.Children[pos]
	}
	ptr.IsWord = true
}

// tc: O(mk + nk), n: length of text, m: words count, k: max length among all words
func indexPairs1(text string, words []string) [][]int {
	root := &Trie1{
		Children: make([]*Trie1, 26),
	}

	// build trie
	for _, str := range words {
		root.Add([]byte(str))
	}

	result := make([][]int, 0)
	for i := range text {
		for ptr, j := root, i; j < len(text); j++ {
			ptr = ptr.Children[text[j]-'a']

			if ptr == nil {
				break
			}

			if ptr.IsWord {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

//	Notes
//	1.	direct compare, tc: O(nmk), n: text length, m: words length, k: average
//		length of word length

//	2.	trie build tc: O(mk), search tc: O(nk)

//	3.	inspired from https://leetcode.com/problems/index-pairs-of-a-string/discuss/319173/Different-Python-solutions
