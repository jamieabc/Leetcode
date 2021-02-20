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
	*h = old[:n-1]
	return x
}

func minimumDeviation(nums []int) int {
	for i := range nums {
		if nums[i]&1 > 0 {
			nums[i] = nums[i] << 1
		}
	}

	low := math.MaxInt32
	for _, n := range nums {
		low = min(low, n)
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, n := range nums {
		heap.Push(maxHeap, n)
	}

	ans := maxHeap.Peek() - low

	// try to reduce deviation by half largest number, if number
	// cannot be half (already even), then stop
	for true {
		n := heap.Pop(maxHeap).(int)
		if n&1 > 0 {
			break
		}

		low = min(low, n>>1)
		ans = min(ans, maxHeap.Peek()-low)
		heap.Push(maxHeap, n>>1)
	}

	return ans
}

func min(i, j int) int {
	if i <= j {
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

//	6.	inspired from https://youtu.be/OTlusbuZX94?t=1371

//		odd => larger, even => smaller, the strategy is to double all odd numbers
//		then use heap to track max numbers, if it's even, half it, need to also check
//		if new half number will be smaller than current minimum

//		I have come up with similar idea to make all even numbers to odd, but I guess
//		this is in wrong direction, because there could have multiple division, where
//		to stop?

//		e.g. 24 -> 12 -> 6 -> 3

//		and start from odd will be a pretty good strategy, because every number
//		doubled is even
