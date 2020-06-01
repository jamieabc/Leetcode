package main

import "math"

// Given a positive 32-bit integer n, you need to find the smallest 32-bit integer which has exactly the same digits existing in the integer n and is greater in value than n. If no such positive 32-bit integer exists, you need to return -1.
//
// Example 1:
//
// Input: 12
// Output: 21
//
//
//
// Example 2:
//
// Input: 21
// Output: -1

func nextGreaterElement(n int) int {
	nums := make([]int, 0)
	for j := n; j > 0; {
		nums = append(nums, j%10)
		j /= 10
	}

	// 1432 -> 2, 3, 4, 1
	// 4321 -> 1, 2, 3, 4
	idx := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			for j := 0; j < i; j++ {
				if nums[j] > nums[i] && nums[j] != nums[j+1] {
					nums[j], nums[i] = nums[i], nums[j]
					idx = i - 1
					break
				}
			}

			// all previous digits are same, choose closest
			if idx == -1 {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				idx = i - 1
			}
			break
		}
	}

	// it's mono increasing
	if idx == -1 {
		return -1
	}

	bucket := make([]int, 10)
	for i := 0; i <= idx; i++ {
		bucket[nums[i]]++
	}

	for j := 0; idx >= 0; {
		if bucket[j] == 0 {
			j++
		} else {
			bucket[j]--
			nums[idx] = j
			idx--
		}
	}

	var num int
	for i := len(nums) - 1; i >= 0; i-- {
		num *= 10
		num += nums[i]
	}

	if num > math.MaxInt32 {
		return -1
	}
	return num
}

//	problems
//	1.	don't know how to solve the problem

//	2.	inspired from https://leetcode.com/problems/next-greater-element-iii/discuss/101824/Simple-Java-solution-(4ms)-with-explanation.

//		the problem I cannot write is wrong logic. to find a larger number
//		means digits cannot be mono-decreasing, so if any increasing happens,
//		that's where to swap last digit.

//	3.	it's not always swap first digit, it's to swap smallest digit

//	4.	when deciding smallest number, it should swap to digits located
//		as closest as possible

//	5.	after decide which location to swap, in order to get smallest
//		number, the digit to swap should be smallest (last digit), but it
//		comes with an exception, if all digits are same with last digit,
//		then it needs to chose closed to swap target

//		1332244 => 1332424
//		1332432 => 1333224
