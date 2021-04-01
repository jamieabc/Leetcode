package main

// You want to form a target string of lowercase letters.
//
// At the beginning, your sequence is target.length '?' marks.  You also have a stamp of lowercase letters.
//
// On each turn, you may place the stamp over the sequence, and replace every letter in the sequence with the corresponding letter from the stamp.  You can make up to 10 * target.length turns.
//
// For example, if the initial sequence is "?????", and your stamp is "abc",  then you may make "abc??", "?abc?", "??abc" in the first turn.  (Note that the stamp must be fully contained in the boundaries of the sequence in order to stamp.)
//
// If the sequence is possible to stamp, then return an array of the index of the left-most letter being stamped at each turn.  If the sequence is not possible to stamp, return an empty array.
//
// For example, if the sequence is "ababc", and the stamp is "abc", then we could return the answer [0, 2], corresponding to the moves "?????" -> "abc??" -> "ababc".
//
// Also, if the sequence is possible to stamp, it is guaranteed it is possible to stamp within 10 * target.length moves.  Any answers specifying more than this number of moves will not be accepted.
//
//
//
// Example 1:
//
// Input: stamp = "abc", target = "ababc"
// Output: [0,2]
// ([1,0,2] would also be accepted as an answer, as well as some other answers.)
//
// Example 2:
//
// Input: stamp = "abca", target = "aabcaca"
// Output: [3,0,1]
//
//
//
// Note:
//
// 1 <= stamp.length <= target.length <= 1000
// stamp and target only contain lowercase letters.

// not workgin
func movesToStamp(stamp string, target string) []int {
	// abc, *bc, ab*, **c, a**
	masks := []string{stamp}
	for i := 1; i < len(stamp); i++ {
		tmp := make([]byte, len(stamp))
		for j := 0; j < i; j++ {
			tmp[j] = '*'
		}
		copy(tmp[i:], stamp[i:])
		masks = append(masks, string(tmp))

		tmp1 := make([]byte, len(stamp))
		for j := 0; j < i; j++ {
			tmp1[len(stamp)-1-j] = '*'
		}
		copy(tmp1[:len(stamp)-i], stamp[:len(stamp)-i])
		masks = append(masks, string(tmp1))
	}

	arr := make([]int, 0)
	str := []byte(target)

	for i := range masks {
		for j := 0; j <= len(target)-len(stamp); j++ {
			if masks[i] == string(str[j:j+len(stamp)]) {
				arr = append(arr, j)
				for k := 0; k < len(stamp); j, k = j+1, k+1 {
					str[j] = '*'
				}

				// make sure j is continuous
				j--
			}
		}
	}

	ans := make([]int, len(arr))
	for i := range ans {
		ans[i] = arr[len(arr)-1-i]
	}

	return ans
}

//	Notes
//	1.	cannot think of any solution, it i do in greedy, as long as left/right
//		not overlap, it's okay

//	2.	if left overlap with right, need to check it's order

//	3.	I though it is lcs, but seems like not the case...

//	4.	inspired from https://leetcode.com/problems/stamping-the-sequence/discuss/189576/C%2B%2B-simple-greedy

//		very interesting technique, encode state and use it to find possible
//		matching...

//		the key point here is that later operation covers previous operation,
//		so if there's a full match in target string, then that position must
//		be last applied

//		the other thing is that later operation full covers previous operations
//		chars, such that when check by *, only starts from beginning or end, but
//		not in between

//		e.g. abcd => *bcd, abc* => **cd, ab** => ***d, a***

//		e.g. 		abcd
//				 abcd
//				abc* + abcd

//				abcd
//				  abcd
//				abcd + **cd
