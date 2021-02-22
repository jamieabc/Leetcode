package main

import (
	"bytes"
	"sort"
)

//  Given a string and a string dictionary, find the longest string in the dictionary that can be formed by deleting some characters of the given string. If there are more than one possible results, return the longest word with the smallest lexicographical order. If there is no possible result, return the empty string.
//
// Example 1:
//
// Input:
// s = "abpcplea", d = ["ale","apple","monkey","plea"]
//
// Output:
// "apple"
//
// Example 2:
//
// Input:
// s = "abpcplea", d = ["a","b","c"]
//
// Output:
// "a"
//
// Note:
//
//     All the strings in the input will only contain lower-case letters.
//     The size of the dictionary won't exceed 1,000.
//     The length of all the strings in the input won't exceed 1,000.

func findLongestWord(s string, d []string) string {
	var longest string

	for _, str := range d {
		if len(s) < len(str) {
			continue
		}

		var i, j int
		for ; i < len(s) && j < len(str); i++ {
			if s[i] == str[j] {
				j++
			}
		}

		if j == len(str) {
			if len(str) > len(longest) {
				longest = str
			} else if len(str) == len(longest) && str < longest {
				longest = str
			}
		}
	}

	return longest
}

// direct compare
// tc: O(nx), n: # of disctionary, x: average string length
func findLongestWord3(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		if len(d[i]) != len(d[j]) {
			return len(d[i]) > len(d[j])
		}
		return d[i] < d[j]
	})

	for _, str := range d {
		// dictionary too long, never could be substring of s
		if len(str) > len(s) {
			continue
		}

		// is substring
		var i, j int
		for ; i < len(s) && j < len(str); i++ {
			if s[i] == str[j] {
				j++
			}
		}

		if j == len(str) {
			return str
		}
	}

	return ""
}

func findLongestWord2(s string, d []string) string {
	size := len(s)

	table := make([][]int, 26)
	for i := 0; i < 26; i++ {
		tmp := make([]int, size)

		// if next character not found, mark as -1
		for j := range tmp {
			tmp[j] = -1
		}

		var store int
		for j := range s {
			if int(s[j]-'a') == i {
				for ; store <= j; store++ {
					tmp[store] = j + 1
				}
			}
		}

		table[i] = tmp
	}

	// tc: O(nx log(n))
	sort.Slice(d, func(i, j int) bool {
		if len(d[i]) != len(d[j]) {
			return len(d[i]) > len(d[j])
		}
		return d[i] < d[j]
	})

	for i := range d {
		// check if s can compose word
		// becareful about index, it's not that obvious

		var j, idx int
		for ; j < len(d[i]) && idx >= 0 && idx < size; j++ {
			idx = table[d[i][j]-'a'][idx]
		}

		if j == len(d[i]) && idx != -1 {
			return d[i]
		}
	}

	return ""
}

type trie struct {
	isWord   bool
	char     byte
	children [26]*trie
}

func (t *trie) insert(str string, idx int) {
	if len(str) == idx {
		t.isWord = true
		return
	}

	if t.children[str[idx]-'a'] == nil {
		t.children[str[idx]-'a'] = &trie{
			char: str[idx],
		}
	}

	t.children[str[idx]-'a'].insert(str, idx+1)
}

func findLongestWord1(s string, d []string) string {
	root := build(d)

	var longest []byte

	recursive(s, root, 0, []byte{}, &longest)

	return string(longest)
}

func build(d []string) *trie {
	root := &trie{}

	for i := range d {
		root.insert(d[i], 0)
	}

	return root
}

func recursive(s string, node *trie, idx int, cur []byte, longest *[]byte) {
	if idx == len(s) {
		return
	}

	for i := idx; i < len(s); i++ {
		if nextNode := node.children[s[i]-'a']; nextNode != nil {
			tmp := make([]byte, len(cur)+1)
			copy(tmp, cur)
			tmp[len(tmp)-1] = s[i]

			if nextNode.isWord {
				smallerString(tmp, longest)
			}

			recursive(s, nextNode, i+1, tmp, longest)
		}
	}
}

func smallerString(cur []byte, longest *[]byte) {
	if len(cur) > len(*longest) {
		*longest = make([]byte, len(cur))
		copy(*longest, cur)
	} else if len(cur) == len(*longest) {
		if bytes.Compare(cur, *longest) < 0 {
			copy(*longest, cur)
		}
	}
}

//	Notes
//	1.	build trie for dictionary, height of trie is longest string
//		but the comparing takes long time, because for s, there are n!
//		combinations

//	2.	there's bytes.Compare(a, b) < 0 means a < b

//	3.	the other way is to build table for s, finding next position of
//		specific char, tc to build table O(26n) = O(n)

//		e.g.  a b c d a  b  c
//		   a: 1 5 5 5 5 -1 -1
//		   b: 2 6 6 6 6  6 -1

//		this table means next starting index, so for example, first char is
//		a, it means next starting index is 1

//		then sort dictionary by length, start from longest string compare
//		backward

//	4.	inspired from solution, sort dictionary by length, compare each
//		char if same length
//
//		n: # of string in dictionary, x: average string length
//
//		sorting takes O(n log(n)), but there might need to comparing each
//		character, in the worst case, every string is same length
//		then it might take O(nx log(n))

//		the other think I didn't think of is check subsequence, it takes
//		O(nx)

//		total tc: O(nx log(n) + nx)

//		it's not that hard...

//	5.	inspired from solution, it's also possible not to sort, just compare
//		each, tc: O(nx)

//	6.	I might walk into dead end, because building trie won't reduce time
//		complexity (deleting chars tc: O(2^n), need to jump out of box

//		normal comparison of substring takes O(n), not that bad
