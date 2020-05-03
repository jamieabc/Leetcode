package main

// Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.
//
// Example:
//
// Input:  [1,2,1,3,2,5]
// Output: [3,5]
//
// Note:
//
//     The order of the result is not important. So in the above example, [5, 3] is also correct.
//     Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?

func singleNumber(nums []int) []int {
	var totalXOR int
	for _, i := range nums {
		totalXOR ^= i
	}

	// get last bit of XOR result
	mask := totalXOR & (-totalXOR)

	var first, second int

	for _, i := range nums {
		if mask&i == 0 {
			first ^= i
		} else {
			second ^= i
		}
	}

	return []int{first, second}
}

//	problems
//	1.	Cannot think of any clue, reference from https://leetcode.com/problems/single-number-iii/discuss/68900/Accepted-C%2B%2BJava-O(n)-time-O(1)-space-Easy-Solution-with-Detail-Explanations

//		First clue comes from appear twice => XOR same number twice become 0,
//		so XOR all numbers get a result of diff1 XOR diff2.

//		With this information, separate all numbers into 2 groups by XOR result,
//		the index of first one (can choose other, 1 means on the position,
//		bits are different so numbers can be grouped)
//		- all numbers has 1 on first index
//		- al numbers has 0 on first index

//		for example, for array of 1, 2, 3, 5, 1, 3
//		all numbers XOR = 2 ^ 5 = 0010 ^ 0101 = 0111
//		separate numbers into 2 groups on right-most bit (cause it's 1):
//		- odd number bits: 1, 1, 3, 3, 5
//		- even number bits: 2
//		XOR first group: 5, XOR second group: 2

//	2.	too slow, I think group array is not needed, can use number to denote

//	3.	too slow, finding last bit can use 2's complement num & -num
