package main

//Given two positive integers n and k, the binary string  Sn is formed as follows:
//
//    S1 = "0"
//    Si = Si-1 + "1" + reverse(invert(Si-1)) for i > 1
//
//Where + denotes the concatenation operation, reverse(x) returns the reversed string x, and invert(x) inverts all the bits in x (0 changes to 1 and 1 changes to 0).
//
//For example, the first 4 strings in the above sequence are:
//
//    S1 = "0"
//    S2 = "011"
//    S3 = "0111001"
//    S4 = "011100110110001"
//
//Return the kth bit in Sn. It is guaranteed that k is valid for the given n.
//
//
//
//Example 1:
//
//Input: n = 3, k = 1
//Output: "0"
//Explanation: S3 is "0111001". The first bit is "0".
//
//Example 2:
//
//Input: n = 4, k = 11
//Output: "1"
//Explanation: S4 is "011100110110001". The 11th bit is "1".
//
//Example 3:
//
//Input: n = 1, k = 1
//Output: "0"
//
//Example 4:
//
//Input: n = 2, k = 3
//Output: "1"
//
//
//
//Constraints:
//
//    1 <= n <= 20
//    1 <= k <= 2n - 1

func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}

	k--
	var invert bool

	total := (1 << n) - 1
	for ; n > 1; n-- {
		mid := total / 2

		// mid
		if k == mid {
			if invert {
				return '0'
			}
			return '1'
		}

		if k > mid {
			k = total - 1 - k
			invert = !invert
		}

		total = mid
	}

	if invert {
		return '1'
	}
	return '0'
}

func findKthBit1(n int, k int) byte {
	if n == 1 {
		return '0'
	}

	k--
	if n == 2 {
		s := "011"
		return s[k]
	}

	total := (1 << n) - 1

	var invert bool

	s := "0111001"
	if n == 3 {
		return s[k]
	}

	for ; n > 3; n, total = n-1, total/2 {
		if k == total/2 {
			if invert {
				return '0'
			}
			return '1'
		}

		if k > total/2 {
			k = total - k - 1
			invert = !invert
		}
	}

	if invert {
		if s[k] == '1' {
			return '0'
		}
		return '1'
	}

	return s[k]
}

//	problems
//	1.	inspired from https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/discuss/780984/Java-Recursive-Solution

//		2 to the power of n can be done by 1<<n

//		It's a beautiful solution, how author finds the form of
//		recursion?

//		I also found some clues to the problem, e.g. checks for
//		k > size/2, but didn't find the form of this problem
