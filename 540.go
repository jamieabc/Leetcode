package main

// You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once. Find this single element that appears only once.
//
// Follow up: Your solution should run in O(log n) time and O(1) space.
//
//
//
// Example 1:
//
// Input: nums = [1,1,2,3,3,4,4,8,8]
// Output: 2
//
// Example 2:
//
// Input: nums = [3,3,7,7,10,11,11]
// Output: 10
//
//
//
// Constraints:
//
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

// 3, 3, 7, 7, 10, 11, 11
func singleNonDuplicate(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	var i, j int
	// 1, 1, 2, 3, 3, 4, 4
	// 1, 1, 2, 2, 3, 3, 4, 5, 5,
	// before single num, increment at even
	// after single num, increment at odd
	for i, j = 0, length-1; i < j; {
		mid := i + (j-i)/2
		if (mid&1 == 0 && nums[mid] == nums[mid+1]) || (mid&1 == 1 && nums[mid] != nums[mid+1]) {
			i = mid + 1
		} else {
			j = mid
		}
	}

	return nums[i]
}

func singleNonDuplicate1(nums []int) int {
	var num int

	for _, n := range nums {
		num ^= n
	}

	return num
}

//	problems
//	1.	inspired from https://leetcode.com/problems/single-element-in-a-sorted-array/discuss/627786/C%2B%2B-O(log-n)-time-O(1)-space-or-Simple-and-clean-or-Use-xor-to-keep-track-of-odd-even-pair

//		unique number only appears at even index, and array length is
//		always odd

//		for even index before unique number: nums[mid] == nums[mid+1]
//		for even index after unique number: nums[mid] != nums[mid+1]

//		using +1 is easier because mid = i + (j-i)/2 tends to go to lower one
//		0 & 1 goes to 0, so plus is more easier to write code

//		also, when decreasing right boundary, next new value of right
//		boundary is mid not mid-1, cause it might cause unique number removed

//		so formula can be further reduce into considering 2 conditions

//		author uses ^1 to denote condition that equal, which is brilliant
