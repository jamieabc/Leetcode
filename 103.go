package main

// Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).
//
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//
//     3
//    / \
//   9  20
//     /  \
//    15   7
//
// return its zigzag level order traversal as:
//
// [
//   [3],
//   [20,9],
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	level := 0
	result := make([][]int, 0)

	for len(queue) > 0 {
		count := len(queue)
		nums := make([]int, 0)

		for i := 0; i < count; i++ {
			q := queue[0]
			queue = queue[1:]

			if q.Left != nil {
				queue = append(queue, q.Left)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
			}

			if level&1 == 1 {
				tmp := append([]int{}, q.Val)
				tmp = append(tmp, nums...)
				nums = tmp
			} else {
				nums = append(nums, q.Val)
			}
		}
		result = append(result, nums)
		level++
	}

	return result
}
