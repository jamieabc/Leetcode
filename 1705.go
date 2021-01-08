package main

// There is a special kind of apple tree that grows apples every day for n days. On the ith day, the tree grows apples[i] apples that will rot after days[i] days, that is on day i + days[i] the apples will be rotten and cannot be eaten. On some days, the apple tree does not grow any apples, which are denoted by apples[i] == 0 and days[i] == 0.
//
// You decided to eat at most one apple a day (to keep the doctors away). Note that you can keep eating after the first n days.
//
// Given two integer arrays days and apples of length n, return the maximum number of apples you can eat.
//
//
//
// Example 1:
//
// Input: apples = [1,2,3,5,2], days = [3,2,1,4,2]
// Output: 7
// Explanation: You can eat 7 apples:
// - On the first day, you eat an apple that grew on the first day.
// - On the second day, you eat an apple that grew on the second day.
// - On the third day, you eat an apple that grew on the second day. After this day, the apples that grew on the third day rot.
// - On the fourth to the seventh days, you eat apples that grew on the fourth day.
//
// Example 2:
//
// Input: apples = [3,0,0,0,0,2], days = [3,0,0,0,0,2]
// Output: 5
// Explanation: You can eat 5 apples:
// - On the first to the third day you eat apples that grew on the first day.
// - Do nothing on the fouth and fifth days.
// - On the sixth and seventh days you eat apples that grew on the sixth day.
//
//
//
// Constraints:
//
//     apples.length == n
//     days.length == n
//     1 <= n <= 2 * 104
//     0 <= apples[i], days[i] <= 2 * 104
//     days[i] = 0 if and only if apples[i] = 0.

type MinHeap [][]int // [ expire, count ]

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() []int   { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap [][]int // [ expire, count ]

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Peek() []int   { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func eatenApples(apples []int, days []int) int {
	size := len(apples)
	h := &MinHeap{}
	heap.Init(h)

	var count int

	for day := 0; day < size || h.Len() > 0; day++ {
		// add apple
		if day < size && apples[day] > 0 {
			heap.Push(h, []int{day + days[day] - 1, apples[day]})
		}

		// remove expired apples
		for h.Len() > 0 && h.Peek()[0] < day {
			heap.Pop(h)
		}

		// eat an apple if possible
		if h.Len() > 0 {
			count++
			h.Peek()[1]--

			if h.Peek()[1] == 0 {
				heap.Pop(h)
			}
		}
	}

	return count
}

func eatenApples1(apples []int, days []int) int {
	size := len(apples)
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	var count int

	for i := 0; i < int(4e4); i++ {
		// remove expired apples
		for minHeap.Len() > 0 && minHeap.Peek()[0] < i {
			heap.Pop(minHeap)
		}

		if i < size {
			// add apples
			if apples[i] > 0 {
				heap.Push(minHeap, []int{i + days[i] - 1, apples[i]})
			}

			if minHeap.Len() == 0 {
				continue
			}

			// eat an apple
			minHeap.Peek()[1]--
			count++

			// remove data if all apples in that range are eaten
			if minHeap.Peek()[1] == 0 {
				heap.Pop(minHeap)
			}
		} else {
			// eat an apple
			if minHeap.Len() > 0 {
				minHeap.Peek()[1]--
				count++

				if minHeap.Peek()[1] == 0 {
					heap.Pop(minHeap)
				}
			}
		}
	}

	return count
}

//	Notes
//	1.	my first intuition is to eat apples rotten earlier, the way to do is
//		sort by end time, then do it greedy

//		e.g. apples = [1, 2, 3, 5, 2]
//			 days =   [3, 2, 1, 4, 2]

//		day 0 ~ 3: 1 apple (idx 0)
//		day 1 ~ 3: 2 apple (idx 1)
//		day 2 ~ 3: 1 apple (idx 2)
// 		day 4 ~ 6: 2 apple (idx 4)
//		day 3 ~ 7: 4 apple (idx 3)

//		day 0: eat 1 from index 0
//		day 1, 2: eat 2 from index 1
//		day 4, 5: eat 2 from index 4
//		day 6: eat 1 from index 3

//		the problem here is that for index 3, it covers day 3, but algorithm
//		didn't find out

//	2.	inspired form https://youtu.be/Q8fjmUmwqRE?t=394

//		idea: eat apples rotten earlier
//		alex provides a very good naming: expire

//		overall behavior:
//		- if new apples grow, add them, then eats a latest expire apple
//		- if not apples grow, if any stored apples exist then eats one

//		it's different than what I think of: find apples exist time range and
//		do calculation

//	3.	inspired from https://leetcode.com/problems/maximum-number-of-eaten-apples/discuss/988328/C%2B%2B-Clean-and-Clear-Easy-to-understand-MinHeap

//		author provides very good logic of code
