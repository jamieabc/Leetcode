package main

import (
	"container/heap"
	"math"
)

// We are given a list schedule of employees, which represents the working time for each employee.
//
// Each employee has a list of non-overlapping Intervals, and these intervals are in sorted order.
//
// Return the list of finite intervals representing common, positive-length free time for all employees, also in sorted order.
//
// (Even though we are representing Intervals in the form [x, y], the objects inside are Intervals, not lists or arrays. For example, schedule[0][0].start = 1, schedule[0][0].end = 2, and schedule[0][0][0] is not defined).  Also, we wouldn't include intervals like [5, 5] in our answer, as they have zero length.
//
//
//
// Example 1:
//
// Input: schedule = [[[1,2],[5,6]],[[1,3]],[[4,10]]]
// Output: [[3,4]]
// Explanation: There are a total of three employees, and all common
// free time intervals would be [-inf, 1], [3, 4], [10, inf].
// We discard any intervals that contain inf as they aren't finite.
// Example 2:
//
// Input: schedule = [[[1,3],[6,7]],[[2,4]],[[2,5],[9,12]]]
// Output: [[5,6],[7,9]]
//
//
// Constraints:
//
// 1 <= schedule.length , schedule[i].length <= 50
// 0 <= schedule[i].start < schedule[i].end <= 10^8

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	sorted := mergeSort(schedule, 0, len(schedule)-1)

	cur := sorted[0]
	freeTime := make([]*Interval, 0)

	for i := 1; i < len(sorted); i++ {
		if cur.End < sorted[i].Start {
			// non-overlap
			freeTime = append(freeTime, &Interval{
				Start: cur.End,
				End:   sorted[i].Start,
			})

			cur = sorted[i]
		} else {
			cur.End = max(cur.End, sorted[i].End)
		}
	}

	return freeTime
}

func mergeSort(arrs [][]*Interval, start, end int) []*Interval {
	if start == end {
		return arrs[start]
	}

	mid := start + (end-start)/2
	left := mergeSort(arrs, start, mid)
	right := mergeSort(arrs, mid+1, end)
	return merge(left, right)
}

func merge(arr1, arr2 []*Interval) []*Interval {
	sorted := make([]*Interval, 0)

	for i, j := 0, 0; i < len(arr1) || j < len(arr2); {
		if i != len(arr1) && (j == len(arr2) || arr1[i].Start < arr2[j].Start || (arr1[i].Start == arr2[j].Start && arr1[i].End <= arr2[j].End)) {
			sorted = append(sorted, arr1[i])
			i++
		} else {
			sorted = append(sorted, arr2[j])
			j++
		}
	}

	return sorted
}

type Interval struct {
	Start int
	End   int
}

type Works []Interval

func (this Works) Len() int { return len(this) }

func (this Works) Less(i, j int) bool {
	if this[i].Start < this[j].Start {
		return true
	} else if this[i].Start > this[j].Start {
		return false
	} else {
		return this[i].End < this[j].End
	}
}

func (this Works) Swap(i, j int)  { this[i], this[j] = this[j], this[i] }
func (this Works) Peek() Interval { return this[0] }

func (this *Works) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*this = append(*this, x.(Interval))
}

func (this *Works) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[:n-1]

	return x
}

func employeeFreeTime2(schedule [][]*Interval) []*Interval {
	w := &Works{}
	heap.Init(w)

	// this part is not necessary, I can only push []{employee's index, this
	// employee's interval index}
	for i := range schedule {
		for j := range schedule[i] {
			heap.Push(w, *schedule[i][j])
		}
	}

	freeTime := make([]*Interval, 0)

	cur := heap.Pop(w).(Interval)

	for w.Len() > 0 {
		top := w.Peek()
		if top.Start > cur.End {
			freeTime = append(freeTime, &Interval{
				Start: cur.End,
				End:   top.Start,
			})
			cur = w.Pop().(Interval)
		} else {
			cur.End = max(cur.End, top.End)
			w.Pop()
		}
	}

	return freeTime
}

func employeeFreeTime1(schedule [][]*Interval) []*Interval {
	if len(schedule) == 0 {
		return []*Interval{}
	}

	// get minimum start time & maximum time among employees
	minStart, maxEnd := math.MaxInt32, math.MinInt32
	if len(schedule) > 0 {
		for i := range schedule {
			minStart = min(minStart, schedule[i][0].Start)
			maxEnd = max(maxEnd, schedule[i][len(schedule[i])-1].End)
		}
	}

	// for every employee, find free time
	// find intersection of every employee free time
	freeTime := make([]*Interval, 0)
	var initialized bool

	for _, employee := range schedule {
		individualFree := make([]*Interval, 0)

		// minStart time
		if len(employee) > 0 && employee[0].Start > minStart {
			individualFree = append(individualFree, &Interval{minStart, employee[0].Start})
		}

		// get individual free time
		end := employee[0].End
		for i := 1; i < len(employee); i++ {
			minStart, maxEnd = min(minStart, employee[i].Start), max(maxEnd, employee[i].End)

			if employee[i].Start < end {
				// overlap
				end = max(end, employee[i].End)
			} else {
				individualFree = append(individualFree, &Interval{end, employee[i].Start})
				end = employee[i].End
			}
		}

		// maxEnd time
		if len(employee) > 0 && employee[len(employee)-1].End < maxEnd {
			individualFree = append(individualFree, &Interval{employee[len(employee)-1].End, maxEnd})
		}

		// find intersection of free time
		if !initialized {
			freeTime = individualFree
			initialized = true
			continue
		}

		tmp := make([]*Interval, 0)

		for i, j := 0, 0; i < len(freeTime) && j < len(individualFree); {
			if freeTime[i].End < individualFree[j].Start {
				i++
			} else if individualFree[j].End < freeTime[i].Start {
				j++
			} else {
				// overlap
				s := max(freeTime[i].Start, individualFree[j].Start)
				e := min(freeTime[i].End, individualFree[j].End)

				if s < e {
					tmp = append(tmp, &Interval{max(freeTime[i].Start, individualFree[j].Start), min(freeTime[i].End, individualFree[j].End)})
				}

				if freeTime[i].End < individualFree[j].End {
					i++
				} else {
					j++
				}
			}
		}

		freeTime = tmp
	}

	return freeTime
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
//	1.	when choosing intersection of free time, need to make sure start < end

//	2.	inspired from https://leetcode.com/problems/employee-free-time/discuss/113134/Simple-Java-Sort-Solution-Using-(Priority-Queue)-or-Just-ArrayList

//		naive solution is to find free time for every employee, then find
//		intersection for every employee free time., tc: O(nm), n: # of employee,
//		m: free time

//		now, free time means no one works, so if sort overall work time, and
//		iterate through it one by one, then can find free time. It acts as
//		big sort for every interval

//		no need to push all intervals into heap, can just push []{employee's
//		index, this employee's interval index}, this tc: O(N log K), N: # of
//		total intervals, K: number of employee

//	3.	inspired from https://leetcode.com/problems/employee-free-time/discuss/113122/Merge-Sort-O(nlgK)-(Java)

//		every employee is already sorted, so it can be seen as merge sort, because
//		during merge, two sub-arrays are all sorted as thie problem
