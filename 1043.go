package main

// Given an integer array A, you partition the array into (contiguous) subarrays of length at most K.  After partitioning, each subarray has their values changed to become the maximum value of that subarray.
//
// Return the largest sum of the given array after partitioning.
//
//
//
// Example 1:
//
// Input: A = [1,15,7,9,2,5,10], K = 3
// Output: 84
// Explanation: A becomes [15,15,15,9,10,10,10]
//
//
//
// Note:
//
//     1 <= K <= A.length <= 500
//     0 <= A[i] <= 10^6

func maxSumAfterPartitioning(arr []int, k int) int {
	size := len(arr)

	// dp[i]: maximum score can get after i
	dp := make([]int, size)

	for i := size - 1; i >= 0; i-- {
		var largest, maxSoFar int

		for j := 0; j < k && i+j < size; j++ {
			largest = max(largest, arr[i+j])

			if i+j+1 < size {
				maxSoFar = max(maxSoFar, largest*(j+1)+dp[i+j+1])
			} else {
				maxSoFar = max(maxSoFar, largest*(j+1))
			}
		}

		dp[i] = maxSoFar
	}

	return dp[0]
}

func maxSumAfterPartitioning5(arr []int, k int) int {
	size := len(arr)

	// dp[i]: maximum score can get before i
	dp := make([]int, size)

	for i := range arr {
		var largest, maxSoFar int

		for j := 0; j < k && i-j >= 0; j++ {
			largest = max(largest, arr[i-j])

			if i-j-1 >= 0 {
				maxSoFar = max(maxSoFar, largest*(j+1)+dp[i-j-1])
			} else {
				maxSoFar = max(maxSoFar, largest*(j+1))
			}
		}

		dp[i] = maxSoFar
	}

	return dp[size-1]
}

func maxSumAfterPartitioning4(A []int, K int) int {
	memo := make([]int, len(A))
	return dfs(A, K, 0, memo)
}

func dfs(nums []int, K, start int, memo []int) int {
	size := len(nums)
	if start >= size {
		return 0
	}

	if memo[start] != 0 {
		return memo[start]
	}

	var maxNumSoFar, maxSum int

	for i := start; i < min(size, start+K); i++ {
		maxNumSoFar = max(maxNumSoFar, nums[i])
		maxSum = max(maxSum, maxNumSoFar*(i-start+1)+dfs(nums, K, i+1, memo))
	}

	memo[start] = maxSum

	return maxSum
}

func maxSumAfterPartitioning3(arr []int, k int) int {
	size := len(arr)

	// dp[i][j]: largest sum from i ~ j
	dp := make([][]int, size+k+1)
	for i := range dp {
		dp[i] = make([]int, size+k+1)
	}

	for i := size - 1; i >= 0; i-- {
		dp[i][i] = arr[i] + dp[i+1][size-1]
		maxSoFar, largestNum := dp[i][i], arr[i]

		for j := 1; j < k && i+j < size; j++ {
			largestNum = max(largestNum, arr[i+j])

			dp[i][i+j] = max(dp[i][i+j], largestNum*(j+1)+dp[i+j+1][size-1])
			maxSoFar = max(maxSoFar, dp[i][i+j])
		}

		if dp[i][size-1] == 0 {
			dp[i][size-1] = maxSoFar
		}
	}

	return dp[0][size-1]
}

func maxSumAfterPartitioning2(A []int, K int) int {
	length := len(A)
	dp := make([]int, length)

	for i := range A {
		m := 0
		for j := 0; j < K && i-j >= 0; j++ {
			m = max(m, A[i-j])

			if i-j > 0 {
				dp[i] = max(dp[i], dp[i-j-1]+m*(j+1))
			} else {
				dp[i] = max(dp[i], m*(j+1))
			}
		}
	}

	return dp[length-1]
}

func maxSumAfterPartitioning1(A []int, K int) int {
	length := len(A)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}

	for i := range A {
		m := 0
		for j := i; j < min(length, i+K); j++ {
			m = max(m, A[j])
			dp[i][j] = (j - i + 1) * m
		}
	}

	var m int
	for d := K; d < length; d++ {
		for i := 0; i+d < length; i++ {
			m = 0
			for j := 0; j <= d; j++ {
				if j == 0 {
					m = max(m, dp[i][i]+dp[i+1][i+d])
				} else if j == i+d {
					m = max(m, dp[j][j]+dp[i][j-1])
				} else {
					m = max(m, dp[i][j]+dp[j+1][i+d])
				}
			}
			dp[i][i+d] = m
		}
	}

	return dp[0][length-1]
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	too slow

//	2.	inspired from https://leetcode.com/problems/partition-array-for-maximum-sum/discuss/290863/JavaC%2B%2BPython-DP

//		I know that partition means different kind of combination,
//		e.g sequence 3, 6, 9, 2, 5 and K = 3

//		max sum = max([3, 6, 9] + [2, 5], [3, 6] + [9, 2, 5],
//					  [3, 6] + [9, 2], [5], ...)

//		this is recursive form, it will get to the answer, but the problem is
//		it didn't see through nature of the problem. what I do is directly
//		partition items, then count them to find max sum.

//		but if I see closely, the form means there are only some combinations
//		determined by K, and if beyond K's former items, max sum keeps the
//		same.

//		e.g. K = 3, index = 9, combinations are:
//		maxSum(8) + arr[9]
//		maxSum(7) + max(arr[8], arr[9]) * 2
//		maxSum(6) + max(arr[7], arr[8], arr[9]) * 3

//		formula can be written as follows:
//		maxSum(i) = max(maxSum(i-1) + a[i],
//						maxSum(i-2) + max(a[i], a[i-1]) * 2,
//		   				maxSum(i-3) + max(a[i], a[i-1], a[i-2]) * 3,
//						...)

// 		dp can reduce down to 1D, also means only care about maximum up to some
//		index, and this value is independent from later sub-problems

//		also, the problem can be partitioned into small independent categories,
//		like counting cares about future appear numbers, only different is in
//		this case cares about previous numbers

//		key point here is to convert consideration of a point influence range
//		from -k ~ +k to -k only, k is actually a parameter controls influence
//		range, later numbers are no influenced by former numbers

//	3.	add reference https://leetcode.com/problems/partition-array-for-maximum-sum/discuss/299443/Java-O(NK).-Faster-than-99.82.-Less-memory-than-100.-With-Explanation.

//		author uses recursive to find solution

//	4.	a month later, cannot solve this problem...

//	5.	inspired from https://leetcode.com/problems/partition-array-for-maximum-sum/discuss/370807/dfs-solution-using-memoization-super-easy-to-understand

//		why dfs work? because maximum sum at index is not changed
//		e.g. [1, 15, 7, 9, 2, 5, 10, 3]
//		start from right most number 3, maximum sum is 3
//		then start from second last number 10, there are two conditions:
//		- [10], [3]
//		- [10, 3]

//		start form third last number 5, there are 3 conditions:
//		- [5], [10], [3]
//		- [5], [10, 3]
//		- [5, 10, 3]

//		sub-problem pattern is revealed, because both [10], [3] & [10, 3] is
//		already considered at second last number

//	6.	after 7 months, cannot solve this problem.

//		dfs, like tree traversal, is a way of solving problem top-down, so
//		the point is to find what kind of result can be stored, and find out
//		transformation relationship

//		the other point is, result only fix when reaches bottom

//		e.g. [1, 2, 9, 3], k = 2
//			   ^            		cut here, 1 + [2, 9, 3]

//			[2, 9, 3], k = 2
//			  ^                 	cut here, 2 + [9, 3]

//			[9, 3], k = 2
//			  ^ 					cut here, 9 + [3] = 12
//			     ^					cut here, 9 + 9 = 18, then return back

//		when thinking this problem, I stuck at how to select numbers,
//		[1, 5, 3, 8], k = 2, if choose first two as group, becomes [5, 5, 3, 8]
//		but since 3 < 5, it would be nice to choose [1, 5, 5, 8], so how can
//		I decide the decision? The trick is go to last number, there's only one
//		way to select, and go back a little, use max to decide which one is
//		better.

//		as I realize last number with no selection, I finally know bottom-up
//		works, because that's how dfs intended to work.

//	7.	in between decision process, there could have 2*k range of selection,
//		so the only way to reduce choices is down to last number, that's also
//		the reason bottom-up work
