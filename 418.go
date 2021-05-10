package main

import "strings"

// Given a rows x cols screen and a sentence represented as a list of strings, return the number of times the given sentence can be fitted on the screen.
//
// The order of words in the sentence must remain unchanged, and a word cannot be split into two lines. A single space must separate two consecutive words in a line.
//
//
//
// Example 1:
//
// Input: sentence = ["hello","world"], rows = 2, cols = 8
// Output: 1
// Explanation:
// hello---
// world---
// The character '-' signifies an empty space on the screen.
//
// Example 2:
//
// Input: sentence = ["a", "bcd", "e"], rows = 3, cols = 6
// Output: 2
// Explanation:
// a-bcd-
// e-a---
// bcd-e-
// The character '-' signifies an empty space on the screen.
//
// Example 3:
//
// Input: sentence = ["i","had","apple","pie"], rows = 4, cols = 5
// Output: 1
// Explanation:
// i-had
// apple
// pie-i
// had--
// The character '-' signifies an empty space on the screen.
//
//
//
// Constraints:
//
// 1 <= sentemce.length <= 100
// 1 <= sentence[i].length <= 10
// sentence[i] consists of lowercase English letters.
// 1 <= rows, cols <= 2 * 104

func wordsTyping(sentence []string, rows int, cols int) int {
	str := strings.Join(sentence, " ") + " "
	size := len(str)
	var start int

	for i := 0; i < rows; i++ {
		start += cols

		if str[start%size] == ' ' {
			// start of each line will always be a character
			start++
		} else {
			// backward to start of a word
			for ; start > 0 && str[(start-1)%size] != ' '; start-- {
			}
		}
	}

	return start / size
}

//	Notes
//	1.	most important observation, start of each row can only be character

//		use this observation to adjust each row maximum position
