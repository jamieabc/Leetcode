package main

// Given a binary tree, count the number of uni-value subtrees.
//
// A Uni-value subtree means all nodes of the subtree have the same value.
//
// Example :
//
// Input:  root = [5,1,5,5,5,null,5]
//
//               5
//              / \
//             1   5
//            / \   \
//           5   5   5
//
// Output: 4

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type info struct {
	count int
	same  bool
}

func countUnivalSubtrees(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	mapping := make(map[*TreeNode]info)
	node := root

	for node != nil || len(stack) != 0 {
		for node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node)
			node = node.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) > 0 && n.Right != nil && n.Right == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			stack = append(stack, n)
			node = n.Right
		} else {
			if n.Left == nil && n.Right == nil {
				mapping[n] = info{
					count: 1,
					same:  true,
				}
			} else if n.Left == nil {
				if mapping[n.Right].same && n.Val == n.Right.Val {
					mapping[n] = info{
						count: mapping[n.Right].count + 1,
						same:  true,
					}
				} else {
					mapping[n] = info{
						count: mapping[n.Right].count,
						same:  false,
					}
				}
			} else if n.Right == nil {
				mapping[n] = info{
					count: mapping[n.Left].count,
					same:  false,
				}

				if mapping[n.Left].same && n.Val == n.Left.Val {
					mapping[n] = info{
						count: mapping[n.Left].count + 1,
						same:  true,
					}
				} else {
					mapping[n] = info{
						count: mapping[n.Left].count,
						same:  false,
					}
				}
			} else {
				mapping[n] = info{
					count: mapping[n.Left].count + mapping[n.Right].count,
					same:  false,
				}

				if mapping[n.Left].same && mapping[n.Right].same && n.Left.Val == n.Right.Val && n.Val == n.Left.Val {
					mapping[n] = info{
						count: mapping[n.Left].count + mapping[n.Right].count + 1,
						same:  true,
					}
				} else {
					mapping[n] = info{
						count: mapping[n.Left].count + mapping[n.Right].count,
						same:  false,
					}
				}
			}
			node = nil
		}
	}

	return mapping[root].count
}

func countUnivalSubtrees1(root *TreeNode) int {
	_, count := postOrder(root)

	return count
}

func postOrder(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}

	if node.Left == nil && node.Right == nil {
		return true, 1
	}

	l, lc := postOrder(node.Left)
	r, rc := postOrder(node.Right)
	sub := lc + rc

	if node.Left == nil {
		if r && node.Right.Val == node.Val {
			return true, sub + 1
		}
		return false, sub
	}

	if node.Right == nil {
		if l && node.Val == node.Left.Val {
			return true, sub + 1
		}
		return false, sub
	}

	if l && r && node.Left.Val == node.Right.Val && node.Val == node.Left.Val {
		return true, sub + 1
	}

	return false, sub
}

//	problems
//	1.	should also compare all children value same, not just direct child
//		value same

//	2.	add reference https://www.geeksforgeeks.org/iterative-postorder-traversal-using-stack/

//		post order iterative traversal, the point here is that store both
//		node.Right & node into stack (if node.Right exist), when popping,
//		check node and stack.top, it they are same, then pop stack top,
//		push node, traverse from popped item

//	3.	inspired from https://leetcode.com/problems/count-univalue-subtrees/discuss/67641/Recursive-and-Iterative-solution-with-detailed-explanation

//		using additional map to store info, write a pretty ugly iterative
//		version
