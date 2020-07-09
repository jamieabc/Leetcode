package main

import "sort"

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)

	var count int
	for i, num := range nums {
		for low, high := i+1, len(nums)-1; low < high; {
			sum := num + nums[low] + nums[high]

			if sum >= target {
				high--
			} else {
				count += high - low
				low++
			}
		}
	}

	return count
}

//	problems
//	1.	this problem needs to find smaller, to using 2 pointers, if high pointer
//		found sum < target, which means all number before high pointer are valid,
//		so count it and move low pointer.

//		this method words because all numbers are sorted, so we can definitely
//		know once pointer moves, sum will become larger
