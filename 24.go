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