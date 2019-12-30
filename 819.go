package main

import (
	"strings"
)

//Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.  It is guaranteed there is at least one word that isn't banned, and that the answer is unique.
//
//Words in the list of banned words are given in lowercase, and free of punctuation.  Words in the paragraph are not case sensitive.  The answer is in lowercase.
//
//
//
//Example:
//
//Input:
//paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
//banned = ["hit"]
//Output: "ball"
//Explanation:
//"hit" occurs 3 times, but it is a banned word.
//"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
//Note that words in the paragraph are not case sensitive,
//that punctuation is ignored (even if adjacent to words, such as "ball,"),
//and that "hit" isn't the answer even though it occurs more because it is banned.
//
//
//
//Note:
//
//    1 <= paragraph.length <= 1000.
//    0 <= banned.length <= 100.
//    1 <= banned[i].length <= 10.
//    The answer is unique, and written in lowercase (even if its occurrences in paragraph may have uppercase symbols, and even if it is a proper noun.)
//    paragraph only consists of letters, spaces, or the punctuation symbols !?',;.
//    There are no hyphens or hyphenated words.
//    Words only consist of letters, never apostrophes or other punctuation symbols.

func mostCommonWord(paragraph string, banned []string) string {
	count := make(map[string]int)

	// count words
	length := len(paragraph)
	var max, i, j int
	var result string

	for i = 0; i < length; {
		// skip non-letter
		for ; i < length; i++ {
			if isLetter(paragraph[i]) {
				break
			}
		}

		if i == length {
			break
		}

		// find word
		for j = i + 1; j < length; j++ {
			if !isLetter(paragraph[j]) {
				break
			}
		}

		if j == length {
			break
		}

		lower := strings.ToLower(paragraph[i:j])

		if isValid(lower, banned) {
			count[lower]++

			// update result
			if count[lower] > max {
				result = lower
				max = count[lower]
			}
		}

		i = j
	}

	if j == length {
		lower := strings.ToLower(paragraph[i:])

		if isValid(lower, banned) {
			count[lower]++

			// update result
			if count[lower] > max {
				result = lower
				max = count[lower]
			}
		}
	}

	return result
}

func isLetter(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
}

func isValid(str string, banned []string) bool {
	for _, b := range banned {
		if str == b {
			return false
		}
	}
	return true
}

// problems
// 1. forget to convert words into lower case
// 2. forget to check if i, j reach strings limit
// 3. for comparison banned, words should transferred into lower case, but it should stored in original format
// 4. point 3 is wrong, cause words coud be BALL and Ball, they are the same, convert all words into lower case
// 5. forget to change logic of transform words into lower case
