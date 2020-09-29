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
	if len(nums) == 0 {
		return 0
	}

	var count int
	product := 1

	for start, end := 0, 0; end < len(nums); end++ {
		product *= nums[end]

		for start <= end && product >= k {
			// shrink
			product /= nums[start]
			start++
		}

		// this if could be omit, since start is top to end  + 1, count will add
		// 0 in this condition
		if start <= end {
			count += end - start + 1
		}
	}

	return count
}

func numSubarrayProductLessThanK2(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	var count int
	product := nums[0]

	for start, end := 0, 0; start < len(nums); {
		// for a start position, expand to maximum possible size
		for end < len(nums)-1 && product*nums[end+1] < k {
			end++
			product *= nums[end]
		}

		// it's possible that nothing valid for this start position
		if product < k {
			count += end - start + 1
		}

		// shrink region size smaller by advance start position
		product /= nums[start]
		start++

		// single number might >= k
		if start > end {
			end++
			if end < len(nums) {
				product = nums[end]
			}
		}
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
//	1.	be careful about conditions to shrink and expand region

//		even single number should be counted, when low == high, that number
//		should be counted if nums[low] < k, which influence the expand condition:
//		low > high should expand

//	2.	inspired from https://leetcode.com/problems/subarray-product-less-than-k/discuss/108861/JavaC%2B%2B-Clean-Code-with-Explanation

//		the idea is to maintain product always < k, and count sub-array size
//		it's perspective focus on end of region, every time expand one size
//		larger, then find it's maximum region meets criteria

//	3.	on 9/29, my thinking is focusing on start of region, find its possible
//		maximum region

//		at first my think was about expand/shrink region size by 1 considering
//		condition, but soon I realize that it's a little tedious and complicated

//	4.	sliding window (two pointers) idea is link searching, fix one point and
//		find region, result relates to region start/end and no permutation related

//		e.g. [10, 5, 200], k = 100
//		fix start position at 10, maximum region meets criteria is [10, 5], so
//		there are 2 conditions: [10] & [10, 5] both meets criteria, thus 2 is
//		added to total count
