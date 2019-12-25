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
	// sort array
	// iterate all elements, find out duplicate and the one that is not continuous
	// edge cases: first & last element missing
	// too slow, n log n

	// sum of all is n(n+1)/2, subtract from existing elements, the find delta of original to duplicate

	// use array to store occurrence of number
	length := len(nums)
	arr := make([]int, length)

	// duplicate will be 2, missing will be 0
	for _, n := range nums {
		arr[n-1]++
	}

	var duplicate, missing int
	for i, n := range arr {
		if n == 2 {
			duplicate = i + 1
		}
		if n == 0 {
			missing = i + 1
		}
	}

	return []int{duplicate, missing}
}

// problems
// 1. too slow, sort of n log n, try use sum & subtract, it will occupy some memory (map)
