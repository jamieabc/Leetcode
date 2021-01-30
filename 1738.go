package main

// You are given a 2D matrix of size m x n, consisting of non-negative integers. You are also given an integer k.
//
// The value of coordinate (a, b) of the matrix is the XOR of all matrix[i][j] where 0 <= i <= a < m and 0 <= j <= b < n (0-indexed).
//
// Find the kth largest value (1-indexed) of all the coordinates of matrix.
//
//
//
// Example 1:
//
// Input: matrix = [[5,2],[1,6]], k = 1
// Output: 7
// Explanation: The value of coordinate (0,1) is 5 XOR 2 = 7, which is the largest value.
//
// Example 2:
//
// Input: matrix = [[5,2],[1,6]], k = 2
// Output: 5
// Explanation: The value of coordinate (0,0) is 5 = 5, which is the 2nd largest value.
//
// Example 3:
//
// Input: matrix = [[5,2],[1,6]], k = 3
// Output: 4
// Explanation: The value of coordinate (1,0) is 5 XOR 1 = 4, which is the 3rd largest value.
//
// Example 4:
//
// Input: matrix = [[5,2],[1,6]], k = 4
// Output: 0
// Explanation: The value of coordinate (1,1) is 5 XOR 2 XOR 1 XOR 6 = 0, which is the 4th largest value.
//
//
//
// Constraints:
//
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 1000
//     0 <= matrix[i][j] <= 106
//     1 <= k <= m * n

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() int          { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// tc: O(mn log(k))
func kthLargestValue(matrix [][]int, k int) int {
	w := len(matrix[0])
	h := &MinHeap{}
	heap.Init(h)

	// prefix[i]: xor from 0 ~ i
	prefix := make([]int, w)

	for i := range matrix {
		var tmp int

		for j := range matrix[0] {
			tmp ^= matrix[i][j]
			prefix[j] ^= tmp

			// easier to implement, but slower
			// heap.Push(h, prefix[j])
			// if h.Len() > k {
			// 	heap.Pop(h)
			// }

			if h.Len() == k {
				if h.Peek() >= prefix[j] {
					continue
				}

				heap.Pop(h)
			}

			heap.Push(h, prefix[j])
		}
	}

	return h.Peek()
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/find-kth-largest-xor-coordinate-value/discuss/1032143/Java-Detailed-Explanation-DP-with-Graph-Demo

//		easier code tofind kth largest if push to heap, pop if length > k
