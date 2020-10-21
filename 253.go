package main

import "sort"

// Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.
//
// Example 1:
//
// Input: [[0, 30],[5, 10],[15, 20]]
// Output: 2
// Example 2:
//
// Input: [[7,10],[2,4]]
// Output: 1

func minMeetingRooms(intervals [][]int) int {
	size := len(intervals)
	if size == 0 {
		return 0
	}

	start, end := make([]int, size), make([]int, size)

	for i, intr := range intervals {
		start[i] = intr[0]
		end[i] = intr[1]
	}

	sort.Ints(start)
	sort.Ints(end)

	var rooms, maxRoom int

	for i, j := 0, 0; i < size && j < size; {
		if start[i] < end[j] {
			i++
			rooms++
			maxRoom = max(maxRoom, rooms)
		} else if start[i] > end[j] {
			rooms--
			j++
		} else {
			// one meeting end, another meeting start, total room used no change
			i++
			j++
		}
	}

	return maxRoom
}

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

func minMeetingRooms1(intervals [][]int) int {
	size := len(intervals)
	if size == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	minEnd := &MinHeap{}
	heap.Init(minEnd)

	maxRoom := 1
	var rooms int

	for i := 0; i < size; i++ {
		if minEnd.Len() == 0 || minEnd.Peek() > intervals[i][0] {
			heap.Push(minEnd, intervals[i][1])
			maxRoom = max(maxRoom, minEnd.Len())
		} else {
			heap.Pop(minEnd)
			i--
		}
	}

	return maxRoom
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	for number of rooms needed, it's the number that how many meetings still
//		going on

//	2.	as long as a room is occupied, it's freed only when meeting is over
