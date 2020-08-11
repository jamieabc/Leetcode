package main

//Given two strings s and t, determine if they are isomorphic.
//
//Two strings are isomorphic if the characters in s can be replaced to get t.
//
//All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.
//
//Example 1:
//
//Input: s = "egg", t = "add"
//Output: true
//
//Example 2:
//
//Input: s = "foo", t = "bar"
//Output: false
//
//Example 3:
//
//Input: s = "paper", t = "title"
//Output: true
//
//Note:
//You may assume both s and t have the same length.
func isIsomorphic(s string, t string) bool {
	// make sure every mapping from s -> t is unique
	mapping := make([]int, 256)
	for i := range mapping {
		mapping[i] = -1
	}

	var idx1, idx2 int
	for i := range s {
		idx1, idx2 = int(s[i]), int(t[i])

		if mapping[idx1] > -1 && mapping[idx1] != idx2 {
			return false
		}

		mapping[idx1] = idx2
	}

	// make sure every mapping occurs only once
	for i := 0; i < len(mapping); i++ {
		if mapping[i] == -1 {
			continue
		}

		if i != mapping[i] && mapping[i] != mapping[mapping[i]] {
			mapping[i], mapping[mapping[i]] = mapping[mapping[i]], mapping[i]
			i--
		}
	}

	for i := range mapping {
		if mapping[i] > -1 && i != mapping[i] {
			return false
		}
	}

	return true
}

func isIsomorphic1(s string, t string) bool {
	mapping1 := make(map[rune]rune)

	for i, c := range s {
		if _, ok := mapping1[c]; !ok {
			mapping1[c] = rune(t[i])
		} else {
			if mapping1[c] != rune(t[i]) {
				return false
			}
		}
	}

	mapping2 := make(map[rune]rune)

	for i, c := range t {
		if _, ok := mapping2[c]; !ok {
			mapping2[c] = rune(s[i])
		} else {
			if mapping2[c] != rune(s[i]) {
				return false
			}
		}
	}

	return true
}

//	problems
//	1.	characters mapping, not only alphabet
