package main

//Given an integer n, return the number of trailing zeroes in n!.
//
//Example 1:
//
//Input: 3
//Output: 0
//Explanation: 3! = 6, no trailing zero.
//
//Example 2:
//
//Input: 5
//Output: 1
//Explanation: 5! = 120, one trailing zero.
//
//Note: Your solution should be in logarithmic time complexity.

func trailingZeroes(n int) int {
	if n <= 4 {
		return 0
	}

	divider := 5
	sum5 := 0

	for divider <= n {
		if divider%5 == 0 {
			sum5 += n / divider
		}
		divider *= 5
	}

	return sum5
}
