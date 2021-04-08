package main

// Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).
//
// Example:
//
// Input: S = "ADOBECODEBANC", T = "ABC"
// Output: "BANC"
//
// Note:
//
//     If there is no such window in S that covers all characters in T, return the empty string "".
//     If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

func minWindow(s string, t string) string {
	target := make([]int, 256)
	for i := range t {
		target[t[i]]++
	}

	counter := make([]int, 256)
	remain := len(t)
	size := len(s)
	shortest := size + 1
	var str string

	for i, j := 0, 0; i < size; {
		if i == j || (j < size && remain > 0) {
			counter[s[j]]++
			if counter[s[j]] <= target[s[j]] {
				remain--
			}
			j++
		} else {
			counter[s[i]]--
			if counter[s[i]] < target[s[i]] {
				remain++
			}
			i++
		}

		if remain == 0 && j-i < shortest {
			shortest = j - i
			str = s[i:j]
		}
	}

	if shortest == size+1 {
		return ""
	}
	return str
}

func minWindow3(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	count := make(map[byte]int)
	for i := range t {
		count[t[i]]++
	}

	filtered := make([]int, 0)
	for i := range s {
		if _, ok := count[s[i]]; ok {
			filtered = append(filtered, i)
		}
	}

	match := make(map[byte]int)
	var matched, end, left, right int
	start := -1

	for left < len(filtered) {
		if matched < len(t) && right < len(filtered) {
			// expand window
			match[s[filtered[right]]]++
			if match[s[filtered[right]]] <= count[s[filtered[right]]] {
				matched++
			}
			right++
		} else {
			// update range if all t is contained
			if matched == len(t) && (start == -1 || end-start > filtered[right-1]+1-filtered[left]) {
				start, end = filtered[left], filtered[right-1]+1
			}

			match[s[filtered[left]]]--
			if match[s[filtered[left]]] < count[s[filtered[left]]] {
				matched--
			}
			left++
		}
	}

	if start == -1 {
		return ""
	}

	return s[start:end]
}

func minWindow2(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	count := make(map[byte]int)
	for i := range t {
		count[t[i]]++
	}

	match := make(map[byte]int)
	var matched, end, left, right int
	start := -1

	for left, right = 0, 0; left < len(s); {
		if matched < len(t) && right < len(s) {
			// not enough, expand window
			if _, ok := count[s[right]]; ok {
				match[s[right]]++
				if match[s[right]] <= count[s[right]] {
					matched++
				}
			}
			right++
		} else {
			// check if criteria met, shrink window
			if matched == len(t) && (start == -1 || end-start > right-left) {
				start, end = left, right
			}

			if _, ok := count[s[left]]; ok {
				match[s[left]]--

				if match[s[left]] < count[s[left]] {
					matched--
				}
			}
			left++
		}
	}

	if start == -1 {
		return ""
	}

	return s[start:end]
}

func minWindow1(s string, t string) string {
	if t == "" || s == "" {
		return ""
	}

	matched, count := make(map[byte][]int), make(map[byte]int)

	// in case t contains multiple same chars
	for i := range t {
		count[t[i]]++
		if _, ok := matched[t[i]]; !ok {
			matched[t[i]] = make([]int, 0)
		}
	}

	var end, match, i, j int
	start := -1

	for i, j = 0, 0; i < len(s); {
		if match == 0 {
			// move i
			for ; i < len(s); i++ {
				if _, ok := matched[s[i]]; ok {
					matched[s[i]] = append(matched[s[i]], i)
					match++
					j = i + 1
					break
				}
			}
		} else if match < len(t) {
			if j == len(s) {
				break
			}

			if _, ok := matched[s[j]]; ok {
				matched[s[j]] = append(matched[s[j]], j)

				// match can only be increased when char number count
				// less or equal to target count
				if len(matched[s[j]]) <= count[s[j]] {
					match++
				}
			}
			j++
		} else {
			// match = target length
			// move i to next char in t
			if start == -1 || end-start > j-1-i {
				start, end = i, j-1
			}

			// move i to char in t, and also reduce match count
			if _, ok := matched[s[i]]; ok {
				matched[s[i]] = matched[s[i]][1:]

				if len(matched[s[i]]) < count[s[i]] {
					match--
				}
			}
			i++

			// successfully find i that makes match reduced,
			// move i to next char in t to let length be minimized
			if match < len(t) {
				for ; i < len(s); i++ {
					if _, ok := matched[s[i]]; ok {
						// in case t is only one char
						if len(matched[s[i]]) <= count[s[i]] {
							break
						} else {
							matched[s[i]] = matched[s[i]][1:]
						}
					}
				}
			}
		}
	}

	if start == -1 {
		return ""
	}

	return s[start : end+1]
}

//	Notes
//	1.	too slow...

//	2.	inspired from solution, sliding can be simplified by following
//		rules:
//		- if not contain all chars in t, expand window
//		- if each char count still valid, shrink window

//	3.	when recording info about s, using count is enough because traverse
//		is in order, so count decreased means a char is just met, and rest
//		still same

//	4.	sliding window algorithm can be further improved by counting only
//		valid char is s, e.g. s = "AXXBXCXXA", t = "ABC"
//		filtered = [A, 0], [B, 3], [C, 5], [A, 8]

//	5.	inspired from https://leetcode.com/problems/minimum-window-substring/discuss/26808/Here-is-a-10-line-template-that-can-solve-most-'substring'-problems

//		a brilliant solution that author use == 0 to check if counter is
//		valid. e.g. if t  == "AABC", A: 2, B: 1, C: 1, when all chars are
//		included, A == B == C == 0

//		in this way, I can remove an variable that tracks count of occurred
//		char in range

//	6.	becareful about terminate condition, it's on left pointer, because
//		there might be a condition that necessary char appears at last, then
//		moving left pointer is kind of shrink (remove duplicates)

//	7.	inspired from solution, if len(s) >>>> len(t), could remove characters
//		not in t

//		use array to store index of s that appears character in t
