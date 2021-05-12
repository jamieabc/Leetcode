package main

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
//
// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
//
//
//
// Example 1:
//
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// Example 2:
//
// Input: digits = ""
// Output: []
//
// Example 3:
//
// Input: digits = "2"
// Output: ["a","b","c"]
//
//
//
// Constraints:
//
// 0 <= digits.length <= 4
// digits[i] is a digit in the range ['2', '9'].

var table = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	combs := make([]string, 0)

	dfs(table, []byte{}, digits, 0, &combs)

	return combs
}

func dfs(table map[byte][]byte, current []byte, digits string, idx int, comb *[]string) {
	size := len(digits)

	if idx == size {
		*comb = append(*comb, string(current))
		return
	}

	for _, b := range table[digits[idx]] {
		tmp := append([]byte{}, current...)
		tmp = append(tmp, b)

		dfs(table, tmp, digits, idx+1, comb)
	}
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/letter-combinations-of-a-phone-number/discuss/8090/Iterative-c%2B%2B-solution-in-0ms
//
//	use array to store possible chars, arr[0] & arr[1] are "", arr[2] = "abc"
