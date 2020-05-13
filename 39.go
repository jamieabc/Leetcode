package main

import "sort"

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
// The same repeated number may be chosen from candidates unlimited number of times.
//
// Note:
//
//     All numbers (including target) will be positive integers.
//     The solution set must not contain duplicate combinations.
//
// Example 1:
//
// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]
//
// Example 2:
//
// Input: candidates = [2,3,5], target = 8,
// A solution set is:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)

	combinations(candidates, []int{}, target, 0, &result)
	return result
}

func combinations(candidates, current []int, target, idx int, result *[][]int) {
	if target == 0 && len(current) != 0 {
		*result = append(*result, current)
		return
	}

	for i := idx; i < len(candidates); i++ {
		if candidates[i] <= target {
			tmp := append([]int{}, current...)
			tmp = append(tmp, candidates[i])
			combinations(candidates, tmp, target-candidates[i], i, result)
		}
	}
}
