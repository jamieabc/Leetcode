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
	quickSort(nums, 0, len(nums)-1)

	return nums
}

func quickSort(nums []int, start, end int) {
	idx := partition(nums, start, end)

	if idx != -1 {
		quickSort(nums, start, idx-1)
		quickSort(nums, idx+1, end)
	}
}

func partition(nums []int, start, end int) int {
	if start >= end {
		return -1
	}

	store, pivot := start, nums[start]
	nums[start], nums[end] = nums[end], nums[start]

	for i := start; i < end; i++ {
		if nums[i] < pivot {
			nums[store], nums[i] = nums[i], nums[store]
			store++
		}
	}

	nums[store], nums[end] = nums[end], nums[store]

	return store
}

func sortArray1(nums []int) []int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) []int {
	if start == end {
		return nums[start : start+1]
	}

	mid := start + (end-start)>>1
	arr1, arr2 := mergeSort(nums, start, mid), mergeSort(nums, mid+1, end)

	ans := make([]int, len(arr1)+len(arr2))

	for idx, p1, p2 := 0, 0, 0; idx < len(ans); idx++ {
		if p2 == len(arr2) || (p1 < len(arr1) && arr1[p1] < arr2[p2]) {
			ans[idx] = arr1[p1]
			p1++
		} else {
			ans[idx] = arr2[p2]
			p2++
		}
	}

	return ans
}

//	Notes
//	1.	merge sort: divide & conquer
//		quick sort: quick select

//	2.	inspired from https://leetcode.com/problems/sort-an-array/discuss/461394/Python-3-(Eight-Sorting-Algorithms)-(With-Explanation)

//		there are also:
//		selection sort (find smallest in remaining array, take it out and put into new array)
//		bubble sort (swap adj number if next number is larger),
//		insertion sort (for every number, checks previous numbers to find it proper location)
//		heap sort (use heap)
//		bucket sort
