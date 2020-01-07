package main

//Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
//For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//
//
//
//But the following [1,2,2,null,3,null,3] is not:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//
//
//
//Note:
//Bonus points if you could solve it both recursively and iteratively.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	// empty
	if root == nil {
		return true
	}

	// single root
	if root.Left == nil && root.Right == nil {
		return true
	}

	return traverse(root.Left, root.Right)
}

func traverse(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	// one side is empty
	if left == nil || right == nil {
		return false
	}

	// value different
	if left.Val != right.Val {
		return false
	}

	// symmetric
	return traverse(left.Left, right.Right) && traverse(left.Right, right.Left)
}

// problems
// 1. symmetric, means value of left children = value of right children
// 2. optimization, when only one node is nil, only need to check any nil, because case of both nil is eliminated
