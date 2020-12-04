package main

import (
	"math/rand"
	"time"
)

//Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.
//
//Follow up:
//What if the linked list is extremely large and its length is unknown to you? Could you solve this efficiently without using extra space?
//
//Example:
//
//// Init a singly linked list [1,2,3].
//ListNode head = new ListNode(1);
//head.next = new ListNode(2);
//head.next.next = new ListNode(3);
//Solution solution = new Solution(head);
//
//// getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
//solution.getRandom();

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	Node *ListNode
	Rand *rand.Rand
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{
		Node: head,
		Rand: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	var val int
	count := 1

	for node := this.Node; node != nil; node = node.Next {
		if this.Rand.Float64() < float64(1)/float64(count) {
			val = node.Val
		}
		count++
	}

	return val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

//	Notes
//	1.	inspired from solution

//		reservoir algorithm that can sample from unknown size

//	2.	reservoir sample needs to sample all possible items

//	3.	inspired from https://leetcode.com/problems/linked-list-random-node/discuss/85659/Brief-explanation-for-Reservoir-Sampling

//		author provides a good proof
