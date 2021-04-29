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

func checkPalindromeFormation2(a string, b string) bool {
	return compare2(a, b) || compare2(b, a)
}

func compare2(a, b string) bool {
	size := len(a)
	var i, j int

	if size&1 > 0 {
		i, j = size/2, size/2
	} else {
		i, j = size/2-1, size/2
	}

	for ; i >= 0; i, j = i-1, j+1 {
		if a[i] != a[j] {
			return isPalindrome2(a[:i+1], b[j:]) || isPalindrome2(b[:i+1], a[j:])
		}
	}

	return true
}

func isPalindrome2(s1, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[len(s2)-1-i] {
			return false
		}
	}

	return true
}

func checkPalindromeFormation1(a string, b string) bool {
	size := len(a)

	if size <= 1 {
		return true
	}

	// odd length
	if size&1 > 0 {
		return compare(a, b, size/2, size/2) || compare(b, a, size/2, size/2)
	}

	return compare(a, b, size/2-1, size/2) || compare(b, a, size/2-1, size/2)
}

func compare(src, dst string, idx1, idx2 int) bool {
	var i int
	size := len(src)

	for ; i+idx2 < size; i++ {
		if src[idx1-i] != src[idx2+i] {
			break
		}
	}

	pos := i
	for ; i+idx2 < size; i++ {
		if dst[idx1-i] != src[idx2+i] {
			break
		}
	}

	if i+idx2 == size {
		return true
	}

	for i = pos; i+idx2 < size; i++ {
		if src[idx1-i] != dst[idx2+i] {
			break
		}
	}

	return i+idx2 == size
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

//	2.	for palindrome, use progressive way to check, start from middle point
//		and compares one by one till boundary

//		the point for this problem is that size is always the same, which makes
//		middle of palindrome either at a or b

//		start from a, try to match as many as same character as possible; if
//		there's any mismatch, try b to see if can reach to end

//	3.	the smart way (alex) to think is from pattern

//		there are 4 conditions to check, can be separated into 2 categories

//		- middle at a
//			a -----					   -----
//
//			b      ---				---

//		- middle at b
//			a ---					     ---
//
//			b   -----				-----

//		the smarter way is to check only for middle at a or middle at b
//		start from that middle, try to expand until left != right, then compare
//		rest of string at a & b separately

//		what i think of yesterday is also O(n), but this one is more clever
