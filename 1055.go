package main

// From any string, we can form a subsequence of that string by deleting some number of characters (possibly no deletions).
//
// Given two strings source and target, return the minimum number of subsequences of source such that their concatenation equals target. If the task is impossible, return -1.
//
//
//
// Example 1:
//
// Input: source = "abc", target = "abcbc"
// Output: 2
// Explanation: The target "abcbc" can be formed by "abc" and "bc", which are subsequences of source "abc".
//
// Example 2:
//
// Input: source = "abc", target = "acdbc"
// Output: -1
// Explanation: The target string cannot be constructed from the subsequences of source string due to the character "d" in target string.
//
// Example 3:
//
// Input: source = "xyz", target = "xzyxz"
// Output: 3
// Explanation: The target string can be constructed as follows "xz" + "y" + "xz".
//
//
//
// Constraints:
//
//     Both the source and target strings consist of only lowercase English letters from "a"-"z".
//     The lengths of source and target string are between 1 and 1000.

func shortestWay(source string, target string) int {
	// create a mapping, at index i, 0 ~ 25 means earliest char occurrence
	// position
	length := len(source)
	mapping := make([][]int, length)
	for i := range mapping {
		mapping[i] = make([]int, 26)
	}

	// for lasts index, find nothing
	for i := range mapping[length-1] {
		mapping[length-1][i] = -1
	}
	mapping[length-1][source[length-1]-'a'] = length - 1

	for i := length - 2; i >= 0; i-- {
		copy(mapping[i], mapping[i+1])
		mapping[i][source[i]-'a'] = i
	}

	var count int

	for i, next := 0, 0; i < len(target); i++ {
		if next == length {
			count++
			next = 0
		}

		if mapping[0][target[i]-'a'] == -1 {
			return -1
		}

		next = mapping[next][target[i]-'a'] + 1

		if next == 0 {
			count++
			i--
		}
	}

	return count + 1
}

func shortestWay3(source string, target string) int {
	mapping := make([][]int, 26)
	for i := range mapping {
		mapping[i] = make([]int, len(source))
	}

	// a's index at 1, 3, 5 => [0, 2, 0, 4, 0, 6]
	for i, c := range source {
		mapping[c-'a'][i] = i + 1
	}

	// update table of a's index 1, 3, 5 => [2, 2, 4, 4, 6, 6]
	for i := range mapping {
		for j, prev := len(source)-1, 0; j >= 0; j-- {
			if mapping[i][j] == 0 {
				mapping[i][j] = prev
			} else {
				prev = mapping[i][j]
			}
		}
	}

	var count int

	for i, next := 0, 0; i < len(target); i++ {
		if next == len(source) {
			next = 0
			count++
		}

		// char not exists
		if mapping[target[i]-'a'][0] == 0 {
			return -1
		}

		next = mapping[target[i]-'a'][next]

		// cannot find char index larger than next
		if next == 0 {
			count++
			i--
		}
	}

	return count + 1
}

func shortestWay2(source string, target string) int {
	mapping := make([][]int, 26)
	for i := range mapping {
		mapping[i] = make([]int, 0)
	}

	for i, c := range source {
		mapping[c-'a'] = append(mapping[c-'a'], i)
	}

	var count, k, l int

	for j := 0; j < len(target); {
		if len(mapping[target[j]-'a']) == 0 {
			return -1
		} else {
			count, j = count+1, j+1
			i := mapping[target[j-1]-'a'][0]

			// binary search
			for j < len(target) {
				tmp := mapping[target[j]-'a']
				if len(tmp) == 0 {
					return -1
				}

				for k, l = 0, len(tmp)-1; k < l; {
					mid := k + (l-k)/2

					if tmp[mid] > i {
						l = mid
					} else {
						k = mid + 1
					}
				}

				if tmp[k] > i {
					j++
					i = tmp[k]
				} else {
					break
				}
			}
		}
	}

	if count == 0 {
		return -1
	}
	return count
}

func shortestWay1(source string, target string) int {
	var count int

	for j := 0; j < len(target); {
		prev := j
		for i := 0; i < len(source) && j < len(target); {
			if source[i] == target[j] {
				j++
			}
			i++
		}

		if prev == j {
			return -1
		} else {
			count++
		}
	}

	if count == 0 {
		return -1
	}
	return count
}

//	problems
//	1.	tc: O(mn)

//	2.	inspired by https://leetcode.com/problems/shortest-way-to-form-string/discuss/330938/Accept-is-not-enough-to-get-a-hire.-Interviewee-4-follow-up

//		additional follow-ups

//	3.	inspired by https://leetcode.com/problems/shortest-way-to-form-string/discuss/304662/Python-O(M-%2B-N*logM)-using-inverted-index-%2B-binary-search-(Similar-to-LC-792)

//		I was thinking about how O(m log n) comes from, it seems like when
//		searching indexes, it's already in sorted order, so search proper
//		index greater than self can be binary search

//		for tc O(n) solution, it further improves the index search by
//		duplicates. e.g. a -> {1, 3, 5, 7, 9}, the reason to search all is
//		because it could have 5 possible locations, and don't know which is
//		the best. This drawback can be improved by dummy data, e.g.
//		a = [1, 3, 3, 5, 5, 7, 7, 9, 9, -1]
//		to know a's position after 6 can be done by a[6]

//	4.	binary search has problem...boundary conditions

//	5.	when using table to store next index, need to take care of start
//		index, when to stop, how to go next index

//	6.	inspired from https://leetcode.com/problems/shortest-way-to-form-string/discuss/330938/Accept-is-not-enough-to-get-a-hire.-Interviewee-4-follow-up

//		author default ith char next to be i+1, means each position is at
//		at least increment by 1, after that update table for each character
//		occurrence index, then updat whole table by occurrence index

//	7.	inspired from https://leetcode.com/problems/shortest-way-to-form-string/discuss/332419/O(M-%2B-N)-Java-solution-with-commented-code-and-detailed-explanation-(Beats-98)

//		add O(n+m) solution
