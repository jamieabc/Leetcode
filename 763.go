package main

//A string S of lowercase English letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.
//
//
//
//Example 1:
//
//Input: S = "ababcbacadefegdehijhklij"
//Output: [9,7,8]
//Explanation:
//The partition is "ababcbaca", "defegde", "hijhklij".
//This is a partition so that each letter appears in at most one part.
//A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
//
//
//
//Note:
//
//    S will have length in range [1, 500].
//    S will consist of lowercase English letters ('a' to 'z') only.

func partitionLabels(S string) []int {
	// create array to find last index of char
	table := make([]int, 26)

	for i := range S {
		table[S[i]-'a'] = i
	}

	// expand region so that every char contains
	result := make([]int, 0)

	start, end := 0, 0

	for end < len(S) {
		for i := start; i <= end && i < len(S); i++ {
			end = max(end, table[S[i]-'a'])
		}

		result = append(result, end-start+1)
		end++
		start = end
	}

	return result
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
