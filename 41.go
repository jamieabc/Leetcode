package main

// Given an unsorted integer array, find the smallest missing positive integer.
//
// Example 1:
//
// Input: [1,2,0]
// Output: 3
// Example 2:
//
// Input: [3,4,-1,1]
// Output: 2
// Example 3:
//
// Input: [7,8,9,11,12]
// Output: 1
// Note:
//
// Your algorithm should run in O(n) time and uses constant extra space.

func firstMissingPositive(nums []int) int {
	for i := range nums {
		// 1 at index 0
		// 2 at index 1
		// skip infinite loop condition [1, 1, 1, 1]
		// do not change if target number is already correct
		for i != nums[i]-1 && nums[i]-1 >= 0 && nums[i]-1 < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := range nums {
		if i != nums[i]-1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func firstMissingPositive1(nums []int) int {
	var one bool

	for i := range nums {
		if nums[i] == 1 {
			one = true
		} else if nums[i] <= 0 {
			nums[i] = 1
		}
	}

	if !one {
		return 1
	}

	size := len(nums)
	for i := range nums {
		idx := abs(nums[i])

		if idx < size && nums[idx] > 0 {
			nums[idx] *= -1
		} else if idx == size && nums[0] > 0 {
			nums[0] *= -1
		}
	}

	for i := 2; i < size; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] < 0 {
		return size + 1
	}

	return size
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	Notes

//	1.	boundary case when array is empty

//	2.	be careful about condition to cause infinite loop, e.g. [1,1]

//		core of this solution is that each number should eventually be placed at
//		it's only valid position (i == nums[i]-1). keep cycling until specific
//		number cannot be exchanged. at this situation, one boundary condition
//		should be considered: target numbers and current number are same as
//		expected value, e.g. [2, 1, 2]

//	3.	inspired from solution, cyclic sort is described by "index as hash key"
//		which is a very precise explanation.

//	4.	after half year, cannot understand why cyclic sort is okay

//	5.	the method of cyclic sort, is to put number that is not at its position
//		to proper location, that's why target swap index is nums[i]-1, because
//		that's where it should be
