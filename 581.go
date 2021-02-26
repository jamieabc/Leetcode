package main

import "math"

//Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.
//
//You need to find the shortest such subarray and output its length.
//
//Example 1:
//
//Input: [2, 6, 4, 8, 10, 9, 15]
//Output: 5
//Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
//
//Note:
//
//    Then length of the input array is in range [1, 10,000].
//    The input array may contain duplicates, so ascending order here means <=.

func findUnsortedSubarray(nums []int) int {
	largest, smallest := math.MinInt32, math.MaxInt32
	start, end := -1, -1
	size := len(nums)

	for i := range nums {
		largest = max(largest, nums[i])
		if nums[i] < largest {
			end = i
		}
	}

	for i := size - 1; i >= 0; i-- {
		smallest = min(smallest, nums[i])
		if nums[i] > smallest {
			start = i
		}
	}

	if start == -1 && end == -1 {
		return 0
	}

	return end - start + 1
}

// tc: O(n), sc: O(1)
func findUnsortedSubarray3(nums []int) int {
	smallest, largest := math.MaxInt32, math.MinInt32
	var flag bool
	size := len(nums)

	// forward find initial point start to fall, also any number greater
	for i := 0; i < size-1; i++ {
		if nums[i] > nums[i+1] {
			flag = true
		}

		// becareful about boundary, since nums[i] > nums[i+1], nums[i+1] is
		// smaller than nums[i]
		if flag {
			smallest = min(smallest, nums[i+1])
		}
	}

	// backward find initial point start to rise, also find any number smaller
	flag = false
	for i := size - 1; i > 0; i-- {
		if nums[i-1] > nums[i] {
			flag = true
		}

		// becarefu about boundary, since nums[i-1] > nums[i], nums[i-1] is
		// larger than nums[i]
		if flag {
			largest = max(largest, nums[i-1])
		}
	}

	// start backward, find any number < largest
	var start, end int
	for end = size - 1; end >= 0 && nums[end] >= largest; end-- {
	}

	// start forward, find any number > largest
	for start = 0; start < size && nums[start] <= smallest; start++ {
	}

	if start > end {
		return 0
	}

	return end - start + 1
}

func findUnsortedSubarray2(nums []int) int {
	stack := make([]int, 0)
	size := len(nums)
	largest := math.MinInt32
	start, end := size+1, -1

	for i := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			end = i
			start = min(start, stack[len(stack)-1])
			largest = max(largest, nums[stack[len(stack)-1]])
			stack = stack[:len(stack)-1]
		}

		if nums[i] < largest {
			end = i
		}

		stack = append(stack, i)
	}

	if start == size+1 {
		return 0
	}
	return end - start + 1
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

// [2, 2, 1, 3, 4]
// [1, 2, 3, 4, 5]
// [1, 2, 4, 5, 3]
// [1, 3, 4, 2, 5]
func findUnsortedSubarray1(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var left, right int

	for left = 0; left < len(nums)-1; left++ {
		// encounter problem
		if nums[left] > nums[left+1] {
			// if it's same number, traverse back to first number
			for left > 0 && nums[left] == nums[left-1] {
				left--
			}
			break
		}
	}

	for right = len(nums) - 1; right >= 1; right-- {
		// encounter problem
		if nums[right] < nums[right-1] {
			// if it's same number, traverse back to first number
			for right < len(nums)-1 && nums[right] == nums[right+1] {
				right++
			}
			break
		}
	}

	min := nums[left]
	max := nums[right]
	for i := left; i <= right; i++ {
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}

	// for ascending, left-1 must be smaller than every number in range
	for left > 0 {
		if nums[left-1] > min {
			left--
		} else {
			break
		}
	}

	for right < len(nums)-1 {
		if nums[right+1] < max {
			right++
		} else {
			break
		}
	}

	if left > right {
		return 0
	}

	return right - left + 1
}

//	Notes
//	1.	I thought it's a stack problem, because if I keep stack in ascending
//		order, any backward number smaller means range should be sorted.

//		so, stack can keep numbers in ascending order, it help to find backward
//		number that are smaller

//		but, I forget to consider popped number, because if popped number is
//		larger than backward numbers, those numbers should also be sorted

//		this problem wants to make sure two things, e.g. [A] [B] [C], B is
//		unsorted range, two properties should be kept:
//		- any number in A should be smaller than B, any number in C should be
//		  larger than B
//		- any number in A, C are sorted

//		only using stack is not enough, need additional checking

//	2.	inspired from solution, there's another way to solve it.

//		to find initial range, start from 0, find position where number starts
//		to fall; start from last, find position where number starts to rise

//		these two points are pivot points, use those points to find any before
//		numbers are larger than rising point, and to find any after numbers are
//		smaller than falling number

//		e.g.
//
//										8
//					7
//				5
//						4
//			3				3		3
//								2
//		1
//								^	3 > 2, start to rise
//					^  4 < 7, first position start to fall

//			^	3 > 2, 3 is the start of unsorted range, because need to make
//				2 in order

//									^	3 < 7, is the end of unsorted range

//		this is a really brilliant way to solve it, breakdown problem into
//		smaller parts, and each part is easy to understand

//		need to check number within range, because if there maximum number in
//		range, it might become right boundary

//	3.	inspired from https://leetcode.com/problems/shortest-unsorted-continuous-subarray/discuss/103057/Java-O(n)-Time-O(1)-Space/106306

//		a beautiful tc O(n) and sc O(1), it observes that a range to be
//		sorted when any number smaller than previous seen max forward from
//		start, and any number larger than previous seen min backward from end

//		very beautiful and elegant...

//	4.	the other thing inspired from solution, it's better to think sequence
//		by graph, like heart beat graph, makes easier to understand
