package main

// Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
//
// Example:
//
// Input:  [1,2,3,4]
// Output: [24,12,8,6]
//
// Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array (including the whole array) fits in a 32 bit integer.
//
// Note: Please solve it without division and in O(n).
//
// Follow up:
// Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i, tmp := 0, 1; i < len(nums); i++ {
		result[i] = tmp
		tmp *= nums[i]
	}

	for i, tmp := len(nums)-1, 1; i >= 0; i-- {
		result[i] *= tmp
		tmp *= nums[i]
	}

	return result
}

//	problems
//	1.	too slow...

//	2.	from reference https://leetcode.com/problems/product-of-array-except-self/discuss/65632/My-solution-beats-100-java-solutions

//		The idea is that order of multiplication doesn't matter. First find
//		multiplication before self from index 0, then find
//		multiplication after self start from index length-1

//		e.g. array of 7, 2, 10, 3
//		normal direction:   1, 7, 14, 140
//		backward direction: 60 , 30 , 3 ,1
//		multiply both: 		60, 210, 42, 140

//		so clever, I spend 30 minutes but cannot think of any clue....the
//		solution gives me the feeling of problem 370
