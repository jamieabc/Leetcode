package main

//Given an array with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.
//
//We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).
//
//Example 1:
//
//Input: [4,2,3]
//Output: True
//Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
//Example 2:
//
//Input: [4,2,1]
//Output: False
//Explanation: You can't get a non-decreasing array by modify at most one element.
//Note: The n belongs to [1, 10,000].

// [3, 4, 2, 3] false
// [2, 3, 3, 2, 4] true
func checkPossibility(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	}

	var changes bool

	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			if changes {
				return false
			}

			if i > 0 {
				if nums[i-1] > nums[i+1] {
					nums[i+1] = nums[i]
				} else {
					nums[i] = nums[i+1]
				}
			} else {
				nums[i] = nums[i+1]
			}
			changes = true
		}
	}
	return true
}

//	problems
//	1.	inspired from https://leetcode.com/problems/non-decreasing-array/discuss/106826/JavaC%2B%2B-Simple-greedy-like-solution-with-explanation

//		the thinking is pretty good, the order relates to 3 numbers:
//		i-1, i, i+1

//		the best strategy is to modify num[i] cause the higher nums[i+1]
//		makes more risk to fail. but when nums[i-1] > nums[i+1] then it's
//		a must to change nums[i+1]

//	2.	add reference https://leetcode.com/problems/non-decreasing-array/discuss/106842/The-easiest-python-solution....

//		the other way is to check if there exist another descending for next
//		number
