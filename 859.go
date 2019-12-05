package main

//Given two strings A and B of lowercase letters, return true if and only if we can swap two letters in A so that the result equals B.
//
//
//Example 1:
//
//Input: A = "ab", B = "ba"
//Output: true
//
//Example 2:
//
//Input: A = "ab", B = "ab"
//Output: false
//
//Example 3:
//
//Input: A = "aa", B = "aa"
//Output: true
//
//Example 4:
//
//Input: A = "aaaaaaabc", B = "aaaaaaacb"
//Output: true
//
//Example 5:
//
//Input: A = "", B = "aa"
//Output: false
//
//
//
//Note:
//
//    0 <= A.length <= 20000
//    0 <= B.length <= 20000
//    A and B consist only of lowercase letters.

func buddyStrings(A string, B string) bool {
	differences := make([]int, 0)

	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			differences = append(differences, i)
		}
	}

	// more than 2 chars different
	if len(differences) > 2 {
		return false
	}

	// 2 chars different
	if len(differences) == 2 {
		return A[differences[0]] == B[differences[1]] && A[differences[1]] == B[differences[0]]
	}

	// exactly same, find any duplicate
	m := make(map[rune]bool)
	for _, r := range A {
		_, ok := m[r]
		if ok {
			return true
		} else {
			m[r] = true
		}
	}
	return false
}
