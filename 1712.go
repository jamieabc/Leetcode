package main

// A split of an integer array is good if:
//
//     The array is split into three non-empty contiguous subarrays - named left, mid, right respectively from left to right.
//     The sum of the elements in left is less than or equal to the sum of the elements in mid, and the sum of the elements in mid is less than or equal to the sum of the elements in right.
//
// Given nums, an array of non-negative integers, return the number of good ways to split nums. As the number may be too large, return it modulo 109 + 7.
//
//
//
// Example 1:
//
// Input: nums = [1,1,1]
// Output: 1
// Explanation: The only good way to split nums is [1] [1] [1].
//
// Example 2:
//
// Input: nums = [1,2,2,2,5,0]
// Output: 3
// Explanation: There are three good ways of splitting nums:
// [1] [2] [2,2,5,0]
// [1] [2,2] [2,5,0]
// [1,2] [2,2] [5,0]
//
// Example 3:
//
// Input: nums = [3,2,1]
// Output: 0
// Explanation: There is no good way to split nums.
//
//
//
// Constraints:
//
//     3 <= nums.length <= 105
//     0 <= nums[i] <= 104

func waysToSplit(nums []int) int {
	mod := int(1e9 + 7)
	size := len(nums)

	sums := make([]int, size)
	sums[0] = nums[0]
	for i := 1; i < size; i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	var count, start, end int

	for i := range nums {
		start = max(start, i+1)

		// sums[i] <= sums[j] - sums[i]
		// 2 * sums[i] <= sums[j]
		for ; start < size-1; start++ {
			if 2*sums[i] <= sums[start] {
				break
			}
		}

		// end is the first index violates condition, so
		// it needs to check from start to make sure at
		// least one index meets condition
		end = min(size-1, max(start, end))

		// total - sums[j] >= sums[j] - sums[i]
		// total + sums[i] >= 2 * sums[j]
		for ; end < size-1; end++ {
			if sums[size-1]+sums[i] < 2*sums[end] {
				break
			}
		}

		// end is the first one fails condition, end should
		// at least > start
		if end > start {
			count = (count + end - start) % mod
		}
	}

	return count
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

func waysToSplit1(nums []int) int {
	size := len(nums)
	sums := make([]int, size)
	sums[0] = nums[0]
	for i := 1; i < size; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	mod := int(1e9 + 7)

	var count int
	for i := range nums {
		start, end := binarySearch(sums, i+1, sums[i]<<1, sums[size-1]+sums[i])

		if end > 0 && end >= start {
			count = (count + end - start + 1) % mod
		}
	}

	return count
}

func binarySearch(sums []int, begin, left, right int) (int, int) {
	var start, end int
	size := len(sums)

	for low, high := begin, size-2; low <= high; {
		mid := low + (high-low)>>1

		if sums[mid] >= left {
			start = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if start == 0 {
		return 0, 0
	}

	for low, high := begin, size-2; low <= high; {
		mid := low + (high-low)>>1

		if sums[mid]<<1 <= right {
			end = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return start, end
}

//	Notes
//	1.	all numbers >= 0, prefix sums become non-decreasing array

//	2.	if i, j are 2 separation points in array, to meets criteria:
//		- s[i] <= s[j] - s[i]
//		- total - s[j] >= s[j] - s[i]

//		above conditions turn into:
//		- s[j] >= 2 * s[i] 			---- (1)
//		- total + s[i] >= 2 * s[j]	---- (2)

//		because all numbers >= 0, prefix sum will be non-decreasing

//		for condition (1), as j increases, s[j] >= s[j-1], so first index meets
//		condition is the start of range

//		for condition (2), as j increases, s[j] >= s[j-1], so first index violates
//		condition is the end of range

//		. . . . . . i . . . . . start . . . . . . . . . . .
//		. . . . . . i . . . . . . . . . . . . end . . . . .

//		start & end forms a range meets conditions

//		as i increases, s[i] increase, previous valid j could also be same or
//		increased (s[j] >= 2 * s[i], i increases, j might also needs to increase)

//		this is a monotonic behavior, acts as a sliding window

//		can use this to find appropriate i, j

//	3.	iterate i from 0 ~ size-2, iterate j from i+1 ~ size-1 takes O(n^2)

//		inspired from https://youtu.be/qW4H7Yc_ynQ?t=2186

//		since it's non-decreasing subsequences, start of j can start from previous i
//		location, because nums[i] >= 0, start of j is either same or increased

//		that reduces problem into sliding window, takes O(n)

//	4.	for implementation, it needs to carefully check boundary conditions,
//		e.g. all numbers are same, etc.
