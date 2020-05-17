package main

// Given a binary tree, return the preorder traversal of its nodes' values.
//
// Example:
//
// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
//
// Output: [1,2,3]
//
// Follow up: Recursive solution is trivial, could you do it iteratively?

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	stack := []*TreeNode{root}
	var node *TreeNode

	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

//	problems
//	1.	algorithm will always put root into slice, it assumes root will always
//		non nil and not check, and this is not correct if root is nil

//	2.	reference from https://leetcode.com/problems/binary-tree-preorder-traversal/discuss/45266/Accepted-iterative-solution-in-Java-using-stack.

//		author provides elegant code, I didnt see through nature of problem.

//		stack is a way to store something need to be processed later, so I
//		need to store what nodes need to be processed later, and since it's
//		N-L-R, right nodes need to be stored.

//		pre-order traverse, current node is processed first, if there's left
//		child, go into left child, then right child. So if I want to put into
//		stack format, store order is right then left child.

//		compare in-order & pre-order, pre-order is easier because self is
//		processed first, then left & right, so self doesn't need to put into
//		stack. but in-order needs to process left first, which means self
//		need to be process later and store in stack.
