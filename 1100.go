package main

// Given a string S, return the number of substrings of length K with no repeated characters.
//
//
//
// Example 1:
//
// Input: S = "havefunonleetcode", K = 5
// Output: 6
// Explanation:
// There are 6 substrings they are : 'havef','avefu','vefun','efuno','etcod','tcode'.
//
// Example 2:
//
// Input: S = "home", K = 5
// Output: 0
// Explanation:
// Notice K can be larger than the length of S. In this case is not possible to find any substring.
//
//
//
// Note:
//
//     1 <= S.length <= 10^4
//     All characters of S are lowercase English letters.
//     1 <= K <= 10^4

func numKLenSubstrNoRepeats(S string, K int) int {
	chars := make([]int, 26)
	lenS := len(S)

	var result int

	if K > lenS || K == 0 {
		return result
	}

	// put first char in
	chars[S[0]-'a']++

	for low, high := 0, 0; high < lenS; {
		if high-low+1 < K {
			high++

			// in case over bound
			if high == lenS {
				break
			}

			chars[S[high]-'a']++
		}

		// make sure sliding window in low-high range has only unique chars
		// the trick here is to use high as a checking, make sure low is moving
		// forward to have no duplicates in range
		for chars[S[high]-'a'] > 1 {
			chars[S[low]-'a']--
			low++
		}

		// check if this sliding windows meets K
		if high-low+1 == K {
			result++

			chars[S[low]-'a']--
			low++
		}
	}

	return result
}

//	problems
//	1.	inspired from reference https://leetcode.com/problems/find-k-length-substrings-with-no-repeated-characters/discuss/322982/Java-Sliding-Window-two-O(n)-codes-w-comments-and-analysis.

//		My thinking is to use as sliding window and another array that stores number of chars in
//		that windows. When every a char is added, loop through array to check if any duplicates.
//		Although it's working, there's a repeated operation of looping array to check.

//		The other solution is reverted the flow, for a given index, find it's
//		range of low - high that makes this sliding window unique.
//		I think this one is better because it's more reasonable to check only once w/o looping.

//		The point for this method is to make sure every high-low range with unique chars.