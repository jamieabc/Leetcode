package main

// Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.
//
//Example 1:
//
//Input: "aba"
//Output: True
//
//Example 2:
//
//Input: "abca"
//Output: True
//Explanation: You could delete the character 'c'.
//
//Note:
//
//    The string will only contain lowercase characters a-z. The maximum length of the string is 50000.

func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return check(s, i+1, j) || check(s, i, j-1)
		}
	}
	return true
}

func check(s string, left, right int) bool {
	for ; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

//	problems
//	1.	optimize, reduce memory usage
//	2.	optimize, simplify code
