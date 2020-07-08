package main

import (
	"math"
	"sort"
)

// sort array, find closest
func threeSumClosest(nums []int, target int) int {
	closest := math.MaxInt32
	sort.Ints(nums)

	var low, high int
	for i := 0; i < len(nums); i++ {
		for low, high = i+1, len(nums)-1; low < high; {
			sum := nums[i] + nums[low] + nums[high]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}

			if sum == target {
				return target
			} else if sum < target {
				low++
			} else {
				high--
			}
		}

		// avoid duplicate nums[i]
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return closest
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
