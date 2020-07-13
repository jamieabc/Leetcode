package main

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	counter := make([]int, 26)
	wordMap := make(map[string]bool)

	var totalMatch int
	for _, word := range words {
		totalMatch += len(word)
		wordMap[word] = true
		for i := range word {
			counter[word[i]-'a']++
		}
	}

	if len(s) < totalMatch {
		return []int{}
	}

	var low, high, match int
	result := make([]int, 0)

	for high < totalMatch {
		counter[s[high]-'a']--
		if counter[s[high]-'a'] >= 0 {
			match++
		}
		high++
	}
	high--

	wordSize := len(words[0])
	if match == totalMatch && matchAnyWord(wordMap, s[low:low+wordSize]) {
		result = append(result, 0)
	}

	for low <= len(s)-totalMatch {
		if high < len(s)-1 && high-low+1 < totalMatch {
			// expand window
			high++
			counter[s[high]-'a']--
			if counter[s[high]-'a'] >= 0 {
				match++
			}
		} else {
			counter[s[low]-'a']++
			if counter[s[low]-'a'] > 0 {
				match--
			}
			low++
		}

		if match == totalMatch && matchAnyWord(wordMap, s[low:low+wordSize]) {
			result = append(result, low)
		}
	}

	return result
}

func matchAnyWord(words map[string]bool, toSearch string) bool {
	return words[toSearch]
}

//	problems
//	1.	concatenation means shrink needs to be words size

//	2.	when shrink, only shrink word size if matches found

//	3.	not only to match char count, start of words should also match

//	4.	don' know how to do it, needs to consider char count and also each
//		word should appear exactly once
