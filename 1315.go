package main

// Given a binary tree, return the sum of values of nodes with even-valued grandparent.  (A grandparent of a node is the parent of its parent, if it exists.)
//
// If there are no nodes with an even-valued grandparent, return 0.
//
//
//
// Example 1:
//
//
//
// Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
// Output: 18
// Explanation: The red nodes are the nodes with even-value grandparent while the blue nodes are the even-value grandparents.
//
//
// Constraints:
//
// The number of nodes in the tree is between 1 and 10^4.
// The value of nodes is between 1 and 100.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Info struct {
	Node                  *TreeNode
	ParentEven, GrandEven bool
}

func sumEvenGrandparent(root *TreeNode) int {
	var sum int
	queue := []Info{{Node: root}}

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		if q.GrandEven {
			sum += q.Node.Val
		}

		var selfEven bool
		if q.Node.Val&1 == 0 {
			selfEven = true
		}

		if q.Node.Left != nil {
			queue = append(queue, Info{
				Node:       q.Node.Left,
				ParentEven: selfEven,
				GrandEven:  q.ParentEven,
			})
		}

		if q.Node.Right != nil {
			queue = append(queue, Info{
				Node:       q.Node.Right,
				ParentEven: selfEven,
				GrandEven:  q.ParentEven,
			})
		}
	}

	return sum
}

func sumEvenGrandparent1(root *TreeNode) int {
	return inOrder(root, false, false)
}

func inOrder(node *TreeNode, parentEven, grandEven bool) int {
	if node == nil {
		return 0
	}

	selfEven := false
	if node.Val&1 == 0 {
		selfEven = true
	}

	childrenSum := inOrder(node.Left, selfEven, parentEven) + inOrder(node.Right, selfEven, parentEven)

	if grandEven {
		return node.Val + childrenSum
	}
	return childrenSum
}

//	problems
//	1.	inspired from https://leetcode.com/problems/sum-of-nodes-with-even-valued-grandparent/discuss/477048/JavaC%2B%2BPython-1-Line-Recursive-Solution

//		could also pass grand parent value or pointer

//	2.	inspired from https://leetcode.com/problems/sum-of-nodes-with-even-valued-grandparent/discuss/482991/Easy-BFS-solution-in-Java

//		could also just check grand child value
