// Reverse a singly linked list.

// Example:

// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
// Follow up:

// A linked list can be reversed either iteratively or recursively. Could you implement both?

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	current := head
	var prev *ListNode
	next := head.Next

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	arr := []int{2, 3, 4, 5}

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
	newHead := reverseList(head)
	printList(newHead)
}

func printList(head *ListNode) {
	i := 0
	current := head
	for nil != current {
		fmt.Printf("current: %v\n", current)
		i++
		current = current.Next
	}
}
