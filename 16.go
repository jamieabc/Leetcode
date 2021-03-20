package main

import (
	"math"
	"sort"
)

// Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
//
//
//
// Example 1:
//
// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
//
//
//
// Constraints:
//
// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4

// tc: O(n^2)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closest := math.MaxInt32
	var ans int
	size := len(nums)

	for i := range nums {
		for j, k := i+1, size-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]

			if cur := abs(sum - target); cur < closest {
				ans = sum
				closest = cur
			}

			if sum == target {
				return sum
			} else if sum > target {
				k--
			} else {
				j++
			}
		}
	}

	return ans
}

// tc: O(n^2 log(n))
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	size := len(nums)

	var ans int
	closest := math.MaxInt32

	for i := range nums {
		for j := i + 1; j < size-1; j++ {
			goal := target - nums[i] - nums[j]

			for low, high := j+1, size-1; low <= high; {
				mid := low + (high-low)>>1

				if nums[mid] == goal {
					return target
				}

				if cur := abs(nums[mid] - goal); cur < closest {
					ans = nums[i] + nums[j] + nums[mid]
					closest = cur
				}

				if nums[mid] > goal {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
		}
	}

	return ans
}

// sort array, find closest
// tc: O(n^2)
func threeSumClosest1(nums []int, target int) int {
	closest := math.MaxInt32
	sort.Ints(nums)

	var low, high int
	for i := 0; i < len(nums); i++ {
		for low, high = i+1, len(nums)-1; low < high; {
			sum := nums[i] + nums[low] + nums[high]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}

			if sum == target {
				return target
			} else if sum < target {
				low++
			} else {
				high--
			}
		}

		// avoid duplicate nums[i]
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return closest
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	Notes
//	1.	inspired from solution, can use binary search to find answer, although
//		it's more time consuming, but it's a good practice

//		fix two numbers, and find the last number, if middle number is larger than
//		goal then go smaller, otherwise, go larger, this always do half of check
