package main

// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
//
// Design an algorithm to find the maximum profit. You may complete at most k transactions.
//
// Notice that you may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//
//
//
// Example 1:
//
// Input: k = 2, prices = [2,4,1]
// Output: 2
// Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
// Example 2:
//
// Input: k = 2, prices = [3,2,6,5,0,3]
// Output: 7
// Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4. Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
//
//
// Constraints:
//
// 0 <= k <= 109
// 0 <= prices.length <= 104
// 0 <= prices[i] <= 1000

func maxProfit(k int, prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}

	// dp[m][n]: maximum profit on mth transaction, sell on nth day
	dp := make([][]int, k+1)
	var i, j int
	for i = range dp {
		dp[i] = make([]int, size)
	}

	var maxProfit int
	for i = 1; i < k+1; i++ {
		for j = 1; j < size; j++ {
			dp[i][j] = max(dp[i][j], maxProfitToDay(dp[i-1], prices, j))
			maxProfit = max(maxProfit, dp[i][j])
		}
	}

	return maxProfit
}

func maxProfitToDay(dp, prices []int, day int) int {
	var maxProfit int

	for i := 0; i < day; i++ {
		// previous transaction ends at day i
		minPrice := prices[i]

		for j := i + 1; j <= day; j++ {
			maxProfit = max(maxProfit, max(0, dp[i]+prices[day]-minPrice))
			minPrice = min(minPrice, prices[j])
		}
	}

	return maxProfit
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
