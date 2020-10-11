package main

// You are given two strings a and b of the same length. Choose an index and split both strings at the same index, splitting a into two strings: aprefix and asuffix where a = aprefix + asuffix, and splitting b into two strings: bprefix and bsuffix where b = bprefix + bsuffix. Check if aprefix + bsuffix or bprefix + asuffix forms a palindrome.
//
// When you split a string s into sprefix and ssuffix, either ssuffix or sprefix is allowed to be empty. For example, if s = "abc", then "" + "abc", "a" + "bc", "ab" + "c" , and "abc" + "" are valid splits.
//
// Return true if it is possible to form a palindrome string, otherwise return false.
//
// Notice that x + y denotes the concatenation of strings x and y.
//
//
//
// Example 1:
//
// Input: a = "x", b = "y"
// Output: true
// Explaination: If either a or b are palindromes the answer is true since you can split in the following way:
// aprefix = "", asuffix = "x"
// bprefix = "", bsuffix = "y"
// Then, aprefix + bsuffix = "" + "y" = "y", which is a palindrome.
//
// Example 2:
//
// Input: a = "abdef", b = "fecab"
// Output: true
//
// Example 3:
//
// Input: a = "ulacfd", b = "jizalu"
// Output: true
// Explaination: Split them at index 3:
// aprefix = "ula", asuffix = "cfd"
// bprefix = "jiz", bsuffix = "alu"
// Then, aprefix + bsuffix = "ula" + "alu" = "ulaalu", which is a palindrome.
//
// Example 4:
//
// Input: a = "xbdef", b = "xecab"
// Output: false
//
//
//
// Constraints:
//
//     1 <= a.length, b.length <= 105
//     a.length == b.length
//     a and b consist of lowercase English letters

func checkPalindromeFormation(a string, b string) bool {
	return mixPalindrome(a, b) || mixPalindrome(b, a)
}

func mixPalindrome(a, b string) bool {
	for low, high := 0, len(a)-1; low < high; low, high = low+1, high-1 {
		if a[low] != b[high] {
			return isPalindrome(a[low:high+1]) || isPalindrome(b[low:high+1])
		}
	}

	return true
}

func isPalindrome(str string) bool {
	for low, high := 0, len(str)-1; low < high; low, high = low+1, high-1 {
		if str[low] != str[high] {
			return false
		}
	}

	return true
}

//	Notes
//	1.	most critical thinking: for a given comparison, there might exist 2
//		conditions to check if remaining forms a palindrome string

//		a: _ _ _ _ x x x x x x x x x
//		b: x x x x o o o o o _ _ _ _

//		a: _ _ _ _ o o o o o x x x x
//		b: x x x x x x x x x _ _ _ _

// 		_: already compared palindrome
//		x: don't care
//		o: check if this range is palindrome
