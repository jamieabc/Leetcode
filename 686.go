package main

import (
	"fmt"
	"strings"
)

//Given two strings A and B, find the minimum number of times A has to be repeated such that B is a substring of it. If no such solution, return -1.
//
//For example, with A = "abcd" and B = "cdabcdab".
//
//Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it; and B is not a substring of A repeated two times ("abcdabcd").
//
//Note:
//The length of A and B will be between 1 and 10000.

func repeatedStringMatch(A string, B string) int {
	lenA := len(A)
	lenB := len(B)

	if lenB == 0 {
		return 0
	}

	if lenB != 0 && lenA == 0 {
		return -1
	}

	if A == B {
		return 1
	}

	var index, count int
	var sb strings.Builder
	var str string
	duplicates := fmt.Sprintf("%s%s", A, A)

	if lenA < lenB {
		for index = 0; index < len(A); index++ {
			if duplicates[index:index+len(A)] == B[:len(A)] {
				break
			}
		}

		if index == len(A) {
			return -1
		}

		remain := lenB - (lenA - index)
		count = remain / lenA
		if remain%lenA != 0 {
			count++
		}
		count++
		sb.WriteString(A[index:])
		for i := 0; i < count-1; i++ {
			sb.WriteString(A)
		}
		str = sb.String()
		if str[:lenB] == B {
			return count
		} else {
			return -1
		}
	} else {
		for index = 0; index < lenA; index++ {
			if B == duplicates[index:index+lenB] {
				break
			}
		}

		if index == lenA {
			return -1
		}

		if index+lenB > lenA {
			count = 2
		} else {
			count = 1
		}
		return count
	}
}
