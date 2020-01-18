package main

//You are climbing a stair case. It takes n steps to reach to the top.
//
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
//Note: Given n will be a positive integer.
//
//Example 1:
//
//Input: 2
//Output: 2
//Explanation: There are two ways to climb to the top.
//1. 1 step + 1 step
//2. 2 steps
//
//Example 2:
//
//Input: 3
//Output: 3
//Explanation: There are three ways to climb to the top.
//1. 1 step + 1 step + 1 step
//2. 1 step + 2 steps
//3. 2 steps + 1 step

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return dp(n)
}

func dp(n int) int {
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2

	for i := 3; i <= n; i++ {
		arr[i] = arr[i-2] + arr[i-1]
	}

	return arr[n]
}

// problems
// 1. forget about boundary, dp needs to run when n >= 3
// 2. in case n is illegal, add criteria
// 3. the problem comes from assumption that array length is larger than 3 (array[2]), but this is not valid if n = 1
