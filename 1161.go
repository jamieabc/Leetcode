package main

import "math"

// Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
//
// Return the smallest level X such that the sum of all the values of nodes at level X is maximal.
//
//
//
// Example 1:
//
// Input: [1,7,0,7,-8,null,null]
// Output: 2
// Explanation:
// Level 1 sum = 1.
// Level 2 sum = 7 + 0 = 7.
// Level 3 sum = 7 + -8 = -1.
// So we return the level with the maximum sum which is level 2.
//
//
//
// Note:
//
//     The number of nodes in the given tree is between 1 and 10^4.
//     -10^5 <= node.val <= 10^5

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	maxLevel := 1
	level := 1
	maxSum := math.MinInt32
	var tmp int

	for len(queue) > 0 {
		tmp = 0

		for idx := len(queue) - 1; idx >= 0; idx-- {
			tmp += queue[0].Val
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}

		if tmp > maxSum {
			maxSum = tmp
			maxLevel = level
		}
		level++
	}

	return maxLevel
}

//	problems
//	1.	too slow, I thinks it's because too many memory allocations

//	2. 	even use same slice runtime is still slow, I think most time comes
//		from space allocation/de-allocation, so I will store each level sum
//		into slice, size of slice is O(log n), complexity is O(n)

//	3.	reference from sample code, I can remove every element used when
//		it's value is calculated

//		This implementation uses less memory because every time a node is
//		processed, it's memory space is released, so it could possibly
//		reduce memory allocation
