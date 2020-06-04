package main

// Say you have an array for which the ith element is the price of a given stock on day i.
//
// If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
//
// Note that you cannot sell a stock before you buy one.
//
// Example 1:
//
// Input: [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//              Not 7-1 = 6, as selling price needs to be larger than buying price.
// Example 2:
//
// Input: [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e. max profit = 0.

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}

	minPrice, maxProfit := prices[0], 0

	for i := 1; i < length; i++ {
		maxProfit = max(maxProfit, prices[i]-minPrice)
		minPrice = min(prices[i], minPrice)
	}

	return maxProfit
}

func maxProfit1(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}

	dp := make([]int, length)
	dp[length-2] = prices[length-1]

	for i := length - 3; i >= 0; i-- {
		dp[i] = max(prices[i+1], dp[i+1])
	}

	var m int
	for i := 0; i < length-1; i++ {
		m = max(m, dp[i]-prices[i])
	}

	return m
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	is max profit < 0, then it should be 0, because it's not necessary to
//		have transaction

//	2.	too slow, the problem can be describe as follows: if a certain buying
//		day is picked, that day's maximum profit is determined. so the dp[i]
//		can be chosen as max number from i+1 to end day. in this way, dp can
//		be reused: dp[i] = max(prices[i], dp[i+1])

//	3.	add reference https://leetcode.com/problems/best-time-to-buy-and-sell-stock/discuss/39039/Sharing-my-simple-and-clear-C%2B%2B-solution

//		author uses bottom-up way to calculate maximum profit, the point
//		is to store minimum price before ith day
