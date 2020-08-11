package main

import (
	"math"
	"sort"
)

//Given an array A of integers, we must modify the array in the following way: we choose an i and replace A[i] with -A[i], and we repeat this process K times in total.  (We may choose the same index i multiple times.)
//
//Return the largest possible sum of the array after modifying it in this way.
//
//
//
//Example 1:
//
//Input: A = [4,2,3], K = 1
//Output: 5
//Explanation: Choose indices (1,) and A becomes [4,-2,3].
//
//Example 2:
//
//Input: A = [3,-1,0,2], K = 3
//Output: 6
//Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].
//
//Example 3:
//
//Input: A = [2,-3,-1,5,-4], K = 2
//Output: 13
//Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
//
//
//
//Note:
//
//    1 <= A.length <= 10000
//    1 <= K <= 10000
//    -100 <= A[i] <= 100

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)

	for i := 0; i < len(A) && A[i] < 0 && K > 0; i, K = i+1, K-1 {
		A[i] = -A[i]
	}

	var sum int
	minNum := math.MaxInt32
	for _, n := range A {
		sum += n
		minNum = min(minNum, n)
	}

	if K&1 > 0 {
		return sum - 2*minNum
	}
	return sum
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() int          { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func largestSumAfterKNegations1(A []int, K int) int {
	h := &MinHeap{}
	heap.Init(h)

	var sum int
	minNum := math.MaxInt32
	for _, n := range A {
		sum += n
		if n <= 0 {
			heap.Push(h, n)
		} else {
			minNum = min(minNum, n)
		}
	}

	for ; K > 0 && h.Len() > 1; K-- {
		popped := heap.Pop(h).(int)
		sum -= 2 * popped
	}

	if K > 0 {
		if h.Len() > 0 {
			if -h.Peek() <= minNum {
				// -1, 2
				if K&1 > 0 {
					sum -= 2 * heap.Peek()
				}
			} else {
				// -2, 1
				sum -= 2 * heap.Peek()
				if K&1 == 0 {
					sum -= 2 * minNum
				}
			}
		} else {
			if K&1 > 0 {
				sum -= 2 * minNum
			}
		}
	}

	return sum
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	problems
//	1.	when negate operation remains, choose between negative number
//		& smallest positive number.
//		There might have 2 conditions:
//		- abs(negative number) <= min positive number
//		- abs(negative number) > min positive number

//	2.	inspired from https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/discuss/252254/JavaC%2B%2BPython-Sort

//		it's more readable because there's no too many if statements.

//		I think it as follows: negate remains, so need to check which
//		number (positive or negative) to negate.

//		lee thinks when negate remains, which means all numbers are
//		positive, so choose minimum number to do operation

//		this is more beautiful not only because of readability, but also
//		my think is step-by-step, but lee is more systematically

//	3.	inspired from https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/discuss/252849/C%2B%2BJava-O(n)-or-O(1)

//		author noticed that number has a fixed range in -100 ~ 100,
//		which can also use bucket sort
