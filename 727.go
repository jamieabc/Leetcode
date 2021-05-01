package main

// Given strings S and T, find the minimum (contiguous) substring W of S, so that T is a subsequence of W.
//
// If there is no such window in S that covers all characters in T, return the empty string "". If there are multiple such minimum-length windows, return the one with the left-most starting index.
//
// Example 1:
//
// Input:
// S = "abcdebdde", T = "bde"
// Output: "bcde"
// Explanation:
// "bcde" is the answer because it occurs before "bdde" which has the same length.
// "deb" is not a smaller window because the elements of T in the window must occur in order.
//
//
//
// Note:
//
// All the strings in the input will only contain lowercase letters.
// The length of S will be in the range [1, 20000].
// The length of T will be in the range [1, 100].

func minWindow(S string, T string) string {
	size := len(S)

	table := make([][]int, 26)
	for i := range table {
		table[i] = make([]int, size)
	}

	for i := 0; i < 26; i++ {
		// because everytime forward, index will plus 1, make it to -2
		// such that -2+1 = -1 still invalid
		prev := -2

		for j := size - 1; j >= 0; j-- {
			if int(S[j]-'a') == i {
				prev = j
			}
			table[i][j] = prev
		}
	}

	start, end := -1, -1

	for cur := table[T[0]-'a'][0]; cur < size; cur++ {
		to := cur

		for i := range T {
			// to already reaches end, not able to find any further char
			if to == len(S) || to == -1 || cur < 0 {
				to = -1
				cur = size
				break
			}

			// for consecutive characters, need to add additional 1 to make
			// index forward
			to = table[T[i]-'a'][to] + 1
		}

		// only update when valid range is found, and is smaller than previous
		if to != -1 && (start == -1 || to-cur < end-start) {
			start, end = cur, to
		}
	}

	if start == -1 {
		return ""
	}
	return S[start:end]
}

//	Notes
//	1.	subsequence, every word need to have same occurrence order

//	2.	becareful about boundary condition, how to end traverse

//	3.	it could be further improved, always start from first occurrence of
//		T's char

//	4.	inspired from https://leetcode.com/problems/minimum-window-subsequence/discuss/512645/Easy-To-Understand-%3A-Sliding-window-2-pointer-Find-then-Improve

//		author tries to solve the problem by two pointer similar to problem
//		https://leetcode.com/problems/minimum-window-substring/

//		although tc is not good O(s^2), but it provides progressive insight
//		borrow from other problems

//	5.	inspired from https://leetcode.com/problems/minimum-window-subsequence/discuss/109354/Python-O(m)-space-complexity-almost-O(n)-time-complexity

//		interesting dp solution

//	6.	i think it could be further improved by searching only valid positions
//		of T[0], kind of similar to two pointer method
