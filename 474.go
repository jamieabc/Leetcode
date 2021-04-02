package main

import "sort"

// You are given an array of binary strings strs and two integers m and n.
//
// Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.
//
// A set x is a subset of a set y if all elements of x are also elements of y.
//
//
//
// Example 1:
//
// Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
// Output: 4
// Explanation: The largest subset with at most 5 0's and 3 1's is {"10", "0001", "1", "0"}, so the answer is 4.
// Other valid but smaller subsets include {"0001", "1"} and {"10", "1", "0"}.
// {"111001"} is an invalid subset because it contains 4 1's, greater than the maximum of 3.
//
// Example 2:
//
// Input: strs = ["10","0","1"], m = 1, n = 1
// Output: 2
// Explanation: The largest subset is {"0", "1"}, so the answer is 2.
//
//
//
// Constraints:
//
//     1 <= strs.length <= 600
//     1 <= strs[i].length <= 100
//     strs[i] consists only of digits '0' and '1'.
//     1 <= m, n <= 100

// tc: O(mnk), m: zero, n: one, k: array size
func findMaxForm(strs []string, m int, n int) int {
	size := len(strs)

	// memo[i][m][n]: start from index i, longest subsest for m zero &
	// n one
	memo := make([][][]int, size)
	for i := range memo {
		memo[i] = make([][]int, m+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, n+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}

	ones, zeros := make([]int, size), make([]int, size)
	for i := range strs {
		for j := range strs[i] {
			if strs[i][j] == '1' {
				ones[i]++
			}
		}
		zeros[i] = len(strs[i]) - ones[i]
	}

	dfs(ones, zeros, memo, 0, m, n)

	return memo[0][m][n]
}

func dfs(ones, zeros []int, memo [][][]int, idx, m, n int) int {
	if idx == len(ones) || m < 0 || n < 0 {
		return 0
	}

	if memo[idx][m][n] != -1 {
		return memo[idx][m][n]
	}

	var tmp int

	if m-zeros[idx] >= 0 && n-ones[idx] >= 0 {
		tmp = max(tmp, 1+dfs(ones, zeros, memo, idx+1, m-zeros[idx], n-ones[idx]))
	}

	tmp = max(tmp, dfs(ones, zeros, memo, idx+1, m, n))

	memo[idx][m][n] = tmp

	return tmp
}

func findMaxForm1(strs []string, m int, n int) int {
	size := len(strs)
	ones, zeros := make([]int, size), make([]int, size)

	for i := range strs {
		for j := range strs[i] {
			if strs[i][j] == '1' {
				ones[i]++
			}
		}
		zeros[i] = len(strs[i]) - ones[i]
	}

	sorted := make([]int, size)
	for i := range sorted {
		sorted[i] = i
	}

	sort.Slice(sorted, func(i, j int) bool {
		if ones[sorted[i]]+zeros[sorted[i]] != ones[sorted[j]]+zeros[sorted[j]] {
			return ones[sorted[i]]+zeros[sorted[i]] < ones[sorted[j]]+zeros[sorted[j]]
		}
		return ones[sorted[i]] < ones[sorted[j]]
	})

	used := make([]bool, size)
	var largest int

	dfs1(ones, zeros, sorted, used, m, n, 0, &largest)

	return largest
}

// tc: O(2^n), n: array length
func dfs1(ones, zeros, sorted []int, used []bool, zero, one, cur int, largest *int) {
	*largest = max(*largest, cur)

	if zero == 0 && one == 0 {
		return
	}

	for i := len(ones) - 1; i >= 0; i-- {
		if used[sorted[i]] {
			continue
		}

		if one-ones[sorted[i]] >= 0 && zero-zeros[sorted[i]] >= 0 {
			used[sorted[i]] = true
			dfs1(ones, zeros, sorted, used, zero-zeros[sorted[i]], one-ones[sorted[i]], cur+1, largest)
			used[sorted[i]] = false
		}
	}
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	backtracking, tc: O(2^n), TLE

//	2.	if it's 1's dimension, it's similar to 416, which is dp problem

//	3.	inspired from sample code, use strings.Count(str, "0") to find out count of 0

//	4.	inspired from https://leetcode.com/problems/ones-and-zeroes/discuss/95807/0-1-knapsack-detailed-explanation.

//		explains details, any also mention this problem is similar to knapsack problem
//		there's dp optimization, not understand
