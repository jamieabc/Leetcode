package main

//Given a singly linked list, determine if it is a palindrome.
//
//Example 1:
//
//Input: 1->2
//Output: false
//
//Example 2:
//
//Input: 1->2->2->1
//Output: true
//
//Follow up:
//Could you do it in O(n) time and O(1) space?

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
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// find half of linked list
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var mid *ListNode
	if fast.Next == nil {
		mid = slow
	} else {
		mid = slow.Next
	}

	// reverse second half of the linked list
	// 1 2 3 3 2 1
	// 1 2 3 2 1
	prev := slow
	slow = slow.Next
	for slow != nil {
		tmp := slow.Next
		slow.Next = prev
		slow, prev = tmp, slow
	}
	slow = prev

	// compare again from head
	for ptr := head; ptr != mid; ptr, slow = ptr.Next, slow.Next {
		if ptr.Val != slow.Val {
			return false
		}
	}

	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	_, ok := traverse(head, head)

	return ok
}

// 1 2 1
// 1 2 2 1
// 1 2 3 3 2 1
func traverse(fast, slow *ListNode) (*ListNode, bool) {
	// odd length
	if fast.Next == nil {
		return slow.Next, true
	}

	// even length
	if fast.Next.Next == nil {
		return slow.Next.Next, slow.Val == slow.Next.Val
	}

	var going *ListNode
	var ok bool

	if fast.Next != nil && fast.Next.Next != nil {
		going, ok = traverse(fast.Next.Next, slow.Next)
	}

	if !ok {
		return nil, false
	}

	return going.Next, slow.Val == going.Val
}

func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var prev, ptr, ptr2, tmp, tmp2 *ListNode
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil {
		ptr2 = slow.Next
		slow = prev
	} else {
		ptr2 = slow.Next
	}

	// 1 - 1
	// 1 - 2 - 2 - 1
	// 1 - 2 - 3 - 2 - 1

	// separate link list
	slow.Next = nil

	// reverse new list
	tmp = ptr2.Next
	prev = ptr2
	prev.Next = nil
	for tmp != nil {
		tmp2 = tmp.Next
		tmp.Next = prev
		prev = tmp
		tmp = tmp2
	}
	ptr = head
	ptr2 = prev

	for ptr != nil {
		if ptr.Val != ptr2.Val {
			return false
		}
		ptr = ptr.Next
		ptr2 = ptr2.Next
	}
	return true
}

//	problems
//	1.	when using single pass, be careful about terminate condition
