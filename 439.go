package main

// Given a string representing arbitrarily nested ternary expressions, calculate the result of the expression. You can always assume that the given expression is valid and only consists of digits 0-9, ?, :, T and F (T and F represent True and False respectively).
//
// Note:
//
//     The length of the given string is ≤ 10000.
//     Each number will contain only one digit.
//     The conditional expressions group right-to-left (as usual in most languages).
//     The condition will always be either T or F. That is, the condition will never be a digit.
//     The result of the expression will always evaluate to either a digit 0-9, T or F.
//
// Example 1:
//
// Input: "T?2:3"
//
// Output: "2"
//
// Explanation: If true, then result is 2; otherwise result is 3.
//
// Example 2:
//
// Input: "F?1:T?4:5"
//
// Output: "4"
//
// Explanation: The conditional expressions group right-to-left. Using parenthesis, it is read/evaluated as:
//
//              "(F ? 1 : (T ? 4 : 5))"                   "(F ? 1 : (T ? 4 : 5))"
//           -> "(F ? 1 : 4)"                 or       -> "(T ? 4 : 5)"
//           -> "4"                                    -> "4"
//
// Example 3:
//
// Input: "T?T?F:5:3"
//
// Output: "F"
//
// Explanation: The conditional expressions group right-to-left. Using parenthesis, it is read/evaluated as:
//
//              "(T ? (T ? F : 5) : 3)"                   "(T ? (T ? F : 5) : 3)"
//           -> "(T ? F : 3)"                 or       -> "(T ? F : 5)"
//           -> "F"                                    -> "F"

func parseTernary(expression string) string {
	length := len(expression)
	if length <= 1 {
		return expression
	}

	stack := make([]byte, 0)

	for i := length - 1; i >= 0; i-- {
		if expression[i] == '?' {
			i--
			t := stack[len(stack)-2]
			f := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			if expression[i] == 'T' {
				stack = append(stack, t)
			} else {
				stack = append(stack, f)
			}
		} else if expression[i] != ':' {
			stack = append(stack, expression[i])
		}
	}

	return string(stack[0])
}

//	problems
//	1.	wrong understanding of question, it's grouping from right-to-left,
//		a ? is paired with a :

//	2.	question already states that each number is only one digit, so it can
//		further reduce moving index

//	3.	inspired from https://leetcode.com/problems/ternary-expression-parser/discuss/92166/Very-easy-1-pass-Stack-Solution-in-JAVA-(NO-STRING-CONCAT)

//		backward traverse, then decide by T or F, but be ware that T & F can
//		be content

//	4.	add reference https://leetcode.com/problems/ternary-expression-parser/discuss/92173/Java-O(n)-using-Binary-Tree

//		author builds tree to traverse, but I didn't take time to go through
