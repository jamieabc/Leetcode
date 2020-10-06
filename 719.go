package main

import (
	"fmt"
	"math"
	"sort"
)

// Given an integer array, return the k-th smallest distance among all the pairs. The distance of a pair (A, B) is defined as the absolute difference between A and B.
//
// Example 1:
// Input:
// nums = [1,3,1]
// k = 1
// Output: 0
// Explanation:
// Here are all the pairs:
// (1,3) -> 2
// (1,1) -> 0
// (3,1) -> 2
// Then the 1st smallest distance pair is (1,1), and its distance is 0.
// Note:
// 2 <= len(nums) <= 10000.
// 0 <= nums[i] < 1000000.
// 1 <= k <= len(nums) * (len(nums) - 1) / 2.

// tc: O(n * logn + n * logn)
// for kth distance, distance sequence may look like: 0, 0, 1, 1, 2, ..., x, x, x+1, ...
// kth smallest distance means there are at most k-1 smaller numbers ahead, and
// smaller_number_count + same_number_count >= k
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	size := len(nums)

	high := nums[size-1] - nums[0]
	low := math.MaxInt32
	for i := 1; i < len(nums); i++ {
		low = min(low, nums[i]-nums[i-1])
	}

	for low < high {
		mid := low + (high-low)/2

		smaller := countSmaller(nums, mid)
		same := countSmaller(nums, mid+1) - smaller

		// dist: 0, 0, 0, ..., mid-1, mid-1, mid, mid, ...
		// if mid is the kth smallest distance
		// k smallest means there are at most k-1 smaller numbers, and at least one number
		// same as guessed
		if smaller < k && smaller+same >= k {
			return mid
		}

		if smaller >= k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

// since numbers are sorted, next number will >= current, so previous right boundary
// will either same or extendr, it means previous boundary can be reused
func countSmaller(nums []int, dist int) int {
	var smaller int

	for i, right := 0, 1; i < len(nums); i++ {
		for right < len(nums) && nums[right]-nums[i] < dist {
			right++
		}
		smaller += right - i - 1
	}

	return smaller
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

// tc: O(n * logn + nk * logk)
// for every number, kth smallest comes within range idx ~ idx+k
// use a max heap size k to track kth smallest
func smallestDistancePair2(nums []int, k int) int {
	sort.Ints(nums)

	h := &MaxHeap{}
	heap.Init(h)

	for i := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			dist := abs(nums[i] - nums[j])

			if h.Len() < k {
				heap.Push(h, dist)
			} else if dist <= h.Peek() {
				heap.Push(h, dist)
				heap.Pop(h)
			}
		}
	}

	return h.Peek()
}

// tc: O(n^2)
func smallestDistancePair1(nums []int, k int) int {
	dist := make([]int, 0)

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			dist = append(dist, abs(nums[i]-nums[j]))
		}
	}

	tmp := nthElement(dist, 0, len(dist)-1, k-1)
	return tmp
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func nthElement(dist []int, start, end, k int) int {
	if start >= end {
		return dist[k]
	}

	pivot, store := dist[start], start
	dist[end], dist[start] = dist[start], dist[end]

	for i := start; i <= end; i++ {
		if dist[i] <= pivot {
			dist[i], dist[store] = dist[store], dist[i]
			store++
		}
	}
	store--

	if store == k {
		return dist[k]
	} else if store > k {
		return nthElement(dist, start, store-1, k)
	}
	return nthElement(dist, store+1, end, k)
}

//	Notes
//	1.	for any given numbers, tc: O(n^2)

//	2.	nth element using quick select, need to move pivot to end of range, such
//		that final situation meets all numbers <= pivot

//	3.	nth element, if start >= end, no need to do

//	4.	for any given numbers, even if sorted, distance among all pairs are still
//		not in order, consider following sequences:

//		1, 2, 3, 4, 5, 100, 1000, 10000, 50000
//		1, 100, 1000, 10000, 10001, 10002, 10003

//		smallest distance might come from anywhere

//		image distance sequence among all pairs like: 0, 0, 1, 2, ..., x-1, x, x, ...
//		if a distance is kth smallest, it meas at most k-1 smaller ahead (because
//		there also might be same number exist)

//		smaller_number_count + same_number_count >= k

//	5.	to count # of same number x, it could be count_smaller(x+1) - count_smaller(x)

//	6.	if numbers are sorted, for any given number i, if maximum range smaller
//		up to j since next number always >= current number, next number range
//		also >= j

//		e.g. numbers: [1, 2, 3, 4, 5, 6, 7]
//		i = 3, maximum distance = 2, so [4, 5] meets criteria
//		for next number 4, right boundary must be >= index 4 (5)

//		as above described, to count # of larger numbers, it takes O(n)

//	7.	for any given distance guess y, it takes O(n) to check if this distance
//		is kth smallest, and can use it to kth smallest distance is larger or
//		smaller than y, which is O(logn)

//	8.	minimum distance can be first calculated
