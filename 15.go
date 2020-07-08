package main

import (
	"fmt"
	"sort"
)

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
//
// Note:
//
// The solution set must not contain duplicate triplets.
//
// Example:
//
// Given array nums = [-1, 0, 1, 2, -1, -4],
//
// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

// two pointer
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		sum := -nums[i]

		for low, high := i+1, len(nums)-1; low < high; {
			if sum == nums[low]+nums[high] {
				result = append(result, []int{nums[i], nums[low], nums[high]})
				for low < high && nums[low] == nums[low+1] {
					low++
				}
				low++
				for low < high && nums[high] == nums[high-1] {
					high--
				}
				high--
			} else if sum < nums[low]+nums[high] {
				high--
			} else {
				low++
			}
		}

		// avoid duplicate nums[i]
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return result
}

func threeSum2(nums []int) [][]int {
	result := make([][]int, 0)
	seen := make(map[string]bool)

	for i, num := range nums {
		partialNums := twoSum(nums, -num, i+1)
		for _, partial := range partialNums {

			smaller, larger := min(num, min(partial[0], partial[1])), max(num, max(partial[0], partial[1]))

			key := fmt.Sprintf("%d-%d", smaller, larger)
			if _, ok := seen[key]; !ok {
				result = append(result, append(partial, num))
				seen[key] = true
			}
		}
	}

	return result
}

func twoSum(nums []int, target, start int) [][]int {
	counter := make(map[int]int)
	result := make([][]int, 0)

	for i := start; i < len(nums); i++ {
		num := nums[i]

		if (target-num == num && counter[num] < 1) || counter[target-num] == 0 {
			counter[num]++
			continue
		}

		result = append(result, []int{num, target - num})
		counter[num]++
	}

	return result
}

func threeSum1(nums []int) [][]int {
	result := make([][]int, 0)
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	keys := make([]int, 0)
	for key := range counter {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for i, num := range keys {
		sum := -num
		counter[num]--

		for j := i; j < len(keys); j++ {
			theOther := sum - keys[j]

			if counter[keys[j]] == 0 || theOther < keys[j] || counter[theOther] == 0 {
				continue
			}

			if (theOther == keys[j] && counter[theOther] >= 2) || (theOther != keys[j] && counter[theOther] >= 1) {
				result = append(result, []int{num, keys[j], theOther})
			}
		}

		counter[num] = 0
	}

	return result
}

//	problems
//	1.	too slow, I spend 1 hour figuring out the solution, yet still slow
//		I am struggling at how to find duplicates, because hashmap can only
//		give answer to existence of number, so it could happen duplicate
//		combinations

//		e.g. -1, 0, 1, 2, -1, -4
//		first choose -1 to be sum, then iterate through rest of number and find
//		[-1, 0, 1]
//		then keep iterating and found [-1, 1, 0] which is exactly duplicates of
//		previous group

//		what I think of is to sort arrays, make sure third number is larger than
//		or equal to previous numbers to make sure no duplicates

//		tc: O(n^2)

//	2.	inspired from https://leetcode.com/problems/3sum/discuss/7380/Concise-O(N2)-Java-solution
//
//		sort numbers and use sliding window (2 pointers) to make sure a number
//		will not be used anymore in this round.

//		so, duplicates simply means I go through same number again, if there's
//		a way of avoiding this, it could work

//	3.	another way from solution, use a hash to check, the tricky part is that
//		since a + b + c = 0, only need 2 numbers for key, e.g. ascending order
