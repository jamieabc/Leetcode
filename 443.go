package main

import (
	"fmt"
)

//Given an array of characters, compress it in-place.
//
//The length after compression must always be smaller than or equal to the original array.
//
//Every element of the array should be a character (not int) of length 1.
//
//After you are done modifying the input array in-place, return the new length of the array.
//
//
//Follow up:
//Could you solve it using only O(1) extra space?
//
//
//Example 1:
//
//Input:
//["a","a","b","b","c","c","c"]
//
//Output:
//Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
//
//Explanation:
//"aa" is replaced by "a2". "bb" is replaced by "b2". "ccc" is replaced by "c3".
//
//
//
//Example 2:
//
//Input:
//["a"]
//
//Output:
//Return 1, and the first 1 characters of the input array should be: ["a"]
//
//Explanation:
//Nothing is replaced.
//
//
//
//Example 3:
//
//Input:
//["a","b","b","b","b","b","b","b","b","b","b","b","b"]
//
//Output:
//Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
//
//Explanation:
//Since the character "a" does not repeat, it is not compressed. "bbbbbbbbbbbb" is replaced by "b12".
//Notice each digit has it's own entry in the array.
//
//
//
//Note:
//
//    All characters have an ASCII value in [35, 126].
//    1 <= len(chars) <= 1000.

// a, a, b, b, c, c, c, d, d, d ,d
func compress(chars []byte) int {
	length := len(chars)
	if length <= 1 {
		return length
	}

	cur := 0
	last := chars[length-1]
	var j int
	for i := 0; i < length-1; i++ {
		if chars[i] == chars[i+1] {
			// find next different char
			for j = i + 1; j < length; j++ {
				if chars[i] != chars[j] {
					break
				}
			}

			replace(chars, &cur, j-i, chars[i])
			i = j - 1
		} else {
			// record different char
			replace(chars, &cur, 1, chars[i])
		}
	}

	// a, b, b, b, c
	// last character is not same
	if j != length {
		replace(chars, &cur, 1, last)
	}

	chars = chars[:cur]

	return len(chars)
}

func replace(chars []byte, cur *int, count int, c byte) {
	chars[*cur] = c
	*cur++

	if count > 1 {
		str := fmt.Sprintf("%d", count)
		for _, s := range str {
			chars[*cur] = byte(s)
			*cur++
		}
	}
}

// summary
// 1. first didn't see problem clearly, forget to mutate slice
// 2. first use previous character to check if in a consecutive row, that makes some special cases
// 3. after use next character to check, last character may not be checked, forget that case
