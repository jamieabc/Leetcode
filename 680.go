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
	length := len(s)
	if length <= 2 {
		return true
	}

	i := 0
	j := length - 1
	deleted := false
	dual := false
	var tmpI, tmpJ int
	for i < j {
		if s[i] != s[j] {
			if deleted {
				if dual {
					i = tmpI
					j = tmpJ - 1
					dual = false
					continue
				}
				return false
			}

			deleted = true
			if s[i+1] == s[j] && s[i] == s[j-1] {
				dual = true
				tmpI = i
				tmpJ = j
				i++
			} else if s[i+1] == s[j] {
				i++
			} else {
				j--
			}
		} else {
			i++
			j--
		}
	}
	return true
}
