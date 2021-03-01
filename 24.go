// Given a linked list, swap every two adjacent nodes and return its head.

// You may not modify the values in the list's nodes, only nodes itself may be changed.

// Example:

// Given 1->2->3->4, you should return the list as 2->1->4->3.

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	prev, cur := dummy, head
	var next *ListNode

	for cur != nil && cur.Next != nil {
		next = cur.Next.Next
		prev.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next
		prev, cur = cur, cur.Next
	}

	return dummy.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	var next, next2 *ListNode
	head = cur.Next
	prev := &ListNode{}

	for cur != nil {
		next = cur.Next
		if next == nil {
			break
		}
		next2 = next.Next

		// swap
		cur.Next = next2
		next.Next = cur
		prev.Next = next

		// loop
		prev = cur
		cur = next2
	}

	return head
}

func swapPairs1(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	newHead := head.Next
	current := head
	var next *ListNode
	var next2 *ListNode
	var prev *ListNode

	for nil != current && nil != current.Next {
		next = current.Next
		next2 = next.Next

		next.Next = current
		current.Next = next2

		if nil != prev {
			prev.Next = next
		}

		prev = current
		current = next2
	}

	return newHead
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	arr := []int{2, 3, 4, 5, 6}

	current := head
	for _, v := range arr {
		item := &ListNode{
			Val:  v,
			Next: nil,
		}
		current.Next = item
		current = item
	}
	printList(head)
	newHead := swapPairs(head)
	printList(newHead)
}

func printList(head *ListNode) {
	current := head
	for nil != current {
		fmt.Printf("item: %v\n", current)
		current = current.Next
	}
}

//	problems
//	1.	takes me roughly 30 minutes to come up a solution, and forget to
//		update prev
