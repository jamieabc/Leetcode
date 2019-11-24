package main

//Given a binary tree, flatten it to a linked list in-place.
//
//For example, given the following tree:
//
//    1
//   / \
//  2   5
// / \   \
//3   4   6
//
//The flattened tree should look like:
//
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	convert(root)
}

func convert(node *TreeNode) (*TreeNode, *TreeNode) {
	if node.Left == nil && node.Right == nil {
		return node, node
	}

	var tmp *TreeNode

	if node.Left == nil && node.Right != nil {
		node.Right, tmp = convert(node.Right)
		return node, tmp
	}

	if node.Left != nil && node.Right == nil {
		node.Right, tmp = convert(node.Left)
		node.Left = nil
		return node, tmp
	}

	left, leftEnd := convert(node.Left)
	right, rightEnd := convert(node.Right)

	node.Right = left
	leftEnd.Right = right
	node.Left = nil

	return node, rightEnd
}
