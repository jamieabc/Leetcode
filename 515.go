package main

import "math"

// You need to find the largest value in each row of a binary tree.
//
// Example:
//
// Input:
//
//           1
//          / \
//         3   2
//        / \   \
//       5   3   9
//
// Output: [1, 3, 9]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		maxVal := math.MinInt32

		for i := 0; i < size; i++ {
			maxVal = max(maxVal, queue[i].Val)

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		ans = append(ans, maxVal)
		queue = queue[size:]
	}

	return ans
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func largestValues1(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	traverse([]*TreeNode{root}, &result)

	return result
}

func traverse(nodes []*TreeNode, result *[]int) {
	length := len(nodes)
	if length == 0 {
		return
	}

	max := math.MinInt32

	for i := 0; i < length; i++ {
		if nodes[i].Val > max {
			max = nodes[i].Val
		}

		if nodes[i].Left != nil {
			nodes = append(nodes, nodes[i].Left)
		}

		if nodes[i].Right != nil {
			nodes = append(nodes, nodes[i].Right)
		}
	}

	*result = append(*result, max)

	traverse(nodes[length:], result)
}

//	problems
//	1.	reference from https://leetcode.com/problems/find-largest-value-in-each-tree-row/discuss/98971/9ms-JAVA-DFS-solution

//		slice can be reused, to avoid some un-necessary memory allocation
