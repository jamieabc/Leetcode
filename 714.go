package main

// You are given an array prices where prices[i] is the price of a given stock on the ith day, and an integer fee representing a transaction fee.
//
// Find the maximum profit you can achieve. You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.
//
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
//
//
//
// Example 1:
//
// Input: prices = [1,3,2,8,4,9], fee = 2
// Output: 8
// Explanation: The maximum profit can be achieved by:
// - Buying at prices[0] = 1
// - Selling at prices[3] = 8
// - Buying at prices[4] = 4
// - Selling at prices[5] = 9
// The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
//
// Example 2:
//
// Input: prices = [1,3,7,5,10,3], fee = 3
// Output: 6
//
//
//
// Constraints:
//
//     1 < prices.length <= 5 * 10^4
//     0 < prices[i], fee < 5 * 10^4

// tc: O(n)
func maxProfit(prices []int, fee int) int {
	var cash int
	hold := -prices[0]

	for i := range prices {
		tmp := cash
		cash = max(cash, hold+prices[i]-fee)
		hold = max(hold, tmp-prices[i])
	}

	return cash
}

// tc: O(n^2)
func maxProfit1(prices []int, fee int) int {
	size := len(prices)

	// dp[i]: maximum profit can achieve start from i
	dp := make([]int, size+1)
	var largest int

	for i := size - 2; i >= 0; i-- {
		dp[i] = dp[i+1]
		for j := i + 1; j < size; j++ {
			if prices[j] > prices[i] {
				dp[i] = max(dp[i], prices[j]-prices[i]-fee+dp[j+1])
				break
			}
		}
		largest = max(largest, dp[i])
	}

	return largest
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	too slow, if i draw the prices as stock prices, then i only want
//		to buy at low & sell at high, so that i might potentially profit
//		this means, find previous higher points

//	2.	inspired from https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/108870/Most-consistent-ways-of-dealing-with-the-series-of-stock-problems

//		author provides a very good explanation of how problems solved
