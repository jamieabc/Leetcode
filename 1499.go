package main

import "math"

// Given an array points containing the coordinates of points on a 2D plane, sorted by the x-values, where points[i] = [xi, yi] such that xi < xj for all 1 <= i < j <= points.length. You are also given an integer k.
//
// Find the maximum value of the equation yi + yj + |xi - xj| where |xi - xj| <= k and 1 <= i < j <= points.length. It is guaranteed that there exists at least one pair of points that satisfy the constraint |xi - xj| <= k.
//
//
//
// Example 1:
//
// Input: points = [[1,3],[2,0],[5,10],[6,-10]], k = 1
// Output: 4
// Explanation: The first two points satisfy the condition |xi - xj| <= 1 and if we calculate the equation we get 3 + 0 + |1 - 2| = 4. Third and fourth points also satisfy the condition and give a value of 10 + -10 + |5 - 6| = 1.
// No other pairs satisfy the condition, so we return the max of 4 and 1.
//
// Example 2:
//
// Input: points = [[0,0],[3,0],[9,2]], k = 3
// Output: 3
// Explanation: Only the first two points have an absolute difference of 3 or less in the x-values, and give the value of 0 + 0 + |0 - 3| = 3.
//
//
//
// Constraints:
//
// 2 <= points.length <= 10^5
// points[i].length == 2
// -10^8 <= points[i][0], points[i][1] <= 10^8
// 0 <= k <= 2 * 10^8
// points[i][0] < points[j][0] for all 1 <= i < j <= points.length
// xi form a strictly increasing sequence.

func findMaxValueOfEquation(points [][]int, k int) int {
	deque := []int{0}
	largest := math.MinInt32
	size := len(points)

	for i := 1; i < size; i++ {
		// remove out of range point
		for len(deque) > 0 && points[i][0]-points[deque[0]][0] > k {
			deque = deque[1:]
		}

		if len(deque) > 0 {
			largest = max(largest, points[i][0]+points[i][1]+points[deque[0]][1]-points[deque[0]][0])
		}

		cur := points[i][1] - points[i][0]
		for len(deque) > 0 {
			last := deque[len(deque)-1]

			if cur > points[last][1]-points[last][0] {
				deque = deque[:len(deque)-1]
			} else {
				break
			}
		}
		deque = append(deque, i)
	}

	return largest
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	didn't think of a solution, aware that x is sorted, so can use two pointers to
//		find valid points in range

//		but i was confused by the equation, yi+yj+|xi-xj|, tend to separate yi+yj and
//		|xi-xj|, then sort cannot meet 2 criteria and fail to find out a solution

//	2.	inspired from https://leetcode.com/problems/max-value-of-equation/discuss/713354/C%2B%2B-or-Deque-or-O(n)-or-Thorough-explanation-with-extensive-comments-or

//		since x is sorted and i < j, equation yi+yj+|xi-xj| can be simplified to
//		yi+yj+xj-xi (x is sorted, j > i, xj > xi), further re-arrange equation
//		(xj+yj)+(yi-xi)

//		this means for a later point j, find all previous valid point i that has
//		maximum value of yi-xi

//		first idea come out is max heap, it could solve the problem

//		but since need to find previous k points, after current point is processed,
//		clean-up out of range points, and get best choice so far

//		if keep a deque in decrease order by yi-xi, if points out of range, pop
//		from head of deque, and if cur point is better than previous, pop deque
//		from end, that could meet out expectation

//		overall, the key point to solve this problem is to aware that sort by two
//		conditions are not practical, need to find only one value to sort. Equation
//		can be further simplified and check only one value, and this value only depends
//		on it self (yi-xi), not correlate to other points
