package main

import "math"

// Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.
//
// Example:
//
// Input: s = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: the subarray [4,3] has the minimal length under the problem constraint.
//
// Follow up:
// If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).

func minSubArrayLen(s int, nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	// accumulative sum to ith index
	sums := make([]int, size)
	sums[0] = nums[0]
	for i := 1; i < size; i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	minL := math.MaxInt32
	var j, k int

	for i := range sums {
		if sums[i] > s {
			diff := sums[i] - s

			for j, k = 0, i; j < k; {
				mid := j + (k-j+1)/2

				if sums[mid] == diff {
					j = mid
					break
				}

				if sums[mid] > diff {
					k = mid - 1
				} else if sums[mid] < diff {
					j = mid
				}
			}

			if i == j {
				minL = 1
			} else {
				minL = min(minL, i-j)
			}
		} else if sums[i] == s {
			minL = min(minL, i+1)
		}
	}

	if minL == math.MaxInt32 {
		return 0
	}
	return minL
}

func minSubArrayLen1(s int, nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	sum := nums[0]
	minL := math.MaxInt32

	for i, j := 0, 1; i < size; {
		if sum < s && j < size {
			sum += nums[j]
			j++
		} else {
			if sum >= s {
				minL = min(minL, j-i)
			}
			sum -= nums[i]
			i++
		}
	}

	if minL == math.MaxInt32 {
		return 0
	}

	return minL
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	problems
//	1.	O(n) uses sliding window

//	2.	O(n log n) uses running sum
