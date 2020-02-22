package main

//Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.
//
//Example:
//
//Input:
//[
//  ["1","0","1","0","0"],
//  ["1","0","1","1","1"],
//  ["1","1","1","1","1"],
//  ["1","0","0","1","0"]
//]
//Output: 6

func maximalRectangle(matrix [][]byte) int {
	yLength := len(matrix)
	if yLength == 0 {
		return 0
	}
	xLength := len(matrix[0])

	height := make([]int, xLength)
	left := make([]int, xLength)
	right := make([]int, xLength)

	// right slice is initialized with length
	for i := range right {
		right[i] = xLength
	}

	maxArea := 0
	for j := range matrix {
		curLeft := 0
		curRight := xLength
		for i := range matrix[0] {
			// update height
			if matrix[j][i] == '0' {
				height[i] = 0
			} else {
				if j == 0 {
					height[i] = 1
				} else {
					height[i]++
				}
			}

			// update left
			if matrix[j][i] == '0' {
				curLeft = i + 1
				left[i] = 0
			} else {
				left[i] = max(left[i], curLeft)
			}

			// update right
			if matrix[j][xLength-1-i] == '0' {
				right[xLength-1-i] = xLength
				curRight = xLength - 1 - i
			} else {
				right[xLength-1-i] = min(right[xLength-1-i], curRight)
			}
		}

		// udpate max area
		for i := range height {
			area := height[i] * (right[i] - left[i])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
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

// problems
// 1. k should be controller by i not xMax
// 2. when calculating xMax, it is absolute index, which means it's x+i
// 3. optimize, use dp, O(m n^2)
// 4. for the first row or column, I just pass it, the problem is that
// 	  if only one row/column, then max is updated which returns error value
// 5. when encounter 0 and then another positive width, the start needs
//    to be updated to the row that has positive number
// 6. optimize, no need to initialize dp first, it can be calculated the
// 	  same time when traversing dp
//	  if matrix[y][x] == '1' then dp[y][x] = dp[y][x-1] + 1
// 7. optimize, when width is 0, no need to calculate any upper elements
// 8. optimize, since width = 0 is no longer considered, start can be
// 	  removed
// 9. optimize, complexity of O(m^2 n) is still too large, the problem
//    comes form the fact that re-calculating rectangle width.
//    From solution article, a clever method has been proposed.
// 	  every rectangle is composed of height, left, right.
//
//	  height[j] = height[j] + 1 if matrix[y][x] == 1
//
//	  left boundary of a rectangle can be considered as maximum index of
// 	  right most zero from current point left part.
//
//	  left[j] = max(left[j], current_left)
//	  current_left starts from 9 and is updated whenever encounter 0
//
//	  right boundary of a rectangle can be considered as minimum index of
//	  right most zero from current point right part.
//
//    be aware that right boundary must start from right most element
//	  (length-1), so that right most zero is considered.
//
//	  right[j] = min(right[j], current_right)
//	  current_right starts from length and is updated whenever
//	  encounter 0
