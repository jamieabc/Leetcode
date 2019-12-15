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

	closed := 0
	_ = recursive(root, &closed)
	return closed
}

func recursive(node *TreeNode, closed *int) int {
	if node == nil {
		return 0
	}

	left := recursive(node.Left, closed)
	// left same
	if node.Left != nil && node.Val == node.Left.Val {
		left++
	}

	// left different, start new path
	if node.Left != nil && node.Val != node.Left.Val {
		left = 0
	}

	right := recursive(node.Right, closed)
	// right same
	if node.Right != nil && node.Val == node.Right.Val {
		right++
	}

	// right different, start new path
	if node.Right != nil && node.Val != node.Right.Val {
		right = 0
	}

	*closed = max(*closed, left+right)

	return max(left, right)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func main() {
	ll2 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	lr2 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	ll1 := &TreeNode{
		Val:   4,
		Left:  ll2,
		Right: lr2,
	}
	rl2 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	rr1 := &TreeNode{
		Val:   5,
		Left:  rl2,
		Right: nil,
	}
	root := &TreeNode{
		Val:   1,
		Left:  ll1,
		Right: rr1,
	}
	fmt.Printf("max: %d\n", longestUnivaluePath(root))
}
