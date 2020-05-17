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
func largestValues(root *TreeNode) []int {
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
