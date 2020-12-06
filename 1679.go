package main

import "sort"

//You are given an integer array nums and an integer k.
//
//In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
//
//Return the maximum number of operations you can perform on the array.
//
//
//
//Example 1:
//
//Input: nums = [1,2,3,4], k = 5
//Output: 2
//Explanation: Starting with nums = [1,2,3,4]:
//- Remove numbers 1 and 4, then nums = [2,3]
//- Remove numbers 2 and 3, then nums = []
//There are no more pairs that sum up to 5, hence a total of 2 operations.
//
//Example 2:
//
//Input: nums = [3,1,3,4,3], k = 6
//Output: 1
//Explanation: Starting with nums = [3,1,3,4,3]:
//- Remove the first two 3's, then nums = [1,4,3]
//There are no more pairs that sum up to 6, hence a total of 1 operation.
//
//
//
//Constraints:
//
//    1 <= nums.length <= 105
//    1 <= nums[i] <= 109
//    1 <= k <= 109

func maxOperations(nums []int, k int) int {
	var ops int
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
		remain := k - num

		if num == remain {
			if counter[num] >= 2 {
				ops++
				counter[num] -= 2
			}
		} else {
			if counter[remain] > 0 {
				ops++
				counter[remain]--
				counter[num]--
			}
		}
	}

	return ops
}

func maxOperations1(nums []int, k int) int {
	sort.Ints(nums)
	size := len(nums)
	var ans int

	for i, j := 0, size-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == k {
			ans++
			i++
			j--
		} else if sum > k {
			j--
		} else {
			i++
		}
	}

	return ans
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/max-number-of-k-sum-pairs/discuss/961351/C%2B%2B-Map-O(n)-and-Two-Pointer-O(nlogn)-easy-solution

//		could also use sort to solve this problem, tc: O(n log(n))
