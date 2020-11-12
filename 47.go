package main

import "sort"

// Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.
//
//
//
// Example 1:
//
// Input: nums = [1,1,2]
// Output:
// [[1,1,2],
//  [1,2,1],
//  [2,1,1]]
//
// Example 2:
//
// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
// Constraints:
//
//     1 <= nums.length <= 8
//     -10 <= nums[i] <= 10

// tc: O(n*n!)
func permuteUnique(nums []int) [][]int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	ans := make([][]int, 0)

	permute(nums, []int{}, counter, &ans)

	return ans
}

func permute(nums, current []int, counter map[int]int, ans *[][]int) {
	if len(nums) == len(current) {
		*ans = append(*ans, current)
		return
	}

	for num, count := range counter {
		if count == 0 {
			continue
		}

		tmp := append([]int{}, current...)
		tmp = append(tmp, num)
		counter[num]--
		permute(nums, tmp, counter, ans)

		counter[num]++
	}
}

// tc: O(n * n!)
func permuteUnique1(nums []int) [][]int {
	// sort number to make same number adjacent
	sort.Ints(nums)
	used := make([]bool, len(nums))
	ans := make([][]int, 0)

	permute(nums, []int{}, used, &ans)

	return ans
}

func permute(nums, current []int, used []bool, ans *[][]int) {
	if len(nums) == len(current) {
		*ans = append(*ans, current)
		return
	}

	var j int
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		used[i] = true

		tmp := append([]int{}, current...)
		tmp = append(tmp, nums[i])
		permute(nums, tmp, used, ans)

		used[i] = false

		for j = i + 1; j < len(nums); j++ {
			if nums[j] != nums[i] {
				break
			}
		}

		i = j - 1
	}
}

//	Notes
//	1.	inspired from solution, using hash is also workable
