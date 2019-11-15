package main

import (
	"fmt"
)

//Given a binary tree, find its maximum depth.
//
//The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
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
//return its depth = 3.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return traverse(root)
}

func traverse(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := 1 + traverse(root.Left)
	right := 1 + traverse(root.Right)

	if right > left {
		return right
	} else {
		return left
	}
}

func main() {
	//[3,9,20,null,null,15,7]
	n6 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	n5 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	n2 := &TreeNode{
		Val:   20,
		Left:  n5,
		Right: n6,
	}
	n1 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	n0 := &TreeNode{
		Val:   3,
		Left:  n1,
		Right: n2,
	}

	fmt.Printf("n0 max depth: %d\n", maxDepth(n0))
}
