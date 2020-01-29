package main

//On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).
//
//Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.
//
//Example 1:
//
//Input: cost = [10, 15, 20]
//Output: 15
//Explanation: Cheapest is start on cost[1], pay that cost and go to the top.
//Example 2:
//
//Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
//Output: 6
//Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
//Note:
//
//cost will have a length in the range [2, 1000].
//Every cost[i] will be an integer in the range [0, 999].

func minCostClimbingStairs(cost []int) int {
	length := len(cost)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return cost[0]
	}

	if length == 2 {
		return min(cost[0], cost[1])
	}

	var prev1, prev2, i int
	for i, prev1, prev2 = 2, cost[1], cost[0]; i < length; i++ {
		tmp := cost[i] + min(prev2, prev1)
		prev2 = prev1
		prev1 = tmp
	}
	return min(prev1, prev2)
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

// problems
// 1. when start from second stair, loop still need to go to end of cost array
// 2. start from second stair has no influence of start from first stair
// 3. no need to save all data, just the previous 1 & previous 2
