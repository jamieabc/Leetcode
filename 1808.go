package main

// You are given a positive integer primeFactors. You are asked to construct a positive integer n that satisfies the following conditions:
//
// The number of prime factors of n (not necessarily distinct) is at most primeFactors.
// The number of nice divisors of n is maximized. Note that a divisor of n is nice if it is divisible by every prime factor of n. For example, if n = 12, then its prime factors are [2,2,3], then 6 and 12 are nice divisors, while 3 and 4 are not.
//
// Return the number of nice divisors of n. Since that number can be too large, return it modulo 109 + 7.
//
// Note that a prime number is a natural number greater than 1 that is not a product of two smaller natural numbers. The prime factors of a number n is a list of prime numbers such that their product equals n.
//
//
//
// Example 1:
//
// Input: primeFactors = 5
// Output: 6
// Explanation: 200 is a valid value of n.
// It has 5 prime factors: [2,2,2,5,5], and it has 6 nice divisors: [10,20,40,50,100,200].
// There is not other value of n that has at most 5 prime factors and more nice divisors.
//
// Example 2:
//
// Input: primeFactors = 8
// Output: 18
//
//
//
// Constraints:
//
// 1 <= primeFactors <= 109

func pow(base, exponent int) int {
	mod := int(1e9 + 7)
	ans := 1

	for exponent > 0 {
		if exponent&1 > 0 {
			ans = (ans * base) % mod
		}
		base = (base * base) % mod
		exponent = exponent >> 1
	}

	return ans
}

func maxNiceDivisors(primeFactors int) int {
	mod := int(1e9 + 7)
	if primeFactors <= 4 {
		return primeFactors
	}

	quo := primeFactors / 3
	remain := primeFactors - quo*3

	if remain == 1 {
		// remain == 1, num = 4, convert 3*1 => 2*2
		return (4 * pow(3, quo-1)) % mod
	} else if remain == 2 {
		// remain = 2, num = 5, convert 3 => 3*2
		// because 3*3 > 2*2*2, so /3 results in the case the 5 only
		// multiply by 3, which is wrong
		return (2 * pow(3, quo) % mod) % mod
	}

	return pow(3, quo) % mod
}

// Notes
//	1.	inspired from https://www.youtube.com/watch?v=5SFXooJnwN8

//		one way to solve this type of constructive problem is brute force and
//		try to find pattern

//		1, 2, 3, 4, 6, 9, 12, 18, 27, 36, 54, 72
//		for n <= 4, f(n) = n
//		for n >= 5, f(n) = 3f(n-3)

//		another way to solve this is to observe relations:
// 		decompose a number, it's prime number can be 2, 3, 5, 7, ...
//		count of 2 = a
//		count of 3 = b
//		count of 5 = c
//		count of 7 = d
//		... etc.

//		a + b + c + d + ... = n
//		total possible permutations are a*b*c*d*...

//		one key to solve this is to observe 3*3 > 2*2*2
//		2+2+2 = 3+3, so whenever count of 2 >= 3, change it to 3

//		the other key to solve the problem is dealing with too large part,
//		3^x = (3^(x/2))^2, and mod the result

//		if x is even, 3^x = (3^(x/2))^2
//		if x is odd, 3^x = (3^(x/2))^2*3

//	2.	inspired from https://www.khanacademy.org/computing/computer-science/cryptography/modarithmetic/a/modular-multiplication

//		(a*b)%c = (a%c * b%c) %c

//	3.	inspired from https://leetcode.com/problems/maximize-number-of-nice-divisors/discuss/1130586/C%2B%2BJava-modpow

//		3^27 = 3 * (3^2)^13 = 3 * 9^13 = 3 * 9 * (9^2)^6 = 3 * 9 * 81^6, etc.

//	4.	if a%b == 1, which means sum to 4, 2*2 > 3*1, that's the reason why
//		when it comes to remainder = 1, change it to 4
