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
	var idx int

	return recursive(expression, &idx)
}

func recursive(exp string, idx *int) string {
	char := exp[*idx]

	if *idx == len(exp)-1 || exp[*idx+1] == ':' {
		*idx += 2
		return string(char)
	}

	*idx += 2

	first, second := recursive(exp, idx), recursive(exp, idx)

	if char == 'T' {
		return first
	}
	return second
}

func parseTernary2(expression string) string {
	size := len(expression)
	stack := []byte{expression[size-1]}

	var tmp byte
	for i := size - 2; i >= 0; i -= 2 {
		stack = append(stack, expression[i-1])

		if expression[i] == '?' {
			if expression[i-1] == 'T' {
				tmp = stack[len(stack)-2]
			} else {
				tmp = stack[len(stack)-3]
			}
			stack = stack[:len(stack)-3]
			stack = append(stack, tmp)
		}
	}

	return string(stack)
}

func parseTernary1(expression string) string {
	size := len(expression)

	if size > 1 {
		for i := size - 1; i >= 0; i-- {
			if expression[i] == '?' {
				if expression[i-1] == 'T' {
					return parseTernary1(expression[:i-1] + expression[i+1:i+2] + expression[i+4:])
				} else {
					return parseTernary1(expression[:i-1] + expression[i+3:])
				}
			}
		}
	}

	return expression
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

//	5.	inspired from https://leetcode.com/problems/ternary-expression-parser/discuss/92166/Very-easy-1-pass-Stack-Solution-in-JAVA-(NO-STRING-CONCAT)/96758

//		there's a one pass recursive solution

//	6.	inspired from https://leetcode.com/problems/ternary-expression-parser/discuss/92164/Easy-and-Concise-5-lines-PythonJava-Solution

//		the point of this problem is to find closest pair of unused (?, : )

//		ans since ? always has a :, if there's a ?, find unused closest : afterwards

//		the trait of this sequence can be reduced by a pair of (?, :), ans use closest, so
//		stack can be used

//	7.	inspired from https://leetcode.com/problems/ternary-expression-parser/discuss/92185/Short-Python-solutions-one-O(n)

//		should provides a O(n^2) solution, at least know how to solve it in normal
//		ways

//		since every ? pair with :, search backward to find last ?
