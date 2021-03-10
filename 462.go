package main

import (
	"container/heap"
	"sort"
)

// Given a non-empty integer array, find the minimum number of moves required to make all array elements equal, where a move is incrementing a selected element by 1 or decrementing a selected element by 1.
//
// You may assume the array's length is at most 10,000.
//
// Example:
//
// Input:
// [1,2,3]
//
// Output:
// 2
//
// Explanation:
// Only two moves are needed (remember each move increments or decrements one element):
//
// [1,2,3]  =>  [2,2,3]  =>  [2,2,2]

func minMoves2(nums []int) int {
	size := len(nums)
	target := (size+1)>>1 - 1

	quickSelect(nums, 0, size-1, target)

	median := nums[target]

	var ans int

	for _, n := range nums {
		ans += abs(median - n)
	}

	return ans
}

func quickSelect(nums []int, start, end, target int) {
	if start >= end {
		return
	}

	store := start
	pivot := nums[start]
	nums[start], nums[end] = nums[end], nums[start]

	for i := start; i <= end; i++ {
		if nums[i] < pivot {
			nums[store], nums[i] = nums[i], nums[store]
			store++
		}
	}

	nums[store], nums[end] = nums[end], nums[store]

	size := len(nums)
	if (size&1 > 0 && store == target) || (size&1 == 0 && (store == target || store == target-1)) {
		return
	}

	if store > target {
		quickSelect(nums, start, store-1, target)
	} else if store < target {
		quickSelect(nums, store+1, end, target)
	}
}

func minMoves2_2(nums []int) int {
	sort.Ints(nums)

	var ans int

	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		// no abs, since already sorted
		ans += nums[r] - nums[l]
	}

	return ans
}

type MaxHeap [][]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() []int        { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMoves2_1(nums []int) int {
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
	}

	mh := &MaxHeap{}
	heap.Init(mh)

	for num, count := range counter {
		heap.Push(mh, []int{count, num})
	}

	var median, count int
	size := len(nums)

	for mh.Len() > 0 {
		popped := heap.Pop(mh).([]int)
		count += popped[0]

		if count >= (size+1)>>1 {
			median = popped[1]
			break
		}
	}

	var ans int

	for _, n := range nums {
		ans += abs(median - n)
	}

	return ans
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	Notes
//	1.	read problem carefully, every move can only increase/decrease by 1

//	2.	inspired from solution, instead of finding median, it means the middle of
//		sorted array, thus start from left & right moving toward middle, also works
