package main

import (
	"container/heap"
	"sort"
)

// You are given an integer array heights representing the heights of buildings, some bricks, and some ladders.
//
// You start your journey from building 0 and move to the next building by possibly using bricks or ladders.
//
// While moving from building i to building i+1 (0-indexed),
//
//     If the current building's height is greater than or equal to the next building's height, you do not need a ladder or bricks.
//     If the current building's height is less than the next building's height, you can either use one ladder or (h[i+1] - h[i]) bricks.
//
// Return the furthest building index (0-indexed) you can reach if you use the given ladders and bricks optimally.
//
//
//
// Example 1:
//
// Input: heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
// Output: 4
// Explanation: Starting at building 0, you can follow these steps:
// - Go to building 1 without using ladders nor bricks since 4 >= 2.
// - Go to building 2 using 5 bricks. You must use either bricks or ladders because 2 < 7.
// - Go to building 3 without using ladders nor bricks since 7 >= 6.
// - Go to building 4 using your only ladder. You must use either bricks or ladders because 6 < 9.
// It is impossible to go beyond building 4 because you do not have any more bricks or ladders.
//
// Example 2:
//
// Input: heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
// Output: 7
//
// Example 3:
//
// Input: heights = [14,3,19,3], bricks = 17, ladders = 0
// Output: 3
//
//
//
// Constraints:
//
//     1 <= heights.length <= 105
//     1 <= heights[i] <= 106
//     0 <= bricks <= 109
//     0 <= ladders <= heights.length

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() int          { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	popped := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return popped
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	size := len(heights)
	var i int

	for ; i < size-1; i++ {
		if diff := heights[i+1] - heights[i]; diff > 0 {
			// push heap
			heap.Push(maxHeap, diff)

			// adjust heap to meet some condition
			for diff > bricks && maxHeap.Len() > 0 && ladders > 0 {
				ladders--
				bricks += heap.Pop(maxHeap).(int)
			}

			// not enough bricks
			if diff > bricks {
				if ladders == 0 {
					// no ladder, cannot go any further
					break
				} else {
					// use ladder
					ladders--
				}
			} else {
				bricks -= diff
			}
		}
	}

	return i
}

func furthestBuilding3(heights []int, bricks int, ladders int) int {
	jumps := make([][2]int, 0)
	size := len(heights)

	for i := 0; i < size-1; i++ {
		if heights[i+1] > heights[i] {
			jumps = append(jumps, [2]int{heights[i+1] - heights[i], i + 1})
		}
	}

	sort.Slice(jumps, func(i, j int) bool {
		if jumps[i][0] != jumps[j][0] {
			return jumps[i][0] < jumps[j][0]
		}

		return jumps[i][1] < jumps[j][1]
	})

	// binary search
	farthest := size - 1 // in case all successful
	for low, high := 0, size-1; low <= high; {
		mid := low + (high-low)/2

		if check(jumps, mid, bricks, ladders) {
			low = mid + 1
		} else {
			high = mid - 1
			farthest = high
		}
	}

	return farthest
}

// tc: O(n)
func check(jumps [][2]int, idx, bricks, ladders int) bool {
	for i := len(jumps) - 1; i >= 0; i-- {
		if jumps[i][1] <= idx {
			if ladders > 0 {
				ladders--
			} else {
				bricks -= jumps[i][0]
			}
		}

		if bricks < 0 {
			return false
		}
	}

	return true
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	popped := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return popped
}

func furthestBuilding2(heights []int, bricks int, ladders int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	var i int

	for ; i < len(heights)-1; i++ {
		if diff := heights[i+1] - heights[i]; diff > 0 {
			heap.Push(minHeap, diff)

			for minHeap.Len() > ladders {
				bricks -= heap.Pop(minHeap).(int)
			}

			// because compute next building, so return current
			if bricks < 0 {
				return i
			}
		}
	}

	return i
}

func furthestBuilding1(heights []int, bricks int, ladders int) int {
	h := &MaxHeap{}
	heap.Init(h)

	size := len(heights)
	if size == 1 {
		return 0
	}

	for i := 1; i < size; i++ {
		diff := heights[i] - heights[i-1]

		if diff > 0 {
			if bricks >= diff {
				bricks -= diff
				heap.Push(h, diff)
			} else {
				if ladders == 0 {
					return i - 1
				}

				ladders--
				heap.Push(h, diff)
				pop := heap.Pop(h).(int)

				// previous crossed gap larger, replace that with ladder and use
				// part of that bricks to cover next jump
				if pop != diff {
					bricks += pop - diff
				}
			}
		}
	}

	return size - 1
}

//	Notes
//	1.	ladders can be treat as single use of infinite gap, so better to use it
//		for as larger gap as possible

//		the whole behavior should use bricks first, if bricks is not enough,
//		replace bricks with ladders

//		what I first write is using deque, although passes tests, but it's not
//		correct, considering following test case:
//		heights = [2, 7, 9, 12]
//		bricks = 5
//		ladders = 1

//	2.	though this problem is greedy and got WA, think over again and recall
//		that this problem is actually max-heap problem

//	3.	for heap problem, it could be a template that push something into
//		heap, adjust heap to make some condition valid, and do after process

//	4.	inspired from solution, it uses min-heap to solve the problem

//		the way is to use ladders first, and when there's not enough ladders,
//		pop from min-heap and use bricks instead

//		very interesting view to see the problem

//	5.	inspired from solution, it's also possible to use binary search to
//		solve the problem

//		given some index, it's possible to sort all jumps, avoid largest number
//		by ladder, compute remaining sums < bricks

//		very smart..., tc is slightly slower, each time sort takes O(n log(n)),
//		there are log(n) times to guess, overall tc O(n log^2(n))

//	6.	to improve tc for binary search, just sort all jumps, attach with
//		index, so that it's possible to identify the jump not in current range

//		tc: O(n log(n))
//		n: check time for each guess
//		log(n): guess count
