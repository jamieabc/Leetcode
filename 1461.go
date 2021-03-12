package main

// Given a binary string s and an integer k.
//
// Return True if every binary code of length k is a substring of s. Otherwise, return False.
//
//
//
// Example 1:
//
// Input: s = "00110110", k = 2
// Output: true
// Explanation: The binary codes of length 2 are "00", "01", "10" and "11". They can be all found as substrings at indicies 0, 1, 3 and 2 respectively.
//
// Example 2:
//
// Input: s = "00110", k = 2
// Output: true
//
// Example 3:
//
// Input: s = "0110", k = 1
// Output: true
// Explanation: The binary codes of length 1 are "0" and "1", it is clear that both exist as a substring.
//
// Example 4:
//
// Input: s = "0110", k = 2
// Output: false
// Explanation: The binary code "00" is of length 2 and doesn't exist in the array.
//
// Example 5:
//
// Input: s = "0000000001011100", k = 4
// Output: false
//
//
//
// Constraints:
//
//     1 <= s.length <= 5 * 10^5
//     s consists of 0's and 1's only.
//     1 <= k <= 20

func hasAllCodes(s string, k int) bool {
	size := len(s)
	if size < k {
		return false
	}

	total := 1 << k
	table := make([]bool, total)
	mask := (1 << k) - 1
	var cur int

	for i := range s {
		cur = ((cur << 1) | int(s[i]-'0')) & mask
		if i >= k-1 {
			if !table[cur] {
				total--
				table[cur] = true
			}
		}
	}

	return total == 0
}

func hasAllCodes2(s string, k int) bool {
	size := len(s)

	if size < k {
		return false
	}

	total := 1 << k
	table := make([]bool, total)
	mask := (1 << k) - 1
	var cur int

	// setup initial status, becareful, loop will do left shift,
	// so here to keep everything before operation
	for i := 0; i < k-1; i++ {
		if s[i] == '1' {
			cur |= 1 << (k - i - 2)
		}
	}

	for i := k - 1; i < len(s); i++ {
		cur = (cur << 1) & mask

		if s[i] == '1' {
			cur |= 1
		}

		if !table[cur] {
			total--
			table[cur] = true
		}
	}

	return total == 0
}

func hasAllCodes1(s string, k int) bool {
	table := make(map[string]bool)

	for i := k - 1; i < len(s); i++ {
		table[s[i-k+1:i+1]] = true
	}

	return len(table) == 1<<k
}

//	Notes
//	1.	inspired from sample code, use []bool should be faster than map

//		also, use bit operation can be faster

//	2.	bit-wise operation needs more check, out of boundary, etc

//		since I assume initial bit should be shift to correct place, so need
//		to make sure s length > k

//	3.	inspired from sample code, use curr = ((curr << 1) | int(s[i] - '0')) & mask
//		to do shift & mask, so clever

//		also, can combine setup part & process part together

//	4.	inspired from https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/discuss/660829/Python-1-line-Follow-Up

//		lee has a good follow-up, change substring to subsequence, which is more broadly
//		applied

//		i don't know how to solved it at the first place, but comment with great solution,
//		key idea is that if all subsequence exist, then every 0/1 pair reduces k by 1

//		0 1 x x x x x
//		no matter what x is, 0 & 1 eliminates k by 1 bit, because 0x 1x makes it work,
//		very smart

//		comment here https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/discuss/660829/Python-1-line-Follow-Up/562031
