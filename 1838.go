package main

import "sort"

// The frequency of an element is the number of times it occurs in an array.
//
// You are given an integer array nums and an integer k. In one operation, you can choose an index of nums and increment the element at that index by 1.
//
// Return the maximum possible frequency of an element after performing at most k operations.
//
//
//
// Example 1:
//
// Input: nums = [1,2,4], k = 5
// Output: 3
// Explanation: Increment the first element three times and the second element two times to make nums = [4,4,4].
// 4 has a frequency of 3.
//
// Example 2:
//
// Input: nums = [1,4,8,13], k = 5
// Output: 2
// Explanation: There are multiple optimal solutions:
// - Increment the first element three times to make nums = [4,4,8,13]. 4 has a frequency of 2.
// - Increment the second element four times to make nums = [1,8,8,13]. 8 has a frequency of 2.
// - Increment the third element five times to make nums = [1,4,13,13]. 13 has a frequency of 2.
//
// Example 3:
//
// Input: nums = [3,9,6], k = 2
// Output: 1
//
//
//
// Constraints:
//
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 105
// 1 <= k <= 105

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	var sum, most int
	size := len(nums)

	for i, j := 0, 0; j < size; j++ {
		// use number at j
		sum += nums[j]

		// make i~j valid, looping if invalid
		for ; i <= j && nums[j]*(j-i+1)-sum > k; i++ {
			sum -= nums[i]
		}

		most = max(most, j-i+1)
	}

	return most
}

func maxFrequency1(nums []int, k int) int {
	sort.Ints(nums)

	size := len(nums)

	mostFreq := 1
	var diff int

	for i, j := 0, 1; j < size; i++ {
		// expand
		for ; j < size && (i == j || diff <= k); j++ {
			diff += (nums[j] - nums[j-1]) * (j - i)
		}

		if j < size || diff > k {
			mostFreq = max(mostFreq, j-i-1)
		} else {
			mostFreq = max(mostFreq, j-i)
		}

		// shrink
		diff -= nums[j-1] - nums[i]
	}

	return mostFreq
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	didn't finish in contest

//	2.	first I though it was about finding median value, and start from median value to
//		solve, but then I figure out it's not median value, because it needs to start
//		form smallest difference

//		then I thought if it's min-heap problem, but everytime a number is increased,
//		need to update the whole heap, because number of cost (difference * duplicate number)
//		is changed

//		then I thought could it be binary search, the search scope is difference, from
//		smallest difference to largest difference, but this is not right, because optimal
//		solution didn't really comes from this way

//		then I re-think from start of the problem, because it says most frequent number,
//		which means groups of same number, and number can only going up, next state can be
//		achieved from previous state, so it looks like a two-pointer problem, the cost can
//		be preserved for next time, and number is always increasing

//		unfortunately, there's some bug and didn't solve it in time, i took about 1 hour
//		for this problem...

//	3.	becareful about expand condition, it's j < size && (i == j || xxx)
//		the reason is because if i == j needs to expand, otherwise cannot going forward

//	4.	the other thing to check is that when j reached end, it's because diff too
//		large or it's because range is reached

//	5.	inspired form https://leetcode.com/problems/frequency-of-the-most-frequent-element/discuss/1175088/C%2B%2B-Maximum-Sliding-Window-Cheatsheet-Template!

//		author provides some templates for two-pointer

//		his solution is really beautiful...the sum to current index, and assumes all numbers
//		are at same height (use the current right index), difference is k

//	6.	inspired from https://leetcode.com/problems/frequency-of-the-most-frequent-element/discuss/1175181/JavaPython-Prefix-Sum-and-Binary-Search-O(NlogN)

//		author provides a very beautiful graph to demonstrate
