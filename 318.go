package main

// Given a string array words, return the maximum value of length(word[i]) * length(word[j]) where the two words do not share common letters. If no such two words exist, return 0.
//
//
//
// Example 1:
//
// Input: words = ["abcw","baz","foo","bar","xtfn","abcdef"]
// Output: 16
// Explanation: The two words can be "abcw", "xtfn".
//
// Example 2:
//
// Input: words = ["a","ab","abc","d","cd","bcd","abcd"]
// Output: 4
// Explanation: The two words can be "ab", "cd".
//
// Example 3:
//
// Input: words = ["a","aa","aaa","aaaa"]
// Output: 0
// Explanation: No such pair of words.
//
//
//
// Constraints:
//
//     2 <= words.length <= 1000
//     1 <= words[i].length <= 1000
//     words[i] consists only of lowercase English letters.

func maxProduct(words []string) int {
	size := len(words)
	table := make([]int, size)

	for i := range words {
		var j int

		for k := range words[i] {
			j |= 1 << int(words[i][k]-'a')
		}
		table[i] = j
	}

	var ans int

	for i := range table {
		for j := i + 1; j < size; j++ {
			if table[i]&table[j] != 0 {
				continue
			}

			ans = max(ans, len(words[i])*len(words[j]))
		}
	}

	return ans
}

func maxProduct2(words []string) int {
	table := make(map[int]int)

	for _, word := range words {
		var i int

		for j := range word {
			i |= 1 << int(word[j]-'a')
		}

		table[i] = max(table[i], len(word))
	}

	var ans int

	for i := range table {
		for j := range table {
			// same char & results larger than 0
			if i&j == 0 {
				ans = max(ans, table[i]*table[j])
			}
		}
	}

	return ans
}

// tc: O(n^2)
func maxProduct1(words []string) int {
	size := len(words)
	tables := make([][]bool, size)
	for i := range tables {
		tables[i] = make([]bool, 26)
	}

	for i := range words {
		for j := range words[i] {
			tables[i][words[i][j]-'a'] = true
		}
	}

	var ans, k int

	for i := range tables {
		for j := i + 1; j < size; j++ {
			if len(words[i])*len(words[j]) <= ans {
				continue
			}

			// compare
			for k = 0; k < 26; k++ {
				if tables[i][k] && tables[j][k] {
					break
				}
			}

			if k == 26 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}

	return ans
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from sample code, use int to denote 26 chars condition

//	2.	inspired from solution, for same group of chars ab & aabb, both contains chars
//		a & b, but only need to consider aabb because it's longer than ab
