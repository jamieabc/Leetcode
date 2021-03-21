package main

// You are given three positive integers n, index and maxSum. You want to construct an array nums (0-indexed) that satisfies the following conditions:
//
// nums.length == n
// nums[i] is a positive integer where 0 <= i < n.
// abs(nums[i] - nums[i+1]) <= 1 where 0 <= i < n-1.
// The sum of all the elements of nums does not exceed maxSum.
// nums[index] is maximized.
//
// Return nums[index] of the constructed array.
//
// Note that abs(x) equals x if x >= 0, and -x otherwise.
//
//
//
// Example 1:
//
// Input: n = 4, index = 2,  maxSum = 6
// Output: 2
// Explanation: The arrays [1,1,2,1] and [1,2,2,1] satisfy all the conditions. There are no other valid arrays with a larger value at the given index.
//
// Example 2:
//
// Input: n = 6, index = 1,  maxSum = 10
// Output: 3
//
//
//
// Constraints:
//
// 1 <= n <= maxSum <= 109
// 0 <= index < n

func maxValue(n int, index int, maxSum int) int {
	ans := 1

	for low, high := 2, maxSum; low <= high; {
		mid := low + (high-low)/2

		if check(mid, maxSum, index, n-1-index) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}

func check(mid, maxSum, left, right int) bool {
	var sum int

	if left >= mid-1 {
		sum += (mid+1)*mid/2 + left - mid + 1
	} else if left > 0 {
		sum += (mid+1)*mid/2 - (mid-left)*(mid-left-1)/2
	}

	if right >= mid-1 {
		sum += (mid+1)*mid/2 + right - mid + 1
	} else if right > 0 {
		sum += (mid+1)*mid/2 - (mid-right)*(mid-right-1)/2
	}

	if left > 0 && right > 0 {
		sum -= mid
	}

	return sum <= maxSum
}

// TLE
func maxValue1(n int, index int, maxSum int) int {
	if n == 1 {
		return maxSum
	}

	diff := maxSum - n
	if diff == 0 {
		return 1
	}

	// distributed n
	var ans int
	for low, high := 2, diff; low <= high; {
		mid := low + (high-low)/2

		sum := mid - 1
		var i, cur int
		for i, cur = index-1, mid-1; i > 0 && cur > 2; i, cur = i-1, cur-1 {
		}

		if index > 0 {
			sum += (mid - 2 + cur - 1) * (index - i) / 2
		}

		for i, cur = index+1, mid-1; i < n-1 && cur > 2; i, cur = i+1, cur-1 {
		}
		if index < n-1 {
			sum += (mid - 2 + cur - 1) * (i - index) / 2
		}

		if sum == diff {
			return mid
		} else if sum > diff {
			high = mid - 1
		} else {
			ans = mid
			low = mid + 1
		}
	}

	return ans
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

//	Notes
//	1.	TLE during contest, because i didn't find out en effective way to count boundary

//	2.	inspired from https://leetcode.com/problems/maximum-value-at-a-given-index-in-a-bounded-array/discuss/1119666/Python-or-O(log-n)-or-Easy-to-understand-explanation

//		the way author check boundary: check if left boundary is large enough,
//		if not, deduct: sum(5,4,3,2,1) - sum(2, 1) = sum(5, 4, 3)

//	3.	even know partial sum, still takes half hour to really pass all test cases...
