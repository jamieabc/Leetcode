package main

import "sort"

// Given a list of intervals, remove all intervals that are covered by another interval in the list.
//
// Interval [a,b) is covered by interval [c,d) if and only if c <= a and b <= d.
//
// After doing so, return the number of remaining intervals.
//
//
//
// Example 1:
//
// Input: intervals = [[1,4],[3,6],[2,8]]
// Output: 2
// Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
//
// Example 2:
//
// Input: intervals = [[1,4],[2,3]]
// Output: 1
//
// Example 3:
//
// Input: intervals = [[0,10],[5,12]]
// Output: 2
//
// Example 4:
//
// Input: intervals = [[3,10],[4,10],[5,11]]
// Output: 2
//
// Example 5:
//
// Input: intervals = [[1,2],[1,4],[3,4]]
// Output: 1
//
//
//
// Constraints:
//
//     1 <= intervals.length <= 1000
//     intervals[i].length == 2
//     0 <= intervals[i][0] < intervals[i][1] <= 10^5
//     All the intervals are unique.

func removeCoveredIntervals(intervals [][]int) int {
	// sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	var covered int

	for base, i := 0, 1; i < len(intervals); i++ {
		if intervals[i][1] <= intervals[base][1] {
			covered++
		} else {
			base = i
		}
	}

	return len(intervals) - covered
}

//	Notes
//	1.	inspired from sample code, since intervals are already sorted by start
//		time, can only check end time

//	2.	sort intervals by start time reduces covered interval check by 1 degree,
//		and if an intervals has larger end time means it can be a new comparison
//		basis

//	3.	inspired from https://leetcode.com/problems/remove-covered-intervals/discuss/451277/JavaC%2B%2BPython-Sort-Solution

//		lee thinks when covered are removed, sequence of each left & right bound
//		will be monotonically increasing, so he directly count what's correct
//		sequence
