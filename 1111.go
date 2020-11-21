package main

// A string is a valid parentheses string (denoted VPS) if and only if it consists of "(" and ")" characters only, and:
//
// It is the empty string, or
// It can be written as AB (A concatenated with B), where A and B are VPS's, or
// It can be written as (A), where A is a VPS.
// We can similarly define the nesting depth depth(S) of any VPS S as follows:
//
// depth("") = 0
// depth(A + B) = max(depth(A), depth(B)), where A and B are VPS's
// depth("(" + A + ")") = 1 + depth(A), where A is a VPS.
// For example,  "", "()()", and "()(()())" are VPS's (with nesting depths 0, 1, and 2), and ")(" and "(()" are not VPS's.
//
//
//
// Given a VPS seq, split it into two disjoint subsequences A and B, such that A and B are VPS's (and A.length + B.length = seq.length).
//
// Now choose any such A and B such that max(depth(A), depth(B)) is the minimum possible value.
//
// Return an answer array (of length seq.length) that encodes such a choice of A and B:  answer[i] = 0 if seq[i] is part of A, else answer[i] = 1.  Note that even though multiple answers may exist, you may return any of them.
//
//
//
// Example 1:
//
// Input: seq = "(()())"
// Output: [0,1,1,1,1,0]
// Example 2:
//
// Input: seq = "()(())()"
// Output: [0,0,0,1,1,0,1,1]
//
//
// Constraints:
//
// 1 <= seq.size <= 10000

func maxDepthAfterSplit(seq string) []int {
	result := make([]int, len(seq))

	for i := range seq {
		if seq[i] == '(' {
			result[i] = i & 1
		} else {
			result[i] = (i - 1) & 1
		}
	}

	return result
}

func maxDepthAfterSplit1(seq string) []int {
	size := len(seq)
	ans := make([]int, size)

	var left int

	for i := range seq {
		if seq[i] == '(' {
			if left&1 == 0 {
				ans[i] = 0
			} else {
				ans[i] = 1
			}
			left++
		} else {
			if left&1 > 0 {
				ans[i] = 0
			} else {
				ans[i] = 1
			}
			left--
		}
	}

	return ans
}

//	Notes
//	1.	initially cannot understand the problem, thanks for reference https://leetcode.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/discuss/358419/Confused-by-this-problem-I-was-too-but-here-is-how-it-became-crystal-clear...

//	2.	stack is not necessary, it can be determined by current char == prev char

//	3.	refactor, initial checking is not necessary

//	4.	reference from https://leetcode.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/discuss/328841/JavaC%2B%2BPython-O(1)-Extra-Space-Except-Output

//		nature of my solution is to put odd level of parentheses into one
//		group, and even level of parentheses into another group, although I
//		didn't aware of that when solving problem, someone mentions this
//		and strikes me.

//		if it's only level related, the problem convert into how to
//		determine parentheses level. this problem already provide valid
//		parentheses, so some kind of pattern can be found:

//		()() - ( at even index into one group
//		(()) - ( at odd index into another group
//		(()())(()) - ( at index 1, 3, 7 into one group
//		             ( at index 0, 6 into another group

//		the reason to have this pattern is because when parenthesis is valid
//		, every left has right. so same level ( will always be even. when
//		parentheses comes one level deep, it becomes odd index, so the
//		solution is really beautiful.
