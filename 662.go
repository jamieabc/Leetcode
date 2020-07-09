package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type item struct {
	id   int
	node *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// bfs
	stack := []item{{
		id:   0,
		node: root,
	}}
	var maxWidth int

	for len(stack) > 0 {
		size := len(stack)
		start, end := math.MaxInt64, math.MinInt64

		for j := 0; j < size; j++ {
			s := stack[j]
			start, end = min(start, s.id), max(end, s.id)

			if s.node.Left != nil {
				stack = append(stack, item{
					id:   s.id*2 + 1,
					node: s.node.Left,
				})
			}

			if s.node.Right != nil {
				stack = append(stack, item{
					id:   s.id*2 + 2,
					node: s.node.Right,
				})
			}
		}

		maxWidth = max(maxWidth, int(end-start+1))

		stack = stack[size:]
	}

	return maxWidth
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//  problems
//  1.  level could exceed int32 size

//	2.	add dfs version https://leetcode.com/problems/maximum-width-of-binary-tree/discuss/106707/Python-Straightforward-BFS-and-DFS-solutions
