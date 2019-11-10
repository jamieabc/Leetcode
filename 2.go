package main

//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Example:
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// both empty
	if l1 == nil && l2 == nil {
		return nil
	}

	//	l1 empty
	if l1 == nil {
		return l2
	}

	// l2 empty
	if l2 == nil {
		return l1
	}

	var carry, sum int
	result := &ListNode{}
	cur := result

	for l1 != nil || l2 != nil || carry != 0 {
		sum = carry
		carry = 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum <= 9 {
			cur.Val = sum
		} else {
			carry = 1
			cur.Val = sum % 10
		}

		if l1 != nil || l2 != nil || carry != 0 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return result
}
