package main

//Given a binary tree, determine if it is height-balanced.
//
//For this problem, a height-balanced binary tree is defined as:
//
//a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
//
//
//
//Example 1:
//
//Given the following tree [3,9,20,null,null,15,7]:
//
//3
/// \
//9  20
///  \
//15   7
//
//Return true.
//
//Example 2:
//
//Given the following tree [1,2,2,3,3,null,null,4,4]:
//
//1
/// \
//2   2
/// \
//3   3
/// \
//4   4
//
//Return false.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return height(root) != -1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	left := height(node.Left)
	if left == -1 {
		return -1
	}

	right := height(node.Right)
	if right == -1 {
		return -1
	}

	// terminate if any subtree height not equal
	if left != right && left+1 != right && left-1 != right {
		return -1
	}

	if left > right {
		return left + 1
	}
	return right + 1
}
