package main

// Given two equal-size strings s and t. In one step you can choose any character of t and replace it with another character.
//
// Return the minimum number of steps to make t an anagram of s.
//
// An Anagram of a string is a string that contains the same characters with a different (or the same) ordering.
//
//
//
// Example 1:
//
// Input: s = "bab", t = "aba"
// Output: 1
// Explanation: Replace the first 'a' in t with b, t = "bba" which is anagram of s.
//
// Example 2:
//
// Input: s = "leetcode", t = "practice"
// Output: 5
// Explanation: Replace 'p', 'r', 'a', 'i' and 'c' from t with proper characters to make t anagram of s.
//
// Example 3:
//
// Input: s = "anagram", t = "mangaar"
// Output: 0
// Explanation: "anagram" and "mangaar" are anagrams.
//
// Example 4:
//
// Input: s = "xxyyzz", t = "xxyyzz"
// Output: 0
//
// Example 5:
//
// Input: s = "friend", t = "family"
// Output: 4
//
//
//
// Constraints:
//
//     1 <= s.length <= 50000
//     s.length == t.length
//     s and t contain lower-case English letters only.

func minSteps(s string, t string) int {
	arr := make([]int, 26)

	for i := range s {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	var steps int
	for _, i := range arr {
		steps += abs(i)
	}

	return steps / 2
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	problems
//	1. 	too slow, the reason is using map with too many operations, since it's only
//		lower case characters, I can have a array with length 26, iterate through
//		2 strings and have their sum of characters and compare.

//		complexity is O(n)

//	2.	from other people's solution, it can be further improved, use single array
//		only. one string operation on array is plus, the other string operation on
//		array is -, thus have the difference count. But this value is duplicated,
//		so it could be summed and divided by 2.

//		reference: https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/discuss/503450/JavaPython-3-Count-occurrences-and-sum-the-difference-w-analysis.
