package main

// For a binary tree T, we can define a flip operation as follows: choose any node, and swap the left and right child subtrees.
//
// A binary tree X is flip equivalent to a binary tree Y if and only if we can make X equal to Y after some number of flip operations.
//
// Write a function that determines whether two binary trees are flip equivalent.  The trees are given by root nodes root1 and root2.
//
//
//
// Example 1:
//
// Input: root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
// Output: true
// Explanation: We flipped at nodes with values 1, 3, and 5.
// Flipped Trees Diagram
//
//
//
// Note:
//
//     Each tree will have at most 100 nodes.
//     Each value in each tree will be a unique integer in the range [0, 99].

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipEquiv(n1, n2 *TreeNode) bool {
	arr1 := []*TreeNode{n1}
	arr2 := []*TreeNode{n2}

	for len(arr1) > 0 && len(arr2) > 0 {
		n1 := arr1[len(arr1)-1]
		arr1 = arr1[:len(arr1)-1]

		n2 := arr2[len(arr2)-1]
		arr2 = arr2[:len(arr2)-1]

		if n1 == nil && n2 == nil {
			continue
		}

		if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}

		arr1 = append(arr1, n1.Left)
		arr1 = append(arr1, n1.Right)

		if n1.Left == n2.Left || (n1.Left != nil && n2.Left != nil && n1.Left.Val == n2.Left.Val) {
			arr2 = append(arr2, n2.Left)
			arr2 = append(arr2, n2.Right)
		} else {
			arr2 = append(arr2, n2.Right)
			arr2 = append(arr2, n2.Left)
		}
	}

	return len(arr1) == 0 && len(arr2) == 0
}

func flipEquiv2(r1, r2 *TreeNode) bool {
	stack := []*TreeNode{r1, r2}

	for len(stack) != 0 {
		n1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n2 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if n1 == nil && n2 == nil {
			continue
		}

		if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}

		if n1.Left == nil && n1.Right == nil && n2.Left == nil && n2.Right == nil {
			continue
		}

		if n1.Left != nil && n2.Left != nil && n1.Right != nil && n2.Right != nil {
			if n1.Left.Val == n2.Left.Val {
				stack = append(stack, n1.Left)
				stack = append(stack, n2.Left)
				stack = append(stack, n1.Right)
				stack = append(stack, n2.Right)
			} else {
				stack = append(stack, n1.Left)
				stack = append(stack, n2.Right)
				stack = append(stack, n1.Right)
				stack = append(stack, n2.Left)
			}
		} else if n1.Left == nil {
			if n2.Left == nil {
				stack = append(stack, n1.Right)
				stack = append(stack, n2.Right)
			} else if n2.Right == nil {
				stack = append(stack, n1.Right)
				stack = append(stack, n2.Left)
			} else {
				return false
			}
		} else if n1.Right == nil {
			if n2.Right == nil {
				stack = append(stack, n1.Left)
				stack = append(stack, n2.Left)
			} else if n2.Left == nil {
				stack = append(stack, n1.Left)
				stack = append(stack, n2.Right)
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func flipEquiv1(n1, n2 *TreeNode) bool {
	return compare(n1, n2)
}

func compare(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		return false
	}

	if n1.Val != n2.Val {
		return false
	}

	return (compare(n1.Left, n2.Left) && compare(n1.Right, n2.Right)) || (compare(n1.Left, n2.Right) && compare(n1.Right, n2.Left))
}

//	problems
//	1.	reference from https://leetcode.com/problems/flip-equivalent-binary-trees/discuss/200514/JavaPython-3-DFS-3-liners-and-BFS-with-explanation-time-and-space%3A-O(n).

//		iterative is much more complex for me... I try to simulate iterative
//		by adding stack. For this problem, the point is to add "correct"
//		nodes & order into stack

//	2.	add single stack solution, it's much more complex because I have to
//		know exact order of each node
