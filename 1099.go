package main

import "sort"

func twoSumLessThanK(A []int, K int) int {
	size := len(A)
	if size <= 1 {
		return -1
	}

	buckets := make([]int, 1001)
	for _, num := range A {
		buckets[num]++
	}

	var maxSum, low, high int
	for low, high = 1, len(buckets)-1; low < high; {
		if buckets[low] == 0 {
			low++
			continue
		}

		if buckets[high] == 0 {
			high--
			continue
		}

		sum := low + high

		if sum >= K {
			high--
		} else {
			maxSum = max(maxSum, sum)
			low++
		}
	}

	if buckets[low] >= 2 && low*2 < K {
		maxSum = max(maxSum, low*2)
	}

	if maxSum == 0 {
		return -1
	}

	return maxSum
}

func twoSumLessThanK1(A []int, K int) int {
	sort.Ints(A)
	var maxSum, sum, low, high int

	for low, high = 0, len(A)-1; low < high; {
		sum = A[low] + A[high]

		if sum >= K {
			high--
		} else {
			maxSum = max(maxSum, sum)
			low++
		}
	}

	if maxSum == 0 {
		return -1
	}

	return maxSum
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	forget condition when not found

//	2	when doing binary search, forget about one condition that high keeps
//		going down, so need to check for low pointer sum

//	3.	inspired from https://leetcode.com/problems/two-sum-less-than-k/discuss/322931/Java-Sort-then-push-from-two-ends.

//		it's slow, because I don't really need to fix one number and search for
//		the other one, just use two pointers to find possible numbers

//	4.	inspired from https://leetcode.com/problems/two-sum-less-than-k/discuss/326486/C%2B%2B-3-solutions

//		since A[i] ranges from 1 ~ 1000, use an array to store, like bucket sort
