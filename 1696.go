package main

import "math"

// You are given a 0-indexed integer array nums and an integer k.
//
// You are initially standing at index 0. In one move, you can jump at most k steps forward without going outside the boundaries of the array. That is, you can jump from index i to any index in the range [i + 1, min(n - 1, i + k)] inclusive.
//
// You want to reach the last index of the array (index n - 1). Your score is the sum of all nums[j] for each index j you visited in the array.
//
// Return the maximum score you can get.
//
//
//
// Example 1:
//
// Input: nums = [1,-1,-2,4,-7,3], k = 2
// Output: 7
// Explanation: You can choose your jumps forming the subsequence [1,-1,4,3] (underlined above). The sum is 7.
//
// Example 2:
//
// Input: nums = [10,-5,-2,4,0,3], k = 3
// Output: 17
// Explanation: You can choose your jumps forming the subsequence [10,4,3] (underlined above). The sum is 17.
//
// Example 3:
//
// Input: nums = [1,-5,-20,4,-1,3,-6,-3], k = 2
// Output: 0
//
//
//
// Constraints:
//
//      1 <= nums.length, k <= 105
//     -104 <= nums[i] <= 104

func maxResult(nums []int, k int) int {
	size := len(nums)
	queue := make([][]int, 0)

	// dp[i]: max score can get up to i
	dp := make([]int, size)
	dp[0] = nums[0]
	queue = append(queue, []int{nums[0], 0})

	// dp[i] = A[i] + max(dp[i-1], dp[i-2], ..., dp[i-k])
	for i := 1; i < size; i++ {
		// remove out of range values
		for len(queue) > 0 && queue[0][1] < i-k {
			queue = queue[1:]
		}

		dp[i] = nums[i] + queue[0][0]

		// make qeuue a mono-decreasing sequence
		for len(queue) > 0 && dp[i] > queue[len(queue)-1][0] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, []int{dp[i], i})
	}

	return dp[size-1]
}

type MaxHeap [][]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
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

// tc: O(n log(n))
func maxResult1(nums []int, k int) int {
	size := len(nums)
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	// dp[i]: max score can get up to i
	dp := make([]int, size)
	dp[0] = nums[0]
	heap.Push(maxHeap, []int{nums[0], 0})

	// dp[i] = A[i] + max(dp[i-1], dp[i-2], ..., dp[i-k])
	for i := 1; i < size; i++ {
		// remove out of range values
		for maxHeap.Len() > 0 && maxHeap.Peek()[1] < i-k {
			heap.Pop(maxHeap)
		}

		dp[i] = nums[i] + maxHeap.Peek()[0]
		heap.Push(maxHeap, []int{dp[i], i})
	}

	return dp[size-1]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from https://www.youtube.com/watch?v=cCNJkk0Etaw

//		use max-heap or max-que to optimize, because next jump depends on max
//		value among some range, no need to compare every time

//		max-heap is straight-forward, although I didn't think of it during contest

//		max-queue is also straight forward, maintain numbers in mono-decreasing
//		order

//	2.	solution provides a good explanation
