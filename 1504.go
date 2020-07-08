package main

import "fmt"

// Given a rows * columns matrix mat of ones and zeros, return how many submatrices have all ones.
//
//
//
// Example 1:
//
// Input: mat = [[1,0,1],
//               [1,1,0],
//               [1,1,0]]
// Output: 13
// Explanation:
// There are 6 rectangles of side 1x1.
// There are 2 rectangles of side 1x2.
// There are 3 rectangles of side 2x1.
// There is 1 rectangle of side 2x2.
// There is 1 rectangle of side 3x1.
// Total number of rectangles = 6 + 2 + 3 + 1 + 1 = 13.
//
// Example 2:
//
// Input: mat = [[0,1,1,0],
//               [0,1,1,1],
//               [1,1,1,0]]
// Output: 24
// Explanation:
// There are 8 rectangles of side 1x1.
// There are 5 rectangles of side 1x2.
// There are 2 rectangles of side 1x3.
// There are 4 rectangles of side 2x1.
// There are 2 rectangles of side 2x2.
// There are 2 rectangles of side 3x1.
// There is 1 rectangle of side 3x2.
// Total number of rectangles = 8 + 5 + 2 + 4 + 2 + 2 + 1 = 24.
//
// Example 3:
//
// Input: mat = [[1,1,1,1,1,1]]
// Output: 21
//
// Example 4:
//
// Input: mat = [[1,0,1],[0,1,0],[1,0,1]]
// Output: 5
//
//
//
// Constraints:
//
//     1 <= rows <= 150
//     1 <= columns <= 150
//     0 <= mat[i][j] <= 1

func numSubmat(mat [][]int) int {
	if len(mat) == 0 {
		return 0
	}

	var matrices int
	consecutiveOnesSoFar := make([]int, len(mat[0]))

	for i := range mat {
		for j := range mat[0] {
			if mat[i][j] == 0 {
				consecutiveOnesSoFar[j] = 0
			} else {
				consecutiveOnesSoFar[j]++
			}
		}

		matrices += countRectangles(consecutiveOnesSoFar)
	}

	return matrices
}

func countRectangles(heights []int) int {
	size := len(heights)
	stack := make([]int, 0)
	rects := make([]int, size)
	var totalRect int

	for i := range heights {
		// keep stack in increasing order
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		// this one is different than 84, it can be seen as rectange one point start from i,
		// rectangle count is sum of rectangles at this height, plus previous height if there
		// exists lower height
		if len(stack) > 0 {
			rects[i] = heights[i]*(i-stack[len(stack)-1]) + rects[stack[len(stack)-1]]
		} else {
			rects[i] = heights[i] * (i + 1)
		}
		totalRect += rects[i]

		stack = append(stack, i)
	}

	return totalRect
}

func countRectangles(heights []int) int {
	size := len(heights)
	stack := make([]int, 0)
	var area int

	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			start := 0
			if len(stack) > 0 {
				start = stack[len(stack)-1] + 1
			}

			area += heights[popped] * (i - start)
		}

		stack = append(stack, i)
	}

	if len(stack) > 0 {
		area += heights[stack[0]] * size
	}

	for i := 1; i < len(stack); i++ {
		area += heights[stack[i]] * (size - stack[i-1] - 1)
	}

	return area
}

//	problems
//	1.	I don't even know how to think....try to start from scratch

//	2.	inspired from https://leetcode.com/problems/count-submatrices-with-all-ones/discuss/720227/Pre-computation-VIDEO-solution-O(m*n*n)

//		most basic concept: rectangle is formed by 2 points, first simplify
//		the problem: if a point if fixed, how many rectangles can be formed by
//		this point, given this point is left-bottom point.

//		e.g. 0 1 0 1	 if choose last row first point to check:
//    		 1 1 0 1  =>   & 1 &     & 1 1 &       &
//			 1 1 1 1     1   1   1 1   1 1   1 1 1   1 1 1 1

//		so rectangles count for last row first point are 6, 6 can be think as
//		follows: first column 1's = 2, second column 1's = 3, but previous is 2
//		so it can only be chosen to 2. third column 1's = 1, third column 1's =
//		3, but previous min is 1
//		6 = 2 * 2 + 1 * 2

//		for a given point, takes O(n^2) (expand from x and y direction) to find
//		rectangles based on that point above, there are n points on that row,
//		so for single row takes O(n^3), and n rows takes O(n^4)

//		this time complexity is pretty high, so improvement is needed. first
//		technique is to reduce finding height of each column. because height
//		can be formed gradually,

//		e.g. first row height:  0 1 0 1
//			 second row height: 1 2 0 2
//			 third row height:  2 3 1 3

//		use another array to store longest consecutive 1's from this row up,
//		tc becomes O(n^3)

//		for a given row, rectangle calculation is as this: for first column,
//		find its width that could form a rectangle (to where a column 1's
//		count > that column's count)

//		e.g for the row height of 2 3 1 3, for idx = 0, since height of index
//		1 is 3 < 2, index 0 height 2 can extend to index 1, but it cannot extend
//		to index 2 because height is 1. given the same rule,
//		6 = 2 * 2 + 1 * 2

//		but this calculation can also be optimized by another way of viewing
//		problem, for every index, number of rectangles formed is decided by
//		previous height, index 1 value 3 > index 0 value 2, so index 1 formed
//		rectangle is 2, index 2 value 1 < index 1 value 2 (already shrink),
//		so index 3 value is min(1, 2) = 1, that is
//		6 = 2 + min(2, 3) + min(2, 1) + min(1, 3) = 2 + 2 + 1 + 1

//		for a given point, takes O(1) to find rectangles, but row has n points,
//		so tc for one row is O(n), overall tc is O(n^3)

//		another technique to reduce complexity is using stack to calculate
//		possible rectangles. number of rectangles formed by one fix point is the
//		max rectangle area that point can have. with stack, tc can be further
//		reduced by O(n^2)

//	3.	inspired from https://leetcode.com/problems/count-submatrices-with-all-ones/discuss/720265/Java-Detailed-Explanation-From-O(MNM)-to-O(MN)-by-using-Stack

//		this one is different from 84, 84 counts maximum rectangle, so it only
//		needs to find width so far, but this problem wants to find all possible
//		rectangles. it's the max rectangle + previous possible rectangle, as
//		illustrated in author's figure
