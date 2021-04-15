package main

import (
	"math"
	"math/rand"
)

// This problem is an interactive problem new to the LeetCode platform.
//
// We are given a word list of unique words, each word is 6 letters long, and one word in this list is chosen as secret.
//
// You may call master.guess(word) to guess a word.  The guessed word should have type string and must be from the original list with 6 lowercase letters.
//
// This function returns an integer type, representing the number of exact matches (value and position) of your guess to the secret word.  Also, if your guess is not in the given wordlist, it will return -1 instead.
//
// For each test case, you have 10 guesses to guess the word. At the end of any number of calls, if you have made 10 or less calls to master.guess and at least one of these guesses was the secret, you pass the testcase.
//
// Besides the example test case below, there will be 5 additional test cases, each with 100 words in the word list.  The letters of each word in those testcases were chosen independently at random from 'a' to 'z', such that every word in the given word lists is unique.
//
// Example 1:
// Input: secret = "acckzz", wordlist = ["acckzz","ccbazz","eiowzz","abcczz"]
//
// Explanation:
//
// master.guess("aaaaaa") returns -1, because "aaaaaa" is not in wordlist.
// master.guess("acckzz") returns 6, because "acckzz" is secret and has all 6 matches.
// master.guess("ccbazz") returns 3, because "ccbazz" has 3 matches.
// master.guess("eiowzz") returns 2, because "eiowzz" has 2 matches.
// master.guess("abcczz") returns 4, because "abcczz" has 4 matches.
//
// We made 5 calls to master.guess and one of them was the secret, so we pass the test case.
//
// Note:  Any solutions that attempt to circumvent the judge will result in disqualification.

/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */
/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */

func findSecretWord(wordlist []string, master *Master) {
	counts := make([][]int, 6)
	for i := range counts {
		counts[i] = make([]int, 26)
	}

	for i := range wordlist {
		for j, c := range wordlist[i] {
			counts[j][int(c-'a')]++
		}
	}

	size := len(wordlist)
	score := make([]int, size)
	for i, w := range wordlist {
		for j := range w {
			score[i] += counts[j][int(w[j]-'a')]
		}
	}

	possible := make([]int, size)
	for i := range possible {
		possible[i] = i
	}

	for count := 0; count < 10; count++ {
		target := freqChoose(score, possible)
		result := master.Guess(wordlist[possible[target]])

		remain := make([]int, 0)
		for i := range possible {
			if i != target && matches(wordlist[possible[i]], wordlist[possible[target]]) == result {
				remain = append(remain, possible[i])
			}
		}

		if len(remain) == 0 {
			return
		}

		possible = remain
	}
}

func freqChoose(score, possible []int) int {
	highest := math.MinInt32
	var idx int

	for i := range possible {
		if score[possible[i]] > highest {
			highest = score[possible[i]]
			idx = i
		}
	}

	return idx
}

func findSecretWord4(wordlist []string, master *Master) {
	size := len(wordlist)
	visited := make([]bool, size)

	for i := 0; i < 10; i++ {
		idx := pick4(wordlist, visited)

		result := master.Guess(wordlist[idx])
		visited[idx] = true

		if result == 6 {
			return
		} else {
			filterWords4(idx, wordlist, visited, result)
		}
	}
}

func pick4(wordlist []string, visited []bool) int {
	freq := make([]int, 26)

	for i, word := range wordlist {
		if visited[i] {
			continue
		}

		for j := range word {
			freq[word[j]-'a']++
		}
	}

	var maxScore, score, idx int
	for i, word := range wordlist {
		if visited[i] {
			continue
		}

		score = 0

		for j := range word {
			score += freq[word[j]-'a']
		}

		if score > maxScore {
			maxScore = score
			idx = i
		}
	}

	return idx
}

func filterWords4(idx int, wordlist []string, visited []bool, result int) {
	for i, str := range wordlist {
		if visited[i] || i == idx {
			continue
		}

		var same int
		for j := range str {
			if str[j] == wordlist[idx][j] {
				same++
			}
		}

		if same != result {
			visited[i] = true
		}
	}
}

func findSecretWord3(wordlist []string, master *Master) {
	size := len(wordlist)
	filtered := make([]bool, size)

	for i := 0; i < 10; i++ {
		idx := pick3(wordlist, filtered)
		filtered[idx] = true

		result := master.Guess(wordlist[idx])

		if result == 6 {
			return
		} else {
			filterWords3(wordlist[idx], wordlist, filtered, result)
		}
	}
}

func pick3(wordlist []string, filtered []bool) int {
	var idx, difference int
	minDifference := math.MaxInt32

	for i := range wordlist {
		if filtered[i] {
			continue
		}
		difference = 0

		for j := range wordlist {
			if filtered[j] || j == i {
				continue
			}

			if similarity3(wordlist[i], wordlist[j]) == 0 {
				difference++
			}

		}

		if difference < minDifference {
			minDifference = difference
			idx = i
		}
	}

	return idx
}

func similarity3(s1, s2 string) int {
	var same int
	for i := range s1 {
		if s1[i] == s2[i] {
			same++
		}
	}

	return same
}

func filterWords3(word string, wordlist []string, filtered []bool, dist int) {
	for i := range wordlist {
		if word == wordlist[i] || filtered[i] {
			continue
		}

		if similarity(word, wordlist[i]) != dist {
			filtered[i] = true
		}
	}
}

func findSecretWord2(wordlist []string, master *Master) {
	size := len(wordlist)
	zMatch := make([][]int, size)
	for i := range zMatch {
		zMatch[i] = make([]int, size)
	}

	// tc: O(n^2)
	for i := range wordlist {
		for j := i + 1; j < size; j++ {
			zMatch[i][j] = matches(wordlist[i], wordlist[j])
			zMatch[j][i] = zMatch[i][j]
		}
	}

	possible := make([]int, size)
	for i := range possible {
		possible[i] = i
	}

	for count := 10; count > 0; count-- {
		target := minMaxChoose(zMatch, possible)
		result := master.Guess(wordlist[possible[target]])

		remain := make([]int, 0)
		for j := range possible {
			if j != target && zMatch[possible[target]][possible[j]] == result {
				remain = append(remain, possible[j])
			}
		}

		if len(remain) == 0 {
			return
		}
		possible = remain
	}
}

func minMaxChoose(zMatch [][]int, possible []int) int {
	minMax := math.MaxInt32
	var idx int

	for i := range possible {
		var tmp int
		for j := 0; j < len(possible); j++ {
			if i != j && zMatch[possible[i]][possible[j]] == 0 {
				tmp++
			}
		}
		if minMax > tmp {
			minMax = tmp
			idx = i
		}
	}
	return idx
}

func findSecretWord1(wordlist []string, master *Master) {
	possible := make([]int, len(wordlist))
	for i := range possible {
		possible[i] = i
	}

	for count := 10; count > 0; count-- {
		target := choose(possible)
		result := master.Guess(wordlist[possible[target]])
		if result == 6 {
			return
		}

		tmp := make([]int, 0)

		for i := 0; i < len(possible); i++ {
			if i != target && matches(wordlist[possible[i]], wordlist[possible[target]]) == result {
				tmp = append(tmp, possible[i])
			}
		}

		possible = tmp
		if len(possible) == 0 {
			break
		}
	}
}

func choose(candidates []int) int {
	return rand.Intn(len(candidates))
}

func matches(src, dst string) int {
	var match int
	for i := range src {
		if src[i] == dst[i] {
			match++
		}
	}

	return match
}

//	problems
//	1.	when guessing words, use src as base, count the difference

//	2.	matches is considered on ith char

//	3.	inspired by https://leetcode.com/problems/guess-the-word/discuss/133862/Random-Guess-and-Minimax-Guess-with-Comparison

//		lee observes the problem and mentions a good point, if a word list
//		is made up of aaaa, bbbb, cccc, etc., then at least 26 tries to
//		get correct answer. 10 tries is to check for reasonable solution

//		most important observation comes from the conclusion that most
//		guesses will end-up to 0, which is crucial to whole discussion.

//		minimax strategy, reduces as many words as possible. the possibility
//		to get 0 match is (25/26)^6 ~ 80%, which means mostly guess will get
//		0 matches. the problem becomes: how to remove most words if there's
//		0 matches. since any words with 0 matches to guess word will be
//		retained, the target becomes: choose a word that has minimum
//		0 matches to all other words.

//		further improvement on O(n^2) is to calculate char occurrence, to
//		eliminates more words, choose the word with more popular chars on
//		each position

//	4.	inspired by https://leetcode.com/problems/guess-the-word/discuss/134251/Optimal-MinMax-Solution-(%2B-extra-challenging-test-cases)

//		author describe guessing word as the distance to secret
//		word. to choose better words, each iteration should re-calculate
//		word score. basic idea is to choose a word among candidates that
//		partitions equally in all distances (0 ~ 6)

//		the idea different from lee is that global optimum is not locally
//		optimum. e.g. a word is good to partition at global could be worst
//		when candidates shrink.

//		it's a really good insight not also describing guess as distance,
//		but also points out globally optimum doesn't guarantee locally
//		optimum

//		due to time limit, I didn't implement this solution

//		the key point is that, elimination not only happens at result = 0,
//		but also for other values (1~5), because different "distance" will not
//		be the answer

//		e.g. secret = "abcde"
//		guess "awxyz" => distance = 1
//		if there's another "kuytr" can be eliminated, because distance of
//		"awxyz" & "kuytr" is 5

//	5.	inspired form https://leetcode.com/problems/guess-the-word/discuss/134087/C%2B%2B-elimination-histogram-beats-Minimax

//		histogram can also be used as a way to find
