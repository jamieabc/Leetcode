package main

import (
	"math"
	"sort"
)

// You are given an array nums of n positive integers.
//
// You can perform two types of operations on any element of the array any number of times:
//
//     If the element is even, divide it by 2.
//         For example, if the array is [1,2,3,4], then you can do this operation on the last element, and the array will be [1,2,3,2].
//     If the element is odd, multiply it by 2.
//         For example, if the array is [1,2,3,4], then you can do this operation on the first element, and the array will be [2,2,3,4].
//
// The deviation of the array is the maximum difference between any two elements in the array.
//
// Return the minimum deviation the array can have after performing some number of operations.
//
//
//
// Example 1:
//
// Input: nums = [1,2,3,4]
// Output: 1
// Explanation: You can transform the array to [1,2,3,2], then to [2,2,3,2], then the deviation will be 3 - 2 = 1.
//
// Example 2:
//
// Input: nums = [4,1,5,20,3]
// Output: 3
// Explanation: You can transform the array after two operations to [4,2,5,5,3], then the deviation will be 5 - 2 = 3.
//
// Example 3:
//
// Input: nums = [2,10,8]
// Output: 3
//
//
//
// Constraints:
//
//     n == nums.length
//     2 <= n <= 105
//     1 <= nums[i] <= 109

func minimumDeviation(nums []int) int {
	size := len(nums)
	largest, smallest := math.MinInt32, math.MaxInt32

	// shrink all numbers
	for i := range nums {
		largest, smallest = max(largest, nums[i]), min(smallest, nums[i])

		for nums[i]&1 == 0 {
			nums[i] = nums[i] >> 1
		}
	}

	sort.Ints(nums)
	deviation := min(largest-smallest, nums[size-1]-nums[0])

	for i := 0; i < size-1; i++ {
		if nums[i] != nums[i+1] {
			// 2 * smallest > largest, new deviation could be
			// 2 * smallest - next number or
			// all numbers are doubled except the number smaller than largest
			if nums[i]<<1 > nums[size-1] {
				return min(deviation, min(nums[size-2]<<1-nums[size-1], nums[i]<<1-nums[i+1]))
			}

			// 2 * smallest <= next number && 2 * smallest < largest,
			// means 2 * smallest becomes new smallest
			if nums[i]<<1 <= nums[i+1] && nums[i]<<1 < nums[size-1] {
				return min(deviation, nums[size-1]-nums[i]<<1)
			}

			// 2 * next number <= largest,
			if nums[i+1]<<1 <= nums[size-1] {
				return min(deviation, nums[size-1]-nums[i+1])
			}

			return min(deviation, (nums[i+1]-nums[i])<<1)
		}
	}

	return deviation
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	odd -> become larger, odd -> become smaller
//		this problem cares only about largest & smallest, to make it simpler, make all
//		numbers to odd

//	2.	sorted by value ascending
//		smallest ... ... ... ... largest

//		to make deviation smaller, at least one number in list should remain same, if
//		all numbers are double, then deviation also double

//		case 1: current_smallest * 2 < next smaller
//
//				deviation shrinks to largest - current_smallest * 2, stop

//		case 2: next number < current_smallest * 2 < largest

//				all numbers are sorted, and it's a non-decreasing sequence
//				new smallest could be current_smallest * 2 (if next number * 2 < largest)
//				or if next number  * 2 > largest, deviation is changed to
//				min(next number * 2 - min(largest, current_smallest * 2,
//				    largest - current_smallest)

//		case 3: current_smallest * 2 > largest

//	3.	it could happen that original array has smallest deviation

//	4.	after operation, largest & smallest may change, needs to sort again

//	5.	find original largest & smallest, even with this, still not be correct

//		original [10, 4, 3] => deviation = 7
//		change   [2, 3, 5]  => deviation = 3

// 		logic is still wrong
