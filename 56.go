package main

import "sort"

//Given a collection of intervals, merge all overlapping intervals.
//
//Example 1:
//
//Input: [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
//
//Example 2:
//
//Input: [[1,4],[4,5]]
//Output: [[1,5]]
//Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//
//NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

type nums struct {
	data [][]int
}

func (n nums) Len() int {
	return len(n.data)
}
func (n nums) Swap(i, j int) {
	n.data[i], n.data[j] = n.data[j], n.data[i]
}

func (n nums) Less(i, j int) bool {
	if n.data[i][0] <= n.data[j][0] {
		return true
	}
	return false
}

func merge(intervals [][]int) [][]int {
	ints := make([][]int, len(intervals))
	for i, n := range intervals {
		ints[i] = []int{n[0], n[1]}
	}

	n := nums{
		data: ints,
	}
	sort.Sort(n)
	return combine(n.data)
}

func combine(data [][]int) [][]int {
	result := make([][]int, 0)
	i := -1

	for _, n := range data {
		if i < 0 {
			result = append(result, n)
			i++
		} else {
			if result[i][0] <= n[0] && result[i][1] >= n[0] {
				// 1, 3 && 2, 4 => 1, 4
				if result[i][1] <= n[1] {
					result[i][1] = n[1]
				}
				// 1, 4 && 2, 3 => 1, 3
			} else {
				result = append(result, n)
				i++
			}
		}
	}
	return result
}

// problems
// 1. 2 situations of overlap:
//    1-3, 2-4 => 1-4
//    1-4, 2-3 => 1-4
