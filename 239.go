package main

// Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.
//
// Follow up:
// Could you solve it in linear time?
//
// Example:
//
// Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
//
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
//
//
// Constraints:
//
// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
// 1 <= k <= nums.length

func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	if size == 0 || k == 0 {
		return []int{}
	}

	maxToLeft, maxFromRight := make([]int, size), make([]int, size)
	for i := 0; i < size; i += k {
		rightBoundary := min(size-1, i+k-1)
		for j := 0; i+j <= rightBoundary; j++ {
			if j == 0 {
				maxToLeft[i] = nums[i]
				maxFromRight[rightBoundary] = nums[rightBoundary]
			} else {
				maxToLeft[i+j] = max(maxToLeft[i+j-1], nums[i+j])
				maxFromRight[rightBoundary-j] = max(maxFromRight[rightBoundary-j+1], nums[rightBoundary-j])
			}
		}
	}

	result := make([]int, 0)
	for i := 0; i <= size-k; i++ {
		result = append(result, max(maxFromRight[i], maxToLeft[i+k-1]))
	}

	return result
}

func maxSlidingWindow1(nums []int, k int) []int {
	size := len(nums)
	if size == 0 || k == 0 {
		return []int{}
	}

	result, dequeue := make([]int, 0), make([]int, 0)

	// 7 nums, idx: 0 - 6, k is 3, last index will be 4
	for i, num := range nums {
		for len(dequeue) > 0 && nums[dequeue[len(dequeue)-1]] < num {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue, i)

		if i >= k-1 {
			result = append(result, nums[dequeue[0]])

			// remove numbers not in range
			if dequeue[0] == i-k+1 {
				dequeue = dequeue[1:]
			}
		}
	}

	return result
}

//	problems
//	1.	inspired form sample code, store index to easier check

//	2.	from solution, dequeue means double-ended queue

//	3.	inspired from solution, use two dp array to find maximum.

//		separates original array into k size sub-array,
//		e.g. size 6 array into sub-array 0-2, 3-5

//		for each sub-array, creates 2 dp arrays:
//		- maximum number from start of each sub-array to index i
//		- maximum number from end of each sub-array to index i

//		e.g. k = 3, sub-array = [1, 3, -1]
//		left: [1] => [1, 3] => [1, 3, 3]
//		right: [-1] => [3, -1] => [3, 3, -1]

// 		w/ 2 dp arrays, for any given index range i - j, max number can come
//		from 2 sub-arrays

//	4.	inspired from https://leetcode.com/problems/sliding-window-maximum/discuss/65881/O(n)-solution-in-Java-with-two-simple-pass-in-the-array

//		author uses better naming of max_left & max_right
