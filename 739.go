package main

//  Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.
//
// For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].
//
// Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].

func dailyTemperatures(T []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(T))

	for i := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			idx := stack[len(stack)-1]
			result[idx] = i - idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}

//	problems
//	1.	inspired by https://leetcode.com/problems/daily-temperatures/discuss/109832/Java-Easy-AC-Solution-with-Stack

//		no need to have if checking, because stack size checking is also done

//		the reason to have j is to know what's the distance from current
//		temperature, so author saves index into stack to reduce additional
//		j checking
