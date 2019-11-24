package main

//Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
//
//Note:
//You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
//
//Example 1:
//
//Input: root = [3,1,4,null,2], k = 1
//   3
//  / \
// 1   4
//  \
//   2
//Output: 1
//
//Example 2:
//
//Input: root = [5,3,6,2,4,null,null,1], k = 3
//       5
//      / \
//     3   6
//    / \
//   2   4
//  /
// 1
//Output: 3
//
//Follow up:
//What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	val, _ := traverse(root, k)
	return val
}

func traverse(node *TreeNode, k int) (int, bool) {
	if node == nil {
		return 0, false
	}

	left, found := traverse(node.Left, k)
	if found {
		return left, found
	}

	if left+1 == k {
		return node.Val, true
	}

	right, found := traverse(node.Right, k-left-1)
	if found {
		return right, found
	}

	return left + right + 1, false
}
