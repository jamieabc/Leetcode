package main

//Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
//
//If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).
//
//The replacement must be in place and use only constant extra memory.
//
//
//
//Example 1:
//
//Input: nums = [1,2,3]
//Output: [1,3,2]
//
//Example 2:
//
//Input: nums = [3,2,1]
//Output: [1,2,3]
//
//Example 3:
//
//Input: nums = [1,1,5]
//Output: [1,5,1]
//
//Example 4:
//
//Input: nums = [1]
//Output: [1]
//
//
//
//Constraints:
//
//1 <= nums.length <= 100
//0 <= nums[i] <= 100

func nextPermutation(nums []int) {
	size := len(nums)

	var j int

	for i := size - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			// find next greater position to swap
			for j = size - 1; j > i; j-- {
				if nums[j] > nums[i] {
					// swap
					nums[j], nums[i] = nums[i], nums[j]
					break
				}
			}

			// j is the place to start reversing
			j = i + 1
			break
		}
	}

	// not found, just revert
	for k := size - 1; j < k; j, k = j+1, k-1 {
		nums[j], nums[k] = nums[k], nums[j]
	}
}

//	Notes
//	1.	inspired form solution, refactor a little
