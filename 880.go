package main

// An encoded string S is given.  To find and write the decoded string to a tape, the encoded string is read one character at a time and the following steps are taken:
//
//     If the character read is a letter, that letter is written onto the tape.
//     If the character read is a digit (say d), the entire current tape is repeatedly written d-1 more times in total.
//
// Now for some encoded string S, and an index K, find and return the K-th letter (1 indexed) in the decoded string.
//
//
//
// Example 1:
//
// Input: S = "leet2code3", K = 10
// Output: "o"
// Explanation:
// The decoded string is "leetleetcodeleetleetcodeleetleetcode".
// The 10th letter in the string is "o".
//
// Example 2:
//
// Input: S = "ha22", K = 5
// Output: "h"
// Explanation:
// The decoded string is "hahahaha".  The 5th letter is "h".
//
// Example 3:
//
// Input: S = "a2345678999999999999999", K = 1
// Output: "a"
// Explanation:
// The decoded string is "a" repeated 8301530446056247680 times.  The 1st letter is "a".
//
//
//
// Constraints:
//
//     2 <= S.length <= 100
//     S will only contain lowercase letters and digits 2 through 9.
//     S starts with a letter.
//     1 <= K <= 10^9
//     It's guaranteed that K is less than or equal to the length of the decoded string.
//     The decoded string is guaranteed to have less than 2^63 letters.

func decodeAtIndex(S string, K int) string {
	var size int
	for i := range S {
		if S[i] >= 'a' && S[i] <= 'z' {
			size++
		} else {
			size *= int(S[i] - '0')
		}
	}

	var i int
	for i = len(S) - 1; i >= 0; i-- {
		K = K % size

		if K == 0 {
			break
		}

		if S[i] >= 'a' && S[i] <= 'z' {
			size--
		} else {
			size /= int(S[i] - '0')
		}
	}

	for ; i >= 0; i-- {
		if S[i] >= 'a' && S[i] <= 'z' {
			return string(S[i])
		}
	}

	return ""
}

func decodeAtIndex(S string, K int) string {
	var idx, size int

	for K > 0 {
		for size, idx = 0, 0; idx < len(S); idx++ {
			if S[idx] >= 'a' && S[idx] <= 'z' {
				size++
			} else {
				next := size * int(S[idx]-'0')

				if next > K {
					K = K % size
					break
				}

				size = next
			}

			if size == K {
				K = 0
				break
			}
		}
	}

	for j := idx; j >= 0; j-- {
		if S[j] >= 'a' && S[j] <= 'z' {
			return string(S[j])
		}
	}

	return ""
}

//	Notes
//	1.	since decoded string could be length of 2^63, it's not practical to
//		generate string an find char

//	2.	inspired from solution, count backward is also workable

//	3.	inspired from https://leetcode.com/problems/decoded-string-at-index/discuss/156747/C%2B%2BPython-O(N)-Time-O(1)-Space

//		can only count size until >= K
