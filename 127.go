package main

import "math"

// Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:
//
//     Only one letter can be changed at a time.
//     Each transformed word must exist in the word list.
//
// Note:
//
//     Return 0 if there is no such transformation sequence.
//     All words have the same length.
//     All words contain only lowercase alphabetic characters.
//     You may assume no duplicates in the word list.
//     You may assume beginWord and endWord are non-empty and are not the same.
//
// Example 1:
//
// Input:
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]
//
// Output: 5
//
// Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
// return its length 5.
//
// Example 2:
//
// Input:
// beginWord = "hit"
// endWord = "cog"
// wordList = ["hot","dot","dog","lot","log"]
//
// Output: 0
//
// Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var found bool
	for _, w := range wordList {
		if w == endWord {
			found = true
			break
		}
	}

	if !found {
		return 0
	}

	mapping := make(map[string][]string)
	for _, w := range wordList {
		keys := toKey(w)
		for _, k := range keys {
			mapping[k] = append(mapping[k], w)
		}
	}

	type info struct {
		level int
		str   string
	}

	visited := make(map[string]int)
	normal, reversed := []info{{1, beginWord}}, []info{{-1, endWord}}

	for len(normal) != 0 {
		end := len(normal)

		// from begin word
		for i := 0; i < end; i++ {
			l, s := normal[i].level, normal[i].str
			keys := toKey(s)

			for _, k := range keys {
				changes := mapping[k]

				for _, j := range changes {
					if j == endWord {
						return l + 1
					} else if v, ok := visited[j]; ok {
						if v < 0 {
							return -v + l
						}
					} else {
						visited[j] = l + 1
						normal = append(normal, info{l + 1, j})
					}
				}
			}
		}
		normal = normal[end:]

		end = len(reversed)
		for i := 0; i < end; i++ {
			l, s := reversed[i].level, reversed[i].str
			keys := toKey(s)

			for _, k := range keys {
				changes := mapping[k]

				for _, j := range changes {
					if v, ok := visited[j]; ok {
						if v > 0 {
							return v - l
						}
					} else {
						visited[j] = l - 1
						reversed = append(reversed, info{l - 1, j})
					}
				}
			}
		}
		reversed = reversed[end:]
	}

	return 0
}

// tc: O(n * m^2), n: list length, m: word length
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	mapping := make(map[string][]string)
	for _, w := range wordList {
		keys := toKey(w)
		for _, k := range keys {
			mapping[k] = append(mapping[k], w)
		}
	}

	type info struct {
		level int
		str   string
	}

	stack := []info{{1, beginWord}}
	visited := make(map[string]bool)

	for len(stack) != 0 {
		end := len(stack)

		for _, s := range stack {
			keys := toKey(s.str)

			for _, k := range keys {
				changes := mapping[k]

				for _, j := range changes {
					if _, ok := visited[j]; !ok {
						visited[j] = true

						if endWord == j {
							return s.level + 1
						} else {
							stack = append(stack, info{s.level + 1, j})
						}
					}
				}
			}
		}

		stack = stack[end:]
	}

	return 0
}

func toKey(str string) []string {
	result := make([]string, 0)
	for i := range str {
		tmp := str[:i] + "*" + str[i+1:]
		result = append(result, tmp)
	}
	return result
}

func ladderLength1(beginWord string, endWord string, wordList []string) int {
	length := len(wordList)
	dst := -1
	for i, w := range wordList {
		if w == endWord {
			dst = i
			break
		}
	}

	// target not in list
	if dst == -1 {
		return 0
	}

	// change list to make target to be last one
	if dst != len(wordList)-1 {
		wordList[dst], wordList[length-1] = wordList[length-1], wordList[dst]
	}

	mapping := make(map[int]int)
	mapping[length-1] = 1
	stack := []int{length - 1}

	for len(stack) != 0 {
		start := len(stack)

		for i := 0; i < start; i++ {
			w := wordList[stack[i]]

			// scan for all list, find changeable
			for j := 0; j < length-1; j++ {
				if _, ok := mapping[j]; j != stack[i] && !ok && changeable(w, wordList[j]) {
					stack = append(stack, j)
					mapping[j] = mapping[stack[i]] + 1
				}
			}
		}
		stack = stack[start:]
	}

	count := math.MaxInt32
	for i := range wordList {
		if c, ok := mapping[i]; ok {
			if beginWord == wordList[i] {
				return c
			} else if changeable(beginWord, wordList[i]) {
				count = min(count, c)
			}
		}
	}

	if count == math.MaxInt32 {
		return 0
	}

	return count + 1
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func changeable(src, dst string) bool {
	var count int
	for i := range src {
		if src[i] != dst[i] {
			count++

			if count > 1 {
				return false
			}
		}
	}

	return count == 1
}

//	problems
//	1.	cannot just check if word is reachable to dst, but also check if
//		src can be transferred from list

//	2.	need to check every word in list, if src can directly into dst, use
//		it

//	3.	tc: O(mn^2), m: each world length, n: list length

//	4.	too slow, use map for checking and memorization, and turns out even
//		slower...

//		map is slower than array is item count not very big...

//	5.	when trying to put src word into list, need to update list length

//	6.	all char are same length, use int to store each word signature, if
//		a word is changeable, than difference will be 1 or -1

//	7.	inspired from solution, it's not dp, and uses BFS as I do, but
//		provides a better way of checking changeable words by mapping a word
//		in n length: dog -> [*og, d*g, do*], saves all changeable words

//		the other thing to aware is that to avoid looping in tree, needs
//		additional map to store visited item

//	8.	when doing bi-directional search, need to check if dst word in list,
//		cause whole process assumes dst is valid

//	9.	for bi-directional, it could be a problem for single char string,
//		list: [a, b, c], src: a, dst: c

//	10.	need to use to stack to traverse: one from begin and the other from
//		end

//	11.	add reference https://leetcode.com/problems/word-ladder/discuss/40707/C%2B%2B-BFS

//		author uses a more clear naming: head & tail
