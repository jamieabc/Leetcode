package main

import "fmt"

// Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

// To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

// Note: Do not modify the linked list.

// Example 1:

// Input: head = [3,2,0,-4], pos = 1
// Output: tail connects to node index 1
// Explanation: There is a cycle in the linked list, where tail connects to the second node.

// Example 2:

// Input: head = [1,2], pos = 0
// Output: tail connects to node index 0
// Explanation: There is a cycle in the linked list, where tail connects to the first node.

// Example 3:

// Input: head = [1], pos = -1
// Output: no cycle
// Explanation: There is no cycle in the linked list.

// Follow up:

// Can you solve it without using extra space?

type ListNode struct {
	Val  int
	Next *ListNode
}

// this problem can be solved by some algorithms (https://en.wikipedia.org/wiki/Floyd%27s_cycle-finding_algorithm)
// what I do refer to indian youtube explanation
// Each round, fast pointer goes 2 steps, slow pointer goes 1 step
// m: distance of cycle start point, relative to head of list
// k: fast and slow pointer meets
// l: length of cycle
// when fast pointer meets slow pointer, fast pointer runs p rounds, slow pointer runs q rounds
// because fast pointer is 2 times fast than slow pointer, the time when two pointers meet, fast pointer
// runs 2 times distance than slow pointer
// 2 * (m + k + p * l) = m + k + q * l
// m + k = (q - 2 * p) * l
// m + k is the multiple of l
// when two pointers meet, slow pointer is at distance far from start of cycle k distance, and m + k is k multiple
// so additional m steps will make slow pointer back to start of cycle, thus why put fast pointer to head
// because m is the distance from head to start of cycle
func detectCycle(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return nil
	}

	if head == head.Next.Next {
		return head
	}

	ptr1 := head
	ptr2 := head

	isCycle := false

	for nil != ptr2.Next && nil != ptr2.Next.Next {
		if ptr1 != head && ptr1 == ptr2 {
			isCycle = true
			break
		}
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next.Next
	}

	if !isCycle {
		return nil
	}

	ptr2 = head

	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}

func main() {
	head := &ListNode{
		Val:  3,
		Next: nil,
	}

	arr := []int{2, 0, -4}

	current := head
	var second *ListNode

	for _, v := range arr {
		item := &ListNode{
			Val:  v,
			Next: nil,
		}

		if nil == head.Next {
			second = item
		}

		current.Next = item
		current = item
	}

	current.Next = second

	fmt.Printf("cycle: %v\n", detectCycle(head))
}
