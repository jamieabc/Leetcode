package main

// Convert a Binary Search Tree to a sorted Circular Doubly-Linked List in place.
//
// You can think of the left and right pointers as synonymous to the predecessor and successor pointers in a doubly-linked list. For a circular doubly linked list, the predecessor of the first element is the last element, and the successor of the last element is the first element.
//
// We want to do the transformation in place. After the transformation, the left pointer of the tree node should point to its predecessor, and the right pointer should point to its successor. You should return the pointer to the smallest element of the linked list.
//
//
//
// Example 1:
//
// Input: root = [4,2,5,1,3]
//
//
// Output: [1,2,3,4,5]
//
// Explanation: The figure below shows the transformed BST. The solid line indicates the successor relationship, while the dashed line means the predecessor relationship.
//
// Example 2:
//
// Input: root = [2,1,3]
// Output: [1,2,3]
//
// Example 3:
//
// Input: root = []
// Output: []
// Explanation: Input is an empty tree. Output is also an empty Linked List.
//
// Example 4:
//
// Input: root = [1]
// Output: [1]
//
//
//
// Constraints:
//
//     -1000 <= Node.val <= 1000
//     Node.left.val < Node.val < Node.right.val
//     All values of Node.val are unique.
//     0 <= Number of Nodes <= 2000

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

// recursive
func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	dummy := &Node{}
	dummy.Right.Left = traverse(root, dummy)
	dummy.Right.Left.Right = dummy.Right

	return dummy.Right
}

func traverse(node, prev *Node) *Node {
	if node.Left != nil {
		node.Left = traverse(node.Left, prev)
		node.Left.Right = node
	} else {
		prev.Right = node
		node.Left = prev
	}

	if node.Right != nil {
		return traverse(node.Right, node)
	}

	return node
}

// iterative
func treeToDoublyList1(root *Node) *Node {
	if root == nil {
		return nil
	}

	node := root
	prev := &Node{}
	dummy := prev
	stack := make([]*Node, 0)

	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		// find next node that needs to process left child
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// make sure node's left is prev (needed when jumping to parent)
		prev.Right = node
		node.Left = prev

		prev = node
		node = node.Right
	}

	prev.Right = dummy.Right
	dummy.Right.Left = prev

	return dummy.Right
}

//  problems
//  1.  make sure when jumping from a node w/o right to it's parent,
//      connect parent & child of left & right child relationships

//	2.	boundary condition needs to check, because when setting head,
//		assumes stack is non-empty which is not ture when root is nil

//	3. inspired from https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/discuss/169965/Python-Inorder-Transverse

//		it's elegant to combine traversing left/right child into same loop

//		the other beautiful thing is author uses a dummy node as head, reduce
//		additional checking in loop

//	4.	when writing recursive, takes me near an hour to come out a solution
//		part of the reason is because I didn't think it clear that how to
//		deal with prev, and what value should be returned

//	5.	add reference  https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/discuss/149151/Concise-Java-solution-Beats-100

//		author also uses recursive, but a little different than mine, didn't
//		take time to look at it

//	6.	add reference https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/discuss/154659/Divide-and-Conquer-without-Dummy-Node-Java-Solution

//		author uses divide-and-conquer, didn't take time to read it
