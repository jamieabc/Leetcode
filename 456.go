package main

import "math"

// Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < nums[j].
//
// Return true if there is a 132 pattern in nums, otherwise, return false.
//
// Follow up: The O(n^2) is trivial, could you come up with the O(n logn) or the O(n) solution?
//
//
// Example 1:
//
//
// Input: nums = [1,2,3,4]
// Output: false
// Explanation: There is no 132 pattern in the sequence.
//
//
// Example 2:
//
//
// Input: nums = [3,1,4,2]
// Output: true
// Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
//
//
// Example 3:
//
//
// Input: nums = [-1,3,2,0]
// Output: true
// Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].
//
//
//
// Constraints:
//
//
// 	n == nums.length
// 	1 <= n <= 10^4
// 	-10^9 <= nums[i] <= 10^9

// tc: O(n)
func find132pattern(nums []int) bool {
	size := len(nums)
	if size <= 2 {
		return false
	}

	smallest := make([]int, size)
	num := math.MaxInt32
	for i := range nums {
		num = min(num, nums[i])
		smallest[i] = num
	}

	stack := make([]int, 0)
	for i := size - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= smallest[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			return true
		}
		stack = append(stack, nums[i])
	}

	return false
}

type intervals struct {
	start, end  int
	left, right *intervals
}

func (i *intervals) add(start, end int) bool {
	if end > i.start && end < i.end {
		return true
	}

	if start < i.start {
		if i.left == nil {
			i.left = &intervals{
				start: start,
				end:   end,
			}
		} else {
			return i.left.add(start, end)
		}
	} else {
		if i.right == nil {
			i.right = &intervals{
				start: start,
				end:   end,
			}
		} else {
			return i.right.add(start, end)
		}
	}

	return false
}

// tc: O(n log(n))
func find132pattern3(nums []int) bool {
	size := len(nums)
	if size <= 2 {
		return false
	}

	smallestToLeft := make([]int, size)
	small := nums[0]
	for i := 1; i < size; i++ {
		smallestToLeft[i] = small
		small = min(small, nums[i])
	}

	var root *intervals

	for i := 1; i < size; i++ {
		start, end := smallestToLeft[i], nums[i]

		if start >= end {
			continue
		}

		if root == nil {
			root = &intervals{
				start: start,
				end:   end,
			}
		} else {
			if root.add(start, end) {
				return true
			}
		}
	}
	return false
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

// tc: O(n^2)
func find132pattern2(nums []int) bool {
	size := len(nums)
	if size <= 2 {
		return false
	}

	smallest := math.MaxInt32
	for i := 0; i < size-1; i++ {
		smallest = min(smallest, nums[i])

		for j := i + 1; j < size; j++ {
			if nums[j] > smallest && nums[j] < nums[i+1] {
				return true
			}
		}
	}

	return false
}

// tc: O(n^3)
func find132pattern1(nums []int) bool {
	size := len(nums)
	if size <= 2 {
		return false
	}

	for i := range nums {
		for j := i + 1; j < size; j++ {
			if nums[i] < nums[j] {
				for k := j + 1; k < size; k++ {
					if nums[k] < nums[j] && nums[k] > nums[i] {
						return true
					}
				}
			}
		}
	}

	return false
}

//	Notes
//	1.	for a give number, the goal is not to find smallest to left of this number
//		and largest to right of this number, the point is about to find if a number
//		on left is smaller than current, smaller that right of this number, and
//		to find if a number to right is smaller than current number and larger
//		than left number

//		e.g.               [42, 43, 6, 12, 3, 4, 6, 11, 20]
//		smallest to left    X   42  42  6  6  3  3   3   3
//		largest to right    43  20  20  20  20  20  20   X

//		if only care about smallest & largest different than self,  than sequence
//		6, 12, 11 will not be considered, because taking 6 (index 6), smallest
//		to left is 3, largest to right is 20, doesn't fits, but it is wrong

//	2.	inspired from solution, when smallest to left is know, problem becomes
//		how to find right number fits into range. the range is smallest to
//		left ~ current, if traversing backward from right, while numbers larger
//		than smallest to left, than keep it, because this number might be a possible
//		one. if this number smaller than smallest to left, means it won't be a
//		possible anymore, remove it

//	3.	index i < j < k
//		value: n[i] < n[k] < n[j]

//		n[i] ~ n[j] defines the interval range, try to find k satisfies this
//		range
//
//		to find k, it's best to increase the search range (smallest n[i] to
//		left), this can be done by one pass iteration

//		after range is found, the point is to find k

//		honestly i am not sure why, but stack can be used to solve the problem

//		start backward, keep array in descending order, for any stack value
//		< n[i] (smallest to left), it's never being used (because this is
//		smallest to left, all other numbers will be larger, and n[k] should be in
//		range of n[i] ~ n[j], so smaller thant current n[i] will never be used
//		again)

//		the goal is to preserve higher number as possible, and stack can be
//		applied

//		for n[k] there are 2 conditions, either n[k] >= n[j] or n[k] < n[j]
//		if n[k] > n[j] => this causes stack to be descending
//		if n[k] > n[j] => goal found

//		so, in general, this operation will maintain stack in descending order

//	4. 	inspired from https://leetcode.com/problems/132-pattern/discuss/94089/Java-solutions-from-O(n3)-to-O(n)-for-%22132%22-pattern-(updated-with-one-pass-slution)

//		author provides a very good explanation of why stack is being used

//		there's also an one-pass solution, but i didn't take time to read it
