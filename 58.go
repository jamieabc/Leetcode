package main

//Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.
//
//If the last word does not exist, return 0.
//
//Note: A word is defined as a character sequence consists of non-space characters only.
//
//Example:
//
//Input: "Hello World"
//Output: 5

// ''
// ' '
// 'abcde  '
// 'abcde'
// '  abc def ghi  '
func lengthOfLastWord(s string) int {
	length := len(s)

	if length == 0 {
		return 0
	}

	// last non-space character
	var trimmedIndex int
	for trimmedIndex = length - 1; trimmedIndex >= 0; trimmedIndex-- {
		if s[trimmedIndex] != ' ' {
			break
		}
	}

	if trimmedIndex < 0 {
		return 0
	}

	if trimmedIndex == 0 {
		return 1
	}

	var i int
	for i = trimmedIndex; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
	}

	return trimmedIndex - i
}
