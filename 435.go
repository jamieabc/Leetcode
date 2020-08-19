package main

import (
	"math"
	"sort"
)

//Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
//
//
//
//Example 1:
//
//Input: [[1,2],[2,3],[3,4],[1,3]]
//Output: 1
//Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
//
//Example 2:
//
//Input: [[1,2],[1,2],[1,2]]
//Output: 2
//Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.
//
//Example 3:
//
//Input: [[1,2],[2,3]]
//Output: 0
//Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
//
//
//
//Note:
//
//    You may assume the interval's end point is always bigger than its start point.
//    Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// sort by end time
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	// dp[i] means max non-overlap interval exist before i
	dp := make([]int, len(intervals))
	dp[0] = 1
	var j int

	for i := 1; i < len(intervals); i++ {
		for j = i - 1; j >= 0; j-- {
			if !isOverlap(intervals[i], intervals[j]) {
				// since time is sorted by end time, when there's no overlap means
				// all previous are non-overlap
				break
			}
		}

		if j == -1 {
			dp[i] = dp[i-1]
		} else {
			dp[i] = max(dp[i-1], dp[j]+1)
		}
	}

	return len(intervals) - dp[len(dp)-1]
}

func eraseOverlapIntervals1(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	var maxNonOverlap int
	end := math.MinInt32

	for _, i := range intervals {
		if i[0] >= end {
			maxNonOverlap++
			end = i[1]
		}
	}

	return len(intervals) - maxNonOverlap
}

// i1 should be interval happens earlier or same than i2
func isOverlap(i1, i2 []int) bool {
	return !(i1[1] <= i2[0])
}

//  Notes
//  1. Thinking process

//      brute force (back tracking) O(n * 2^n)
//      just like permutation, each interval w/ 2 options: stay or
//     remove, after selection, check if remain intervals are non-overlap

//      ---                - 1
//        -----            - 2
//            -----        - 3

//      remove 2, makes rest intervals non-overlap
//      remove 1 & (2 or 3) makes rest intervals non-overlap
//      remove 2 is best choice since it's minimal removal

//      heuristics: remove intervals w/ most overlaps to others
//      - generate a table to store interval overlap count O(n^2)

//      - need to recalculate rest of non-overlaps O(n^2)
//          o need to main a list to decide which one goes next
//          o max heap, but how to update each interval overlap count
//            when one interval is removed?

//      - how to decide which one to remove if several intervals w/ same
//        amount of overlap count?
//          o just remove them one by one, because each overlap means
//            one interval should be removed, so it doesn't matter which
//            one goes first

//      sort intervals O(n log n)
//      create an array to store each interval overlap count O(n^2)
//      loop, scan array to find max overlap interval, remove it, update
//      its overlap intervals

//      -----
//        --------
//            ---------
//                 -------

//      -----
//         --------
//                --------
//                  ---------

//  2.  Fail on test case:

//      ---         0 - 2   - (1)
//       ---        1 - 3   - (2)
//        ---       2 - 4   - (3)
//         ---      3 - 5   - (4)
//          ---     4 - 6   - (5)

//      previously I assume when overlap count is same, removing which
//      interval first doesn't affect final result, but it's not true

//      this counter example demonstrate differences. all 2, 3, 4 w/
//      same overlap count, remove (2), (3), (4) makes (1) & (5) non-
//      overlap

//      remove (2), (4) makes (1), (3), (5) non-overlap

//      for same overlap count, need to try all possibility to find
//      optimal  solution. the reason greedy not work is similar to
//      problem 1547, when deciding interval to remove, not considering
//      relationships among intervals is the main reason

//      so this is a dp/dfs problem

//	3.	fails at test case [[0,2],[1,3],[1,3],[2,4],[3,5],[3,5],[4,6]]
//		overlaps: [2 3 3 4 3 3 2], since [2,4] overlaps most intervals, by
//		algorithm, it will be removed first

//		but there are some duplicate intervals, those duplicate intervals
//		should be removed first because there will be removed anyway

//		remove most overlap intervals is a heuristic, but I didn't aware that
//		if there are duplicates, this heuristic won't work

//	4.	no need to update overlap count every time, because overlaps are fixed,
//		so removing an interval, updates count and it's overlap interval

//	5.	remove minimum overlap => maximum non-overlap that can achieve, it's
//		similar to another problem: there are number of courses, I can only
//		attend one course at a time, what is the maximum courses I can attend?

//		maximum non-overlap can be found by choosing earliest dead line =>
//		O(n log n) to sort, and O(n) to iterate

//	6.	each interval is choose or not-choose, which is similar to knapsack
//		problem
