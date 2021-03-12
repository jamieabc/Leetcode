package main

// You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Follow up:
// What if you cannot modify the input lists? In other words, reversing the lists is not allowed.
//
// Example:
//
// Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 8 -> 0 -> 7

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var n1, n2 int

	for ptr := l1; ptr != nil; ptr = ptr.Next {
		n1++
	}

	for ptr := l2; ptr != nil; ptr = ptr.Next {
		n2++
	}

	head := &ListNode{}

	for ptr1, ptr2 := l1, l2; n1 > 0 || n2 > 0; {
		if n1 == n2 {
			head.Val = ptr1.Val + ptr2.Val
			ptr1, ptr2 = ptr1.Next, ptr2.Next
			n1, n2 = n1-1, n2-1
		} else if n1 > n2 {
			head.Val = ptr1.Val
			ptr1 = ptr1.Next
			n1--
		} else {
			head.Val = ptr2.Val
			ptr2 = ptr2.Next
			n2--
		}

		next := &ListNode{
			Next: head,
		}
		head = next
	}

	var overflow int
	var prev, cur, next *ListNode

	for prev, cur = nil, head.Next; cur != nil; {
		cur.Val += overflow

		if cur.Val >= 10 {
			cur.Val -= 10
			overflow = 1
		} else {
			overflow = 0
		}

		next = cur.Next
		cur.Next = prev
		cur, prev = next, cur
	}

	if overflow == 1 {
		return &ListNode{
			Val:  1,
			Next: prev,
		}
	}

	return prev
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := make([]*ListNode, 0), make([]*ListNode, 0)

	for ptr := l1; ptr != nil; ptr = ptr.Next {
		stack1 = append(stack1, ptr)
	}

	for ptr := l2; ptr != nil; ptr = ptr.Next {
		stack2 = append(stack2, ptr)
	}

	var overflow int
	cur := &ListNode{}

	for i, j := len(stack1)-1, len(stack2)-1; i >= 0 || j >= 0; {
		if i >= 0 && j >= 0 {
			cur.Val = stack1[i].Val + stack2[j].Val + overflow
			i, j = i-1, j-1
		} else if i >= 0 {
			cur.Val = stack1[i].Val + overflow
			i--
		} else {
			cur.Val = stack2[j].Val + overflow
			j--
		}

		if cur.Val >= 10 {
			cur.Val -= 10
			overflow = 1
		} else {
			overflow = 0
		}

		tmp := &ListNode{
			Next: cur,
		}
		cur = tmp
	}

	if overflow == 1 {
		cur.Val = 1
		return cur
	}
	return cur.Next
}

//	Notes
//	1.	becareful when all nodes are processed, and overflow exist

//	2.	>= 10, deduct 10

//	3.	inspired from sample code, can first align two linked lists, and add
//		values w/o overflow and link nodes from highest -> lowest (head will be
//		at lowest), then reverse process linked list

//		this is reverse of reverse if normal, kind of technique to solve this

//	4.	linked list technique:
//		- add dummy (not for this one)
//		- use prev, cur, next to reverse
//		- properly handle start/end condition (infinite loop)
