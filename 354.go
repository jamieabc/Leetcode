package main

import "sort"

// You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi] represents the width and the height of an envelope.
//
// One envelope can fit into another if and only if both the width and height of one envelope is greater than the width and height of the other envelope.
//
// Return the maximum number of envelopes can you Russian doll (i.e., put one inside the other).
//
// Note: You cannot rotate an envelope.
//
//
//
// Example 1:
//
// Input: envelopes = [[5,4],[6,4],[6,7],[2,3]]
// Output: 3
// Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).
//
// Example 2:
//
// Input: envelopes = [[1,1],[1,1],[1,1]]
// Output: 1
//
//
//
// Constraints:
//
// 1 <= envelopes.length <= 5000
// envelopes[i].length == 2
// 1 <= wi, hi <= 104

// tc: O(n log(n))
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}

		// first number in asc, if first number is same, second number
		// sort in desc, to avoid edge condition about [1, 2], [1, 3], [1, 4]
		// count as 3 but actuallly 1
		return envelopes[i][1] > envelopes[j][1]
	})

	arr := make([]int, 0)

	for _, e := range envelopes {
		if len(arr) > 0 {
			if e[1] > arr[len(arr)-1] {
				arr = append(arr, e[1])
			} else {
				var idx int
				for low, high := 0, len(arr)-1; low <= high; {
					mid := low + (high-low)/2

					if arr[mid] == e[1] {
						idx = mid
						break
					} else if arr[mid] > e[1] {
						idx = mid
						high = mid - 1
					} else {
						low = mid + 1
					}
				}

				arr[idx] = e[1]
			}
		} else {
			arr = append(arr, e[1])
		}
	}

	return len(arr)
}

// tc: O(n^2)
func maxEnvelopes1(env [][]int) int {
	sort.Slice(env, func(i, j int) bool {
		if env[i][0] != env[j][0] {
			return env[i][0] < env[j][0]
		}
		return env[i][1] < env[j][1]
	})

	size := len(env)

	// dp[i]: maximum env to i
	dp := make([]int, size)
	for i := range dp {
		dp[i] = 1
	}

	ans := 1
	for i := 1; i < size; i++ {
		tmp := 1
		for j := i - 1; j >= 0; j-- {
			if env[j][0] < env[i][0] && env[j][1] < env[i][1] {
				tmp = max(tmp, dp[j]+1)
			}
		}
		dp[i] = tmp
		ans = max(ans, dp[i])
	}

	return ans
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from solution, this problem is about LIS
