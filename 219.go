package main

//Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.
//
//Example 1:
//
//Input: nums = [1,2,3,1], k = 3
//Output: true
//
//Example 2:
//
//Input: nums = [1,0,1,1], k = 1
//Output: true
//
//Example 3:
//
//Input: nums = [1,2,3,1,2,3], k = 2
//Output: false

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	mapping := make(map[int]bool)
	length := len(nums)
	for i, n := range nums {
		if _, ok := mapping[n]; !ok {
			mapping[n] = true
		} else {
			start, end := findIndexRange(i, length, k)
			for j := start; j <= end; j++ {
				if n == nums[j] && j != i {
					return true
				}
			}
		}
	}
	return false
}

func findIndexRange(i, length, k int) (int, int) {
	start := i - k
	end := i + k
	if start < 0 {
		start = 0
	}

	if end >= length {
		end = length - 1
	}

	return start, end
}
