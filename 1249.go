package main

// Given a string s of '(' , ')' and lowercase English characters.
//
// Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting parentheses string is valid and return any valid string.
//
// Formally, a parentheses string is valid if and only if:
//
// It is the empty string, contains only lowercase characters, or
// It can be written as AB (A concatenated with B), where A and B are valid strings, or
// It can be written as (A), where A is a valid string.
//
//
//
// Example 1:
//
// Input: s = "lee(t(c)o)de)"
// Output: "lee(t(c)o)de"
// Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
//
// Example 2:
//
// Input: s = "a)b(c)d"
// Output: "ab(c)d"
//
// Example 3:
//
// Input: s = "))(("
// Output: ""
// Explanation: An empty string is also valid.
//
// Example 4:
//
// Input: s = "(a(b(c)d)"
// Output: "a(b(c)d)"
//
//
//
// Constraints:
//
// 1 <= s.length <= 10^5
// s[i] is one of  '(' , ')' and lowercase English letters.

func minRemoveToMakeValid(s string) string {
	size := len(s)
	table := make([]bool, size)
	var open int

	for i := range s {
		if s[i] == '(' {
			open++
			table[i] = true
		} else if s[i] == ')' {
			if open > 0 {
				// paired with (, reduce open by 1
				table[i] = true
				open--
			} else {
				// not enougn ( to pair
				table[i] = false
			}
		} else {
			table[i] = true
		}
	}

	// mark orpahn ( as remove
	for i := size - 1; open > 0; i-- {
		if s[i] == '(' {
			open--
			table[i] = false
		}
	}

	ans := make([]byte, 0)

	for i := range s {
		if table[i] {
			ans = append(ans, s[i])
		}
	}

	return string(ans)
}

// tc: O(n)
func minRemoveToMakeValid2(s string) string {
	open := make([]int, 0)
	toRemove := make(map[int]bool)

	for i := range s {
		if s[i] == '(' {
			open = append(open, i)
		} else if s[i] == ')' {
			if len(open) == 0 {
				toRemove[i] = true
			} else {
				// paired, remove from list
				open = open[:len(open)-1]
			}
		}
	}

	for _, i := range open {
		toRemove[i] = true
	}

	ans := make([]byte, 0)

	for i := range s {
		if _, ok := toRemove[i]; !ok {
			ans = append(ans, s[i])
		}
	}

	return string(ans)
}

// tc: O(n)
func minRemoveToMakeValid1(s string) string {
	// index in left/right means need to be removed
	left, right := make([]int, 0), make([]int, 0)

	for i := range s {
		if s[i] == '(' {
			left = append(left, i)
		} else if s[i] == ')' {
			if len(left) > 0 {
				// use closest pair
				left = left[:len(left)-1]
			} else {
				right = append(right, i)
			}
		}
	}

	ans := make([]byte, 0)

	for i := range s {
		if s[i] == '(' && len(left) > 0 && left[0] == i {
			left = left[1:]
		} else if s[i] == ')' && len(right) > 0 && right[0] == i {
			right = right[1:]
		} else {
			ans = append(ans, s[i])
		}
	}

	return string(ans)
}

//	Notes
//	1.	to have minimum deletion, it means to use as many pairs of parenthesis as
//		possible, for ( can pair with ), but ) cannot exist alone

//		to make sure ( and ) are paired, checks can only happen when ) encountered

//	2.	inspired form solution, could use one map & stack to store orphan
//		parenthesis

//	3.	inspired from sample code, since ) can match to any (, but ( might not

//		e.g.
//		() is ok
//		(() need to remove one (, first or last both okay
//		()( need to remove one (, must be last one

//		from above example, if there's un-balanced parenthesis, remove from last
//		guarantees always work
