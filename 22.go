package main

import "strings"

//Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
//For example, given n = 3, a solution set is:
//
//[
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	permutation(n, n, []byte{}, &result)

	return result
}

// from solution: tc: O(4^n / n^0.5)
func permutation(open, closed int, current []byte, result *[]string) {
	if open == 0 && closed == 0 {
		*result = append(*result, string(current))
		return
	}

	if open == closed {
		permutation(open-1, closed, append(current, '('), result)
	} else {
		permutation(open, closed-1, append(current, ')'), result)
		if open > 0 {
			permutation(open-1, closed, append(current, '('), result)
		}
	}
}

func generateParenthesis1(n int) []string {
	result := make([]string, 0)
	brute([]byte{}, 0, 0, n, &result)
	return result
}

func brute(strs []byte, left, right, target int, result *[]string) {
	if left == target && right == target {
		var sb strings.Builder
		for _, c := range strs {
			sb.WriteByte(c)
		}
		*result = append(*result, sb.String())
		return
	}

	length := len(strs)
	if left < target {
		strs1 := make([]byte, length)
		copy(strs1, strs)
		strs1 = append(strs1, '(')
		brute(strs1, left+1, right, target, result)
	}

	if right < target && right < left {
		strs2 := make([]byte, length)
		copy(strs2, strs)
		strs2 = append(strs2, ')')
		brute(strs2, left, right+1, target, result)
	}
}

//	Notes
//	1.	too slow, it's brute force, time complexity is O( (2n!) * n), so
//		one way to improve is always generating valid parenthesis

//	2.	inspired from https://leetcode.com/problems/generate-parentheses/discuss/10105/Concise-recursive-C%2B%2B-solution

//		instead of providing n, left, right, author uses a clever way: left &
//		right starts from n

//	3.	another dp solution https://leetcode.com/problems/generate-parentheses/discuss/10369/Clean-Python-DP-Solution
