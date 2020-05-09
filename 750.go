package main

// Given a grid where each entry is only 0 or 1, find the number of corner rectangles.
//
// A corner rectangle is 4 distinct 1s on the grid that form an axis-aligned rectangle. Note that only the corners need to have the value 1. Also, all four 1s used must be distinct.
//
//
//
// Example 1:
//
// Input: grid =
// [[1, 0, 0, 1, 0],
//  [0, 0, 1, 0, 1],
//  [0, 0, 0, 1, 0],
//  [1, 0, 1, 0, 1]]
// Output: 1
// Explanation: There is only one corner rectangle, with corners grid[1][2], grid[1][4], grid[3][2], grid[3][4].
//
//
//
// Example 2:
//
// Input: grid =
// [[1, 1, 1],
//  [1, 1, 1],
//  [1, 1, 1]]
// Output: 9
// Explanation: There are four 2x2 rectangles, four 2x3 and 3x2 rectangles, and one 3x3 rectangle.
//
//
//
// Example 3:
//
// Input: grid =
// [[1, 1, 1, 1]]
// Output: 0
// Explanation: Rectangles must have four distinct corners.
//
//
//
// Note:
//
//     The number of rows and columns of grid will each be in the range [1, 200].
//     Each grid[i][j] will be either 0 or 1.
//     The number of 1s in the grid will be at most 6000.

func countCornerRectangles(grid [][]int) int {
	y := len(grid)
	if y == 0 {
		return 0
	}
	x := len(grid[0])
	if x == 0 {
		return 0
	}

	var count int

	dp := make([][]int, x)
	for i := range dp {
		dp[i] = make([]int, x)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				for k := j + 1; k < x; k++ {
					if grid[i][k] == 1 {
						count += dp[j][k]
						dp[j][k]++
					}
				}
			}
		}
	}

	return count
}

func countCornerRectangles2(grid [][]int) int {
	y := len(grid)
	if y <= 1 {
		return 0
	}
	x := len(grid[0])
	if x <= 1 {
		return 0
	}

	ones := make([][]int, y)
	for i := range ones {
		ones[i] = make([]int, 0)
	}

	// find out all 1s
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				ones[i] = append(ones[i], j)
			}
		}
	}

	var count int

	// for each row, find intersect to other rows
	for i := range ones {
		for j := i + 1; j < y; j++ {
			common := intersect(ones[i], ones[j])
			if common >= 2 {
				// find combinations
				count += (common * (common - 1)) >> 1
			}
		}
	}

	return count
}

// arrays are already sorted
func intersect(row1, row2 []int) int {
	var count int
	for i, j := 0, 0; i < len(row1) && j < len(row2); {
		if row1[i] == row2[j] {
			count++
			i++
			j++
		} else if row1[i] < row2[j] {
			i++
		} else {
			j++
		}
	}

	return count
}

//	problems
//	1.	this complexity is O(n^4), runtime reaches limit

//	2.	combination count only relates to number of 1's on same index, so
//		goal is to find intersect count of 1 between two rows. To find out
//		intersect of 1, so I use another array to store 1's index on each
//		row.

//	3.	reference from https://leetcode.com/problems/number-of-corner-rectangles/discuss/110200/Summary-of-three-solutions-based-on-three-different-ideas

//		/2 can be done by >> 1

//	4.	reference from https://leetcode.com/problems/number-of-corner-rectangles/discuss/148531/Very-fast-Java-DP-solution

//		This is a really clever solution, whenever in a row two points are
//		1s, then it is possibly half part of rectangle. It can use dp to
//		store how many # of half rectangles, the length of dp is [n][n],
//		e.g. [1][4] means idx 1 & 4 in a row are 1s, [4][8] means idx
//		4 & 8 in a row are 1s. Whenever a new row comes, it check itself
//		to find 1s (current row half rectangle), then check dp for another
//		half rectangle to form one. And dp is total count of all idx i-j
//		1s, so add this number into total count.

//		When first time sees code, I think it should be dp[m][n], but
//		actually it's only relationships between same row, so first
//		length is n not m.

//		e.g. 1 0 1 0 <- dp[0][2] = 1
//		     1 0 1 1 <- dp[0][2] += 1, dp[0][3] = 1, dp[2][3] = 1
//			 1 1 1 0 <- dp[0][1] = 1, dp[0][2] += 1, dp[1][2] = 1

//		before processing second row, dp[0][2] = 1 means half rectangle
//		already exist, and second row has another dp[0][2] means count += 1

//		before processing third row, dp[0][2] = 2 means there are 2 half
//		rectangle exist, the third row exists another dp[0][2], so count
//		+= 2, which will be 3

//	5.	reference from https://leetcode.com/problems/number-of-corner-rectangles/discuss/188581/Google-follow-up-question.-A-general-case-solution.

//		google has follow-up questions, find general form of all numbers
//		not just 1.

//		I think it's still possible to use dp, but with map[int][][]int
