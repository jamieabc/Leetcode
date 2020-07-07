package main

import "math"

// Given an array of integers arr and an integer target.
//
// You have to find two non-overlapping sub-arrays of arr each with sum equal target. There can be multiple answers so you have to find an answer where the sum of the lengths of the two sub-arrays is minimum.
//
// Return the minimum sum of the lengths of the two required sub-arrays, or return -1 if you cannot find such two sub-arrays.
//
//
//
// Example 1:
//
// Input: arr = [3,2,2,4,3], target = 3
// Output: 2
// Explanation: Only two sub-arrays have sum = 3 ([3] and [3]). The sum of their lengths is 2.
//
// Example 2:
//
// Input: arr = [7,3,4,7], target = 7
// Output: 2
// Explanation: Although we have three non-overlapping sub-arrays of sum = 7 ([7], [3,4] and [7]), but we will choose the first and third sub-arrays as the sum of their lengths is 2.
//
// Example 3:
//
// Input: arr = [4,3,2,6,2,3,4], target = 6
// Output: -1
// Explanation: We have only one sub-array of sum = 6.
//
// Example 4:
//
// Input: arr = [5,5,4,4,5], target = 3
// Output: -1
// Explanation: We cannot find a sub-array of sum = 3.
//
// Example 5:
//
// Input: arr = [3,1,1,1,5,1,2,1], target = 3
// Output: 3
// Explanation: Note that sub-arrays [1,2] and [2,1] cannot be an answer because they overlap.
//
//
//
// Constraints:
//
// 1 <= arr.length <= 10^5
// 1 <= arr[i] <= 1000
// 1 <= target <= 10^8

func minSumOfLengths(arr []int, target int) int {
	size := len(arr)
	// dp[i] means min size of intervals that sums to target
	validMinSizeBefore := make([]int, size)
	validMinSizeBefore[0] = math.MaxInt32

	sum := arr[0]
	minSumSize := math.MaxInt32
	var left, right int
	for left, right = 0, 0; left < size; {
		if sum == target {
			if validMinSizeBefore[left] != math.MaxInt32 {
				minSumSize = min(minSumSize, validMinSizeBefore[left]+right-left+1)
			}

			if right < size-1 {
				validMinSizeBefore[right+1] = min(validMinSizeBefore[right], right-left+1)
				right++
				sum += arr[right]
			}

			sum -= arr[left]
			left++
		} else if right < size-1 && sum < target {
			right++
			sum += arr[right]
			validMinSizeBefore[right] = validMinSizeBefore[right-1]
		} else {
			sum -= arr[left]
			left++
		}
	}

	if minSumSize == math.MaxInt32 {
		return -1

	}

	return minSumSize
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	problems
//	1.	algorithm is wrong, cannot guarantee first found number is minimum, cannot guarantee that found next non-overlap
//		interval is optimal

//	2.	too slow, tc: O(n^2)

//	3.	inspired from https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/discuss/686105/JAVA-or-Sliding-window-with-only-one-array-or-No-HasMap

//		use dp to store non-overlaping intervals, dp[i] means minimum size for
//		interval ends before i (not including i)

//	4.	inspired from https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/discuss/685486/JAVA-O(N)-Time-Two-Pass-Solution-using-HashMap.

//		author uses a hasp map to store running sum, since every number > 0, sum
//		is increasing, so if an interval meets criteria, the other interval
//		exists is: running sum + target

//	5.	add reference https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/discuss/685548/Java-Sliding-Window-with-dp-O(N)-20-lines

//		I actually don't understand other people solution, until read this one
