package main

//Given a m x n matrix mat and an integer threshold. Return the maximum side-length of a square with a sum less than or equal to threshold or return 0 if there is no such square.
//
//
//
//Example 1:
//
//Input: mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
//Output: 2
//Explanation: The maximum side length of square with sum less than 4 is 2 as shown.
//
//Example 2:
//
//Input: mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]], threshold = 1
//Output: 0
//
//Example 3:
//
//Input: mat = [[1,1,1,1],[1,0,0,0],[1,0,0,0],[1,0,0,0]], threshold = 6
//Output: 3
//
//Example 4:
//
//Input: mat = [[18,70],[61,1],[25,85],[14,40],[11,96],[97,96],[63,45]], threshold = 40184
//Output: 2
//
//
//
//Constraints:
//
//    1 <= m, n <= 300
//    m == mat.length
//    n == mat[i].length
//    0 <= mat[i][j] <= 10000
//    0 <= threshold <= 10^5

func maxSideLength(mat [][]int, threshold int) int {
	y := len(mat)
	if y == 0 {
		return 0
	}

	x := len(mat[0])
	if x == 0 {
		return 0
	}

	accumulation := make([][]int, y)
	for i := 0; i < y; i++ {
		accumulation[i] = make([]int, x)
	}

	// an array that sums all previous numbers mat[0...i][0...j]
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				accumulation[i][j] = mat[i][j]
			} else if i == 0 {
				accumulation[i][j] = accumulation[i][j-1] + mat[i][j]
			} else if j == 0 {
				accumulation[i][j] = accumulation[i-1][j] + mat[i][j]
			} else {
				accumulation[i][j] = accumulation[i][j-1] + accumulation[i-1][j] - accumulation[i-1][j-1] + mat[i][j]
			}
		}
	}

	side := -1
	for i := 0; i <= min(x, y); i++ {
		found := false
		for j := y - i - 1; j >= 0; j-- {
			for k := x - i - 1; k >= 0; k-- {
				if sum(accumulation, k, j, k+i, j+i) <= threshold {
					found = true
					side = i
					goto loop
				}
			}
		}

	loop:
		if !found {
			return side + 1
		}
	}

	return side + 1
}

// needs to make sure not exceed boundaries
func sum(accumulation [][]int, x1, y1, x2, y2 int) int {
	if x1 == 0 && y1 == 0 {
		return accumulation[y2][x2]
	} else if x1 == 0 {
		return accumulation[y2][x2] - accumulation[y1-1][x1]
	} else if y1 == 0 {
		return accumulation[y2][x2] - accumulation[y1][x1-1]
	}

	return accumulation[y2][x2] - accumulation[y2][x1-1] - accumulation[y1-1][x2] + accumulation[y1-1][x1-1]
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

// problems
// 1. too slow, the complexity is O(m^3 n^3), use another accumulation array
//	  to reduce operation
// 2. optimize, not only 1 dimension accumulation, but also 2 dimension
// 	  accumulation
