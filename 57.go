package main

// Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
//
// You may assume that the intervals were initially sorted according to their start times.
//
// Example 1:
//
// Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
// Output: [[1,5],[6,9]]
//
// Example 2:
//
// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// Output: [[1,2],[3,10],[12,16]]
// Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
//
// NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

func insert(intervals [][]int, newInterval []int) [][]int {
	size := len(intervals)

	if size == 0 {
		return [][]int{newInterval}
	}

	arr := make([][]int, 0)
	var inserted bool
	start, end := newInterval[0], newInterval[1]

	for i := 0; i < size; i++ {
		intr := intervals[i]

		if intr[0] > end {
			if !inserted {
				arr = append(arr, []int{start, end})
				inserted = true
			}

			arr = append(arr, intr)
		} else if intr[1] < start {
			arr = append(arr, intr)
		} else {
			// overlap
			start = min(start, intr[0])
			end = max(end, intr[1])
		}
	}

	if !inserted {
		arr = append(arr, []int{start, end})
	}

	return arr
}

func insert2(intervals [][]int, newInterval []int) [][]int {
	start, end := newInterval[0], newInterval[1]
	result := make([][]int, 0)
	var i int

	// add left non-overlapping intervals
	for i = 0; i < len(intervals) && intervals[i][1] < start; i++ {
	}
	result = append(result, intervals[:i]...)

	// merge overlapping intervals
	for ; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			start, end = min(start, intervals[i][0]), max(end, intervals[i][1])
		} else {
			break
		}
	}

	result = append(result, []int{start, end})

	// add right side of non-overlapping intervals
	result = append(result, intervals[i:]...)

	return result
}

func insert1(intervals [][]int, newInterval []int) [][]int {
	size := len(intervals)
	if size == 0 {
		return [][]int{newInterval}
	}

	result := make([][]int, 0)

	// earlier than first, insert to head
	if newInterval[1] < intervals[0][0] {
		result = append(result, newInterval)
		result = append(result, intervals...)
		return result
	}

	// later than last, insert at last
	last := intervals[size-1]
	if last[1] < newInterval[0] {
		intervals = append(intervals, newInterval)
		return intervals
	}

	// middle, find where intersection begins
	var i, start, end int
	for i = 0; i < size; i++ {
		// |----| interval
		//   |------| newInterval

		// |-----| newInterval
		//     |-----| interval
		if overlap(intervals[i], newInterval) {
			start, end = min(intervals[i][0], newInterval[0]), max(intervals[i][1], newInterval[1])
			break
		} else if newInterval[1] < intervals[i][0] {
			result = append(result, newInterval)
			result = append(result, intervals[i:]...)
			return result
		} else {
			result = append(result, intervals[i])
		}
	}

	// start to merge
	for i = i + 1; i < size; i++ {
		// |----| interval
		//   |-----| newInterval

		// |------| newInterval
		//     |-------| interval
		if overlap(intervals[i], []int{start, end}) {
			start, end = min(start, intervals[i][0]), max(end, intervals[i][1])
		} else {
			break
		}
	}

	result = append(result, []int{start, end})

	if i < size {
		result = append(result, intervals[i:]...)
	}

	return result
}

func overlap(interval1, interval2 []int) bool {
	if interval1[0] > interval2[0] {
		interval1, interval2 = interval2, interval1
	}

	return !(interval2[0] > interval1[1])
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

//	problems
//	1.	when deciding start & end of overlapping regions, need to make sure
//		compare previous and current

//	2.	new interval may not necessarily overlap

//	3.	inspired from solution, use greedy solution to merge new intervals when
//		overlaps

//		This is not hard, I should have thought of it...

//	4.	inspired from https://leetcode.com/problems/insert-interval/discuss/21602/Short-and-straight-forward-Java-solution

//		the process is quite straight forward, for left intervals not overlap,
//		add to result, for overlapping intervals, merge, then put rest intervals
//		into result

//		When I see the solution, suddenly think it should be clear as that, but
//		why didn't I think of it?
