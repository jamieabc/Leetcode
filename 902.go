package main

import (
	"math"
	"sort"
	"strconv"
)

// Given an array of digits, you can write numbers using each digits[i] as many times as we want.  For example, if digits = ['1','3','5'], we may write numbers such as '13', '551', and '1351315'.
//
// Return the number of positive integers that can be generated that are less than or equal to a given integer n.
//
//
//
// Example 1:
//
// Input: digits = ["1","3","5","7"], n = 100
// Output: 20
// Explanation:
// The 20 numbers that can be written are:
// 1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.
//
// Example 2:
//
// Input: digits = ["1","4","9"], n = 1000000000
// Output: 29523
// Explanation:
// We can write 3 one digit numbers, 9 two digit numbers, 27 three digit numbers,
// 81 four digit numbers, 243 five digit numbers, 729 six digit numbers,
// 2187 seven digit numbers, 6561 eight digit numbers, and 19683 nine digit numbers.
// In total, this is 29523 integers that can be written using the digits array.
//
// Example 3:
//
// Input: digits = ["7"], n = 8
// Output: 1
//
//
//
// Constraints:
//
//     1 <= digits.length <= 9
//     digits[i].length == 1
//     digits[i] is a digit from '1' to '9'.
//     All the values in digits are unique.
//     1 <= n <= 109

// tc: O(k log(n)), k: n's digit count, log(n) because we care only about n's
// digit count
func atMostNGivenDigitSet(digits []string, n int) int {
	str := strconv.Itoa(n)
	size := len(digits)

	// dp[i]: count from ith digit ~ last digit
	dp := make([]int, len(str)+1)
	dp[len(dp)-1] = 1 // right most digit always count

	combCount := 1
	var count int
	for i := len(dp) - 2; i >= 0; i-- {
		for j := range digits {
			b := []byte(digits[j])[0]
			if b < str[i] {
				dp[i] += combCount
			} else if b == str[i] {
				dp[i] += dp[i+1]
			}
		}
		combCount *= size

		// for any number digit count less than n, combination count adds all
		if i > 0 {
			count += combCount
		}
	}

	return count + dp[0]
}

// there are some improvements can be made, binary search is not needed becase
// at most 10 digits, which means log(10) not really big impact

// convert digits from string to number is also not needed, instead, convert n from
// int to string can also do big/small comparison

// in order to do binary search, numbers need to be sorted, but as in dp solution,
// it's not needed to do so
func atMostNGivenDigitSet2(digits []string, n int) int {
	var count int
	size := len(digits)
	nums := make([]int, 0)

	// find every digit of n
	for tmp := n; tmp > 0; tmp /= 10 {
		nums = append(nums, tmp%10)
	}

	// for every number with digits number < n's digit number, all combinations
	// count
	prev := 1
	for i := 0; i < len(nums)-1; i++ {
		prev *= size
		count += prev
	}

	ds := make([]int, size)
	for i := range digits {
		ds[i] = int([]byte(digits[i])[0] - '0')
	}
	sort.Ints(ds)

	// compute combination count when digit count same as n
	for i := len(nums) - 1; i >= 0; i-- {
		// find closed number <= digit
		idx := binarySearch(ds, nums[i])

		if idx == -1 {
			break
		}

		if ds[idx] == nums[i] && i != 0 {
			count += idx * prev
		} else {
			count += (idx + 1) * prev
			break
		}
		prev /= size
	}

	return count
}

func binarySearch(digits []int, n int) int {
	idx := -1

	for low, high := 0, len(digits)-1; low <= high; {
		mid := low + (high-low)>>1

		if digits[mid] == n {
			idx = mid
			break
		} else if digits[mid] > n {
			high = mid - 1
		} else {
			idx = mid
			low = mid + 1
		}
	}

	return idx
}

func atMostNGivenDigitSet1(digits []string, n int) int {
	size := len(digits)
	nums := make([]int, size)
	for i := range digits {
		nums[i] = int([]byte(digits[i])[0] - '0')
	}
	sort.Ints(nums)

	var count int

	recursive(nums, 0, n, &count)

	return count
}

func recursive(nums []int, current, n int, count *int) {
	for i := range nums {
		tmp := current*10 + nums[i]

		if tmp > n {
			return
		} else {
			*count++
		}

		recursive(nums, tmp, n, count)
	}
}

//	Notes
//	1.	naive solution is to use backtracking, but this is not efficient

//	2.	combination count relates to digit count

//	3.	inspired from solution, not really needs to convert digits from string
//		to number, because the point is to compare big/small, so it should be
//		okay to compare in string

//		binary search is not needed, at most for log(10)

//	4.	inspired from solution, a very good abstract of the problem: dp[i] means
//		combination count for i ~ last

//		e.g. n = 5678
//		dp[0]: count of 5678
//		dp[1]: count of 678
//		dp[2]: count of 78
