package main

//Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
//
//For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
//Example:
//
//Given the sorted array: [-10,-3,0,5,9],
//
//One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
//
//0
/// \
//-3   9
///   /
//-10  5

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)

	if length == 0 {
		return nil
	}

	if length == 1 {
		return &TreeNode{Val: nums[0]}
	}

	return balanceTree(nums, 0, length-1)
}

func balanceTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	if left == right {
		return &TreeNode{Val: nums[left]}
	}

	middle := (left + right) / 2
	return &TreeNode{
		Val:   nums[middle],
		Left:  balanceTree(nums, left, middle-1),
		Right: balanceTree(nums, middle+1, right),
	}
}
