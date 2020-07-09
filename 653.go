package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	counter := make(map[int]bool)

	// in-order traversal
	stack := make([]*TreeNode, 0)
	ptr := root

	for len(stack) > 0 || ptr != nil {
		for ptr != nil {
			theOtherNum := k - ptr.Val
			if counter[theOtherNum] {
				return true
			}

			counter[ptr.Val] = true

			stack = append(stack, ptr)
			ptr = ptr.Left
		}

		ptr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ptr = ptr.Right
	}

	return false
}
