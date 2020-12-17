package main

import (
	"math"
	"sort"
)

//Given a wooden stick of length n units. The stick is labelled from 0 to n. For example, a stick of length 6 is labelled as follows:
//
//Given an integer array cuts where cuts[i] denotes a position you should perform a cut at.
//
//You should perform the cuts in order, you can change the order of the cuts as you wish.
//
//The cost of one cut is the length of the stick to be cut, the total cost is the sum of costs of all cuts. When you cut a stick, it will be split into two smaller sticks (i.e. the sum of their lengths is the length of the stick before the cut). Please refer to the first example for a better explanation.
//
//Return the minimum total cost of the cuts.
//
//
//
//Example 1:
//
//Input: n = 7, cuts = [1,3,4,5]
//Output: 16
//Explanation: Using cuts order = [1, 3, 4, 5] as in the input leads to the following scenario:
//
//The first cut is done to a rod of length 7 so the cost is 7. The second cut is done to a rod of length 6 (i.e. the second part of the first cut), the third is done to a rod of length 4 and the last cut is to a rod of length 3. The total cost is 7 + 6 + 4 + 3 = 20.
//Rearranging the cuts to be [3, 5, 1, 4] for example will lead to a scenario with total cost = 16 (as shown in the example photo 7 + 4 + 3 + 2 = 16).
//
//Example 2:
//
//Input: n = 9, cuts = [5,6,1,4,2]
//Output: 22
//Explanation: If you try the given cuts ordering the cost will be 25.
//There are much ordering with total cost <= 25, for example, the order [4, 6, 5, 2, 1] has total cost = 22 which is the minimum possible.
//
//
//
//Constraints:
//
//    2 <= n <= 10^6
//    1 <= cuts.length <= min(n - 1, 100)
//    1 <= cuts[i] <= n - 1
//    All the integers in cuts array are distinct.

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	size := len(cuts)

	// dp[i][j]: min cost from i ~ j
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	for d := 0; d < size; d++ {
		for L := 0; L+d < size; L++ {
			R := L + d

			for i := L; i <= R; i++ {
				var left, right int
				tmp := cuts[i]

				if R+1 < size {
					tmp = cuts[R+1]
				} else {
					tmp = n
				}

				if L > 0 {
					tmp -= cuts[L-1]
				}

				if i > L {
					left = dp[L][i-1]
				}

				if i+1 <= R {
					right = dp[i+1][R]
				}

				dp[L][R] = min(dp[L][R], tmp+left+right)
			}
		}
	}
	return dp[0][size-1]
}

func minCost1(n int, cuts []int) int {
	sort.Ints(cuts)

	// try all possibilities
	minCost := math.MaxInt32
	selected := make([]bool, len(cuts))

	dfs(cuts, selected, 0, 0, n, &minCost)

	return minCost
}

func dfs(cuts []int, selected []bool, currentCost, cutCount, n int, minCost *int) {
	if cutCount == len(cuts) {
		*minCost = min(*minCost, currentCost)
		return
	}

	for i := 0; i < len(cuts); i++ {
		// skip already selected
		if selected[i] {
			continue
		}

		// cut here
		selected[i] = true
		dfs(cuts, selected, currentCost+findCost(cuts, selected, i, n), cutCount+1, n, minCost)
		selected[i] = false
	}
}

func findCost(cuts []int, selected []bool, idx, n int) int {
	left, right := 0, n

	for i := idx - 1; i >= 0; i-- {
		if selected[i] {
			left = cuts[i]
			break
		}
	}

	for i := idx + 1; i < len(selected); i++ {
		if selected[i] {
			right = cuts[i]
			break
		}
	}

	return right - left
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	something wrong in binary search

//		the goal is to find some index that cuts[index] >= target,
//		so binary direction & for loop condition should be carefully
//		determined

//		mid = low (high-low)/2 chooses smaller index, to make index
//		valid, low can never be larger than high (think of boundary high
//		that causes array out of bound)

//	2.	wrong answer on 30 [18,15,13,7,5,26,25,29]
//		expected: 92, my: 94

//		I think using greedy (always choose cut closest to mid) might
//		not be appropriate, take test case for example:

//		n = 30, cuts = [18,15,13,7,5,26,25,29]

//		by greedy algorithm, it chooses cut 15 cause it's the middle
//		cut for 30, by that way cost is 94

//		but minimum cost comes from choose cut 18 first, so it's either
//		it's not okay to find minimum by greedy, or I haven't fully
//		think through the problem

//		I will try to use brute force + memoization to solve this one
//		, then next problem occurs: how to do memoization? The nature
//		of this problem is that exact order doesn't matter,
//		relationships matters.

//		for example, [18 7 5 13 15 25 26 29] & [18 7 25 13 5 26 29 15]
//		with same cost

//	3.	inspired from https://leetcode.com/problems/minimum-cost-to-cut-a-stick/discuss/780880/DP-with-picture-(Burst-Balloons)

//		I was stuck at a place that how to design memoization, because
//		I was thinking it's a set, how to represent a set as key to
//		map? I come up with a very naive solution: concat each number
//		with -. And even with same group of cuts, might have different
//		cost under different parent cut, so need another string to
//		denote range of stick start-end. Overall data structure
//		as map[string]map[string]int

//		author sort first, so each number will have a fix position, then
//		this fix index can be used as map key.

//		the other point is to add 0 & n into cuts, make it as sentinel

//	4.	after solving burst balloon (https://leetcode.com/problems/burst-balloons/)

//		this problem is similar to that one
