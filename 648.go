package main

// In English, we have a concept called root, which can be followed by some other word to form another longer word - let's call this word successor. For example, when the root "an" is followed by the successor word "other", we can form a new word "another".
//
// Given a dictionary consisting of many roots and a sentence consisting of words separated by spaces, replace all the successors in the sentence with the root forming it. If a successor can be replaced by more than one root, replace it with the root that has the shortest length.
//
// Return the sentence after the replacement.
//
//
//
// Example 1:
//
// Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// Output: "the cat was rat by the bat"
//
// Example 2:
//
// Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
// Output: "a a b c"
//
// Example 3:
//
// Input: dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
// Output: "a a a a a a a a bbb baba a"
//
// Example 4:
//
// Input: dictionary = ["catt","cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// Output: "the cat was rat by the bat"
//
// Example 5:
//
// Input: dictionary = ["ac","ab"], sentence = "it is abnormal that this solution is accepted"
// Output: "it is ab that this solution is ac"
//
//
//
// Constraints:
//
//     1 <= dictionary.length <= 1000
//     1 <= dictionary[i].length <= 100
//     dictionary[i] consists of only lower-case letters.
//     1 <= sentence.length <= 10^6
//     sentence consists of only lower-case letters and spaces.
//     The number of words in sentence is in the range [1, 1000]
//     The length of each word in sentence is in the range [1, 1000]
//     Each two consecutive words in sentence will be separated by exactly one space.
//     sentence does not have leading or trailing spaces.

type Trie struct {
	IsWord   bool
	Children [26]*Trie
}

func (t *Trie) build(word string, idx int) {
	if idx == len(word) {
		t.IsWord = true
		return
	}

	loc := int(word[idx] - 'a')
	if t.Children[loc] == nil {
		t.Children[loc] = &Trie{}
	}

	t.Children[loc].build(word, idx+1)
}

func (t *Trie) search(word string, idx int) string {
	if idx == len(word) {
		return word
	}

	loc := int(word[idx] - 'a')

	if t.IsWord {
		return word[:idx]
	}

	if t.Children[loc] == nil {
		return word
	}

	return t.Children[loc].search(word, idx+1)
}

func replaceWords(dictionary []string, sentence string) string {
	root := &Trie{}

	for _, word := range dictionary {
		root.build(word, 0)
	}

	ans := make([]byte, 0)
	size := len(sentence)

	var j int
	for i := 0; i < size; i++ {
		if sentence[i] == ' ' {
			ans = append(ans, byte(' '))
		} else {
			for j = i; j < size && sentence[j] != ' '; j++ {
			}

			ans = append(ans, []byte(root.search(sentence[i:j], 0))...)
			i = j - 1
		}
	}

	return string(ans)
}

//	Notes
//	1.	when search, there could be no match, which means need to check index out
//		of bound condition

//	2.	inspired from https://leetcode.com/problems/replace-words/discuss/105767/Java-SimpleClassical-Trie-questionsolution-(Beat-96)

//		could use iterative instead of recursive to search, not implement

//	3.	inspired from https://leetcode.com/problems/replace-words/discuss/105755/Python-Straightforward-with-Explanation-(Prefix-hash-Trie-solutions)

//		could also use hashmap to store all prefix string combinations

//	4.	inspired from https://leetcode.com/problems/replace-words/discuss/105855/C%2B%2B-trie-with-optimizations-(50-ms)

//		not need to build whole trie, if there's already one, stop there
