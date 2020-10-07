package main

import "math"

// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
//
// Follow-up: Could you solve the problem in linear time and in O(1) space?
//
//
//
// Example 1:
//
// Input: nums = [3,2,3]
// Output: [3]
// Example 2:
//
// Input: nums = [1]
// Output: [1]
// Example 3:
//
// Input: nums = [1,2]
// Output: [1,2]
//
//
// Constraints:
//
// 1 <= nums.length <= 5 * 104
// -109 <= nums[i] <= 109

func majorityElement(nums []int) []int {
	var n1, n2, count1, count2 int

	// remove every 3 different numbers
	for _, num := range nums {
		if n1 == num {
			count1++
		} else if n2 == num {
			count2++
		} else if count1 == 0 {
			n1 = num
			count1++
		} else if count2 == 0 {
			n2 = num
			count2++
		} else {
			count1--
			count2--
		}
	}

	// count candidates again
	counter := make([]int, 2)
	for _, num := range nums {
		if num == n1 {
			counter[0]++
		} else if num == n2 {
			counter[1]++
		}
	}

	result := make([]int, 0)

	if counter[0] > len(nums)/3 {
		result = append(result, n1)
	}

	if counter[1] > len(nums)/3 {
		result = append(result, n2)
	}

	return result
}

//	Notes
//	1.	similar to original idea, I think it's like finding a way to make
//		majority numbers in majority status. What I think of is when 3 different
//		numbers are met, remove 1 from all of them, thus, numbers with more than
//		1/3 still in majority status.

//		but I still can't distinguish situation if there are total n numbers,
//		a with n/3+1 & b with n/3+1, the other situation a with n/3-1 & b with
//		n/3-1, so need to iterate through all number again to make sure found
//		numbers really more than n/3

//	2.	inspired from solution, same technique, but uses variables instead of
//		hash

//	3.	can use count as an index to decide change a number or not
