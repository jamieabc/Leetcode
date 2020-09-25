package main

import (
	"sort"
	"strconv"
)

// Given a list of non negative integers, arrange them such that they form the largest number.
//
// Example 1:
//
// Input: [10,2]
// Output: "210"
//
// Example 2:
//
// Input: [3,30,34,5,9]
// Output: "9534330"
//
// Note: The result may be very large, so you need to return a string instead of an integer.

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return "0"
	}

	strs := make([]string, 0)
	for _, num := range nums {
		str := strconv.Itoa(num)
		strs = append(strs, str)
	}

	sort.Slice(strs, func(i, j int) bool {

		// for m, n := 0, 0; true; m, n = (m+1)%len(strs[i]), (n+1)%len(strs[j]) {
		// 	if m == len(strs[i])-1 && n == len(strs[j])-1 {
		// 		// 12, 1212, doesn't matter which one goes first
		// 		return strs[i][m] >= strs[j][n]
		// 	} else if strs[i][m] > strs[j][n] {
		// 		return true
		// 	} else if strs[i][m] < strs[j][n] {
		// 		return false
		// 	}
		// }

		first, second := strs[i]+strs[j], strs[j]+strs[i]
		return first > second
	})

	// first digit is the largest first digit among all numbers
	if strs[0][0] == '0' {
		return "0"
	}

	var largest string
	for i := range strs {
		largest += strs[i]
	}

	return largest
}

//	Notes
//	1.	to get larger digit, one situation might happens: part of number is
//		repeated, e.g. 12 & 1212 or 12 & 121

//	2.	another corner case is that "00" should only be "0"

//	3.	inspired from solution, another smarter way to do is concat two strings
//		and compare

//	4.	inspired from https://leetcode.com/problems/largest-number/discuss/53158/My-Java-Solution-to-share
//
//		to check if number should be 0, check first number first digit, because
//		first number first digit is the largest digit among all numbers
