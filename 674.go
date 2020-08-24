package main

//Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).
//
//Example 1:
//
//Input: [1,3,5,4,7]
//Output: 3
//Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3.
//Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4.
//
//Example 2:
//
//Input: [2,2,2,2,2]
//Output: 1
//Explanation: The longest continuous increasing subsequence is [2], its length is 1.
//
//Note: Length of the array will not exceed 10,000. ]

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLength, continuous := 1, 1

	for i, prev := 1, nums[0]; i < len(nums); i++ {
		if prev < nums[i] {
			continuous++
		} else {
			continuous = 1
		}

		maxLength = max(maxLength, continuous)
		prev = nums[i]
	}

	return maxLength
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
