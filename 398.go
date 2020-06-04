package main

import "math/rand"

// Given an array of integers with possible duplicates, randomly output the index of a given target number. You can assume that the given target number must exist in the array.
//
// Note:
// The array size can be very large. Solution that uses too much extra space will not pass the judge.
//
// Example:
//
// int[] nums = new int[] {1,2,3,3,3};
// Solution solution = new Solution(nums);
//
// // pick(3) should return either index 2, 3, or 4 randomly. Each index should have equal probability of returning.
// solution.pick(3);
//
// // pick(1) should return 0. Since in the array only nums[0] is equal to 1.
// solution.pick(1);

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (this *Solution) Pick(target int) int {
	var count, result int
	for i := range this.nums {
		if this.nums[i] == target {
			count++
			if rand.Intn(count) == 0 {
				result = i
			}
		}
	}
	return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

//	problems
//	1.	wrong return, it's index value

//	2.	too slow

//	3.	add reference https://leetcode.com/problems/random-pick-index/discuss/88080/What-on-earth-is-meant-by-too-much-memory

//		author explains background of this problem, when a data stream comes,
//		cannot know how long for that data, but we still want to equally
//		show target? and every time pick will be a new data stream, so it
//		shouldn't store any data, because it will be new one next time.

//		so the problem becomes: how to decide a target is chosen or not
//		chosen , based on previous information?

//	4.	add reference https://leetcode.com/problems/random-pick-index/discuss/88072/Simple-Reservoir-Sampling-solution

//		I have never know this problem is called reservoir sampling, when
//		input is too large and not able to know its size, how to equally
//		get some target?

//		if first item comes, it has 100% to be chosen
//		if second item comes, 50% of second element is chosen, then both
//		  1 & 2 has same possibility
//		if third item comes, 1/3 possibility to replace to third, 2/3
//		  possibility to keep existing, and since 1 & 2 has 50%, all items
//		  are equally chosen
