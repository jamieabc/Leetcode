package main

// Given a string S of '(' and ')' parentheses, we add the minimum number of parentheses ( '(' or ')', and in any positions ) so that the resulting parentheses string is valid.
//
// Formally, a parentheses string is valid if and only if:
//
//     It is the empty string, or
//     It can be written as AB (A concatenated with B), where A and B are valid strings, or
//     It can be written as (A), where A is a valid string.
//
// Given a parentheses string, return the minimum number of parentheses we must add to make the resulting string valid.
//
//
//
// Example 1:
//
// Input: "())"
// Output: 1
//
// Example 2:
//
// Input: "((("
// Output: 3
//
// Example 3:
//
// Input: "()"
// Output: 0
//
// Example 4:
//
// Input: "()))(("
// Output: 4
//
//
//
// Note:
//
//     S.length <= 1000
//     S only consists of '(' and ')' characters.

func minAddToMakeValid(S string) int {
	var balance, count int

	for i := range S {
		if S[i] == '(' {
			// waiting for )
			balance++
		} else {
			if balance > 0 {
				// pair (
				balance--
			} else {
				count++

				// invalid ) blocks previous valid (
				balance = 0
			}
		}
	}

	return balance + count
}

func minAddToMakeValid1(S string) int {
	stack := make([]rune, 0)
	var add int

	for _, s := range S {
		if s == '(' {
			stack = append(stack, s)
		} else {
			if len(stack) == 0 {
				add++
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return add + len(stack)
}

//	Notes
//	1.	() is a pair, which means ( is awaiting for ), and ) alone is not
//		allowed. Due to this reason, two variables are needed to show status
