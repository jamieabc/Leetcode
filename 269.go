package main

import "math"

// There is a new alien language which uses the latin alphabet. However, the order among letters are unknown to you. You receive a list of non-empty words from the dictionary, where words are sorted lexicographically by the rules of this new language. Derive the order of letters in this language.
//
// Example 1:
//
// Input:
// [
//   "wrt",
//   "wrf",
//   "er",
//   "ett",
//   "rftt"
// ]
//
// Output: "wertf"
//
// Example 2:
//
// Input:
// [
//   "z",
//   "x"
// ]
//
// Output: "zx"
//
// Example 3:
//
// Input:
// [
//   "z",
//   "x",
//   "z"
// ]
//
// Output: ""
//
// Explanation: The order is invalid, so return "".
//
// Note:
//
//     You may assume all letters are in lowercase.
//     If the order is invalid, return an empty string.
//     There may be multiple valid order of letters, return any one of them is fine.

func alienOrder(words []string) string {
	graph, inDegree := buildGraph(words)

	//     for key, val := range graph {
	//         m := make([]string, 0)
	//         for _, s := range val {
	//             m = append(m, string(s))
	//         }
	//         fmt.Printf("%s: %v\n", string(key), m)
	//     }

	// for i := range inDegree {
	//     fmt.Printf("%s: %d\n", string(byte('a'+i)), inDegree[i])
	// }

	var wordCount int
	for i := range inDegree {
		if inDegree[i] != -1 {
			wordCount++
		}
	}

	order := topologicalSort(graph, inDegree)

	if len(order) == wordCount {
		return string(order)
	}
	return ""
}

func topologicalSort(graph map[byte][]byte, inDegree []int) []byte {
	sources := make([]byte, 0)
	for i := range inDegree {
		if inDegree[i] == 0 {
			sources = append(sources, byte(i+'a'))
		}
	}

	result := make([]byte, 0)

	for len(sources) > 0 {
		next := sources[0]
		sources = sources[1:]

		result = append(result, next)

		for _, n := range graph[next] {
			inDegree[n-'a']--
			if inDegree[n-'a'] == 0 {
				sources = append(sources, n)
			}
		}
	}

	return result
}

func buildGraph(words []string) (map[byte][]byte, []int) {
	graph := make(map[byte][]byte)
	inDegree := make([]int, 26)
	for i := range inDegree {
		inDegree[i] = -1
	}

	// reset every existing char in-degree to 0
	for _, word := range words {
		for i := range word {
			inDegree[word[i]-'a'] = 0
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i][0] == words[j][0] {
				// compare to find first difference
				var k int
				for k = 1; k < len(words[i]) && k < len(words[j]); k++ {
					if words[i][k] != words[j][k] {
						graph[words[i][k]] = append(graph[words[i][k]], words[j][k])
						addDegree(inDegree, words[j][k])
						k++
						break
					}
				}
				// words[i] longer than words[j] and part of them are same,
				// since it's lexical order, this should not happen
				// e.g. ["abc", "ab"]
				if k == len(words[j]) && len(words[i]) > len(words[j]) {
					inDegree[words[i][0]-'a'] = math.MaxInt32
				}
			} else {
				// change word, add another relationship since list in lexical
				// order
				graph[words[i][0]] = append(graph[words[i][0]], words[j][0])
				addDegree(inDegree, words[j][0])

				i = j - 1
				break
			}
		}
	}

	return graph, inDegree
}

func addDegree(inDegree []int, b byte) {
	idx := b - 'a'
	if inDegree[idx] == -1 {
		inDegree[idx] = 1
	} else {
		inDegree[idx]++
	}
}

//	problems
//	1.	for every word first char, need to put it into graph

//	2.	for char after first difference, also need to add them into graph

//	3.	inspired from https://leetcode.com/problems/alien-dictionary/discuss/545020/%22abc%22%22ab%22-expected-%22%22

//		word is in lexical order, so if all char are same and one reaches end,
//		shorter one should be earlier

//	4.	initialize all existing char to 0

//	5.	clarification https://leetcode.com/problems/alien-dictionary/discuss/70111/The-description-is-wrong
