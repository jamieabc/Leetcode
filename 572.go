package main

// Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also be considered as a subtree of itself.
//
//Example 1:
//Given tree s:
//
//     3
//    / \
//   4   5
//  / \
// 1   2
//
//Given tree t:
//
//   4
//  / \
// 1   2
//
//Return true, because t has the same structure and node values with a subtree of s.
//
//Example 2:
//Given tree s:
//
//     3
//    / \
//   4   5
//  / \
// 1   2
//    /
//   0
//
//Given tree t:
//
//   4
//  / \
// 1   2
//
//Return false.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}

	return equals(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func equals(original, sub *TreeNode) bool {
	if original == nil && sub == nil {
		return true
	} else if original == nil || sub == nil {
		return false
	}

	if original.Val != sub.Val {
		return false
	}

	return equals(original.Left, sub.Left) && equals(original.Right, sub.Right)
}

func isSubtree1(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	return iterate(s, t)
}

func iterate(node, t *TreeNode) bool {
	if node == nil {
		return false
	}

	var middle bool
	if node.Val == t.Val {
		middle = compare(node.Left, t.Left) && compare(node.Right, t.Right)
	}

	if middle {
		return true
	}

	if node.Left == nil && node.Right == nil {
		return false
	}

	var left, right bool

	if node.Left != nil {
		left = iterate(node.Left, t)
	}

	if node.Right != nil {
		right = iterate(node.Right, t)
	}

	return left || right
}

func compare(src, dst *TreeNode) bool {
	if src == nil && dst == nil {
		return true
	}

	if (src == nil && dst != nil) || (src != nil && dst == nil) {
		return false
	}

	left := compare(src.Left, dst.Left)
	right := compare(src.Right, dst.Right)

	return src.Val == dst.Val && left && right
}

//	Notes
//	1. event if value same, traversing still needs to go on
//	2. only check left, forget to check right

//	3.	inspired from solution, traverse tree in some order (pre-oder,
//		in-order, post-order), convert tree into string and check if
//		t-string is in s-string

//	4.	inspired from https://leetcode.com/problems/subtree-of-another-tree/discuss/102724/Java-Solution-tree-traversal

//		author simplifies condition check, first check both nil, then
//		check one nil (||), then check value same

//	5.	inspired from https://leetcode.com/problems/subtree-of-another-tree/discuss/102741/Python-Straightforward-with-Explanation-(O(ST)-and-O(S%2BT)-approaches)

//		interesting thinking to hash nodes' children
