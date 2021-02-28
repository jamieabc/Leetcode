package main

import "container/heap"

// You are given two arrays of integers nums1 and nums2, possibly of different lengths. The values in the arrays are between 1 and 6, inclusive.
//
// In one operation, you can change any integer's value in any of the arrays to any value between 1 and 6, inclusive.
//
// Return the minimum number of operations required to make the sum of values in nums1 equal to the sum of values in nums2. Return -1 if it is not possible to make the sum of the two arrays equal.
//
//
//
// Example 1:
//
// Input: nums1 = [1,2,3,4,5,6], nums2 = [1,1,2,2,2,2]
// Output: 3
// Explanation: You can make the sums of nums1 and nums2 equal with 3 operations. All indices are 0-indexed.
// - Change nums2[0] to 6. nums1 = [1,2,3,4,5,6], nums2 = [6,1,2,2,2,2].
// - Change nums1[5] to 1. nums1 = [1,2,3,4,5,1], nums2 = [6,1,2,2,2,2].
// - Change nums1[2] to 2. nums1 = [1,2,2,4,5,1], nums2 = [6,1,2,2,2,2].
// Example 2:
//
// Input: nums1 = [1,1,1,1,1,1,1], nums2 = [6]
// Output: -1
// Explanation: There is no way to decrease the sum of nums1 or to increase the sum of nums2 to make them equal.
// Example 3:
//
// Input: nums1 = [6,6], nums2 = [1]
// Output: 3
// Explanation: You can make the sums of nums1 and nums2 equal with 3 operations. All indices are 0-indexed.
// - Change nums1[0] to 2. nums1 = [2,6], nums2 = [1].
// - Change nums1[1] to 2. nums1 = [2,2], nums2 = [1].
// - Change nums2[0] to 4. nums1 = [2,2], nums2 = [4].
//
//
// Constraints:
//
// 1 <= nums1.length, nums2.length <= 105
// 1 <= nums1[i], nums2[i] <= 6

// tc: O(n+m)
func minOperations(nums1 []int, nums2 []int) int {
	counter1, counter2 := make([]int, 7), make([]int, 7)
	var sum1, sum2 int

	for _, n := range nums1 {
		sum1 += n
		counter1[n]++
	}

	for _, n := range nums2 {
		sum2 += n
		counter2[n]++
	}

	if sum1 == sum2 {
		return 0
	}

	// impossible
	m, n := len(nums1), len(nums2)
	if m > 6*n || m*6 < n {
		return -1
	}

	var arr1, arr2 []int

	// make condition more general
	if sum1 >= sum2 {
		arr1, arr2 = counter1, counter2
	} else {
		arr1, arr2 = counter2, counter1
		sum1, sum2 = sum2, sum1
	}

	var ans int
	diff := sum1 - sum2

	for i := 5; i >= 1 && diff > 0; i-- {
		delta := i - 1
		total := arr1[i+1] + arr2[6-i]

		if reduction := delta * total; reduction < diff {
			ans += total
			diff -= reduction
		} else {
			remain := diff % delta

			ans += diff / delta
			if remain > 0 {
				ans++
			}
			break
		}
	}

	return ans
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

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int          { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// tc: O((n+m) log(n+m))
func minOperations1(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)

	// impossible
	if m*6 < n || n*6 < m {
		return -1
	}

	var sum1, sum2 int
	mh := &MaxHeap{}
	heap.Init(mh)

	for _, i := range nums1 {
		sum1 += i
	}

	for _, i := range nums2 {
		sum2 += i
	}

	var arr1, arr2 []int
	if sum1 >= sum2 {
		arr1, arr2 = nums1, nums2
	} else {
		arr1, arr2 = nums2, nums1
		sum1, sum2 = sum2, sum1
	}

	diff := sum1 - sum2

	if diff == 0 {
		return 0
	}

	// larger array reduces sum, 6 -> 1, 5 -> 1, etc.
	for _, i := range arr1 {
		heap.Push(mh, i-1)
	}

	// smaller array increases sum, 1 -> 6, 2 -> 6, etc.
	for _, i := range arr2 {
		heap.Push(mh, 6-i)
	}

	var ans int

	for diff > 0 {
		n := heap.Pop(mh).(int)
		ans++
		diff -= n
	}

	return ans
}

//	Notes
//	1.	didn't think of solution during contest

//	2.	inspired from alex and https://leetcode.com/problems/equal-sum-arrays-with-minimum-number-of-operations/discuss/1085847/Python-3-simple-greedy-solution

//		this could be a greedy problem, choose maximum difference as possible

//		the beautiful part is to convert minimum changes to let every number is either
//		1 or 6, which has smallest changes

//		2, 2, 2, 3, 3 	=> sum 12
//		3, 3		 	=> sum 6

//		to make smallest change, let every modification has biggest impact

//		- to have difference -5, comes from larger array from 6 -> 1 or smaller array 1 -> 6
//		  since there's no 1 or 6 digits, skip

//		- to have difference -4, comes from larger array from 5 -> 1 or smaller array 2 -> 6
//		  since larger array has no 5 and smaller has no 2, skip

//		- to have difference -3, comes from larger array from 4 -> 1 or smaller array 3 -> 6
//		  larger array has no 4, but smaller array has 2 3, reduce difference
