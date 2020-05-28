package main

import "sort"

// Given an array of integers nums, sort the array in ascending order.
//
//
//
// Example 1:
//
// Input: nums = [5,2,3,1]
// Output: [1,2,3,5]
// Example 2:
//
// Input: nums = [5,1,1,2,0,0]
// Output: [0,0,1,1,2,5]
//
//
// Constraints:
//
// 1 <= nums.length <= 50000
// -50000 <= nums[i] <= 50000

func sortArray(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}

	if length == 2 {
		if nums[0] <= nums[1] {
			return []int{nums[0], nums[1]}
		} else {
			return []int{nums[1], nums[0]}
		}
	}

	// split
	mid := (length - 1) / 2
	a := mergeSort(nums[:mid])
	b := mergeSort(nums[mid:])

	result := make([]int, 0)

	// merge
	for i, j := 0, 0; i < len(a) || j < len(b); {
		if i != len(a) && (j == len(b) || a[i] <= b[j]) {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	return result
}

func sortArray1(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}
