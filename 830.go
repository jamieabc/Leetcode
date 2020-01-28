package main

//In a string S of lowercase letters, these letters form consecutive groups of the same character.
//
//For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and "yy".
//
//Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.
//
//The final answer should be in lexicographic order.
//
//
//
//Example 1:
//
//Input: "abbxxxxzzy"
//Output: [[3,6]]
//Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.
//Example 2:
//
//Input: "abc"
//Output: []
//Explanation: We have "a","b" and "c" but no large group.
//Example 3:
//
//Input: "abcdddeeeeaabbbcd"
//Output: [[3,5],[6,9],[12,14]]
//
//
//Note:  1 <= S.length <= 1000

func largeGroupPositions(S string) [][]int {
	result := make([][]int, 0)
	length := len(S)

	if length <= 2 {
		return result
	}

	var i, j int
	for i = 1; i < length; i++ {
		if S[i] == S[i-1] {
			for j = i + 1; j < length; j++ {
				if S[j] != S[i] {
					break
				}
			}

			if j-i+1 >= 3 {
				result = append(result, []int{i - 1, j - 1})
			}
			i = j
		}
	}

	return result
}
