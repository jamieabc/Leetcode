package main

import (
	"math"
	"sort"
)

// A sequence of numbers is called arithmetic if it consists of at least two elements, and the difference between every two consecutive elements is the same. More formally, a sequence s is arithmetic if and only if s[i+1] - s[i] == s[1] - s[0] for all valid i.
//
// For example, these are arithmetic sequences:
//
// 1, 3, 5, 7, 9
// 7, 7, 7, 7
// 3, -1, -5, -9
//
// The following sequence is not arithmetic:
//
// 1, 1, 2, 5, 7
//
// You are given an array of n integers, nums, and two arrays of m integers each, l and r, representing the m range queries, where the ith query is the range [l[i], r[i]]. All the arrays are 0-indexed.
//
// Return a list of boolean elements answer, where answer[i] is true if the subarray nums[l[i]], nums[l[i]+1], ... , nums[r[i]] can be rearranged to form an arithmetic sequence, and false otherwise.
//
//
//
// Example 1:
//
// Input: nums = [4,6,5,9,3,7], l = [0,0,2], r = [2,3,5]
// Output: [true,false,true]
// Explanation:
// In the 0th query, the subarray is [4,6,5]. This can be rearranged as [6,5,4], which is an arithmetic sequence.
// In the 1st query, the subarray is [4,6,5,9]. This cannot be rearranged as an arithmetic sequence.
// In the 2nd query, the subarray is [5,9,3,7]. This can be rearranged as [3,5,7,9], which is an arithmetic sequence.
//
// Example 2:
//
// Input: nums = [-12,-9,-3,-12,-6,15,20,-25,-20,-15,-10], l = [0,1,6,4,8,7], r = [4,4,9,7,9,10]
// Output: [false,true,false,false,true,true]
//
//
//
// Constraints:
//
// n == nums.length
// m == l.length
// m == r.length
// 2 <= n <= 500
// 1 <= m <= 500
// 0 <= l[i] < r[i] < n
// -105 <= nums[i] <= 105

// tc: O(mn), m: nums size, n: query size
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool, len(l))

	for i := range l {
		if r[i]-l[i] <= 1 {
			ans[i] = true
			continue
		}

		// find max & min
		maxVal, minVal := math.MinInt32, math.MaxInt32
		for j := l[i]; j <= r[i]; j++ {
			maxVal, minVal = max(maxVal, nums[j]), min(minVal, nums[j])
		}

		if maxVal == minVal {
			ans[i] = true
			continue
		}

		// check if every arithmetic numbers exist
		diff := (maxVal - minVal) / (r[i] - l[i])

		// diff might be 0, divide by 0 is not allowed
		if diff == 0 {
			ans[i] = false
			continue
		}

		size := r[i] - l[i] + 1
		checks := make([]bool, size)
		found := true

		for j := l[i]; j <= r[i]; j++ {
			idx := (nums[j] - minVal) / diff

			// might not be arithmetic, idx could out of bound
			if idx >= size || nums[j] != minVal+diff*idx || checks[idx] {
				found = false
				break
			}
			checks[idx] = true
		}

		ans[i] = found
	}

	return ans
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

// tc: O(m n log(n)), m: nums size, n: range size & sort
func checkArithmeticSubarrays1(nums []int, l []int, r []int) []bool {
	ans := make([]bool, len(l))

	for i := range l {
		if r[i]-l[i] <= 1 {
			ans[i] = true
			continue
		}

		tmp := append([]int{}, nums[l[i]:r[i]+1]...)
		sort.Ints(tmp)

		arithmatic := true
		diff := tmp[1] - tmp[0]

		for j := 2; j < len(tmp); j++ {
			if tmp[j]-tmp[j-1] != diff {
				arithmatic = false
			}
		}

		ans[i] = arithmatic
	}

	return ans
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/arithmetic-subarrays/discuss/910421/C%2B%2B-Two-Approaches-(140-vs-40-ms)

//		arithmetic sequence means knowing each number difference can check, but
//		if getting numbers not adjacent, it might be wrong difference.

//		the brilliant is that finding max & min for such range, and pre-fill
//		all possible numbers to check if all exist.
