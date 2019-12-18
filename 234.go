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

func isPalindrome(head *ListNode) bool {
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
