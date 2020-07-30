package main

// Find the sum of all left leaves in a given binary tree.
//
// Example:
//
//     3
//    / \
//   9  20
//     /  \
//    15   7
//
// There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumOfLeftLeaves(root *TreeNode) int {
	// this one is critical, because after this checking, loop always assumes
	// data is not nil
	if root == nil {
		return 0
	}

	var sum int

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				sum += node.Left.Val
			} else {
				stack = append(stack, node.Left)
			}
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return sum
}

type Info struct {
	Node   *TreeNode
	IsLeft bool
}

func sumOfLeftLeaves2(root *TreeNode) int {
	var sum int

	info := Info{
		Node:   root,
		IsLeft: false,
	}
	stack := make([]Info, 0)

	for info.Node != nil || len(stack) > 0 {
		for info.Node != nil {
			stack = append(stack, info)

			info = Info{
				Node:   info.Node.Left,
				IsLeft: true,
			}
		}

		info = stack[len(stack)-1]
		node := info.Node
		stack = stack[:len(stack)-1]

		if node.Left == nil && node.Right == nil && info.IsLeft {
			sum += node.Val
		}

		info = Info{
			Node:   node.Right,
			IsLeft: false,
		}
	}

	return sum
}

func sumOfLeftLeaves1(root *TreeNode) int {
	return inOrder(root, false)
}

func inOrder(node *TreeNode, isLeft bool) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil && isLeft {
		return node.Val
	}

	l, r := inOrder(node.Left, true), inOrder(node.Right, false)
	return l + r
}

//	problems
//	1.	be careful about leaf checking, because this checking also fits to root

//	2.	need a way to node to determine left is left or not
