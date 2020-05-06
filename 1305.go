package main

// Given two binary search trees root1 and root2.
//
// Return a list containing all the integers from both trees sorted in ascending order.
//
//
//
// Example 1:
//
// Input: root1 = [2,1,4], root2 = [1,0,3]
// Output: [0,1,1,2,3,4]
//
// Example 2:
//
// Input: root1 = [0,-10,10], root2 = [5,1,7,0,2]
// Output: [-10,0,0,1,2,5,7,10]
//
// Example 3:
//
// Input: root1 = [], root2 = [5,1,7,0,2]
// Output: [0,1,2,5,7]
//
// Example 4:
//
// Input: root1 = [0,-10,10], root2 = []
// Output: [-10,0,10]
//
// Example 5:
//
// Input: root1 = [1,null,8], root2 = [8,1]
// Output: [1,1,8,8]
//
//
//
// Constraints:
//
//     Each tree has at most 5000 nodes.
//     Each node's value is between [-10^5, 10^5].

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	q1 := make([]*TreeNode, 0)
	q2 := make([]*TreeNode, 0)

	// initial setup
	traverse(root1, &q1)
	traverse(root2, &q2)

	result := make([]int, 0)
	for len(q1) != 0 || len(q2) != 0 {
		var n1, n2 *TreeNode

		if len(q1) > 0 {
			n1 = q1[len(q1)-1]
		}

		if len(q2) > 0 {
			n2 = q2[len(q2)-1]
		}

		if (n1 != nil && n2 != nil && n1.Val <= n2.Val) || n2 == nil {
			result = append(result, n1.Val)
			q1 = q1[:len(q1)-1]
			if n1.Right != nil {
				traverse(n1.Right, &q1)
			}
		} else {
			result = append(result, n2.Val)
			q2 = q2[:len(q2)-1]
			if n2.Right != nil {
				traverse(n2.Right, &q2)
			}
		}
	}

	return result
}

func traverse(node *TreeNode, queue *[]*TreeNode) {
	for ; node != nil; node = node.Left {
		*queue = append(*queue, node)
	}
}

//	problems
//	1.	wrong criteria of check i, j is valid

//	2.	wrong condition, when j is increased, should not do checking cause it
//		might breaks loop condition

//	3.	too slow, I guess it's from allocating two slices to store traverse
//		order.

//	4.	cannot assume both tree at same height
