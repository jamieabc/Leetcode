package main

//Given a binary tree, return the inorder traversal of its nodes' values.
//
//Example:
//
//Input: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//Output: [1,3,2]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := []*TreeNode{root}
	next := root.Left

	for len(stack) > 0 || next != nil {
		// keep traverse left
		for next != nil {
			stack = append(stack, next)
			next = next.Left
		}

		next = stack[len(stack)-1]
		result = append(result, next.Val)
		stack = stack[:len(stack)-1]

		next = next.Right
	}

	return result
}

func inorderTraversal1(root *TreeNode) []int {
	result := make([]int, 0)
	traverse(root, &result)
	return result
}

// inorder: left-root-right
func traverse(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		traverse(node.Left, result)
	}

	*result = append(*result, node.Val)

	if node.Right != nil {
		traverse(node.Right, result)
	}
}

//	problems
//	1. 	recursive is easy, I should try looping. The difference of
//		recursive & looping is that recursive is stack based (function
//		call), so first thing to do is to used stack.
//		w/ stack cannot implement traversed, the reason is because
//		program doesn't know where to start.

//		Recall from recursive, it keeps calling it self, and return to
//		the location that is called, so if I change it to looping,
//		then I need a way to know which part of logic to start.

//		so I created two variables left & right to denote if left &
//		right child are traversed.

//		If left child is not traversed, add this node to stack.
//		If left child is traversed, add self value to result, proceed
//		to right child.
//		If both left & right child are traversed, remove the node.

//	2.	optimize, program can be further improved by removing flags,
//		since left & right is to denote left & right child is
//		traversed or not, if I can find another way to only keep
//		nodes that need to be traversed, then left & right can be
//		eliminated

//		The other point of in-order traversal, it's mechanism is keep
//		traversing left, then to right. So stack means something to
//		deal with.

//		The problem of using iteration is "how to know a node is
//		processed and can be abandoned? A loop can be used to always
//		finding left node if exist. And then save the value to
//		result, remove the node, keep on going to right child.

//		But the problem here is tricky, since a node in stack means
//		something to deal with, how do I know it's left is already
//		processed?

//		So my thinking here is changed, "something" needs to be
//		checked to know iteration keeps going.

//		Overall process as follows: put nodes that haven't been
//		processed into stack. Keep traversing left nodes. If node's
//		left child is empty, store value, remove it from stack,
//		keep traversing right child.

//		And since a node is removed first then check right child,
//		so loop criteria needs to changed: both stack & pointer are
//		empty that program can end.
