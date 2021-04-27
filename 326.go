package main

import "math"

//Given an integer, write a function to determine if it is a power of three.
//
//Example 1:
//
//Input: 27
//Output: true
//
//Example 2:
//
//Input: 0
//Output: false
//
//Example 3:
//
//Input: 9
//Output: true
//
//Example 4:
//
//Input: 45
//Output: false
//
//Follow up:
//Could you do it without using any loop / recursion?

func isPowreOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

func isPowerOfThree2(n int) bool {
	val := math.Log(float64(n)) / math.Log(float64(3))

	var epsilon float64 = 0.0000000001

	return math.Mod(val+epsilon, float64(1)) <= 2*epsilon
}

func isPowerOfThree1(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}

//	Notes
//	1.	forget about minus number

//	2.	inspired from solution,
//		mathematical representataion of log(a)(b) = log(10)(a)/log(10)(b)

//		if k = 3^m
//		m = log(3)(k), m should be an integer
//		but this is problematic, because to rely on this needs precision of
//		division, and floating point division introduces inaccuracy, to make
//		sure a number is integer or not, need to compare with smallest difference

//		go's % can only deal with integer, to compare floating point number,
//		need to use math.Mod

//		inspired from https://stackoverflow.com/a/22185792
//		use math.Nextafter(1.0, 2.0) - 1.0 to get epsilon

//		it turns out previous equation is not able to pass largest 3^19
//		(largest number) in 32-bits

//		inspired from https://gist.github.com/cevaris/bc331cbe970b03816c6b
//		use epsilon for 0.0000000001

//		the most important concept here is that when dealing with small
//		inaccuracy, need to compare by in-equality, instead of equality (==)

//	3.	inspired form solution, use largest power of three number 3^19
//		since 3 is a prime number, the only condition other number can be fully
//		divide is when other number is also power of three
