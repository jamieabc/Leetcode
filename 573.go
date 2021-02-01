package main

import "math"

// There's a tree, a squirrel, and several nuts. Positions are represented by the cells in a 2D grid. Your goal is to find the minimal distance for the squirrel to collect all the nuts and put them under the tree one by one. The squirrel can only take at most one nut at one time and can move in four directions - up, down, left and right, to the adjacent cell. The distance is represented by the number of moves.
//
// Example 1:
//
// Input:
// Height : 5
// Width : 7
// Tree position : [2,2]
// Squirrel : [4,4]
// Nuts : [[3,0], [2,5]]
// Output: 12
// Explanation:
// ​​​​​
//
// Note:
//
//     All given positions won't overlap.
//     The squirrel can take at most one nut at one time.
//     The given positions of nuts have no order.
//     Height and width are positive integers. 3 <= height * width <= 10,000.
//     The given positions contain at least one nut, only one tree and one squirrel.

func minDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	var dist int
	saved := math.MinInt32

	for i := range nuts {
		tmp := hamming(tree, nuts[i])
		dist += tmp << 1
		saved = max(saved, tmp-hamming(squirrel, nuts[i]))
	}

	return dist - saved
}

func hamming(src, dst []int) int {
	return abs(src[0]-dst[0]) + abs(src[1]-dst[1])
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	if there's no squirrel, total distance = 2 * dist(tree ~ nuts[i])
//		distance can be saved because the first nut to tree, no need to go
//		back

//		so saved = dist(tree ~ nuts[i]) - dist(squirrel ~ nuts[i])
//		to get smallest travel distance, then saved should be maximized

//	2.	saved distance could be negative number, so start of the number should
//		math.MinInt32, not 0, otherwise, negative number will not be count
