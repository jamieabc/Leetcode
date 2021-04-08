package main

import (
	"fmt"
	"math"
)

//Given a binary tree, determine if it is a valid binary search tree (BST).
//
//Assume a BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than the node's key.
//The right subtree of a node contains only nodes with keys greater than the node's key.
//Both the left and right subtrees must also be binary search trees.
//
//
//
//Example 1:
//
//2
/// \
//1   3
//
//Input: [2,1,3]
//Output: true

// Thought
// For every right child, make sure all right children bigger then self
// For every left child, make sure all left children smaller than self
// Iterate from root to leaf
// Complexity: O(nlogn)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	node := root
	stack := make([]*TreeNode, 0)
	var prev *int

	for node != nil || len(stack) > 0 {
		for ; node != nil; node = node.Left {
			stack = append(stack, node)
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil && *prev >= node.Val {
			return false
		}
		prev = &node.Val

		node = node.Right
	}

	return true
}

func isValidBST3(root *TreeNode) bool {
	ok, _ := inOrder3(root, nil)
	return ok
}

func inOrder3(node *TreeNode, prev *int) (bool, *int) {
	if node == nil {
		return true, prev
	}

	if prev != nil && node.Val <= *prev {
		return false, prev
	}

	ok, l := inOrder3(node.Left, prev)

	if !ok || (l != nil && node.Val <= *l) {
		return false, prev
	}

	return inOrder3(node.Right, &node.Val)
}

// can be further optimized, prev is used on right subtree, return value is
// used for left subtree
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ok, _ := inOrder2(root, nil)
	return ok
}

func inOrder2(node *TreeNode, prev *int) (bool, int) {
	if node.Left == nil && node.Right == nil {
		if prev == nil {
			return true, node.Val
		}
		return node.Val > *prev, node.Val
	}

	if node.Left != nil {
		ok, l := inOrder2(node.Left, prev)

		if !ok || node.Val <= l {
			return false, node.Val
		}
	}

	if node.Right != nil {
		return inOrder2(node.Right, &node.Val)
	}
	return true, node.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return validateBST(root.Left, math.MinInt64, root.Val) && validateBST(root.Right, root.Val, math.MaxInt64)
}

func validateBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	if node.Right != nil && (node.Right.Val <= node.Val || node.Right.Val <= min) {
		return false
	}

	if node.Left != nil && (node.Left.Val >= node.Val || node.Left.Val >= max) {
		return false
	}

	return validateBST(node.Left, min, node.Val) && validateBST(node.Right, node.Val, max)
}

func main() {
	n6 := &TreeNode{
		Val:   13,
		Left:  nil,
		Right: nil,
	}
	n5 := &TreeNode{
		Val:   11,
		Left:  nil,
		Right: nil,
	}
	n4 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	n3 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	n2 := &TreeNode{
		Val:   12,
		Left:  n5,
		Right: n6,
	}
	n1 := &TreeNode{
		Val:   8,
		Left:  n3,
		Right: n4,
	}
	n0 := &TreeNode{
		Val:   10,
		Left:  n1,
		Right: n2,
	}
	fmt.Printf("n0 is BST: %t\n", isValidBST(n0))

	y1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	y0 := &TreeNode{
		Val:   1,
		Left:  y1,
		Right: nil,
	}
	fmt.Printf("y0 is BST: %t\n", isValidBST(y0))

	// [5,14,null,1]
	z3 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	z1 := &TreeNode{
		Val:   14,
		Left:  z3,
		Right: nil,
	}
	z0 := &TreeNode{
		Val:   5,
		Left:  z1,
		Right: nil,
	}
	fmt.Printf("z0 is BST: %t\n", isValidBST(z0))

	// [34,-6,null,-21]
	w3 := &TreeNode{
		Val:   -21,
		Left:  nil,
		Right: nil,
	}
	w1 := &TreeNode{
		Val:   -6,
		Left:  w3,
		Right: nil,
	}
	w0 := &TreeNode{
		Val:   34,
		Left:  w1,
		Right: nil,
	}
	fmt.Printf("w0 is BST: %t\n", isValidBST(w0))
}

//	Notes
//	1.	this problem is hard for me, even today
//		because not just in-order traversal, but also boundary conditions

//		need to check largest from left subtree, which means return children
//		value to parent

//		need to check smallest from right subtree, which means passing parent
//		value to right subtree

//		initially, put math.MinInt32 for dfs, but it encounters a problem:
//		the node value could be math.MinInt32, which makes it hard to detect
//		is smallest node is valid or not. To prevent this, use pointer to check

//	2.	inspired from solution, instead of returning and passing values, it
//		uses min~max range to check
