package main

// The set S originally contains numbers from 1 to n. But unfortunately, due to the data error, one of the numbers in the set got duplicated to another number in the set, which results in repetition of one number and loss of another number.
//
//Given an array nums representing the data status of this set after the error. Your task is to firstly find the number occurs twice and then find the number that is missing. Return them in the form of an array.
//
//Example 1:
//
//Input: nums = [1,2,2,4]
//Output: [2,3]
//
//Note:
//
//    The given array size will in the range [2, 10000].
//    The given array's numbers won't have any order.

func findErrorNums(nums []int) []int {
	length := len(nums)
	var duplicate, missing, xor int

	for i := 0; i < length; i++ {
		if nums[abs(nums[i])-1] < 0 {
			duplicate = abs(nums[i])
		} else {
			nums[abs(nums[i])-1] *= -1
		}
		xor = xor ^ (i + 1) ^ abs(nums[i])
	}

	missing = xor ^ duplicate

	return []int{duplicate, missing}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// problems
// 	1. 	too slow, sort of n log n, try use sum & subtract, it will occupy
//		some memory (map)
//	2.	use additional array index to denote if a number already exists, and
//		not updated one is the missing one
//	3.	optimization, additional array is used, just to store if a number
//		exist/missing. If there's another way to denote this information,
// 		then additional memory can be saved.
//
//		This is really smart, use index to denote if a number exist as
//		previous way, but by negate existing array. If a number is already
//		negative, then that index is the duplicate one.
//		And, for the number that is positive, that index is missing one.
//		So clever.......
//	4.	using XOR, same number XOR will become 0 (no effect)
//		a ^ b ^ b = a
//		(1 ^ 2 ^ 3 ^ ... ^ n) ^ (1 ^ 2 ^ 3 ^ ... ^ n) = 0
//   	if some number b is change to c, after XOR, final result = b ^ c
//		if I can find duplicates, missing is result ^ duplicate
