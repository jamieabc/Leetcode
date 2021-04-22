package main

// There is a brick wall in front of you. The wall is rectangular and has several rows of bricks. The bricks have the same height but different width. You want to draw a vertical line from the top to the bottom and cross the least bricks.
//
// The brick wall is represented by a list of rows. Each row is a list of integers representing the width of each brick in this row from left to right.
//
// If your line go through the edge of a brick, then the brick is not considered as crossed. You need to find out how to draw the line to cross the least bricks and return the number of crossed bricks.
//
// You cannot draw a line just along one of the two vertical edges of the wall, in which case the line will obviously cross no bricks.
//
//
//
// Example:
//
// Input: [[1,2,2,1],
// [3,1,2],
// [1,3,2],
// [2,4],
// [3,1,2],
// [1,3,1,1]]
//
// Output: 2
//
// Explanation:
//
//
//
// Note:
//
// The width sum of bricks in different rows are the same and won't exceed INT_MAX.
// The number of bricks in each row is in range [1,10,000]. The height of wall is in range [1,10,000]. Total number of bricks of the wall won't exceed 20,000.

// tc: O(n), n: total bricks
func leastBricks(wall [][]int) int {
	h := len(wall)

	table := make(map[int]int)
	var length, largest int

	for i := range wall {
		length = 0

		for j := 0; j < len(wall[i])-1; j++ {
			length += wall[i][j]
			table[length]++

			largest = max(largest, table[length])
		}
	}

	return h - largest
}

// tc: O(mn), width & height of wall
func leastBricks1(wall [][]int) int {
	h := len(wall)
	least := h

	var width int
	for i := range wall[0] {
		width += wall[0][i]
	}

	index := make([]int, h)
	current := make([]int, h)
	for i := range index {
		current[i] = wall[i][0]
	}

	var count int
	for i := 1; i < width; i++ {
		count = 0

		for j := range wall {
			if current[j] == i {
				index[j]++
				current[j] += wall[j][index[j]]
			} else {
				count++
			}
		}

		least = min(least, count)
	}
	return least
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	the point is not to traverse all numbers, because it only breaks at
//		certain point, only focus on that point is enought

//	2.	inspired from solution, the hasmap way tc should be slightly modify,
//		traverse all bricks

//	3.	inspired from https://leetcode.com/problems/brick-wall/discuss/888577/IntuitionC%2B%2BWith-PicturesHashMapDetailed-ExplanationCommentsSolutionCode

//		author provides serires of good graph

//	4.	inspired from https://leetcode.com/problems/brick-wall/discuss/137777/Java-Map-solution

//		the solutino can be further reduced, no need to traver hashmap twice
