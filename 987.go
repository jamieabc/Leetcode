package main

import (
	"fmt"
	"sort"
)

// Given the root of a binary tree, calculate the vertical order traversal of the binary tree.
//
// For each node at position (x, y), its left and right children will be at positions (x - 1, y - 1) and (x + 1, y - 1) respectively.
//
// The vertical order traversal of a binary tree is a list of non-empty reports for each unique x-coordinate from left to right. Each report is a list of all nodes at a given x-coordinate. The report should be primarily sorted by y-coordinate from highest y-coordinate to lowest. If any two nodes have the same y-coordinate in the report, the node with the smaller value should appear earlier.
//
// Return the vertical order traversal of the binary tree.
//
//
//
// Example 1:
//
// Input: root = [3,9,20,null,null,15,7]
// Output: [[9],[3,15],[20],[7]]
// Explanation: Without loss of generality, we can assume the root node is at position (0, 0):
// The node with value 9 occurs at position (-1, -1).
// The nodes with values 3 and 15 occur at positions (0, 0) and (0, -2).
// The node with value 20 occurs at position (1, -1).
// The node with value 7 occurs at position (2, -2).
//
// Example 2:
//
// Input: root = [1,2,3,4,5,6,7]
// Output: [[4],[2],[1,5,6],[3],[7]]
// Explanation: The node with value 5 and the node with value 6 have the same position according to the given scheme.
// However, in the report [1,5,6], the node with value 5 comes first since 5 is smaller than 6.
//
//
//
// Constraints:
//
//     The number of nodes in the tree is in the range [1, 1000].
//     0 <= Node.val <= 1000

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func verticalTraversal(root *TreeNode) [][]int {
	// table[i]: nodes when x-axis = i
	table := make(map[int][][2]int)
	minX, maxX := 1000, -1000

	dfs(root, 0, 0, &minX, &maxX, table)

	ans := make([][]int, 0)

	for i := minX; i <= maxX; i++ {
		if arr, ok := table[i]; ok {
			sort.Slice(arr, func(j, k int) bool {

				// higher y, or same y but smaller value
				return arr[j][0] > arr[k][0] ||
					(arr[j][0] == arr[k][0] && arr[j][1] < arr[k][1])
			})

			tmp := make([]int, len(arr))

			for j := range arr {
				tmp[j] = arr[j][1]
			}

			ans = append(ans, tmp)
		}
	}

	return ans
}

func dfs(node *TreeNode, x, y int, minX, maxX *int, table map[int][][2]int) {
	if node == nil {
		return
	}

	*minX, *maxX = min(*minX, x), max(*maxX, x)
	table[x] = append(table[x], [2]int{y, node.Val})

	dfs(node.Left, x-1, y-1, minX, maxX, table)
	dfs(node.Right, x+1, y-1, minX, maxX, table)
}

// tc: worst case O(n log(n)), beacause each sort contains number of node for
// same x & same y
func verticalTraversal2(root *TreeNode) [][]int {
	table := make()
	table := make([][][]int, 2000)
	for i := range table {
		table[i] = make([][]int, 1000)
	}

	var minX, maxX int

	inOrder(root, 0, 0, &minX, &maxX, table)

	ans := make([][]int, 0)

	// order matters, want x from smallest to largest, y from highest to lowest
	// in this table, vertical is x, horizontal is y
	// table[0][0] means x is -999, y is 0
	// table[999][10] means x is 0, y is -10
	for i := minX; i <= maxX; i++ {
		tmp := make([]int, 0)

		for j := range table[0] {
			if len(table[i][j]) > 0 {
				sort.Ints(table[i][j])

				tmp = append(tmp, table[i][j]...)
			}
		}

		// need to check if any nodes resides in this specific x
		if len(tmp) > 0 {
			ans = append(ans, tmp)
		}
	}

	return ans
}

func inOrder(node *TreeNode, x, y int, minX, maxX *int, table [][][]int) {
	if node == nil {
		return
	}

	// convert x ranges from -999 ~ 999 to 0 ~ 1999, so offset if 999
	*minX, *maxX = min(*minX, 999+x), max(*maxX, 999+x)
	table[999+x][-y] = append(table[999+x][-y], node.Val)

	inOrder(node.Left, x-1, y-1, minX, maxX, table)
	inOrder(node.Right, x+1, y-1, minX, maxX, table)
}

type MinHeap [][]int // [x, y, val]

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i][0] != h[j][0] {
		return h[i][0] < h[j][0]
	} else if h[i][1] != h[j][1] {
		return h[i][1] > h[j][1]
	} else {
		return h[i][2] < h[j][2]
	}
}

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() []int   { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// tc: O(n log(n))
func verticalTraversal1(root *TreeNode) [][]int {
	h := &MinHeap{}
	heap.Init(h)

	inOrderTraverse(root, 0, 0, h)

	ans := make([][]int, 0)
	x := -1000

	tmp := make([]int, 0)

	for h.Len() > 0 {
		i := heap.Pop(h).([]int)

		if i[0] != x {
			x = i[0]
			if len(tmp) > 0 {
				ans = append(ans, tmp)
			}
			tmp = []int{i[2]}
		} else {
			tmp = append(tmp, i[2])
		}
	}

	// since loop terminates when heap is empty, there might be some left
	// nodes, dont' forget to push
	if len(tmp) > 0 {
		ans = append(ans, tmp)
	}

	return ans
}

func inOrderTraverse(node *TreeNode, x, y int, h *MinHeap) {
	if node == nil {
		return
	}

	heap.Push(h, []int{x, y, node.Val})

	inOrderTraverse(node.Left, x-1, y-1, h)
	inOrderTraverse(node.Right, x+1, y-1, h)
}

//	Notes
//	1.	it takes me long time to understand the problem, how node should be
//		sorted

//	2.	inspired form https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/discuss/231113/C%2B%2B-hashmap-vs.-map

//		since there will at most 1000 nodes, both x & y ranges from -1000 ~ 1000
//		just create [1000][1000][]int, somehow borrow concept of bucket sort

//		I think this one is faster than mine, because mine take O(n log(n)), but
//		the sorted part here is the longest of position (x,y)

//	3.	inspired from sample code, when sort, return multiple bool value meets
//		order criteria

//	4.	though static array [2000][1000]int is easy to allocate, it takes too
//		much time to compare each value, use map[int][][2]int to sort

//	5.	actually, verticalTraversal1 with same runtime as verticalTraversal,
//		it's not bad as I imagined
