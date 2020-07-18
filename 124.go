package main

import "math"

// Given a non-empty binary tree, find the maximum path sum.
//
// For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.
//
// Example 1:
//
// Input: [1,2,3]
//
//        1
//       / \
//      2   3
//
// Output: 6
//
// Example 2:
//
// Input: [-10,9,20,null,null,15,7]
//
//    -10
//    / \
//   9  20
//     /  \
//    15   7
//
// Output: 42

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	if root == nil {
		return 0
	}

	pathSum(root, &maxSum)

	return maxSum
}

// iterate through every node
// find maximum path sum
func pathSum(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	l := pathSum(node.Left, maxSum)
	r := pathSum(node.Right, maxSum)

	curMax := max(node.Val, node.Val+max(l, r))
	*maxSum = max(*maxSum, max(curMax, node.Val+l+r))

	return curMax
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	max sum path for a node may come from itself only

//	2.	since no need to pass root and any node, max sum could be 3 conditions:
//		- self only
//		- self + left
//		- self + right
//		- self + left + right

//	3.	in this case, memo is not needed, because every node is traversed once
