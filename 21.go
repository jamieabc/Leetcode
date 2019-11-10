package main

//Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
//Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, cur *ListNode
	node1 := l1
	node2 := l2

	// empty
	if node1 == nil && node2 == nil {
		return nil
	}

	// node 1 empty
	if node1 == nil {
		return node2
	}

	// node 2 empty
	if node2 == nil {
		return node1
	}

	// setup head
	if node1.Val <= node2.Val {
		result = node1
		node1 = node1.Next
	} else {
		result = node2
		node2 = node2.Next
	}

	cur = result

	for node1 != nil || node2 != nil {
		if node1 == nil {
			cur, node2 = nextNode(cur, node2)
		} else if node2 == nil {
			cur, node1 = nextNode(cur, node1)
		} else {
			// choose by smaller
			if node1.Val <= node2.Val {
				cur, node1 = nextNode(cur, node1)
			} else {
				cur, node2 = nextNode(cur, node2)
			}
		}
	}

	return result
}

func nextNode(cur *ListNode, node *ListNode) (*ListNode, *ListNode) {
	cur.Next = node
	return cur.Next, node.Next
}
