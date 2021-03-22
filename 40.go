package main

import "sort"

// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.
//
// Each number in candidates may only be used once in the combination.
//
// Note: The solution set must not contain duplicate combinations.
//
//
//
// Example 1:
//
// Input: candidates = [10,1,2,7,6,1,5], target = 8
// Output:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
//
// Example 2:
//
// Input: candidates = [2,5,2,1,2], target = 5
// Output:
// [
// [1,2,2],
// [5]
// ]
//
//
//
// Constraints:
//
//     1 <= candidates.length <= 100
//     1 <= candidates[i] <= 50
//     1 <= target <= 30

// tc: O(2^n), each number is chosen or not chosen
func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)

	sort.Ints(candidates)
	dfs(candidates, []int{}, target, 0, &ans)

	return ans
}

func dfs(candidates, cur []int, target, idx int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, cur)
		return
	}

	size := len(candidates)
	if idx == size || target < 0 {
		return
	}

	var count int
	for i := idx; i < size && candidates[i] == candidates[idx]; i, count = i+1, count+1 {
	}

	// not include self, skip duplicates
	dfs(candidates, cur, target, idx+count, ans)

	// include self and duplicates
	tmp := append([]int{}, cur...)
	for i := 0; i < count; i++ {
		tmp = append(tmp, candidates[idx])
		target -= candidates[idx]
		dfs(candidates, tmp, target, idx+count, ans)
	}
}

//	Notes
//	1.	to have unique sequence, for every duplicates numbers, control their
//		occurrence to make sure duplicate sequence won't appear

//		e.g. [1, 1, 1, 2, 3]
//		      ^ ------->
//			  ^  ^ ---->
//			  ^  ^  ^ ->

//		[1], [1, 1], [1, 1, 1] and go on
