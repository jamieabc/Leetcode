package main

import "math"

// Given two strings s1, s2, find the lowest ASCII sum of deleted characters to make two strings equal.
//
// Example 1:
//
// Input: s1 = "sea", s2 = "eat"
// Output: 231
// Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to the sum.
// Deleting "t" from "eat" adds 116 to the sum.
// At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum possible to achieve this.
//
// Example 2:
//
// Input: s1 = "delete", s2 = "leet"
// Output: 403
// Explanation: Deleting "dee" from "delete" to turn the string into "let",
// adds 100[d]+101[e]+101[e] to the sum.  Deleting "e" from "leet" adds 101[e] to the sum.
// At the end, both strings are equal to "let", and the answer is 100+101+101+101 = 403.
// If instead we turned both strings into "lee" or "eet", we would get answers of 433 or 417, which are higher.
//
// Note:
// 0 < s1.length, s2.length <= 1000.
// All elements of each string will have an ASCII value in [97, 122].

func minimumDeleteSum(s1 string, s2 string) int {
	n1, n2 := len(s1), len(s2)

	// dp[i][j]: min cost to make s1[:i] == s2[:j], not including i & j
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	// first line, s1 is empty, remove all s2, iterate on s2
	for j := 1; j <= n2; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}

	// iterate on s1
	for i := 1; i <= n1; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])

		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}

	return dp[n1][n2]
}

func minimumDeleteSum1(s1 string, s2 string) int {
	n1, n2 := len(s1), len(s2)
	longest := max(n1, n2)

	memo := make([][]int, longest)
	for i := range memo {
		memo[i] = make([]int, longest)
		for j := range memo[0] {
			memo[i][j] = math.MaxInt32
		}
	}

	return dfs(s1, s2, 0, 0, memo)
}

func dfs(s1, s2 string, idx1, idx2 int, memo [][]int) int {
	if idx1 == len(s1) {
		var cost int
		for i := idx2; i < len(s2); i++ {
			cost += int(s2[i])
		}

		return cost
	}

	if idx2 == len(s2) {
		var cost int
		for i := idx1; i < len(s1); i++ {
			cost += int(s1[i])
		}

		return cost
	}

	if memo[idx1][idx2] == math.MaxInt32 {
		if s1[idx1] == s2[idx2] {
			memo[idx1][idx2] = dfs(s1, s2, idx1+1, idx2+1, memo)
		} else {
			cost := int(s1[idx1]) + dfs(s1, s2, idx1+1, idx2, memo)
			cost = min(cost, int(s2[idx2])+dfs(s1, s2, idx1, idx2+1, memo))

			memo[idx1][idx2] = cost
		}
	}

	return memo[idx1][idx2]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/discuss/108828/C%2B%2B-DP-with-explanation

//		dp[i][j]: minimum cost to make s1[:i] == s2[:j]
//		if s1[i] == s2[j], dp[i][j] = dp[i-1][j-1]
//		if s1[i] != s2[j], dp[i][j] = min(dp[i-1][j] + s1[i], dp[i][j-1] + s2[j])

//	2.	inspired from https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/discuss/108811/JavaDP(With-Explanation)

//		author has same feeling, similar to longest common subsequence

//	3.	add reference https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/discuss/136346/C%2B%2B-DP
