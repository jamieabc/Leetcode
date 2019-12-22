package main

//Given an array consisting of n integers, find the contiguous subarray of given length k that has the maximum average value. And you need to output the maximum average value.
//
//Example 1:
//
//Input: [1,12,-5,-6,50,3], k = 4
//Output: 12.75
//Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
//
//
//
//Note:
//
//    1 <= k <= n <= 30,000.
//    Elements of the given array will be in the range [-10,000, 10,000].

func findMaxAverage(nums []int, k int) float64 {
	length := len(nums)
	var sum, maxSum int

	if length <= k {
		for _, n := range nums {
			maxSum += n
		}
		return float64(maxSum) / float64(k)
	}

	// find initial sum
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum = sum

	for i := k; i < length; i++ {
		sum -= nums[i-k]
		sum += nums[i]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

// problem
// 1. end of i should be length-1
