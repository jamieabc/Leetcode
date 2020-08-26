package main

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
	Val      byte
	Children []*Trie
}

func (t *Trie) Add(bytes []byte) {
	var ptr *Trie
	var idx int

	for ptr, idx = t, 0; idx < len(bytes); idx++ {
		pos := bytes[idx] - 'a'
		if ptr.Children[pos] == nil {
			ptr.Children[bytes[idx]-'a'] = &Trie{
				IsWord:   idx == len(bytes)-1,
				Children: make([]*Trie, 26),
			}
		}
		ptr = ptr.Children[pos]
	}
	ptr.IsWord = true
}

// tc: O(mk + nk), n: length of text, m: words count, k: max length among all words
func indexPairs(text string, words []string) [][]int {
	root := &Trie{
		Children: make([]*Trie, 26),
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
