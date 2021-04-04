package main

// Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
//
//
//
// Example 1:
//
// Input: s = "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()".
// Example 2:
//
// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".
// Example 3:
//
// Input: s = ""
// Output: 0
//
//
// Constraints:
//
// 0 <= s.length <= 3 * 104
// s[i] is '(', or ')'.

func longestValidParentheses(s string) int {
	var longest, left, right int

	for i := range s {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			longest = max(longest, left+right)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}

		if left == right {
			longest = max(longest, left+right)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return longest
}

func longestValidParentheses3(s string) int {
	// store previous un-paired (, if all matched, store previous invalid length
	// initial -1 pop means invalid condition
	stack := []int{-1}
	var longest int

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				longest = max(longest, i-prev)
			}
		}
	}

	return longest
}

// tc: O(n)
// sc: O(n)
func longestValidParentheses2(s string) int {
	size := len(s)

	// dp[i]: longest continuous valid parenthesis length
	dp := make([]int, size)
	var longest int

	for i := range s {
		if s[i] == ')' {
			if i > 0 {
				if s[i-1] == '(' {
					if i >= 2 {
						dp[i] = dp[i-2] + 2
					} else {
						dp[i] = 2
					}
				} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1] >= 2 {
						dp[i] = 2 + dp[i-1] + dp[i-dp[i-1]-2]
					} else {
						dp[i] = 2 + dp[i-1] + dp[i-dp[i-1]]
					}

				}
			}
			longest = max(longest, dp[i])
		}
	}

	return longest
}

// tc: O(n)
// sc: O(n)
func longestValidParentheses1(s string) int {
	size := len(s)

	// dp[i]: longest continuous valid parenthesis length
	dp := make([]int, size)
	stack := make([]int, 0)
	var longest int

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 {
				start := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if start-1 >= 0 {
					longest = max(longest, i-start+1+dp[start-1])
					dp[i] = i - start + 1 + dp[start-1]
				} else {
					longest = max(longest, i-start+1)
					dp[i] = i - start + 1
				}
			} else {
				// reset all existing data, previous will never be valid
				stack = stack[:0]
			}
		}
	}

	return longest
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	parenthesis is completed when ) encountered

//	2.	in order to get longest valid parenthesis, need to have a way to store
//		previous longest parenthesis length => dp

//	3.	although I solve it, but not the optimal code, i think follow-up would
//		be use stack or dp only, or even how to do it would extra memory

//	4.	inspired from solution, can only use dp to solve

//		parenthesis closes at ), only ) may have longest length

//		() => check previous, it could be ()(), need to concat

//		()) => check previous, it could be (()), so moving backward to find
//		start of last ), and try to forward by -2, because it could be ()(())

//	5.	inspired form solution, use stack to keep record of "longest invalid length"

//		stack is used to track previous valid position, so if encounter ), pop one from
//		stack and found there's still something, it means this is the valid parenthesis,
//		update longest

//		if after popping, there's nothing in the stack, it means this is invalid ) because
//		cannot found any previous unused (, update stack (at this time, this means invalid
//		parenthesis index)

//		that's the reason to prepend a -1 into stack, because it denotes that valid to now

//		it's kind of like finding some range using two pointers

//	6.	inspired from solution, swiping two turns: left -> right and right -> left

//		count ( and ) separately, reset counter when ) is more than ( because it's
//		never going to be valid

//		the reason is that extra ( is ok, because there might be more ) afterwards
//		but extra ) is not okay, it means invalid

//		with only one scan left -> right, it might not be sufficient in follow case
//		((((())), it's still valid, but since ( is always more than ), cannot find

//		so scan right -> left help to detect this kind of condition
