package main

import (
	"fmt"
)

//Given a binary tree, find the length of the longest path where each node in the path has the same value. This path may or may not pass through the root.
//
//The length of path between two nodes is represented by the number of edges between them.
//
//
//
//Example 1:
//
//Input:
//
//              5
//             / \
//            4   5
//           / \   \
//          1   1   5
//
//Output: 2
//
//
//
//Example 2:
//
//Input:
//
//              1
//             / \
//            4   5
//           / \   \
//          4   4   5
//
//Output: 2
//
//
//
//Note: The given binary tree has not more than 10000 nodes. The height of the tree is not more than 1000.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var longest int
	_ = recursive(root, root.Val, &longest)

	// longest is # of elements, edges are deducted by 1
	return longest - 1
}

func recursive(node *TreeNode, target int, longest *int) int {
	if node == nil {
		return 0
	}

	left := recursive(node.Left, node.Val, longest)
	right := recursive(node.Right, node.Val, longest)

	*longest = max(*longest, left+right+1)

	if node.Val != target {
		return 0
	}

	return max(left, right) + 1
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	programs
//	1.	rewrite, I think it's more easier to understand
//	2.	refactor, use max instead of checking
//	3.	optimize, I didn't think it thoroughly.
//		I use stack to store next node to check. The complexity is O(n), and
//		with memory overhead.

//		The trick is to use post-order, L-R-N, with left & right traversed,
//		a node N can decide it's longest path. The return value is longest
//		left / right path. So it needs to pass N's value to left & right, so
//		that when traversing, a node can know how to return value.
//	4.	need to check if value should update, then return 0 if value not match
