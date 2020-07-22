package main

import (
	"container/heap"
	"math"
)

// Given a non-empty string s and an integer k, rearrange the string such that the same characters are at least distance k from each other.
//
// All input strings are given in lowercase letters. If it is not possible to rearrange the string, return an empty string "".
//
// Example 1:
//
// Input: s = "aabbcc", k = 3
// Output: "abcabc"
// Explanation: The same letters are at least distance 3 from each other.
// Example 2:
//
// Input: s = "aaabc", k = 3
// Output: ""
// Explanation: It is not possible to rearrange the string.
// Example 3:
//
// Input: s = "aaadbbcc", k = 2
// Output: "abacabcd"
// Explanation: The same letters are at least distance 2 from each other.

func rearrangeString(s string, k int) string {
	size := len(s)
	count, nextIndex := make([]int, size), make([]int, size)

	for i := range s {
		count[s[i]-'a']++
	}

	str := make([]byte, size)
	for i := 0; i < size; i++ {
		idx := nextChar(count, nextIndex, i)

		// cannot find any char meet criteria
		if idx == -1 {
			return ""
		}

		str[i] = byte('a' + idx)
		count[idx]--
		nextIndex[idx] = i + k
	}

	return string(str)
}

func nextChar(count, nextIndex []int, index int) int {
	maxCount := math.MinInt32
	idx := -1

	for i := range count {
		if count[i] > 0 && count[i] > maxCount && nextIndex[i] <= index {
			maxCount = count[i]
			idx = i
		}
	}

	return idx
}

type MaxHeap struct {
	Count int
	Char  byte
}

type MaxHeaps []MaxHeap

func (h MaxHeaps) Len() int { return len(h) }
func (h MaxHeaps) Less(i, j int) bool {
	if h[i].Count > h[j].Count {
		return true
	} else if h[i].Count < h[j].Count {
		return false
	}

	// make sure every time output is same order
	return h[i].Char < h[j].Char
}

func (h MaxHeaps) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeaps) Peek() MaxHeap { return h[0] }

func (h *MaxHeaps) Push(x interface{}) {
	*h = append(*h, x.(MaxHeap))
}

func (h *MaxHeaps) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func rearrangeString1(s string, k int) string {
	size := len(s)
	if k <= 1 || size <= 1 {
		return s
	}

	// count char occurrence
	counter := make([]int, 26)
	var maxOccurrence, maxOccurrenceCount int
	for i := range s {
		counter[s[i]-'a']++
		if counter[s[i]-'a'] > maxOccurrence {
			maxOccurrence = counter[s[i]-'a']
			maxOccurrenceCount = 1
		} else if counter[s[i]-'a'] == maxOccurrence {
			maxOccurrenceCount++
		}
	}

	// max character count cannot exceed limit
	if (maxOccurrence-1)*k+maxOccurrenceCount > size {
		return ""
	}

	// store char & occurrence count into max heap
	h := &MaxHeaps{}
	heap.Init(h)
	for idx := range counter {
		heap.Push(h, MaxHeap{
			Count: counter[idx],
			Char:  byte('a' + idx),
		})
	}

	str := make([]byte, size)
	queue := make([]MaxHeap, 0)

	// write char by occurrence count
	var idx int
	var popped MaxHeap
	for i, chunks := 0, size/k; i < chunks; i++ {
		for j := 0; j < k; j++ {
			popped = heap.Pop(h).(MaxHeap)
			str[idx] = popped.Char
			idx++
			popped.Count--

			// put un-finished char into queue
			if popped.Count > 0 {
				queue = append(queue, popped)
			}
		}

		// put un-finished char back to heap
		for len(queue) > 0 {
			heap.Push(h, queue[0])
			queue = queue[1:]
		}
	}

	// write remaining
	for ; idx < size; idx++ {
		popped = heap.Pop(h).(MaxHeap)
		str[idx] = popped.Char
	}

	return string(str)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	boundary condition: k <= 0, because there's a division, need to take
//		zero into consideration

//	2.	need to clarify distant k meaning

//	3.	inspired from https://leetcode.com/problems/rearrange-string-k-distance-apart/discuss/278269/Python-O(N)-solution

//		fail condition: (maxOccurrence-1) * k + (how many chars with same maxCount) > size
//		e.g. a _ _ a _ _ a
//			 a b _ b _ _ a b

//	4.	boundary condition: string length = 1 can fit into every k

//	5.	cannot think of a solution

//	6.	inspired from https://leetcode.com/problems/rearrange-string-k-distance-apart/discuss/83199/Greedy-Solution-Beats-95

//		in order to have every character separate k distance, the procedure as
//		follows:

//		separate string into several chunks (size / k)

//		for every chunk, write order from highest count to top k highest count,
//		then put char back to heap if still need to write additional

//		e.g. aaaabbbcccddde, k = 4
//		a: 4, b: 3, c: 3, d: 3, e: 1

//		_ _ _ _ , _ _ _ _ , _ _ _ _ , _ _
//		a b c d , _ _ _ _ , _ _ _ _ , _ _
//		a b c d , a b c d , _ _ _ _ , _ _
//		a b c d , a b c d , a b c d , _ _
//		a b c d , a b c d , a b c d , a e

//		another post https://leetcode.com/problems/rearrange-string-k-distance-apart/discuss/161489/Why-does-it-work-EDF-in-disguise!

// 		EDF (earliest deadline first) for scheduling, same as here

//		tc: O(n log n)

//	7.	inspired from sample code, a brilliant solution

//		use two arrays to store count & next valid index, everytime sweep find
//		max count as next character to write
//		e.g. aaaabbbcccddde, k = 4

//				a	b	c	d	e
//		count	4	3	3	3	1
//		valid	0	0	0	0	0

//		result	a
//				a	b	c	d	e
//		count	3	3	3	3	1
//		valid	4	0	0	0	0

//		result	ab
//				a	b	c	d	e
//		count	3	2	3	3	1
//		valid	4	5	0	0	0

//		result	abc
//				a	b	c	d	e
//		count	3	2	2	3	1
//		valid	4	5	6	0	0

//		result	abcd
//				a	b	c	d	e
//		count	3	2	2	2	1
//		valid	4	5	6	7	0

//		result	abcda
//				a	b	c	d	e
//		count	2	2	2	2	1
//		valid	8	5	6	7	0

//		result	abcdab
//				a	b	c	d	e
//		count	2	1	2	2	1
//		valid	8	9	6	7	0

//		result	abcdabc
//				a	b	c	d	e
//		count	2	1	1	2	1
//		valid	8	9	10	7	0

//		result	abcdabcd
//				a	b	c	d	e
//		count	2	1	1	1	1
//		valid	8	9	10	11	0

//		result	abcdabcda
//				a	b	c	d	e
//		count	1	1	1	1	1
//		valid	12	9	10	11	0

//		result	abcdabcdab
//				a	b	c	d	e
//		count	1	0	1	1	1
//		valid	12	13	10	11	0

//		result	abcdabcdabc
//				a	b	c	d	e
//		count	1	0	0	1	1
//		valid	12	13	14	11	0

//		result	abcdabcdabcd
//				a	b	c	d	e
//		count	1	0	0	0	1
//		valid	12	13	14	15	0

//		result	abcdabcdabcda
//				a	b	c	d	e
//		count	0	0	0	0	1
//		valid	16	13	14	15	0

//		result	abcdabcdabcdae
//				a	b	c	d	e
//		count	0	0	0	0	0
//		valid	16	13	14	15	17
