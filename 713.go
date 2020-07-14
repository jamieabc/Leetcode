package main

// Your are given an array of positive integers nums.
//
// Count and print the number of (contiguous) subarrays where the product of all the elements in the subarray is less than k.
//
// Example 1:
//
// Input: nums = [10, 5, 2, 6], k = 100
// Output: 8
// Explanation: The 8 subarrays that have product less than 100 are: [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].
// Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.
//
// Note:
// 0 < nums.length <= 50000.
// 0 < nums[i] < 1000.
// 0 <= k < 10^6.

func numSubarrayProductLessThanK(nums []int, k int) int {
	size := len(nums)
	if size == 0 || k == 0 {
		return 0
	}

	var count int
	product := 1

	for low, high := 0, 0; high < size; high++ {
		product *= nums[high]

		for low <= high && product >= k {
			// shrink
			product /= nums[low]
			low++
		}
		count += high - low + 1
	}

	return count
}

func numSubarrayProductLessThanK1(nums []int, k int) int {
	size := len(nums)
	if size == 0 || k == 0 {
		return 0
	}

	var count int
	product := nums[0]

	for low, high := 0, 0; low < size; {
		if high < size-1 && (low > high || product*nums[high+1] < k) {
			// expand
			high++
			product *= nums[high]
		} else {
			// shrink
			if product < k {
				count += high - low + 1
			}
			product /= nums[low]
			low++
		}
	}

	return count
}

//	problems
//	1.	be careful condition to shrink and expand

//		even single number should be counted, when low == high, that number
//		should be counted if nums[low] < k, which influence the expand condition:
//		low > high should expand

//	2.	inspired from https://leetcode.com/problems/subarray-product-less-than-k/discuss/108861/JavaC%2B%2B-Clean-Code-with-Explanation

//		the idea is to maintain product always < k, and count sub-array size
