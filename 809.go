package main

// Sometimes people repeat letters to represent extra feeling, such as "hello" -> "heeellooo", "hi" -> "hiiii".  In these strings like "heeellooo", we have groups of adjacent letters that are all the same:  "h", "eee", "ll", "ooo".
//
// For some given string S, a query word is stretchy if it can be made to be equal to S by any number of applications of the following extension operation: choose a group consisting of characters c, and add some number of characters c to the group so that the size of the group is 3 or more.
//
// For example, starting with "hello", we could do an extension on the group "o" to get "hellooo", but we cannot get "helloo" since the group "oo" has size less than 3.  Also, we could do another extension like "ll" -> "lllll" to get "helllllooo".  If S = "helllllooo", then the query word "hello" would be stretchy because of these two extension operations: query = "hello" -> "hellooo" -> "helllllooo" = S.
//
// Given a list of query words, return the number of words that are stretchy.
//
//
//
// Example:
// Input:
// S = "heeellooo"
// words = ["hello", "hi", "helo"]
// Output: 1
// Explanation:
// We can extend "e" and "o" in the word "hello" to get "heeellooo".
// We can't extend "helo" to get "heeellooo" because the group "ll" is not size 3 or more.
//
//
//
// Notes:
//
//     0 <= len(S) <= 100.
//     0 <= len(words) <= 100.
//     0 <= len(words[i]) <= 100.
//     S and all words in words consist only of lowercase letters

func expressiveWords(S string, words []string) int {
	var match int

	for _, w := range words {
		if stretchy(w, S) {
			match++
		}
	}

	return match
}

func stretchy(src, dst string) bool {
	if len(src) == len(dst) {
		return src == dst
	}

	if len(src) > len(dst) || (len(src) == 0 && len(dst) != 0) || src[0] != dst[0] {
		return false
	}

	var i, j int
	minDuplicate := 3
	countSrc, countDst := 0, 0

	for ; i < len(src) && j < len(dst); i, j = i+1, j+1 {
		for countSrc = 1; i < len(src)-1 && src[i] == src[i+1]; i, countSrc = i+1, countSrc+1 {
		}

		for countDst = 1; j < len(dst)-1 && dst[j] == dst[j+1]; j, countDst = j+1, countDst+1 {
		}

		// every char should exist
		if src[i] != dst[j] {
			return false
		}

		if countDst < countSrc {
			return false
		}

		if countDst > countSrc && countDst < minDuplicate {
			return false
		}
	}

	// char mismatch
	if i != len(src) || j != len(dst) {
		return false
	}

	return true
}

//	problems
//	1.	should compare from start

//	2.	the algorithm add count when found next char is same, so no need
//		extra checking when index reaches length-1
