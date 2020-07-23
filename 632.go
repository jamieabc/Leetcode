package main

import (
	"container/heap"
	"math"
)

// You have k lists of sorted integers in ascending order. Find the smallest range that includes at least one number from each of the k lists.
//
// We define the range [a,b] is smaller than range [c,d] if b-a < d-c or a < c if b-a == d-c.
//
//
//
// Example 1:
//
// Input: [[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
// Output: [20,24]
// Explanation:
// List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
// List 2: [0, 9, 12, 20], 20 is in range [20,24].
// List 3: [5, 18, 22, 30], 22 is in range [20,24].
//
//
// Note:
//
// The given list may contain duplicates, so ascending order means >= here.
// 1 <= k <= 3500
// -105 <= value of elements <= 105.

type Num struct {
	Val        int
	ArrayIndex int
	ItemIndex  int
}

type MinHeap []Num

func (h MinHeap) Len() int      { return len(h) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() Num     { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Num))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func smallestRange(nums [][]int) []int {
	h := &MinHeap{}
	heap.Init(h)

	maxNum := math.MinInt32
	for i := range nums {
		heap.Push(h, Num{
			Val:        nums[i][0],
			ArrayIndex: i,
			ItemIndex:  0,
		})
		maxNum = max(maxNum, nums[i][0])
	}

	// find minimum range
	minRange := maxNum - h.Peek().Val
	start, end := h.Peek().Val, maxNum

	for h.Len() == len(nums) {
		// find minimum among all array
		popped := heap.Pop(h).(Num)

		// find minimum range & update start/end
		if popped.ItemIndex < len(nums[popped.ArrayIndex])-1 {
			popped.ItemIndex++
			popped.Val = nums[popped.ArrayIndex][popped.ItemIndex]
			maxNum = max(maxNum, nums[popped.ArrayIndex][popped.ItemIndex])
			heap.Push(h, popped)

			if maxNum-h.Peek().Val < minRange {
				minRange = maxNum - h.Peek().Val
				start, end = h.Peek().Val, maxNum
			}
		}
	}

	return []int{start, end}
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	inspired from https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/discuss/104905/Python-Straightforward-with-Explanation
