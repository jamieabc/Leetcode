package main

// You are given two strings a and b that consist of lowercase letters. In one operation, you can change any character in a or b to any lowercase letter.
//
// Your goal is to satisfy one of the following three conditions:
//
//     Every letter in a is strictly less than every letter in b in the alphabet.
//     Every letter in b is strictly less than every letter in a in the alphabet.
//     Both a and b consist of only one distinct letter.
//
// Return the minimum number of operations needed to achieve your goal.
//
//
//
// Example 1:
//
// Input: a = "aba", b = "caa"
// Output: 2
// Explanation: Consider the best way to make each condition true:
// 1) Change b to "ccc" in 2 operations, then every letter in a is less than every letter in b.
// 2) Change a to "bbb" and b to "aaa" in 3 operations, then every letter in b is less than every letter in a.
// 3) Change a to "aaa" and b to "aaa" in 2 operations, then a and b consist of one distinct letter.
// The best way was done in 2 operations (either condition 1 or condition 3).
//
// Example 2:
//
// Input: a = "dabadd", b = "cda"
// Output: 3
// Explanation: The best way is to make condition 1 true by changing b to "eee".
//
//
//
// Constraints:
//
//     1 <= a.length, b.length <= 105
//     a and b consist only of lowercase letters.

func minCharacters(a string, b string) int {
	if a == b {
		return 0
	}

	sizeA, sizeB := len(a), len(b)

	// find char counter for both strings
	counter1, counter2 := make([]int, 26), make([]int, 26)

	for i := range a {
		counter1[a[i]-'a']++
	}

	for i := range b {
		counter2[b[i]-'a']++
	}

	minOps := max(sizeA, sizeB)

	// a == b: iterate through all chars c, sum characters not equal to c
	for i := range counter1 {
		minOps = min(minOps, sizeA+sizeB-counter1[i]-counter2[i])
	}

	// a < b: iterate from smallest char c,
	// sum char count that a >= c, b < c
	// a > b: iterate from smallest char c,
	// sum char count that a < c, b >= c

	sum1, sum2 := counter1[0], counter2[0]
	for i := 1; i < len(counter1); i++ {
		minOps = min(minOps, sizeA-sum1+sum2)
		minOps = min(minOps, sizeB-sum2+sum1)
		sum1, sum2 = sum1+counter1[i], sum2+counter2[i]
	}

	return minOps
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
