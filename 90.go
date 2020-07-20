package main

import "sort"

// Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).
//
// Note: The solution set must not contain duplicate subsets.
//
// Example:
//
// Input: [1,2,2]
// Output:
// [
//   [2],
//   [1],
//   [1,2,2],
//   [2,2],
//   [1,2],
//   []
// ]

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)
	recursive(nums, 0, []int{}, &result)

	return result
}

func recursive(nums []int, start int, data []int, result *[][]int) {
	*result = append(*result, data)

	for i := start; i < len(nums); i++ {
		tmp := append([]int{}, data...)
		tmp = append(tmp, nums[i])
		recursive(nums, i+1, tmp, result)

		var j int
		for j = i + 1; j < len(nums); j++ {
			if nums[j] != nums[i] {
				break
			}
		}
		i = j - 1
	}
}
