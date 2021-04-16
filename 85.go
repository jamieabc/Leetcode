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

// tc: O(mn), scan row with O(n), calculate max area also takes O(n)
// overall O(m*2n) = O(mn)
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([]int, len(matrix[0]))
	stack := []int{0}

	var rect int
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == '0' {
				dp[j] = 0
			} else {
				dp[j] += 1
			}

			// calculate maximum rectangle area till now
			for len(stack) > 1 && dp[stack[len(stack)-1]] >= dp[j] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				end := stack[len(stack)-1]

				if dp[end] < dp[top] {
					end++
				}
				rect = max(rect, dp[top]*(j-end))
			}
			stack = append(stack, j)
		}

		// calculate remaining area
		for start, j := stack[len(stack)-1], len(stack)-1; j >= 1; j-- {
			end := stack[j-1]
			if dp[end] < dp[stack[j]] {
				end++
			}

			rect = max(rect, dp[stack[j]]*(start-end+1))
		}
		stack = stack[:1]
	}

	return rect
}

// tc: O(mn)
func maximalRectangle2(matrix [][]byte) int {
	h := len(matrix)

	if h == 0 {
		return 0
	}

	w := len(matrix[0])

	// dp[i]: vertical consecutive 1s
	dp := make([]int, w)
	var largest int

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '0' {
				dp[j] = 0
			} else {
				dp[j]++
			}
		}

		// count maximum rectangle so far
		largest = max(largest, maxArea(dp))
	}

	return largest
}

func maxArea(ones []int) int {
	var area int
	stack := []int{-1}

	for i := range ones {
		for len(stack) > 1 && ones[stack[len(stack)-1]] > ones[i] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			area = max(area, ones[prev]*(i-stack[len(stack)-1]-1))
		}

		stack = append(stack, i)
	}

	// process remaining data
	for len(stack) > 1 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		area = max(area, ones[cur]*(len(ones)-stack[len(stack)-1]-1))
	}

	return area
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	length := len(matrix[0])
	h, l, r := make([]int, length), make([]int, length), make([]int, length)
	for i := range r {
		r[i] = length
	}

	var maxArea int
	for i := range matrix {
		cl, cr := 0, length

		for j := range h {
			if matrix[i][j] == '1' {
				h[j]++
			} else {
				h[j] = 0
			}
		}

		for j := range l {
			if matrix[i][j] == '1' {
				l[j] = max(l[j], cl)
			} else {
				l[j] = 0
				cl = j + 1
			}
		}

		for j := length - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				r[j] = min(r[j], cr)
			} else {
				r[j] = length
				cr = j
			}
		}

		for j := range h {
			maxArea = max(maxArea, h[j]*(r[j]-l[j]))
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

//	Notes
//	1.	k should be controller by i not xMax
//	2.	when calculating xMax, it is absolute index, which means it's x+i
//	3.	optimize, use dp, O(m n^2)
//	4.	for the first row or column, I just pass it, the problem is that
//		if only one row/column, then max is updated which returns error value
//	5.	when encounter 0 and then another positive width, the start needs
//		to be updated to the row that has positive number
//	6.	optimize, no need to initialize dp first, it can be calculated the
//		same time when traversing dp
//		if matrix[y][x] == '1' then dp[y][x] = dp[y][x-1] + 1
//	7.	optimize, when width is 0, no need to calculate any upper elements
//	8.	optimize, since width = 0 is no longer considered, start can be
//		removed
//	9.	optimize, complexity of O(m^2 n) is still too large, the problem
//		comes form the fact that re-calculating rectangle width.
//		From solution article, a clever method has been proposed.
//		every rectangle is composed of height, left, right.
//
//		height[j] = height[j] + 1 if matrix[y][x] == 1
//
//		left boundary of a rectangle can be considered as maximum index of
//		right most zero from current point left part.
//
//		left[j] = max(left[j], current_left)
//		current_left starts from 9 and is updated whenever encounter 0
//
//		right boundary of a rectangle can be considered as minimum index of
//		right most zero from current point right part.
//
//		be aware that right boundary must start from right most element
//		(length-1), so that right most zero is considered.
//
//		right[j] = min(right[j], current_right)
//		current_right starts from length and is updated whenever
//		encounter 0

//	10.	this problem relates to lc 84, calculate maximum rectangle
//		tc: O(mn), because worst case in stack-based traverse is (2n), so for
//		each line it's still n

//	11.	solution for finding h, l, r is to use smart way to store iterated
//		data
