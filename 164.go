package main

import "math"

// Given an integer array nums, return the maximum difference between two successive elements in its sorted form. If the array contains less than two elements, return 0.
//
// You must write an algorithm that runs in linear time and uses linear extra space.
//
//
//
// Example 1:
//
// Input: nums = [3,6,9,1]
// Output: 3
// Explanation: The sorted form of the array is [1,3,6,9], either (3,6) or (6,9) has the maximum difference 3.
//
// Example 2:
//
// Input: nums = [10]
// Output: 0
// Explanation: The array contains less than 2 elements, therefore return 0.
//
//
//
// Constraints:
//
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 109

func maximumGap(nums []int) int {
	size := len(nums)

	if size < 2 {
		return 0
	}

	var largest int
	smallest := math.MaxInt32
	distinct := 1

	for i, n := range nums {
		largest = max(largest, n)
		smallest = min(smallest, n)
		if i > 0 && nums[i] != nums[i-1] {
			distinct++
		}
	}

	if smallest == largest {
		return 0
	}

	gap := (largest - smallest + 1) / (distinct - 1)

	buckets := make([][2]int, size)
	for i := range buckets {
		buckets[i] = [2]int{math.MaxInt32, math.MinInt32} // smallest, largest
	}

	for _, n := range nums {
		idx := (n - smallest) / gap
		buckets[idx][0] = min(buckets[idx][0], n)
		buckets[idx][1] = max(buckets[idx][1], n)
	}

	var maxDiff, prev int

	for i := range buckets {
		if i > 0 {
			if buckets[i][0] != math.MaxInt32 {
				maxDiff = max(maxDiff, max(buckets[i][0]-buckets[prev][1], buckets[i][1]-buckets[i][0]))
				prev = i
			}
		} else {
			maxDiff = max(maxDiff, buckets[i][1]-buckets[i][0])
			prev = 0
		}
	}

	return maxDiff
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	i thinks this problems relates to pigeon hole, maximum difference either
//		occurs at each bucket, or occurs at boundary
