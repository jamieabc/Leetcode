package main

import "container/heap"

// Given an n x n matrix where each of the rows and columns are sorted in ascending order, return the kth smallest element in the matrix.
//
// Note that it is the kth smallest element in the sorted order, not the kth distinct element.
//
//
//
// Example 1:
//
// Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
// Output: 13
// Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13
//
// Example 2:
//
// Input: matrix = [[-5]], k = 1
// Output: -5
//
//
//
// Constraints:
//
// n == matrix.length
// n == matrix[i].length
// 1 <= n <= 300
// -109 <= matrix[i][j] <= -109
// All the rows and columns of matrix are guaranteed to be sorted in non-degreasing order.
// 1 <= k <= n2

func kthSmallest(matrix [][]int, k int) int {
	w, h := len(matrix[0]), len(matrix)

	var ans int

	for low, high := matrix[0][0], matrix[h-1][w-1]; low <= high; {
		mid := low + (high-low)>>1
		tmp := count(matrix, mid)

		// becareful, since count smaller or equal, there might be
		// possibility that mid is too large, so record value until
		// not met
		if tmp >= k {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}

func count(matrix [][]int, target int) int {
	w, h := len(matrix[0]), len(matrix)

	var count int

	for i, j := h-1, 0; i >= 0 && j < w; {
		if matrix[i][j] <= target {
			count += i + 1
			j++
		} else if matrix[i][j] > target {
			i--
		}
	}

	return count
}

type MinHeap [][]int // row, column, number

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() []int        { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	(*h) = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	popped := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return popped
}

// tc: O(k log(row))
func kthSmallest1(matrix [][]int, k int) int {
	mh := &MinHeap{}
	heap.Init(mh)

	for i := range matrix {
		heap.Push(mh, []int{i, 0, matrix[i][0]})
	}

	w := len(matrix[0])
	for i := 0; i < k-1; i++ {
		popped := heap.Pop(mh).([]int)
		row, col := popped[0], popped[1]

		if col < w-1 {
			heap.Push(mh, []int{row, col + 1, matrix[row][col+1]})
		}
	}

	return mh.Peek()[2]
}

//	Notes
//	1.	inspired form solution, binary search can be used to solve this problem

//	2.	key point for each row/column in ascending order is that, for
//		bottom-left number, it's smallest in that row & largest in that column
//		top-right number, it's largest in that row & smallest in the column
//		use this condition to do some searching
