package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//There are n different online courses numbered from 1 to n. You are given an array courses where courses[i] = [durationi, lastDayi] indicate that the ith course should be taken continuously for durationi days and must be finished before or on lastDayi.
//
//You will start on the 1st day and you cannot take two or more courses simultaneously.
//
//Return the maximum number of courses that you can take.
//
//
//
//Example 1:
//
//Input: courses = [[100,200],[200,1300],[1000,1250],[2000,3200]]
//Output: 3
//Explanation:
//There are totally 4 courses, but you can take 3 courses at most:
//First, take the 1st course, it costs 100 days so you will finish it on the 100th day, and ready to take the next course on the 101st day.
//Second, take the 3rd course, it costs 1000 days so you will finish it on the 1100th day, and ready to take the next course on the 1101st day.
//Third, take the 2nd course, it costs 200 days so you will finish it on the 1300th day.
//The 4th course cannot be taken now, since you will finish it on the 3300th day, which exceeds the closed date.
//
//Example 2:
//
//Input: courses = [[1,2]]
//Output: 1
//
//Example 3:
//
//Input: courses = [[3,2],[4,3]]
//Output: 0
//
//
//
//Constraints:
//
//1 <= courses.length <= 104
//1 <= durationi, lastDayi <= 104

type MaxHeap []int // duration

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

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	maxHeap := &MaxHeap{}
	var attend int
	day := 1

	for _, course := range courses {
		if day+course[0]-1 <= course[1] {
			attend++
			heap.Push(maxHeap, course[0])
			day += course[0]
		} else {
			// try to replace previous taken longer course with shorter course
			if maxHeap.Len() > 0 && maxHeap.Peek() > course[0] {
				day += course[0] - maxHeap.Peek()
				heap.Pop(maxHeap)
				heap.Push(maxHeap, course[0])
			}
		}
	}

	return attend
}

// TLE
// tc: O(mn), m: course length, n: largest day course could reach
func scheduleCourse1(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	var longest int
	for _, course := range courses {
		longest = max(longest, course[0]+course[1])
	}

	size := len(courses)

	// dp[i][j]: maximum course can attend after day i for course j
	dp := make([][]int, longest+1)
	for i := range dp {
		dp[i] = make([]int, size)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	dfs1(courses, 1, 0, dp)

	var largest int
	for i := range dp[1] {
		largest = max(largest, dp[1][i])
	}

	return largest
}

func dfs1(courses [][]int, day, idx int, dp [][]int) int {
	size := len(courses)

	if idx == size {
		return 0
	}

	var attend int
	if dp[day][idx] == -1 {
		if courses[idx][0]+day-1 <= courses[idx][1] {
			attend = 1 + dfs1(courses, courses[idx][0]+day, idx+1, dp)
		}

		dp[day][idx] = max(attend, dfs1(courses, day, idx+1, dp))
	}

	return dp[day][idx]
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	at first I thought it's greedy problem, but after 2 WA, i found counter
//		examples:

//		- sort by duration
//		  e.g. [2, 5], [3, 4]

//		  if start from [2, 5], end at day 3 (1+2-1), then second interval cannot
//		  meet criteria

//		- sort by end day
//		  e.g. [4, 5], [2, 6]

//		  if start from [4, 5], end at day 4 (1+4-1), then second interval cannot
//		  meet criteria

//		  so, this is a dp problem

//	2.	recurring problem becomes: on day i, what's the maximum attendance count?

//	3.	for dp, can only update if all possibilities are calculated

//	4.	inspired from solution, even for dp, it still need to take course with
//		proper order.

//		e.g. courses (a, x) & (b, y)

//		- if a > x & b > y, no course can be take

//		- if a+b < min(x, y), only one course could be take, but there's some conditions
//		  to discuss, but in general, it's always beneficial to take course that
//		  ends earlier (this is the reason why using dp, still need to sort array)

//		- if a+b < x && a+b < y, two courses can be taken

//		also, course taken on different might also affect, so dp size is 2D,
//		represents course taken on specific day

//	5.	inspired from solution, sort courses by end day, and make sure each time
//		a course is taken is valid (day + duration < deadline)

//		now, for any coming course doesn't meet deadline, could try to remove
//		previous longer duration course and see if that helps to take this course.

//		the reason is because there are 2 factors that influence result, duration
//		and deadline, sort by deadline could exist a condition that longer duration
//		course is taken first, but it's also same course amount of taken course
//		with smaller duration

//		to do this, for use a heap to store previous course that is compared by
//		course duration

//	6.	inspired from https://leetcode.com/problems/course-schedule-iii/discuss/104847/Python-Straightforward-with-Explanation

//		since interval has no overlap, it's easier to just compare total duration with
//		end time
