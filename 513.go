package main

// Given a binary tree, find the leftmost value in the last row of the tree.
//
// Example 1:
//
// Input:
//
//     2
//    / \
//   1   3
//
// Output:
// 1
//
// Example 2:
//
// Input:
//
//         1
//        / \
//       2   3
//      /   / \
//     4   5   6
//        /
//       7
//
// Output:
// 7
//
// Note: You may assume the tree (i.e., the given root node) is not NULL.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	maxLevel := -1
	var val int
	traverse(root, 0, &maxLevel, &val)

	return val
}

// DFS
func traverse(node *TreeNode, level int, maxLevel, val *int) {
	if node == nil {
		return
	}

	if level > *maxLevel {
		*maxLevel = level
		*val = node.Val
	}

	traverse(node.Left, level+1, maxLevel, val)
	traverse(node.Right, level+1, maxLevel, val)
}

// BFS
func findBottomLeftValue1(root *TreeNode) int {
	stack := []*TreeNode{root}
	val := root.Val

	for len(stack) > 0 {
		size := len(stack)
		var found bool

		for i := 0; i < size; i++ {
			if stack[i].Left != nil {
				if !found {
					val = stack[i].Left.Val
					found = true
				}
				stack = append(stack, stack[i].Left)
			}

			if stack[i].Right != nil {
				if !found {
					val = stack[i].Right.Val
					found = true
				}
				stack = append(stack, stack[i].Right)
			}
		}

		stack = stack[size:]
	}

	return val
}

//	Notes
//	1.	reference from https://leetcode.com/problems/find-bottom-left-tree-value/discuss/98779/Right-to-Left-BFS-(Python-%2B-Java)

//		BFS can also do this, but aware that the order traverse is right then
//		left, because queue is first in first out, so right first means last
//		one will be left-most node
