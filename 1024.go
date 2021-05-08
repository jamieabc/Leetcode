package main

import "sort"

// You are given a series of video clips from a sporting event that lasted T seconds.  These video clips can be overlapping with each other and have varied lengths.
//
// Each video clip clips[i] is an interval: it starts at time clips[i][0] and ends at time clips[i][1].  We can cut these clips into segments freely: for example, a clip [0, 7] can be cut into segments [0, 1] + [1, 3] + [3, 7].
//
// Return the minimum number of clips needed so that we can cut the clips into segments that cover the entire sporting event ([0, T]).  If the task is impossible, return -1.
//
//
//
// Example 1:
//
// Input: clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
// Output: 3
// Explanation:
// We take the clips [0,2], [8,10], [1,9]; a total of 3 clips.
// Then, we can reconstruct the sporting event as follows:
// We cut [1,9] into segments [1,2] + [2,8] + [8,9].
// Now we have segments [0,2] + [2,8] + [8,10] which cover the sporting event [0, 10].
//
// Example 2:
//
// Input: clips = [[0,1],[1,2]], T = 5
// Output: -1
// Explanation:
// We can't cover [0,5] with only [0,1] and [1,2].
//
// Example 3:
//
// Input: clips = [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]], T = 9
// Output: 3
// Explanation:
// We can take clips [0,4], [4,7], and [6,9].
//
// Example 4:
//
// Input: clips = [[0,4],[2,8]], T = 5
// Output: 2
// Explanation:
// Notice you can have extra video after the event ends.
//
//
//
// Constraints:
//
// 1 <= clips.length <= 100
// 0 <= clips[i][0] <= clips[i][1] <= 100
// 0 <= T <= 100

// tc: O(T)
func videoStitching(clips [][]int, T int) int {
	if T == 0 {
		return 0
	}
	// dp[i]: furthest can go
	dp := make([]int, 101) // T max to 100

	for _, clip := range clips {
		dp[clip[0]] = max(dp[clip[0]], clip[1])
	}

	size := len(dp)
	farthest := dp[0]
	jump := 1
	var next int

	for i := 1; i < size && farthest < T; jump++ {
		for ; i < size && i <= farthest; i++ {
			next = max(next, dp[i])
		}

		if next == farthest {
			return -1
		}

		farthest = next
	}

	if farthest < T {
		return -1
	}

	return jump
}

// tc: O(n log(n))
func videoStitching2(clips [][]int, T int) int {
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})

	// end: previous round best selection
	// best: next round best selection
	var best, end, count int
	size := len(clips)

	for i := 0; i < size && end < T; count++ {
		for ; i < size && clips[i][0] <= end; i++ {
			best = max(best, clips[i][1])
		}

		// cannot find continous one
		if best == end {
			return -1
		}
		end = best
	}

	// overall cannot meet criteria
	if end < T {
		return -1
	}

	return count
}

// tc: O(n log(n))
func videoStitching1(clips [][]int, T int) int {
	if T == 0 {
		return 0
	}

	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] != clips[j][0] {
			return clips[i][0] < clips[j][0]
		}
		return clips[i][1] < clips[j][1]
	})

	if clips[0][0] > 0 {
		return -1
	}

	var i, j, tmp int
	size := len(clips)
	count := 1

	// make sure starting point is best choice
	for ; i < size && clips[i][0] == 0; i++ {
	}
	i--

	for j = i + 1; j < size && clips[i][1] < T; count++ {
		// find the interval non-overlap
		for tmp = j; j < size; j++ {
			if !overlap(clips[i], clips[j]) {
				break
			} else if clips[j][1] > clips[tmp][1] {
				// for intervals already overlap, select the best one which
				// extends largest covered range
				tmp = j
			}
		}

		// next interval is non-overlap, cannot be continuous
		if j == i+1 && j < size {
			return -1
		}

		i = tmp
	}

	if clips[i][1] < T {
		return -1
	}
	return count
}

func overlap(i, j []int) bool {
	return i[1] >= j[0]
}

//	Notes
//	1.	it takes me 2 hours to figure out this is 2-ptr problem
//		at first I try to sort intervals by start time, if start time are same,
//		sort end time desc

//		e.g.	[1, 5], [1, 9] => [1, 9], [1,5]

//		i thought this can be greedily solved by always selection first
//		interval extends existing covered range

//		there's a problem that greedy doesn't consider previous previous interval

//		e.g. [1, 4], [2, 6], [4, 9]
//		first iteration select [1, 4]
//		second iteration select [2, 6], since 6 > 4
//		third iteration, select [4, 9] since 9 > 6

//		but the optimal solution is to un-select [2, 6] and select [1, 4]

//	2.	then i try to use max-heap to solve, max-heap always returns longest
//		interval, but it doesn't help on example in note 1

//		then i thought about dp, but since dp always scan previous selection,
//		it seems not necessary

//		finally i think of 2-pointer, each selection preserves state of selection,
//		if intervals sort start asc, same start sort by end asc, then each
//		iteration can find best choice by expanding next selection, until
//		next selection has no overlap with existing selection

//		e.g. [1, 4], [2, 6], [4, 9], [5, 10]

//		first select [1, 4]
//		second go to [5, 10] because [2, 6] & [4, 9] are overlap by [1, 4], best
//			selection is [4, 9] which makes covered range continous, an also
//			enlarge most

//	3.	there's another problem to first select number 1 interval,

//		e.g. [0, 2], [0, 11]

//		it's better to select [o, 11] since it covers more range

//	4.	becareful about boundary condition, if T == 0, no need to select anything

//	5.	for selection, need to choose the one that extends most, it's not
//		necessarily the last interval

//	6.	last year, i wrote O(n^2) solution, which is really bad...

//	7.	inspired from https://leetcode.com/problems/video-stitching/discuss/269988/C%2B%2BJava-6-lines-O(n-log-n)

//		it's much prettier

//		there are two variables store states:
//		- end: current iteration furthest interval to go
//		- st: previous iteration furthest interval to go

//		each iteration, make sure interval is able to expand, otherwise,
//		return -1

//	8.	this problem can be solved by 2-pointer is because previous state
//		can be used by next state, and only cares about maximum value, no
//		other constraints

//	9.	inspired from https://leetcode.com/problems/video-stitching/discuss/484877/Python-(24-ms-beats-99)-Jump-Game-II-O(N)-time-O(1)-memory

//		after convert to array that each position can jump to where, this problem
//		becomes jump game
