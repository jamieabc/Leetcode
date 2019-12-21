package main

//Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Given the below binary tree and sum = 22,
//
//      5
//     / \
//    4   8
//   /   / \
//  11  13  4
// /  \      \
//7    2      1
//
//return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return traverse(root, sum, 0)
}

func traverse(node *TreeNode, target, tmp int) bool {
	if node == nil {
		return target == tmp
	}

	current := tmp + node.Val

	if current == target && node.Left == nil && node.Right == nil {
		return true
	}

	left, right := false, false

	if node.Left != nil {
		left = traverse(node.Left, target, current)
	}

	if node.Right != nil {
		right = traverse(node.Right, target, current)
	}

	return left || right
}

// problems
// 1. didn't think clear that every path go back to root, or each node check status
// 2. while traversing, when sum is equal, forget to check if it's a leaf
// 3. add a wrong condition, if current sum > target, forget to think that sum might be minus
