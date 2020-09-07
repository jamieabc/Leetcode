package main

//You have a list of words and a pattern, and you want to know which words in words matches the pattern.
//
//A word matches the pattern if there exists a permutation of letters p so that after replacing every letter x in the pattern with p(x), we get the desired word.
//
//(Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.)
//
//Return a list of the words in words that match the given pattern.
//
//You may return the answer in any order.
//
//
//
//Example 1:
//
//Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
//Output: ["mee","aqq"]
//Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
//"ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation,
//since a and b map to the same letter.
//
//
//
//Note:
//
//    1 <= words.length <= 50
//    1 <= pattern.length = words[i].length <= 20

func findAndReplacePattern(words []string, pattern string) []string {
	result := make([]string, 0)
	for i := range words {
		if isSamePattern(words[i], pattern) {
			result = append(result, words[i])
		}
	}

	return result
}

func isSamePattern(w1, w2 string) bool {
	from, to := make([]int, 256), make([]int, 256)

	for i := range w1 {
		if from[w1[i]] != 0 && from[w1[i]] != int(w2[i]) {
			return false
		}

		if to[w2[i]] != 0 && to[w2[i]] != int(w1[i]) {
			return false
		}

		from[w1[i]] = int(w2[i])
		to[w2[i]] = int(w1[i])
	}

	return true
}
