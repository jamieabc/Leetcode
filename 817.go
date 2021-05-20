package main

// We are given head, the head node of a linked list containing unique integer values.
//
// We are also given the list nums, a subset of the values in the linked list.
//
// Return the number of connected components in nums, where two values are connected if they appear consecutively in the linked list.
//
// Example 1:
//
// Input:
// head: 0->1->2->3
// nums = [0, 1, 3]
// Output: 2
// Explanation:
// 0 and 1 are connected, so [0, 1] and [3] are the two connected components.
//
// Example 2:
//
// Input:
// head: 0->1->2->3->4
// nums = [0, 3, 1, 4]
// Output: 2
// Explanation:
// 0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3, 4] are the two connected components.
//
// Note:
//
// If n is the length of the linked list given by head, 1 <= n <= 10000.
// The value of each node in the linked list will be in the range [0, n - 1].
// 1 <= nums.length <= 10000.
// nums is a subset of all values in the linked list.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func numComponents(head *ListNode, nums []int) int {
	var size int
	for node := head; node != nil; node = node.Next {
		size++
	}

	buckets := make([]bool, size+1)

	for _, n := range nums {
		buckets[n] = true
	}

	var count int

	for node := head; node != nil; {
		// forward until match found
		for ; node != nil && !buckets[node.Val]; node = node.Next {
		}

		if node == nil {
			break
		}

		count++

		// forward until non-match found
		for ; node != nil && buckets[node.Val]; node = node.Next {
		}
	}

	return count
}

func numComponents1(head *ListNode, nums []int) int {
	table := make(map[int]bool)
	for _, n := range nums {
		table[n] = true
	}

	var connected int

	for node := head; node != nil; {
		// forward until match found
		for node != nil && !table[node.Val] {
			node = node.Next
		}

		if node == nil {
			break
		}

		connected++

		// forward until non-match found
		for node != nil && table[node.Val] {
			node = node.Next
		}
	}

	return connected
}

//	Notes
//	1.	inspired from sample code, then I notice that each value in nodes
//		is in [0, n-1], which makes it possible to solve by bucket sort
