package main

import "sort"

// Design a search autocomplete system for a search engine. Users may input a sentence (at least one word and end with a special character '#'). For each character they type except '#', you need to return the top 3 historical hot sentences that have prefix the same as the part of sentence already typed. Here are the specific rules:
//
//     The hot degree for a sentence is defined as the number of times a user typed the exactly same sentence before.
//     The returned top 3 hot sentences should be sorted by hot degree (The first is the hottest one). If several sentences have the same degree of hot, you need to use ASCII-code order (smaller one appears first).
//     If less than 3 hot sentences exist, then just return as many as you can.
//     When the input is a special character, it means the sentence ends, and in this case, you need to return an empty list.
//
// Your job is to implement the following functions:
//
// The constructor function:
//
// AutocompleteSystem(String[] sentences, int[] times): This is the constructor. The input is historical data. Sentences is a string array consists of previously typed sentences. Times is the corresponding times a sentence has been typed. Your system should record these historical data.
//
// Now, the user wants to input a new sentence. The following function will provide the next character the user types:
//
// List<String> input(char c): The input c is the next character typed by the user. The character will only be lower-case letters ('a' to 'z'), blank space (' ') or a special character ('#'). Also, the previously typed sentence should be recorded in your system. The output will be the top 3 historical hot sentences that have prefix the same as the part of sentence already typed.
//
//
// Example:
// Operation: AutocompleteSystem(["i love you", "island","ironman", "i love leetcode"], [5,3,2,2])
// The system have already tracked down the following sentences and their corresponding times:
// "i love you" : 5 times
// "island" : 3 times
// "ironman" : 2 times
// "i love leetcode" : 2 times
// Now, the user begins another search:
//
// Operation: input('i')
// Output: ["i love you", "island","i love leetcode"]
// Explanation:
// There are four sentences that have prefix "i". Among them, "ironman" and "i love leetcode" have same hot degree. Since ' ' has ASCII code 32 and 'r' has ASCII code 114, "i love leetcode" should be in front of "ironman". Also we only need to output top 3 hot sentences, so "ironman" will be ignored.
//
// Operation: input(' ')
// Output: ["i love you","i love leetcode"]
// Explanation:
// There are only two sentences that have prefix "i ".
//
// Operation: input('a')
// Output: []
// Explanation:
// There are no sentences that have prefix "i a".
//
// Operation: input('#')
// Output: []
// Explanation:
// The user finished the input, the sentence "i a" should be saved as a historical sentence in system. And the following input will be counted as a new search.
//
//
// Note:
//
//     The input sentence will always start with a letter and end with '#', and only one blank space will exist between two words.
//     The number of complete sentences that to be searched won't exceed 100. The length of each sentence including those in the historical data won't exceed 100.
//     Please use double-quote instead of single-quote when you write test cases even for a character input.
//     Please remember to RESET your class variables declared in class AutocompleteSystem, as static/class variables are persisted across multiple test cases. Please see here for more details.

type Node struct {
	// children   map[byte]*Node
	children   []*Node
	candidates []string
}

type AutocompleteSystem struct {
	root, cur *Node
	typed     []byte
	counts    map[string]int
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	root := &Node{
		children: make([]*Node, 27),
	}

	counts := make(map[string]int)

	// for all sentences
	for i, s := range sentences {
		// for all chars in one sentence
		add(root, s)
		counts[s] = times[i]
	}

	return AutocompleteSystem{
		root:   root,
		cur:    root,
		counts: counts,
		typed:  make([]byte, 0),
	}
}

func add(root *Node, str string) {
	cur := root

	for j := range str {
		idx := toInt(str[j])
		if cur.children[idx] == nil {
			cur.children[idx] = &Node{
				children:   make([]*Node, 27),
				candidates: make([]string, 0),
			}
		}
		cur = cur.children[idx]

		// make sure added string is not duplicated
		cur.candidates = append(cur.candidates, str)
	}
}

func toInt(c byte) int {
	if c == ' ' {
		return 26
	}
	return int(c - 'a')
}

func (this *AutocompleteSystem) Input(c byte) []string {
	var result []string
	str := string(this.typed)
	if c == '#' {
		if _, ok := this.counts[str]; !ok {
			add(this.root, str)
			this.counts[str] = 1
		} else {
			this.counts[str]++
		}

		this.cur = this.root
		this.typed = this.typed[:0]
		return result
	}

	if this.cur != nil {
		this.cur = this.cur.children[toInt(c)]
		result = sorted(this.cur, this.counts)
	}

	this.typed = append(this.typed, c)

	return result
}

func sorted(node *Node, counts map[string]int) []string {
	if node == nil {
		return []string{}
	}

	sort.Slice(node.candidates, func(i, j int) bool {
		s1 := node.candidates[i]
		s2 := node.candidates[j]
		if counts[s1] == counts[s2] {
			return s1 < s2
		}
		return counts[s1] >= counts[s2]
	})

	if len(node.candidates) <= 3 {
		return node.candidates
	}
	return node.candidates[:3]
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */

//	problems
//	1.	inspired from solution, easiest way is to brute force with a
//		map(string)int to denote sentences and times

//		store every input char and compares to map keys. if keys match first
//		n characters, add to result list. then sort result list in
//		alphabetical order

//		constructor tc: O(kl), k: sentences average length k, l: # of
//		sentences

//		input tc: O(n + m log m), n: # of sentences, m: sort possible list
//		of size m

//	2.	inspired from solution, indexing a map by its first char,
//		[26]map(string)int

//		constructor tc: O(kl + 26), k: sentences average length, l: # of
//		sentences

//		input tc: O(s + m log m), s: sentences that match first character,
//		m: sort possible list of size m

//	3.	add reference https://leetcode.com/problems/design-search-autocomplete-system/discuss/105386/Python-Clean-Solution-Using-Trie

//		comment points out result is sorted by times in descending order,
//		if occur time is same, then sentence is sorted in char ascending
//		order

//	4.	inspired from https://leetcode.com/problems/design-search-autocomplete-system/discuss/105376/Java-solution-Trie-and-PriorityQueue

//		author stores sentences at each level, so that when input comes,
//		it's quicker to find all possibilities

//	5.	sort fails, e.g. count: abc: 2, def: 2 turn out to be [abc def
//		abc def], because there are 2 2's

//	6.	inspired from sample code, use [27] to store nodes. also, it could
//		exists only one global count for sentences, no need to store in trie

//		also, for string comparison, can use s1 < s2
