package main

import "math"

// There are a row of n houses, each house can be painted with one of the three colors: red, blue or green. The cost of painting each house with a certain color is different. You have to paint all the houses such that no two adjacent houses have the same color.
//
// The cost of painting each house with a certain color is represented by a n x 3 cost matrix. For example, costs[0][0] is the cost of painting house 0 with color red; costs[1][2] is the cost of painting house 1 with color green, and so on... Find the minimum cost to paint all houses.
//
// Note:
// All costs are positive integers.
//
// Example:
//
// Input: [[17,2,17],[16,16,5],[14,3,19]]
// Output: 10
// Explanation: Paint house 0 into blue, paint house 1 into green, paint house 2 into blue.
//              Minimum cost: 2 + 5 + 3 = 10.

func minCost(costs [][]int) int {
	length := len(costs)
	if length == 0 {
		return 0
	}
	colors := len(costs[0])
	dp := make([]int, colors)
	copy(dp, costs[0])

	for i := 1; i < length; i++ {
		tmp1 := min(dp[0], dp[1])
		tmp2 := min(dp[1], dp[2])
		tmp3 := min(dp[0], dp[2])
		dp[0] = costs[i][0] + tmp2
		dp[1] = costs[i][1] + tmp3
		dp[2] = costs[i][2] + tmp1
	}

	return min(dp[0], min(dp[1], dp[2]))
}

func minCost1(costs [][]int) int {
	// dp[i][j] mean jth house chose ith color, i = 0, 1, 2
	length := len(costs)

	if length == 0 {
		return 0
	}

	colors := len(costs[0])
	dp := make([][]int, colors)
	for i := range dp {
		dp[i] = make([]int, length)
	}

	dp[0][0] = costs[0][0]
	dp[1][0] = costs[0][1]
	dp[2][0] = costs[0][2]

	for j := 1; j < length; j++ {
		for i := 0; i < colors; i++ {
			m := math.MaxInt32
			for k := 0; k < colors; k++ {
				if k == i {
					continue
				}
				m = min(m, dp[k][j-1])
			}
			dp[i][j] = m + costs[j][i]
		}
	}

	return min(min(dp[0][length-1], dp[1][length-1]), dp[2][length-1])
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	problems
//	1.	tc: O(n), every house choose min cost from previous, and n houses to calculate

//	2.	from solution, it reuses original array, so space complexity is O(1)
