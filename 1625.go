package main

import "strings"

// You are given a string s of even length consisting of digits from 0 to 9, and two integers a and b.
//
// You can apply either of the following two operations any number of times and in any order on s:
//
//     Add a to all odd indices of s (0-indexed). Digits post 9 are cycled back to 0. For example, if s = "3456" and a = 5, s becomes "3951".
//     Rotate s to the right by b positions. For example, if s = "3456" and b = 1, s becomes "6345".
//
// Return the lexicographically smallest string you can obtain by applying the above operations any number of times on s.
//
// A string a is lexicographically smaller than a string b (of the same length) if in the first position where a and b differ, string a has a letter that appears earlier in the alphabet than the corresponding letter in b. For example, "0158" is lexicographically smaller than "0190" because the first position they differ is at the third letter, and '5' comes before '9'.
//
//
//
// Example 1:
//
// Input: s = "5525", a = 9, b = 2
// Output: "2050"
// Explanation: We can apply the following operations:
// Start:  "5525"
// Rotate: "2555"
// Add:    "2454"
// Add:    "2353"
// Rotate: "5323"
// Add:    "5222"
// Add:    "5121"
// Rotate: "2151"
// Add:    "2050"
// There is no way to obtain a string that is lexicographically smaller then "2050".
//
// Example 2:
//
// Input: s = "74", a = 5, b = 1
// Output: "24"
// Explanation: We can apply the following operations:
// Start:  "74"
// Rotate: "47"
// Add:    "42"
// Rotate: "24"
// There is no way to obtain a string that is lexicographically smaller then "24".
//
// Example 3:
//
// Input: s = "0011", a = 4, b = 2
// Output: "0011"
// Explanation: There are no sequence of operations that will give us a lexicographically smaller string than "0011".
//
// Example 4:
//
// Input: s = "43987654", a = 7, b = 3
// Output: "00553311"
//
//
//
// Constraints:
//
//     2 <= s.length <= 100
//     s.length is even.
//     s consists of digits from 0 to 9 only.
//     1 <= a <= 9
//     1 <= b <= s.length - 1

// tc: O(mnd), m: # of rotations, n: # of add operations for a char(at most 10),
// d: # of odd digits
// for every rotation, try every digit's add operation
func findLexSmallestString(s string, a int, b int) string {
	size := len(s)
	idx := size - b
	visited := make(map[string]bool)
	queue := [][]byte{[]byte(s)}

	smallest := make([]byte, size)
	copy(smallest, s)

	for len(queue) > 0 {
		bs := queue[0]
		queue = queue[1:]

		if visited[string(bs)] {
			continue
		}
		visited[string(bs)] = true

		next1 := append([]byte{}, bs[idx:]...)
		next1 = append(next1, bs[:idx]...)
		next2 := add(bs, a)

		if strings.Compare(string(next1), string(smallest)) < 0 {
			copy(smallest, next1)
		}

		if strings.Compare(string(next2), string(smallest)) < 0 {
			copy(smallest, next2)
		}

		queue = append(queue, next1, next2)
	}

	return string(smallest)
}

func add(tmp []byte, a int) []byte {
	size := len(tmp)
	for i := 1; i < size; i += 2 {
		tmp[i] = byte('0' + (int(tmp[i]-'0')+a)%10)
	}

	return tmp
}

//	Notes
//	1.	add operation is applied to all odd digits, not just one digit

//	2.	lexicographical order cares left most character, because it determines
//		the smallest order

//		e.g. 55556

//	3.	recursive call sometimes hard to check for time complexity

//	4.	string length is even, if b is even number, only odd index are changeable

//		e.g. 1 2 3 4 5 6, b = 2
// 		       ^   ^   ^  changeable
//      rotation 1: 5 6 1 2 3 4
//                    ^   ^   ^  changeable
//		rotation 2: 3 4 5 6 1 2
//		              ^   ^   ^  changeable

//		if b is odd number, both even and odd are changeable

//		e.g. 1 2 3 4 5 6 7 8 9 10, b = 3
//		       ^   ^   ^   ^   ^  changeable
//		rotation 1: 8 9 10 1 2 3 4 5 6 7
//                    ^    ^   ^   ^   ^  changeable
//		rotation 2: 5 6 7 8 9 10 1 2 3 4
//                    ^   ^   ^    ^   ^  changeable
//		rotation 3: 2 3 4 5 6 7 8 9 10 1
//					  ^   ^   ^   ^    ^  changeable

//	5.	inspired from https://leetcode.com/problems/lexicographically-smallest-string-after-applying-operations/discuss/899708/Essential-View%3A-Group-Theory-or-Java-14ms-100-or-O(s)-time-O(1)-space

//		a better explanation
