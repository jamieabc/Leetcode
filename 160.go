package main

//Write a program to find the node at which the intersection of two singly linked lists begins.
//
//For example, the following two linked lists:
//
//begin to intersect at node c1.
//
//
//
//Example 1:
//
//Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
//Output: Reference of the node with value = 8
//Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
//
//
//
//Example 2:
//
//Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
//Output: Reference of the node with value = 2
//Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [0,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.
//
//
//
//Example 3:
//
//Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
//Output: null
//Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
//Explanation: The two lists do not intersect, so return null.
//
//
//
//Notes:
//
//    If the two linked lists have no intersection at all, return null.
//    The linked lists must retain their original structure after the function returns.
//    You may assume there are no cycles anywhere in the entire linked structure.
//    Your code should preferably run in O(n) time and use only O(1) memory.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	ptr1, ptr2 := headA, headB

	for ptr1 != ptr2 {
		if ptr1 == nil {
			ptr1 = headB
		} else {
			ptr1 = ptr1.Next
		}

		if ptr2 == nil {
			ptr2 = headA
		} else {
			ptr2 = ptr2.Next
		}
	}

	return ptr1
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil {
		return nil
	}

	// only one node for headA
	if headA.Next == nil {
		// check if any B visited
		ptr := headB
		for ptr != nil {
			if ptr == headA {
				return ptr
			}
			ptr = ptr.Next
		}
		return nil
	}

	// go to end of linked list
	end := headA
	for end.Next != nil {
		end = end.Next
	}

	// create cycle
	end.Next = headB

	// find cycle start
	slow, fast := headA.Next, headA.Next.Next
	for slow != nil && fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// if any pointer go to nil, there's no cycle
	if slow == nil || fast == nil || fast.Next == nil {
		end.Next = nil
		return nil
	}

	fast = headA
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	// remove cycle
	end.Next = nil
	return slow
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}

	lengthA := listLength(headA)
	lengthB := listLength(headB)

	var longer, shorter *ListNode
	var diff int

	// remove un-necessary nodes, since intersection means last few parts of
	// elements are same
	if lengthA >= lengthB {
		longer = headA
		shorter = headB
		diff = lengthA - lengthB
	} else {
		longer = headB
		shorter = headA
		diff = lengthB - lengthA
	}

	for ; diff > 0; diff-- {
		longer = longer.Next
	}

	for longer != nil {
		if longer == shorter {
			return longer
		}
		longer = longer.Next
		shorter = shorter.Next
	}

	return nil
}

func listLength(head *ListNode) int {
	length := 0
	tmp := head
	for tmp != nil {
		length++
		tmp = tmp.Next
	}
	return length
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/intersection-of-two-linked-lists/discuss/49785/Java-solution-without-knowing-the-difference-in-len!

//		the problem is about different length of linked list. If two
//		linked list with same length, then start from two heads, compare
//		each node one by one to find start of intersection.

//		author provides a brilliant way, say two linked list with size
//		a & b, a+b = b+a, which means start from headA and to headB, will
//		intersect to another one start from headB and to headA.

//		another small trick is that when traversing node, check null
//		at next iteration, which helps to determine non-intersection
//		condition
