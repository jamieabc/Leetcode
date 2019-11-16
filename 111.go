package main

import "fmt"

//Given a binary tree, find its minimum depth.
//
//The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Given binary tree [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//return its minimum depth = 2.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return traverse(root)
}

func traverse(node *TreeNode) int {
	if node.Left == nil && node.Right == nil {
		return 1
	}

	if node.Left != nil && node.Right == nil {
		return traverse(node.Left) + 1
	}

	if node.Left == nil && node.Right != nil {
		return traverse(node.Right) + 1
	}

	left := traverse(node.Left) + 1
	right := traverse(node.Right) + 1

	if left < right {
		return left
	} else {
		return right
	}
}

func main() {
	n1 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	n0 := &TreeNode{
		Val:   1,
		Left:  n1,
		Right: nil,
	}

	fmt.Printf("n0 min depth: %d\n", minDepth(n0))
}
