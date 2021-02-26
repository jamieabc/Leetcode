package main

// Given a balanced parentheses string S, compute the score of the string based on the following rule:
//
//     () has score 1
//     AB has score A + B, where A and B are balanced parentheses strings.
//     (A) has score 2 * A, where A is a balanced parentheses string.
//
//
//
// Example 1:
//
// Input: "()"
// Output: 1
//
// Example 2:
//
// Input: "(())"
// Output: 2
//
// Example 3:
//
// Input: "()()"
// Output: 2
//
// Example 4:
//
// Input: "(()(()))"
// Output: 6
//
//
//
// Note:
//
//     S is a balanced parentheses string, containing only ( and ).
//     2 <= S.length <= 50

// tc: O(n), sc: O(1)
func scoreOfParentheses(S string) int {
	var ans, level int

	for i := range S {
		if S[i] == '(' {
			level++
		} else {
			level--
		}

		if S[i] == ')' && S[i-1] == '(' {
			ans += 1 << level
		}
	}

	return ans
}

// tc: O(n), sc: O(n)
func scoreOfParentheses2(S string) int {
	stack := make([]int, 0)

	for i := range S {
		if S[i] == '(' {
			stack = append(stack, 0)
		} else {
			var tmp int

			// encounter 0, means it should be doulbed
			// encounter any number > 0, means +
			for len(stack) > 0 && stack[len(stack)-1] > 0 {
				tmp += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			stack[len(stack)-1] = max(1, tmp*2)
		}
	}

	var ans int

	for _, n := range stack {
		ans += n
	}

	return ans
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

// tc: O(n^2)
func scoreOfParentheses1(s string) int {
	return recursive(s, 0, len(s)-1)
}

func recursive(s string, start, end int) int {
	var ans, bal int

	for i := start; i <= end; i++ {
		if s[i] == '(' {
			bal++
		} else {
			bal--
		}

		if bal == 0 {
			if i == start+1 {
				ans++
			} else {
				ans += 2 * recursive(s, start+1, i-1)
			}
			start = i + 1
		}
	}

	return ans
}

//	Notes
//	1.	didn't think of way to solve it

//	2.	inspired from solution, divide and conquer, find balanced parenthesis,
//		then every parenthesis group becomes parallel, which means plus
//		relationship

//	3.	need to reset start position after each balanced pair found, because
//		recursion depends on this

//	4.	inspired from solution, use stack to solve it, but I don't really see
//		the pattern of this

//		but somehow I realized that, pushing 0 into stack means this value
//		will be updated later, it fits the problem target: value is determined
//		by later parenthesis

//		0 means outer parenthesis, > 0 means already parsed parenthesis

//	5.	inspired from https://leetcode.com/problems/score-of-parentheses/discuss/141777/C%2B%2BJavaPython-O(1)-Space

//		lee has a very good insight about this problem, whenever popping,
//		value is max of (1, *2), and new level will reset calculation and put
//		back to stack. stack is a kind of history...but, I have difficulty
//		understand his solution, will keep record for now

//		and, how to solve it in O(1)...

//		lee also provides a very brilliant solution by tc O(n) & sp O(1), key
//		point is to know how many levels of deep from start, and avoid )) case,
//		because it's already considered
