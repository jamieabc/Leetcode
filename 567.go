package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	counter := make([]int, 26)
	for i := range s1 {
		counter[s1[i]-'a']++
	}

	for i := range counter {
		if counter[i] == 0 {
			counter[i] = -1
		}
	}

	var low, high, match int
	if counter[s2[0]-'a'] > 0 {
		match++
		counter[s2[0]-'a']--
	}

	if match == len(s1) {
		return true
	}

	for low <= len(s2)-len(s1) {
		if high < len(s2)-1 && high-low+1 < len(s1) {
			// expand window to same size of s1
			high++
			counter[s2[high]-'a']--

			if counter[s2[high]-'a'] >= 0 {
				match++
			}
		} else {
			// shrink window
			counter[s2[low]-'a']++
			if counter[s2[low]-'a'] > 0 {
				match--
			}
			low++
		}

		if match == len(s1) {
			return true
		}
	}

	return false
}

//	problems
//	1.	boundary condition, e.g. empty string, first character match

//	2.	random permutation means size & char count in that range should match,
//		so expand window to that size, and shrink if match count not right

//		most important thing is that range in s2 need to match with s1 length

//	3.	add reference https://leetcode.com/problems/permutation-in-string/discuss/102590/8-lines-slide-window-solution-in-Java

//		author first subtract character from counter, and make sure if all
//		char count are 0 means exactly match

//	4	spend 1 hour on this problem, first 40 minutes is trying to adjust
//		window w/o using property of length of range need to be same as s1

//		I need to be more sensitive about hint from problem, since it will
//		reduce complexity of the solution

//		sliding window problem is all about how to expand/shrink & terminate
//		conditions
