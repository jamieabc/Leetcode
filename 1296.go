package main

import "sort"

// Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into sets of k consecutive numbers
// Return True if its possible otherwise return False.
//
//
//
// Example 1:
//
// Input: nums = [1,2,3,3,4,4,5,6], k = 4
// Output: true
// Explanation: Array can be divided into [1,2,3,4] and [3,4,5,6].
//
// Example 2:
//
// Input: nums = [3,2,1,2,3,4,3,4,5,9,10,11], k = 3
// Output: true
// Explanation: Array can be divided into [1,2,3] , [2,3,4] , [3,4,5] and [9,10,11].
//
// Example 3:
//
// Input: nums = [3,3,2,2,1,1], k = 3
// Output: true
//
// Example 4:
//
// Input: nums = [1,2,3,4], k = 3
// Output: false
// Explanation: Each array should be divided in subarrays of size 3.
//
//
//
// Constraints:
//
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= nums.length
//
// Note: This question is the same as 846: https://leetcode.com/problems/hand-of-straights/

// is size of nums or k too large, sort is not practical, iterate through
// nums to get count, tc: O(n + m log m), n: size of nums, m: size of
// different number
func isPossibleDivide(nums []int, k int) bool {
	// each number count
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
	}

	numbers := make([]int, 0)
	for num := range counter {
		numbers = append(numbers, num)
	}
	sort.Ints(numbers)

	var group int
	events := make([]int, 0)

	for _, num := range numbers {
		if len(events) == 0 {
			for j := 0; j < counter[num]; j++ {
				events = append(events, num+k-1)
			}
			group = counter[num]
			continue
		}

		if counter[num] < group {
			return false
		}

		if counter[num] > group {
			for j := 0; j < counter[num]-group; j++ {
				events = append(events, num+k-1)
			}
			group = counter[num]
		}

		for len(events) > 0 {
			if events[0] == num {
				events = events[1:]
				group--
			} else {
				break
			}
		}
	}

	return len(events) == 0
}

func isPossibleDivide1(nums []int, k int) bool {
	size := len(nums)
	if size < k || size%k != 0 {
		return false
	}

	sort.Ints(nums)

	// a queue to store event when group should decrease
	events := []int{nums[0] + k - 1}

	var duplicates int
	group := 1
	prev := -1

	for i := 0; i < size; i++ {
		// find duplicates
		for duplicates = 1; i < size-1 && nums[i+1] == nums[i]; i++ {
			duplicates++
		}

		// check consecutive
		if prev != -1 && len(events) > 0 && nums[i]-prev != 1 {
			return false
		}

		// not enough duplicate number to distribute into groups
		if duplicates < group {
			return false
		}

		// additional groups are needed
		if duplicates > group {
			extra := duplicates - group
			for j := 0; j < extra; j++ {
				events = append(events, nums[i]+k-1)
			}
			group += extra
		}

		// if any group reaches k count, remove it
		for len(events) > 0 {
			if nums[i] == events[0] {
				events = events[1:]
				group--
			} else {
				break
			}
		}

		prev = nums[i]
	}

	return len(events) == 0
}

//  problems
//  1.  for this algorithm to work, initial status of events & group should
//      be setup

//      also, use nums[i] to check group reach limit, instead of i

//	2.	consecutive, could be +1 or -1

//	3.	when split into sub-array, order could be changed, as long as
//		maintains k size

//		this question should be ask before writing code, it's really
//		crucial

//	4.	inspired from https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/discuss/470238/JavaC%2B%2BPython-Exactly-Same-as-846.-Hand-of-Straights

//		if k or number size is too large, how to do faster?

//		if size too large, then sort takes lots of time, use hashmap to
//		store all number occurrence count, sort distinct number, and
//		use same technique as events & decrease while iterating numbers

//	5.	inspired from https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/discuss/457728/O(n)-Python-and-C%2B%2B

//		author provides an interesting idea, when removing number, no need
//		to sort distinct number, randomly pick a number and keep decreasing
//		until no count, the that number + 1 is the root.

//		e.g. [3,2,1,5,6,7,7,8,9], pick 8, since 7, 6, 5 all exist but 4
//		not exist, so start from 5 and remove k elements (5, 6, 7)
