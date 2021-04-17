package main

import "fmt"

// A Range Module is a module that tracks ranges of numbers. Your task is to design and implement the following interfaces in an efficient manner.
//
// addRange(int left, int right) Adds the half-open interval [left, right), tracking every real number in that interval. Adding an interval that partially overlaps with currently tracked numbers should add any numbers in the interval [left, right) that are not already tracked.
//
// queryRange(int left, int right) Returns true if and only if every real number in the interval [left, right) is currently being tracked.
//
// removeRange(int left, int right) Stops tracking every real number currently being tracked in the interval [left, right).
//
// Example 1:
//
// addRange(10, 20): null
// removeRange(14, 16): null
// queryRange(10, 14): true (Every number in [10, 14) is being tracked)
// queryRange(13, 15): false (Numbers like 14, 14.03, 14.17 in [13, 15) are not being tracked)
// queryRange(16, 17): true (The number 16 in [16, 17) is still being tracked, despite the remove operation)
//
// Note:
// A half open interval [left, right) denotes all real numbers left <= x < right.
// 0 < left < right < 10^9 in all calls to addRange, queryRange, removeRange.
// The total number of calls to addRange in a single test case is at most 1000.
// The total number of calls to queryRange in a single test case is at most 5000.
// The total number of calls to removeRange in a single test case is at most 1000.

type RangeModule struct {
	Intervals [][]int
}

func Constructor() RangeModule {
	return RangeModule{
		Intervals: make([][]int, 0),
	}
}

func (this *RangeModule) AddRange(left int, right int) {
	size := len(this.Intervals)

	// linear scan existing
	arr := make([][]int, 0)
	var inserted bool

	for i := 0; i < size; i++ {
		intr := this.Intervals[i]

		if independent(intr, []int{left, right}) {
			// in case inserted interval missed
			if !inserted && intr[0] >= right {
				arr = append(arr, []int{left, right})
				inserted = true
			}

			arr = append(arr, intr)
		} else {
			for ; i < size && !independent(this.Intervals[i], []int{left, right}); i++ {
				left = min(left, this.Intervals[i][0])
				right = max(right, this.Intervals[i][1])
			}
			i--

			arr = append(arr, []int{left, right})
			inserted = true
		}
	}

	if !inserted {
		arr = append(arr, []int{left, right})
	}

	this.Intervals = arr
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	size := len(this.Intervals)

	if size == 0 || this.Intervals[0][0] >= right || this.Intervals[size-1][1] <= left {
		return false
	}

	// binary search
	var idx int
	for low, high := 0, size-1; low <= high; {
		mid := low + (high-low)/2

		if this.Intervals[mid][0] == left {
			idx = mid
			break
		} else if this.Intervals[mid][0] < left {
			idx = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// check contain
	if intr := this.Intervals[idx]; intr[0] <= left && intr[1] >= right {
		return true
	}
	return false
}

func (this *RangeModule) RemoveRange(left int, right int) {
	// this.Idx++
	if len(this.Intervals) == 0 {
		return
	}

	arr := make([][]int, 0)
	for _, intr := range this.Intervals {
		if independent(intr, []int{left, right}) {
			arr = append(arr, intr)
		} else {
			if intr[0] >= left && intr[1] <= right {
				// inside removed range, skip this one
				continue
			} else {
				// partial match
				if intr[0] >= left {
					arr = append(arr, []int{right, intr[1]})
				} else if intr[1] <= right {
					arr = append(arr, []int{intr[0], left})
				} else {
					arr = append(arr, []int{intr[0], left})
					arr = append(arr, []int{right, intr[1]})
				}
			}
		}
	}

	this.Intervals = arr
}

func independent(intr1, intr2 []int) bool {
	return intr1[1] < intr2[0] || intr2[1] < intr1[0] || intr1[0] > intr2[1] || intr2[0] > intr1[1]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	insert: linear scan, like merge into intervals

//		remove: linear scan

//		query: binary search by start time

//	2.	becareful about overlap, == can be viewed as overlap

//	3.	inspired from https://leetcode.com/problems/range-module/discuss/108912/C%2B%2B-vector-O(n)-and-map-O(logn)-compare-two-solutions

//	4.	for add range, need to check if inserted interval is not overlap with
//		any existing one but already crossed
