package main

import (
	"container/heap"
	"math"
	"sort"
)

// Given an array nums, you are allowed to choose one element of nums and change it by any value in one move.
//
// Return the minimum difference between the largest and smallest value of nums after perfoming at most 3 moves.
//
//
//
// Example 1:
//
// Input: nums = [5,3,2,4]
// Output: 0
// Explanation: Change the array [5,3,2,4] to [2,2,2,2].
// The difference between the maximum and minimum is 2-2 = 0.
//
// Example 2:
//
// Input: nums = [1,5,0,10,14]
// Output: 1
// Explanation: Change the array [1,5,0,10,14] to [1,1,0,1,1].
// The difference between the maximum and minimum is 1-0 = 1.
//
// Example 3:
//
// Input: nums = [6,6,0,1,1,4,6]
// Output: 2
//
// Example 4:
//
// Input: nums = [1,5,6,14,15]
// Output: 1
//
//
//
// Constraints:
//
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

func minDifference(nums []int) int {
	size := len(nums)
	if size <= 4 {
		return 0
	}

	sort.Ints(nums)

	minimum := nums[size-1] - nums[3]

	for i := 1; i <= 3; i++ {
		minimum = min(minimum, nums[size-1-i]-nums[3-i])
	}

	return minimum
}

func min(i, j int) int {
	if i <= j {
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
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() int          { return h[0] }
func (h *MinHeap) Push(x interface{}) {
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minDifference1(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}

	maxNums, minNums := &MaxHeap{}, &MinHeap{}
	heap.Init(maxNums)
	heap.Init(minNums)
	sizeLimit := 4

	// find max 4 & min 4 numbers
	for _, n := range nums {
		// not full: push
		// full and top number >= n: pop then push
		if maxNums.Len() == sizeLimit && maxNums.Peek() >= n {
			heap.Pop(maxNums)
		}
		if maxNums.Len() < sizeLimit {
			heap.Push(maxNums, n)
		}

		// not full: push
		// full and top number <= n: pop ten push
		if minNums.Len() == sizeLimit && minNums.Peek() <= n {
			heap.Pop(minNums)
		}
		if minNums.Len() < sizeLimit {
			heap.Push(minNums, n)
		}
	}

	top4 := make([]int, sizeLimit)
	for i := sizeLimit - 1; minNums.Len() > 0; i-- {
		top4[i] = heap.Pop(minNums).(int)
	}

	bottom4 := make([]int, sizeLimit)
	for i := sizeLimit - 1; maxNums.Len() > 0; i-- {
		bottom4[i] = heap.Pop(maxNums).(int)
	}

	// find all possibilities and select minimum difference
	minDiff := math.MaxInt32
	for i := 0; i < sizeLimit; i++ {
		minDiff = min(minDiff, abs(top4[i]-bottom4[sizeLimit-1-i]))
	}

	return minDiff
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	Notes
//	1.	after finding max4 & min4, the problem becomes dp, not greedy

//		even choose min difference among top4 & min4, still not guarantee global
//		minimum difference

//	2.	only remove either largest or smallest to reduce difference, so it's all about
//		removing consecutive numbers from boundary
