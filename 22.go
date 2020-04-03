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

//	problems
//	1.	too slow, it's brute force, time complexity is O( (2n!) * n), so
//		one way to improve is always generating valid parenthesis
