package main

import "strings"

// S and T are strings composed of lowercase letters. In S, no letter occurs more than once.
//
// S was sorted in some custom order previously. We want to permute the characters of T so that they match the order that S was sorted. More specifically, if x occurs before y in S, then x should occur before y in the returned string.
//
// Return any permutation of T (as a string) that satisfies this property.
//
// Example :
// Input:
// S = "cba"
// T = "abcd"
// Output: "cbad"
// Explanation:
// "a", "b", "c" appear in S, so the order of "a", "b", "c" should be "c", "b", and "a".
// Since "d" does not appear in S, it can be at any position in T. "dcba", "cdba", "cbda" are also valid outputs.
//
//
// Note:
//
// S has length at most 26, and no character is repeated in S.
// T has length at most 200.
// S and T consist of lowercase letters only.

func customSortString(S string, T string) string {
	arr := make([]int, 26)
	var sb strings.Builder

	for i := range T {
		arr[T[i]-'a']++
	}

	for i := range S {
		for arr[S[i]-'a'] > 0 {
			sb.WriteByte(S[i])
			arr[S[i]-'a']--
		}
	}

	for i := range arr {
		for arr[i] > 0 {
			sb.WriteByte(byte('a' + i))
			arr[i]--
		}
	}

	return sb.String()
}

//	problems
//	1.	from reference https://leetcode.com/problems/custom-sort-string/discuss/116573/Java-Bucket-sort-solution-O(N%2BM)-with-follow-up-questions

//		arr1 is not needed, because it's already sorted.

//		also, there's some follow ups worth thinking:

//		- If the custom order S is too large, how to solve it efficiently?

//			if S is too large, and S is guaranteed to be non-repeated, so it's possible to cut S into
//		parts, then T is sorted by that part, and combined them together

//			e.g. S = "abcdef", S1 = "abc", S2 = "def"
//			T = "aabbccddeeffzx", T1 sorted by S1 = abc + abcddeeffzx
//			T1 sorted by S2 = abc + def + abcdefzx

//		- If the string T is too large, how to solve it efficiently?

//			if T is too large, separate T into parts, each part sorted by S, then combine them in order

//		- What if repetitions are allowed in S such that relative ordering in S and final answer remain the same?

//			e.g. S = "cbac", T = "abcdabcd", result should be "cabcabdd", ab should reside between c

//		thanks for sharing, it helps a lot.
