package main

// You are given a string s representing a list of words. Each letter in the word has one or more options.
//
//     If there is one option, the letter is represented as is.
//     If there is more than one option, then curly braces delimit the options. For example, "{a,b,c}" represents options ["a", "b", "c"].
//
// For example, if s = "a{b,c}", the first character is always 'a', but the second character can be 'b' or 'c'. The original list is ["ab", "ac"].
//
// Return all words that can be formed in this manner, sorted in lexicographical order.
//
//
//
// Example 1:
//
// Input: s = "{a,b}c{d,e}f"
// Output: ["acdf","acef","bcdf","bcef"]
//
// Example 2:
//
// Input: s = "abcd"
// Output: ["abcd"]
//
//
//
// Constraints:
//
//     1 <= s.length <= 50
//     s consists of curly brackets '{}', commas ',', and lowercase English letters.
//     s is guaranteed to be a valid input.
//     There are no nested curly brackets.
//     All characters inside a pair of consecutive opening and ending curly brackets are different.

func expand(s string) []string {
	strs := make([][]string, 0)
	size := len(s)

	// convert string into groups of string
	for i := 0; i < size; i++ {
		if s[i] == '{' {
			tmp := make([]string, 0)
			j := i + 1
			for ; j < size && s[j] != '}'; j++ {
				if s[j] != ',' {
					tmp = append(tmp, s[j:j+1])
				}
			}
			strs = append(strs, tmp)
			i = j
		} else {
			// pure char
			strs = append(strs, []string{s[i : i+1]})
		}
	}

	// make sure all chars are in order
	for i := range strs {
		sort.Slice(strs[i], func(j, k int) bool {
			return strs[i][j][0] < strs[i][k][0]
		})
	}

	ans := make([]string, 0)
	idx := make([]int, len(strs))

	for true {
		tmp := make([]byte, len(strs))

		for i := range strs {
			tmp[i] = strs[i][idx[i]][0]
		}

		ans = append(ans, string(tmp))

		// find next
		var found bool
		for i := len(strs) - 1; i >= 0; i-- {
			if idx[i] == len(strs[i])-1 {
				idx[i] = 0
			} else {
				idx[i]++
				found = true
				break
			}
		}

		if !found {
			break
		}
	}

	return ans
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/brace-expansion/discuss/314308/Python-3-line-Using-Product

//		lee replace { => ' ', } => ' ' , then split string by ' '
//		very clever...
