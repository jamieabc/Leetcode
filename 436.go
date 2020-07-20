package main

// Given a set of intervals, for each of the interval i, check if there exists an interval j whose start point is bigger than or equal to the end point of the interval i, which can be called that j is on the "right" of i.
//
// For any interval i, you need to store the minimum interval j's index, which means that the interval j has the minimum start point to build the "right" relationship for interval i. If the interval j doesn't exist, store -1 for the interval i. Finally, you need output the stored value of each interval as an array.
//
// Note:
//
//     You may assume the interval's end point is always bigger than its start point.
//     You may assume none of these intervals have the same start point.
//
//
//
// Example 1:
//
// Input: [ [1,2] ]
//
// Output: [-1]
//
// Explanation: There is only one interval in the collection, so it outputs -1.
//
//
//
// Example 2:
//
// Input: [ [3,4], [2,3], [1,2] ]
//
// Output: [-1, 0, 1]
//
// Explanation: There is no satisfied "right" interval for [3,4].
// For [2,3], the interval [3,4] has minimum-"right" start point;
// For [1,2], the interval [2,3] has minimum-"right" start point.
//
//
//
// Example 3:
//
// Input: [ [1,4], [2,3], [3,4] ]
//
// Output: [-1, 2, -1]
//
// Explanation: There is no satisfied "right" interval for [1,4] and [3,4].
// For [2,3], the interval [3,4] has minimum-"right" start point.
//
// NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

type intr struct {
	idx, val int
}

type MaxHeap []intr

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Peek() interface{}  { return h[0] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(intr))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findRightInterval(intervals [][]int) []int {
	maxStart, maxEnd := &MaxHeap{}, &MaxHeap{}
	heap.Init(maxStart)
	heap.Init(maxEnd)

	for i := range intervals {
		heap.Push(maxStart, intr{
			idx: i,
			val: intervals[i][0],
		})

		heap.Push(maxEnd, intr{
			idx: i,
			val: intervals[i][1],
		})
	}

	result := make([]int, len(intervals))
	result[heap.Pop(maxEnd).(intr).idx] = -1

	var prev intr
	for maxEnd.Len() > 0 {
		popped := heap.Pop(maxEnd).(intr)

		for prev == (intr{}) || (maxStart.Len() > 0 && maxStart.Peek().(intr).val >= popped.val) {
			prev = heap.Pop(maxStart).(intr)
		}

		if prev.val >= popped.val {
			result[popped.idx] = prev.idx
		} else {
			result[popped.idx] = -1
		}
	}

	return result
}

//	problems
//	1.	use max heap & min heap to keep array sorted by start order, keep doing
//		this to find next larger interval takes O(n^2 log n)

//	2.	intuition: this problem is to find closest and larger start of current
//		end. so basically there needs 2 sorted sequence: start & end.

//		so it is possible to use 2 max-heap or 2 min-heap of start & end to
//		find next interval

//	3.	inspired from sample code, can also sort intervals by start time and use
//		binary search to find next interval

//	4.	start time: 8:33, end time: 10:27

//	5.	inspired from solution, since start time is unique, it can be used as
//		an index => use hash to store original start: index, and sort intervals,
//		then use binary search to find next interval
