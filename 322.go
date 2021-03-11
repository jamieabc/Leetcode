package main

import (
	"math"
	"sort"
)

// You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
//
// You may assume that you have an infinite number of each kind of coin.
//
//
//
// Example 1:
//
// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
//
// Example 2:
//
// Input: coins = [2], amount = 3
// Output: -1
//
// Example 3:
//
// Input: coins = [1], amount = 0
// Output: 0
//
// Example 4:
//
// Input: coins = [1], amount = 1
// Output: 1
//
// Example 5:
//
// Input: coins = [1], amount = 2
// Output: 2
//
//
//
// Constraints:
//
//     1 <= coins.length <= 12
//     1 <= coins[i] <= 231 - 1
//     0 <= amount <= 104

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	sort.Ints(coins)
	for _, coin := range coins {
		if coin <= amount {
			dp[coin] = 1
		}
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c >= 0 {
				dp[i] = min(dp[i], dp[i-c]+1)
			} else {
				break
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func coinChange1(coins []int, amount int) int {
	// dp default to 0, and final checking condition is also 0, need to check additionally
	if amount == 0 {
		return 0
	}

	// dp[i]: min # of coins to reach that value
	dp := make([]int, amount+1)

	for i := range coins {
		if coins[i] <= amount {
			dp[coins[i]] = 1
		}
	}

	for i := 1; i <= amount; i++ {
		if dp[i] == 0 {
			continue
		}

		for j := range coins {
			target := coins[j] + i

			if target <= amount {
				if dp[target] == 0 {
					dp[target] = dp[i] + 1
				} else {
					dp[target] = min(dp[target], dp[i]+1)
				}
			}
		}
	}

	if dp[amount] == 0 {
		return -1
	}

	return dp[amount]
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	fill dp default math.Maxint32 is just as a mark ot make min work, can
//		also be done by check if dp[i] == 0

//	2.	becareful about boundary conditions, anything that might exceed
//		array
