package main

// There are n uniquely-sized sticks whose lengths are integers from 1 to n. You want to arrange the sticks such that exactly k sticks are visible from the left. A stick is visible from the left if there are no longer sticks to the left of it.
//
// For example, if the sticks are arranged [1,3,2,5,4], then the sticks with lengths 1, 3, and 5 are visible from the left.
//
// Given n and k, return the number of such arrangements. Since the answer may be large, return it modulo 109 + 7.
//
//
//
// Example 1:
//
// Input: n = 3, k = 2
// Output: 3
// Explanation: [1,3,2], [2,3,1], and [2,1,3] are the only arrangements such that exactly 2 sticks are visible.
// The visible sticks are underlined.
//
// Example 2:
//
// Input: n = 5, k = 5
// Output: 1
// Explanation: [1,2,3,4,5] is the only arrangement such that all 5 sticks are visible.
// The visible sticks are underlined.
//
// Example 3:
//
// Input: n = 20, k = 11
// Output: 647427950
// Explanation: There are 647427950 (mod 109 + 7) ways to rearrange the sticks such that exactly 11 sticks are visible.
//
//
//
// Constraints:
//
// 1 <= n <= 1000
// 1 <= k <= n

// tc: O(nk)
var mod = int(1e9 + 7)

func rearrangeSticks(n int, k int) int {
	// dp[n][k]: # of ways to form n, k
	// transition: if longest at last: n-1, k-1; if longest not at last: n-1, k
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[1][1] = 1

	dfs(n, k, dp)

	return dp[n][k]
}

func dfs(n, k int, dp [][]int) int {
	if n == 0 || k == 0 || k > n {
		return 0
	}

	if dp[n][k] == -1 {
		// longest at last, not longest at last
		dp[n][k] = (dfs(n-1, k-1, dp) + (n-1)*dfs(n-1, k, dp)) % mod
	}

	return dp[n][k]
}

//	Notes
//	1.	i know this is dp problem, but cannot finished during contest

//	2.	inspired from https://leetcode.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/discuss/1211070/C%2B%2B-Detailed-Explanation-with-Thought-Process-or-DP and https://www.youtube.com/watch?v=O761YBjGxGA

//		the smarter way to solve this is to consider from last position, there will be
//		2 conditions:
//		- longest at last, since it's longest, it's always visible, remain becomes dfs(n-1, k-1)
//		- longest not at last, since longest is ahead, it's not visible, remain
//		  becomes (n-1) * dfs(n-1, k)

//		another small optimization when n < k, no need to proceed because there's not enough
//		number to pick to make it valid
