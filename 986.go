package main

// Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.
//
// Return the intersection of these two interval lists.
//
// (Formally, a closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.  The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented as a closed interval.  For example, the intersection of [1, 3] and [2, 4] is [2, 3].)
//
//
//
// Example 1:
//
// Input: A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
// Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
// Reminder: The inputs and the desired output are lists of Interval objects, and not arrays or lists.
//
//
//
// Note:
//
// 0 <= A.length < 1000
// 0 <= B.length < 1000
// 0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9
//
// NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := make([][]int, 0)
	for i, j := 0, 0; i < len(A) && j < len(B); {
		if overlap(A[i], B[j]) {
			result = append(result, intersection(A[i], B[j]))
		}

		// for the earlier interval, jump to next one
		if A[i][1] < B[j][1] {
			i++
		} else if A[i][1] > B[j][1] {
			j++
		} else {
			i++
			j++
		}
	}

	return result
}

func overlap(i1, i2 []int) bool {
	if i1[1] < i2[0] || i2[1] < i1[0] {
		return false
	}

	return true
}

func intersection(i1, i2 []int) []int {
	// if overlap, return max(start) ~ min(end)
	return []int{max(i1[0], i2[0]), min(i1[1], i2[1])}
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

func intervalIntersection1(A [][]int, B [][]int) [][]int {
	result := make([][]int, 0)

	if len(A) == 0 || len(B) == 0 {
		return result
	}

	for i, j := 0, 0; i < len(A) && j < len(B); {
		if A[i][1] < B[j][0] {
			i++
		} else if A[i][0] > B[j][1] {
			j++
		} else {
			result = append(result, []int{max(A[i][0], B[j][0]), min(A[i][1], B[j][1])})
			if A[i][1] < B[j][1] {
				i++
			} else if A[i][1] > B[j][1] {
				j++
			} else {
				i++
				j++
			}
		}
	}

	return result
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
//	1.	inspired by https://leetcode.com/problems/interval-list-intersections/discuss/231108/C%2B%2B-O(n)-%22merge-sort%22

//		when checking intersect, some condition can already be used to
//		decide i or j should increment

//		also, then checking of start point can use max to decide

//	2.	inspired from https://leetcode.com/problems/interval-list-intersections/discuss/231122/Java-two-pointers-O(m-%2B-n)

//		length invalid doesn't need to check

//	3.	reference from sample code, when both ends are same, increase both
