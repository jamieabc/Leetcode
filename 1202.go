package main

// You are given a string s, and an array of pairs of indices in the string pairs where pairs[i] = [a, b] indicates 2 indices(0-indexed) of the string.
//
// You can swap the characters at any pair of indices in the given pairs any number of times.
//
// Return the lexicographically smallest string that s can be changed to after using the swaps.
//
//
//
// Example 1:
//
// Input: s = "dcab", pairs = [[0,3],[1,2]]
// Output: "bacd"
// Explaination:
// Swap s[0] and s[3], s = "bcad"
// Swap s[1] and s[2], s = "bacd"
// Example 2:
//
// Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
// Output: "abcd"
// Explaination:
// Swap s[0] and s[3], s = "bcad"
// Swap s[0] and s[2], s = "acbd"
// Swap s[1] and s[2], s = "abcd"
// Example 3:
//
// Input: s = "cba", pairs = [[0,1],[1,2]]
// Output: "abc"
// Explaination:
// Swap s[0] and s[1], s = "bca"
// Swap s[1] and s[2], s = "bac"
// Swap s[0] and s[1], s = "abc"
//
//
// Constraints:
//
// 1 <= s.length <= 10^5
// 0 <= pairs.length <= 10^5
// 0 <= pairs[i][0], pairs[i][1] < s.length
// s only contains lower case English letters.

func smallestStringWithSwaps(s string, pairs [][]int) string {
	size := len(s)
	groups := make([]int, size)
	for i := range groups {
		groups[i] = i
	}

	// union
	for _, p := range pairs {
		p1, p2 := find(groups, p[0]), find(groups, p[1])
		groups[p1] = p2
	}

	// table[root]: chars
	table := make(map[int][]int)
	for i := range s {
		root := find(groups, i)

		if _, ok := table[root]; !ok {
			table[root] = make([]int, 26)
		}

		table[root][s[i]-'a']++
	}

	bytes := make([]byte, size)

	for i := range s {
		root := find(groups, i)
		for j := range table[root] {
			if table[root][j] > 0 {
				bytes[i] = byte(j + 'a')
				table[root][j]--
				break
			}
		}
	}

	return string(bytes)
}

func find(groups []int, idx int) int {
	if groups[idx] != idx {
		groups[idx] = find(groups, groups[idx])
	}
	return groups[idx]
}

//	Notes
//	1.	use rank (how many items belong to this group) to faster union process

//	2.	inspired from https://leetcode.com/problems/smallest-string-with-swaps/discuss/388257/C%2B%2B-with-picture-union-find

//		voturbac directly append char w/ same group and sort them
