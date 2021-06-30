package main

// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.



// Example 1:

// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
// Explanation: [1,1,1,0,0,1,1,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

// Example 2:

// Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
// Output: 10
// Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.



// Constraints:

//     1 <= nums.length <= 105
//     nums[i] is either 0 or 1.
//     0 <= k <= nums.length

func longestOnes(nums []int, k int) int {
	var longest int
	remain := k
	n := len(nums)

	for i, j := 0, 0; j < n; j++ {
		for ; j < n; j++ {
			if nums[j] == 0 {
				remain--

				if remain < 0 {
					break
				}
			}
		}

		longest = max(longest, j-i)

		// forward i either to next position of 0 or to j
		for ; i < j && nums[i] == 1; i++ {}

		// no matter how, i should advance
		// e.g. [0, 0, 0], k = 0
		remain++
		i++
	}

	return longest
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
//	Notes
//	1.	initially i change the problem into pick k from n zeros, tc becomes O(n^k)
//
//		e.g. [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], k = 2
//
//		zeros at [3, 4, 5, 10], pick 2 from 4
//
//		so i though about using dp, but there are 10^5 numbers in an array, which is too big for a key
//		with 2^(10^5) possibilities, didn't think of a viable solution to solve this problem
//
//	2.	after reference others solution, found that this is sliding window problem
//
//		longest consecutive still occurs from some index, use all possible flips
//		previous position reachable range can be further used by later position, which reduces computation
//
//		e.g.	0 0 0 1 1 0 0, k = 2
//
//				<-> start from 0, extend to 1
//
//				  <-> start from 1, extend to 2
//
//					<---> start from 2, extend to 5
//
//					  <-----> start from 3, extend to 6
