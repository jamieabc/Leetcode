package main

import (
	"strings"
)

//Given two binary strings, return their sum (also a binary string).
//
//The input strings are both non-empty and contains only characters 1 or 0.
//
//Example 1:
//
//Input: a = "11", b = "1"
//Output: "100"
//
//Example 2:
//
//Input: a = "1010", b = "1011"
//Output: "10101"

func addBinary(a string, b string) string {
	base := a
	add := b
	if len(a) <= len(b) {
		base = b
		add = a
	}

	result := make([]byte, len(base)+1) // +1 in case exist carry
	length := len(result)

	c := byte('0')
	for i := range base {
		src1 := base[len(base)-1-i]
		if i < len(add) {
			// add
			src2 := add[len(add)-1-i]
			result[length-1-i] = xor(xor(src1, src2), c)
			c = carry(src1, src2, c)
		} else if c == '1' {
			// write carry
			result[length-1-i] = xor(src1, c)
			c = carry(src1, c, 0)
		} else {
			// write remaining
			result[length-1-i] = src1
		}
	}

	start := 1

	// check if carry exist
	if c == '1' {
		result[0] = '1'
		start = 0
	}

	var sb strings.Builder
	for i := start; i < length; i++ {
		sb.WriteByte(result[i])
	}

	return sb.String()
}

func xor(a, b byte) byte {
	tmp1 := a - '0'
	tmp2 := b - '0'
	if tmp1 == tmp2 {
		return byte('0')
	}
	return byte('1')
}

func carry(a, b, c byte) byte {
	tmp1 := a - '0'
	tmp2 := b - '0'
	tmp3 := c - '0'
	if tmp1&tmp2 == 1 || tmp2&tmp3 == 1 || tmp1&tmp3 == 1 {
		return byte('1')
	}
	return byte('0')
}

// problems
// 1. when doing operation, convert '0' to 0, but when write back to string, forget to convert 0 to '0'
// 2. too complicate conversion => unify, all use '0' or '1' as input
