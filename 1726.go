package main

import "sort"

// Given an array nums of distinct positive integers, return the number of tuples (a, b, c, d) such that a * b = c * d where a, b, c, and d are elements of nums, and a != b != c != d.
//
//
//
// Example 1:
//
// Input: nums = [2,3,4,6]
// Output: 8
// Explanation: There are 8 valid tuples:
// (2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
// (3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)
//
// Example 2:
//
// Input: nums = [1,2,4,5,10]
// Output: 16
// Explanation: There are 16 valids tuples:
// (1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)
// (2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)
// (2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,4,5)
// (4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)
//
// Example 3:
//
// Input: nums = [2,3,4,6,8,12]
// Output: 40
//
// Example 4:
//
// Input: nums = [2,3,5,7]
// Output: 0
//
//
//
// Constraints:
//
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 104
//     All elements in nums are distinct.

// tc: O(n^2)
func tupleSameProduct(nums []int) int {
	size := len(nums)
	if size <= 3 {
		return 0
	}

	table := make(map[int]int)

	for i := range nums {
		for j := i + 1; j < size; j++ {
			table[nums[i]*nums[j]]++
		}
	}

	var ans int

	for _, count := range table {
		ans += count * (count - 1) * 4
	}

	return ans
}

// tc: O(n^3)
func tupleSameProduct1(nums []int) int {
	size := len(nums)
	if size <= 3 {
		return 0
	}

	sort.Ints(nums)

	var count int

	for i := 0; i <= size-4; i++ {
		for j := size - 1; j > i; j-- {
			target := nums[i] * nums[j]

			for k, l := i+1, j-1; k < l; {
				if k == i {
					k++
				} else if l == j {
					j--
				} else {
					tmp := nums[k] * nums[l]

					if tmp == target {
						count += 8
						k++
						l--
					} else if tmp > target {
						l--
					} else {
						k++
					}
				}
			}
		}
	}

	return count
}

//	Notes
//	1.	all numbers are distinct, if same result of multiplication, then numbers
//		must be different

//		inspired from https://www.youtube.com/watch?v=E3XV8qXdQDs

//		[1, 3, 4, 12], 1*12 = 3*4 = 12, there are 2 groups,

//		combination(2, 2) * 2 * 2 * 2
//							^ - swap of 1 & 12
//								^ - swap of 3 & 4
//									^ - swap of (1, 12) & (3, 4)
