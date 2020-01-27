package main

//Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
//
//Example:
//Given a binary tree
//
//          1
//         / \
//        2   3
//       / \
//      4   5
//
//Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
//
//Note: The length of path between two nodes is represented by the number of edges between them.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	traverse(root, &max)
	return max
}

func traverse(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}

	l := traverse(node.Left, max)
	r := traverse(node.Right, max)

	sum := l + r
	if sum > *max {
		*max = sum
	}

	if l >= r {
		return l + 1
	} else {
		return r + 1
	}
}

// problems
// 1. longest path means only space between nodes
// 2. it's not necessary leaf-to-leaf, any 2 nodes can do
