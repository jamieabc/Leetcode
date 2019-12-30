package main

//Given a non-empty special binary tree consisting of nodes with the non-negative value, where each node in this tree has exactly two or zero sub-node. If the node has two sub-nodes, then this node's value is the smaller value among its two sub-nodes. More formally, the property root.val = min(root.left.val, root.right.val) always holds.
//
//Given such a binary tree, you need to output the second minimum value in the set made of all the nodes' value in the whole tree.
//
//If no such second minimum value exists, output -1 instead.
//
//Example 1:
//
//Input:
//    2
//   / \
//  2   5
//     / \
//    5   7
//
//Output: 5
//Explanation: The smallest value is 2, the second smallest value is 5.
//
//
//
//Example 2:
//
//Input:
//    2
//   / \
//  2   2
//
//Output: -1
//Explanation: The smallest value is 2, but there isn't any second smallest value.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	// 2 conditions:
	// 1 - all same, return same value
	// 2 - one different, return different
	// determination - closest to root value
	result := recursive(root, root.Val)
	if result == root.Val || result == -1 {
		return -1
	}

	return result
}

//      2
//   2     2
// 5   2 2   7

//      2
//   2     5
//       5   7
func recursive(node *TreeNode, target int) int {
	if node == nil {
		return -1
	}

	if node.Left == nil && node.Right == nil {
		return node.Val
	}

	var left, right int

	if node.Left != nil {
		left = recursive(node.Left, target)
	}

	if node.Right != nil {
		right = recursive(node.Right, target)
	}

	// all same value
	if node.Val == left && node.Val == right {
		return node.Val
	}

	// one different
	if node.Val != target {
		return node.Val
	}
	if left == target {
		return right
	}
	if right == target {
		return left
	}

	if left < right {
		return left
	}
	return right
}

// problems
// 1. last condition compare left & right is not correct, need to exclude condition one side equals -1
// 2. second largest number might exist at the end of tree, the number cannot be decided by middle of traversing
// 3. think about wrong algorithm, when node, left, right are all different, need to check if node value different from target, if yes, return node value
// 4. for children value if any value equals target, return the other one; otherwise, return the smallest one
