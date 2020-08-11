package main

//Given a binary tree, collect a tree's nodes as if you were doing this: Collect and remove all leaves, repeat until the tree is empty.
//
//
//
//Example:
//
//Input: [1,2,3,4,5]
//
//          1
//         / \
//        2   3
//       / \
//      4   5
//
//Output: [[4,5,3],[2],[1]]
//
//
//
//Explanation:
//
//1. Removing the leaves [4,5,3] would result in this tree:
//
//          1
//         /
//        2
//
//
//
//2. Now removing the leaf [2] would result in this tree:
//
//          1
//
//
//
//3. Now removing the leaf [1] would result in the empty tree:
//
//          []
//
//[[3,5,4],[2],[1]], [[3,4,5],[2],[1]], etc, are also consider correct answers since per each level it doesn't matter the order on which elements are returned.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findLeaves(root *TreeNode) [][]int {
	result := make([][]int, 0)
	inOrder(root, &result)

	return result
}

func inOrder(node *TreeNode, result *[][]int) int {
	if node == nil {
		return 0
	}

	l := inOrder(node.Left, result)
	r := inOrder(node.Right, result)
	node.Left, node.Right = nil, nil

	level := max(l, r) + 1

	// expand result
	if len(*result) < level {
		*result = append(*result, make([][]int, level-len(*result))...)
	}

	(*result)[level-1] = append((*result)[level-1], node.Val)

	return level
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/find-leaves-of-binary-tree/discuss/83778/10-lines-simple-Java-solution-using-recursion-with-explanation

//		a node's distance to leaf is determined if both left & right are
//		traversed, so no need extra map to store, just traverse both
//		children and find level

//		I have noticed that a node's distance to leaf may be changed, but
//		I didn't think of this value can be determined when all children
//		are visited

//	2.	read carefully, need to remove leaves
