package main

import "strings"

// Given a binary string s (a string consisting only of '0' and '1's) and a positive integer n, return true if and only if for every integer x from 1 to n, the binary representation of x is a substring of s.
//
//
//
// Example 1:
//
// Input: s = "0110", n = 3
// Output: true
//
// Example 2:
//
// Input: s = "0110", n = 4
// Output: false
//
//
//
// Note:
//
// 1 <= s.length <= 1000
// 1 <= n <= 109

// tc: O(ns), for n ~ n/2, each time check length s
func queryString(s string, n int) bool {
	if n > 1024 {
		return false
	}

	for i := n; i > n/2; i-- {
		if !strings.Contains(s, toString(i)) {
			return false
		}
	}

	return true
}

func toString(n int) string {
	var str string

	for n > 0 {
		if n&1 > 0 {
			str = "1" + str
		} else {
			str = "0" + str
		}

		n = n >> 1
	}

	return str
}

// tc: O(nk)
func queryString1(s string, n int) bool {
	size := len(s)

	// find longest length of window
	k := 1
	for num := 1; num < n; num = num << 1 {
		k++
	}

	count := n
	arr := make([]bool, n+1)
	var num int

	for i := 0; i < size; i++ {
		// skip heading 0, cause it won't change anything
		for ; i < size && s[i] == '0'; i++ {
		}

		// find number in range i ~ min(i+k, size-1)
		num = 0
		for j := i; j < min(i+k, size); j++ {
			num = num << 1
			if s[j] == '1' {
				num++
			}

			if num <= n && n > 0 && !arr[num] {
				count--
				arr[num] = true
			}
		}
	}

	return count == 0
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	initially i thought it's 2-pointer, from index i ~ index j is computed,
//		i -> i+1, j -> j+1, only consider additional states but this is not true
//		in following example

//		e.g.	1 1 0 1 1
//				|-> 1, 3, 6, 13, 27
//				  |-> 1, 2, 5, 11

//		but it removes heading 1, [1, 3, 6, 13, 27] all remove 1 bits becomes
//		[0, 1, 2, 5, 11]

//		if heading is 0, removing 0 means do nothing, e.g. 011 => 11,

//		if i keep records of previous possible numbers, then use XOR can generate
//		next numbers, width of the sliding window at most 32 (because n <= 10^9)

//		tc should be O(nk)
//		sc should be O(n)

//		but it might not necessary to store previous numbers then dor XOR operation
//		it could also use i ~ min(i+k, size) and find all numbers

//		this way, tc should be O(nk), sc O(1)

//	2.	inspired from https://leetcode.com/problems/binary-string-with-substrings-representing-1-to-n/discuss/260847/JavaC%2B%2BPython-O(S)
//
//		instead of searching all possible combinations, try to find number from n ~ n/2
//		represents in string

//		the reason is because for interval gap from i ~ 2*i, all digits are different,

//		e.g.	i = 8
//		1000
//		1001
//		1010
//		1011
//		1100
//		1101
//		1110
//		1111

//		as long as 9~16 exist, no need to check reset of number 1~8 because, removing leading
//		1 forms numbers 1~8

//		inspired from https://leetcode.com/problems/binary-string-with-substrings-representing-1-to-n/discuss/526301/Pigeon-hole

//		for len(s) <= 1000, if the digits length is 10, maximum non-overlap substrings are
//		991, but 10 digits with 2^10 = 1024 combinations which is not enough

//		tc becomes O(n*s)

//		this is really brilliant, because construct s from n is really hard, so use some
//		calculation to reduce limit of n, and also faster to verify
