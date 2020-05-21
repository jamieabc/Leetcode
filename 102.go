package main

// Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
//
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its level order traversal as:
//
// [
//   [3],
//   [9,20],
//   [15,7]
// ]
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// DFS
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	traverse(&result, 0, root)
	return result
}

func traverse(result *[][]int, level int, node *TreeNode) {
	if node == nil {
		return
	}

	// make sure result has enough space
	length := len(*result)
	if level+1 > length {
		for i := level + 1 - length; i > 0; i-- {
			*result = append(*result, []int{})
		}
	}

	(*result)[level] = append((*result)[level], node.Val)
	traverse(result, level+1, node.Left)
	traverse(result, level+1, node.Right)
}

// BFS
func levelOrder1(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	stack := []*TreeNode{root}

	for length := len(stack); true; length = len(stack) {
		tmp := make([]int, 0)
		for i := 0; i < length; i++ {
			tmp = append(tmp, stack[i].Val)
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}

			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		if len(tmp) == 0 {
			break
		}

		result = append(result, tmp)
		stack = stack[length:]
	}

	return result
}

//	problems
//	1.	 cannot assume root is non-nil
