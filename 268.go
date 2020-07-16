package main

// Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
//
// Example 1:
//
// Input: [3,0,1]
// Output: 2
// Example 2:
//
// Input: [9,6,4,2,3,5,7,0,1]
// Output: 8
// Note:
// Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?

func missingNumber(nums []int) int {
	size := len(nums)

	for i := 0; i < size; {
		if nums[i] == size {
			i++
		} else {
			if nums[i] != i {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			} else {
				i++
			}
		}
	}

	for i := range nums {
		if nums[i] != i {
			return i
		}
	}

	return size
}

//	problems
//	1.	from solution, xor same number twice becomes 0, so i ^ nums[i] to get the
//		missing number
