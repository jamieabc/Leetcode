package main

// The gray code is a binary numeral system where two successive values differ in only one bit.
//
// Given an integer n representing the total number of bits in the code, return any sequence of gray code.
//
// A gray code sequence must begin with 0.
//
//
//
// Example 1:
//
// Input: n = 2
// Output: [0,1,3,2]
// Explanation:
// 00 - 0
// 01 - 1
// 11 - 3
// 10 - 2
// [0,2,3,1] is also a valid gray code sequence.
// 00 - 0
// 10 - 2
// 11 - 3
// 01 - 1
//
// Example 2:
//
// Input: n = 1
// Output: [0,1]
//
//
//
// Constraints:
//
// 1 <= n <= 16

func grayCode(n int) []int {
	ans := make([]int, 0)
	ans = append(ans, 0, 1)

	for i := 2; i <= n; i++ {
		size := len(ans)
		for j := size - 1; j >= 0; j-- {
			ans = append(ans, ans[j]|(1<<(i-1)))
		}
	}

	return ans
}

func grayCode1(n int) []int {
	visited := make(map[int]bool)
	ans := []int{0}
	visited[0] = true

	for len(ans) < (1 << n) {
		num := ans[len(ans)-1]
		for j := 0; j < n; j++ {
			tmp := num ^ (1 << j)
			if !visited[tmp] {
				ans = append(ans, tmp)
				visited[tmp] = true
				break
			}
		}
	}

	return ans
}

//	Notes
//	1.	using visited to store if a number is visited, it's kind of slow because
//		additional checks are needed

//	2.	inspired from https://leetcode.com/problems/gray-code/discuss/29891/Share-my-solution

//		backward append 1 to next bit
//		0 -> 1
//		11 -> 10 (from previous 1, 0)
//		111 -> 110 -> 101 -> 100 (from previous 01, 11, 1, 0)

//		very clever solution, like combination, find a way that matches rule
