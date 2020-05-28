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
	mapping := make(map[int]int)
	maxFreq := -1

	postOrder(root, mapping, &maxFreq)
	result := make([]int, 0)

	for k, v := range mapping {
		if v == maxFreq {
			result = append(result, k)
		}
	}

	return result
}

func postOrder(node *TreeNode, mapping map[int]int, maxFreq *int) int {
	if node == nil {
		return 0
	}

	l := postOrder(node.Left, mapping, maxFreq)
	r := postOrder(node.Right, mapping, maxFreq)

	sum := l + r + node.Val
	mapping[sum]++

	if mapping[sum] > *maxFreq {
		*maxFreq = mapping[sum]
	}

	return sum
}
