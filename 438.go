package main

func findAnagrams(s string, p string) []int {
	if len(s) == 0 {
		return []int{}
	}

	counter := make([]int, 26)
	for i := range p {
		counter[p[i]-'a']++
	}

	result := make([]int, 0)

	var low, high, match int
	counter[s[0]-'a']--
	if counter[s[0]-'a'] >= 0 {
		match++
	}

	if match == len(p) {
		result = append(result, 0)
	}

	for low <= len(s)-len(p) {
		if high < len(s)-1 && high-low+1 < len(p) {
			// expand window
			high++
			counter[s[high]-'a']--
			if counter[s[high]-'a'] >= 0 {
				match++
			}
		} else {
			// shrink window
			counter[s[low]-'a']++
			if counter[s[low]-'a'] > 0 {
				match--
			}
			low++
		}

		if match == len(p) {
			result = append(result, low)
		}
	}

	return result
}

//	problems
//	1.	add reference https://leetcode.com/problems/find-all-anagrams-in-a-string/discuss/92007/Sliding-Window-algorithm-template-to-solve-all-the-Leetcode-substring-search-problem.

//		author lists many problems related to sliding window
