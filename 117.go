package main

// Given a binary tree
//
// struct Node {
// int val;
// Node *left;
// Node *right;
// Node *next;
// }
// Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
//
// Initially, all next pointers are set to NULL.
//
//
//
// Follow up:
//
// You may only use constant extra space.
// Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.
//
//
// Example 1:
//
//
//
// Input: root = [1,2,3,4,5,null,7]
// Output: [1,#,2,3,#,4,5,7,#]
// Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
//
//
// Constraints:
//
// The number of nodes in the given tree is less than 6000.
// -100 <= node.val <= 100

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	for node := root; node != nil; {
		var next, prev *Node

		for cur := node; cur != nil; cur = cur.Next {
			if prev != nil {
				if cur.Left != nil {
					prev.Next = cur.Left
					prev = cur.Left
				}

				if cur.Right != nil {
					prev.Next = cur.Right
					prev = cur.Right
				}
			} else {
				if cur.Left == nil && cur.Right == nil {
					continue
				} else {
					if cur.Left != nil {
						next = cur.Left
						prev = cur.Left
					}

					if cur.Right != nil {
						if prev != nil {
							prev.Next = cur.Right
						} else {
							next = cur.Right
						}
						prev = cur.Right
					}
				}
			}
		}

		node = next
	}

	return root
}

func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		size := len(stack)

		for i := 0; i < size; i++ {
			node := stack[i]

			if i < size-1 {
				node.Next = stack[i+1]
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}

		stack = stack[size:]
	}

	return root
}

//	Notes
//	1.	it's possible to solve in sc: O(1), because if next left relationship is
//		built properly, then use next to find same level next node
