package main

//Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.
//
//Example 1:
//
//Input: 2
//Output: [0,1,1]
//Example 2:
//
//Input: 5
//Output: [0,1,1,2,1,2]
//Follow up:
//
//It is very easy to come up with a solution with run time O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a single pass?
//Space complexity should be O(n).
//Can you do it like a boss? Do it without using any builtin function like __builtin_popcount in c++ or in any other language.

// https://www.youtube.com/watch?v=QjEyO1137cM
func countBits(num int) []int {
	dp := make([]int, num+1)

	if num >= 1 {
		dp[1] = 1
	}

	if num >= 2 {
		dp[2] = 1
	}

	if num >= 3 {
		for i, j := 2, 1; i+j <= num; {
			if j == i {
				dp[i*2] = 1
				j = 1
				i *= 2
			} else {
				dp[i+j] = dp[j] + 1
				j++
			}
		}
	}

	return dp
}

//	1. 	optimization, use dp to reduce duplicated operation
