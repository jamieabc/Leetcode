package main

//  Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.
//
// For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].
//
// Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].

func dailyTemperatures(T []int) []int {
	size := len(T)
	ans := make([]int, size)

	for i := size - 1; i >= 0; i-- {
		for j := i - 1; j >= 0 && T[j] < T[i]; j-- {
			ans[j] = i - j
		}
	}

	return ans
}

func dailyTemperatures1(T []int) []int {
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

//	Notes
//	1.	inspired by https://leetcode.com/problems/daily-temperatures/discuss/109832/Java-Easy-AC-Solution-with-Stack

//		no need to have if checking, because stack size checking is also done

//		the reason to have j is to know what's the distance from current
//		temperature, so author saves index into stack to reduce additional
//		j checking

//	2.	inspired from https://leetcode.com/problems/daily-temperatures/discuss/121787/C%2B%2B-Clean-code-with-explanation%3A-O(n)-time-and-O(1)-space-(beats-99.13)

//		tc remains O(n) and sc can be further reduced down to O(1), the trick is that finding
//		next larger number is a mono-increasing. viewing backward if mono-decreasing, start
//		backward, once larger number is found, all number after that will be determined by
//		that number

//		e.g. 5, 8, 9, 4, 3, 2, 10, 6
//		ans	 0, 0, 0, 0, 0, 0,  0, 0

//								v  ^		start from 6 backward, 10 is larger
//			 0, 0, 0, 0, 0, 0,  0, 1

//			 v				    ^			start from 10 backward, no any larger, update
//			 6, 5, 4, 3, 2, 1,  0, 1

//			 			 v	^    			start from 2 backward, 3 is larger
//			 6, 5, 4, 3, 2, 1,  0, 1

//			 		  v	 ^	     			start from 3 backward, 4 is larger
//			 6, 5, 4, 3, 2, 1,  0, 1

//			 	   v  ^	  	     			start from 4 backward, 9 is larger
//			 6, 5, 4, 3, 2, 1,  0, 1

//			 v	   ^   	  	     			start from 9 backward, no any larger, update
//			 2, 1, 4, 3, 2, 1,  0, 1
