package main

import (
	"container/heap"
	"math/rand"
)

// Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
//
// Example 1:
//
// Input: [3,2,1,5,6,4] and k = 2
// Output: 5
//
// Example 2:
//
// Input: [3,2,3,1,2,4,5,5,6] and k = 4
// Output: 4
//
// Note:
// You may assume k is always valid, 1 ≤ k ≤ array's length.

// tc: average O(n), worst O(n^2)
func findKthLargest(nums []int, k int) int {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	// kth from 1 ~ n, so convert it to index base
	quickSelect(nums, 0, len(nums)-1, k-1)

	return nums[k-1]
}

func quickSelect(nums []int, target, start, end int) int {
	store := start
	nums[start], nums[end] = nums[end], nums[start]

	// partition
	for i := start; i < end; i++ {
		// becareful, it's kth largest, order is desc
		if nums[i] > nums[end] {
			nums[store], nums[i] = nums[i], nums[store]
			store++
		}
	}

	nums[store], nums[end] = nums[end], nums[store]

	if store == target {
		return nums[store]
	} else if store > target {
		return quickSelect(nums, target, start, store-1)
	}
	return quickSelect(nums, target, store+1, end)
}

// min heap
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() int          { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// tc: O(n log(k))
func findKthLargest1(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, n := range nums {
		if h.Len() == k && h.Peek() < n {
			heap.Pop(h)
		}

		if h.Len() < k {
			heap.Push(h, n)
		}
	}

	return h.Peek()
}

//	Notes
//	1.	when finding kth largest number, I should use min heap to make sure
//		all other number in the heap are larger, means root number is kth
//		largest

//	2.	inspired from https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/60300/Java-Quick-Select

//		quick-sort like mechanism can be used here
//		time complexity is O(n), it is reduced half each time,
//		n + n/2 + n/4 + n/8 + ... + 1 = n + n-1
