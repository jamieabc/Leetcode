package main

// Given an integer n, return the number of strings of length n that consist only of vowels (a, e, i, o, u) and are lexicographically sorted.
//
// A string s is lexicographically sorted if for all valid i, s[i] is the same as or comes before s[i+1] in the alphabet.
//
//
//
// Example 1:
//
// Input: n = 1
// Output: 5
// Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].
//
// Example 2:
//
// Input: n = 2
// Output: 15
// Explanation: The 15 sorted strings that consist of vowels only are
// ["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
// Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.
//
// Example 3:
//
// Input: n = 33
// Output: 66045
//
//
//
// Constraints:
//
//     1 <= n <= 50

// tc: O(n)
func countVowelStrings(n int) int {
	// dp[i] means count of count of words ends at a, e, i, o, u
	dp := make([]int, 5)
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < n-1; i++ {
		next := make([]int, 5)

		for j := 0; j < 5; j++ {
			for k := 0; k <= j; k++ {
				next[j] += dp[k]
			}

		}

		dp = next
	}

	var count int
	for i := range dp {
		count += dp[i]
	}
	return count
}

// tc: O(n^4)
func countVowelStrings1(n int) int {
	return perm(n, 0)
}

// a, e, i, o, u
func perm(n, last int) int {
	if n == 0 {
		return 1
	}

	var count int
	for i := last; i < 5; i++ {
		count += perm(n-1, i)
	}

	return count
}

//	Notes
//	1.	for a words length i ends at c, total count = length i-1 and ending
//		<= c

//		count(length = 5, ends at e) = count(length = 4, ends at a) +
//									   count(length = 4, ends at e)

//		this relationships causes dp
