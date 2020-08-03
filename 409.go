package main

// Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

// This is case sensitive, for example "Aa" is not considered a palindrome here.

// Note:
// Assume the length of given string will not exceed 1,010.

// Example:

// Input:
// "abccccdd"

// Output:
// 7

// Explanation:
// One longest palindrome that can be built is "dccaccd", whose length is 7.

// for every even word count, must be able to generate palindrome
// if odd char exist, plus 1 as center of palindrome
// aabbccdde => a:2, b:2, c:2, d:2, e:1 => max: 9
// aaa => a:3 => max: 3
// abc => a:1, b:1, c:1 => max: 1
// aaac => a:3, c:1 => max: 3
// aaabbb => a:3, b:3 => max: 5 (bbaaabb)
func longestPalindrome(s string) int {
	counter := make([]int, 52)

	// calculate each char count
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			counter[s[i]-'a']++
		} else {
			counter[26+s[i]-'A']++
		}
	}

	var maxSize int
	var oddCount bool

	for _, count := range counter {
		if count&1 > 0 {
			oddCount = true
			maxSize += count - 1
		} else {
			maxSize += count
		}
	}

	if oddCount {
		return maxSize + 1
	}

	return maxSize
}

//	problems
//	1.	it contains both upper & lower case
