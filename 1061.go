package main

import "strings"

// Given strings A and B of the same length, we say A[i] and B[i] are equivalent characters. For example, if A = "abc" and B = "cde", then we have 'a' == 'c', 'b' == 'd', 'c' == 'e'.
//
// Equivalent characters follow the usual rules of any equivalence relation:
//
//     Reflexivity: 'a' == 'a'
//     Symmetry: 'a' == 'b' implies 'b' == 'a'
//     Transitivity: 'a' == 'b' and 'b' == 'c' implies 'a' == 'c'
//
// For example, given the equivalency information from A and B above, S = "eed", "acd", and "aab" are equivalent strings, and "aab" is the lexicographically smallest equivalent string of S.
//
// Return the lexicographically smallest equivalent string of S by using the equivalency information from A and B.
//
//
//
// Example 1:
//
// Input: A = "parker", B = "morris", S = "parser"
// Output: "makkek"
// Explanation: Based on the equivalency information in A and B, we can group their characters as [m,p], [a,o], [k,r,s], [e,i]. The characters in each group are equivalent and sorted in lexicographical order. So the answer is "makkek".
//
// Example 2:
//
// Input: A = "hello", B = "world", S = "hold"
// Output: "hdld"
// Explanation:  Based on the equivalency information in A and B, we can group their characters as [h,w], [d,e,o], [l,r]. So only the second letter 'o' in S is changed to 'd', the answer is "hdld".
//
// Example 3:
//
// Input: A = "leetcode", B = "programs", S = "sourcecode"
// Output: "aauaaaaada"
// Explanation:  We group the equivalent characters in A and B as [a,o,e,r,s,c], [l,p], [g,t] and [d,m], thus all letters in S except 'u' and 'd' are transformed to 'a', the answer is "aauaaaaada".
//
//
//
// Note:
//
//     String A, B and S consist of only lowercase English letters from 'a' - 'z'.
//     The lengths of string A, B and S are between 1 and 1000.
//     String A and B are of the same length.

func smallestEquivalentString(A string, B string, S string) string {
	mapping := make([]int, 26)
	for i := range mapping {
		mapping[i] = i
	}

	for i := range A {
		a, b := find(mapping, toIndex(A[i])), find(mapping, toIndex(B[i]))
		if a <= b {
			mapping[b] = a
		} else {
			mapping[a] = b
		}
	}

	var sb strings.Builder

	for i := range S {
		sb.WriteByte(byte(find(mapping, toIndex(S[i])) + 'a'))
	}

	return sb.String()
}

func toIndex(b byte) int {
	return int(b - 'a')
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func find(mapping []int, i int) int {
	if mapping[i] != i {
		mapping[i] = find(mapping, mapping[i])
	}
	return mapping[i]
}
