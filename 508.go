package main

//  Given the root of a tree, you are asked to find the most frequent subtree sum. The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself). So what is the most frequent subtree sum value? If there is a tie, return all the values with the highest frequency in any order.
//
// Examples 1
// Input:
//
//   5
//  /  \
// 2   -3
//
// return [2, -3, 4], since all the values happen only once, return all of them in any order.
//
// Examples 2
// Input:
//
//   5
//  /  \
// 2   -5
//
// return [2], since 2 happens twice, however -5 only occur once.
//
// Note: You may assume the sum of values in any subtree is in the range of 32-bit signed integer.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findFrequentTreeSum(root *TreeNode) []int {
	counter := make(map[int]int)
	var maxOccurrence int

	recursive(root, counter, &maxOccurrence)

	ans := make([]int, 0)

	for key, val := range counter {
		if val == maxOccurrence {
			ans = append(ans, key)
		}
	}

	return ans
}

func recursive(node *TreeNode, counter map[int]int, maxOccurrence *int) int {
	if node == nil {
		return 0
	}

	cur := node.Val + recursive(node.Left, counter, maxOccurrence) + recursive(node.Right, counter, maxOccurrence)

	counter[cur]++

	if counter[cur] > *maxOccurrence {
		*maxOccurrence = counter[cur]
	}

	return cur
}
