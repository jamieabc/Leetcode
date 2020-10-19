package main

import (
	"fmt"
	"sort"
)

// You are the manager of a basketball team. For the upcoming tournament, you want to choose the team with the highest overall score. The score of the team is the sum of scores of all the players in the team.
//
// However, the basketball team is not allowed to have conflicts. A conflict exists if a younger player has a strictly higher score than an older player. A conflict does not occur between players of the same age.
//
// Given two lists, scores and ages, where each scores[i] and ages[i] represents the score and age of the ith player, respectively, return the highest overall score of all possible basketball teams.
//
//
//
// Example 1:
//
// Input: scores = [1,3,5,10,15], ages = [1,2,3,4,5]
// Output: 34
// Explanation: You can choose all the players.
//
// Example 2:
//
// Input: scores = [4,5,6,5], ages = [2,1,2,1]
// Output: 16
// Explanation: It is best to choose the last 3 players. Notice that you are allowed to choose multiple people of the same age.
//
// Example 3:
//
// Input: scores = [1,2,3,5], ages = [8,9,10,1]
// Output: 6
// Explanation: It is best to choose the first 3 players.
//
//
//
// Constraints:
//
//     1 <= scores.length, ages.length <= 1000
//     scores.length == ages.length
//     1 <= scores[i] <= 106
//     1 <= ages[i] <= 1000
func bestTeamScore(scores []int, ages []int) int {
	size := len(scores)

	table := make([][]int, size)
	for i := range scores {
		table[i] = []int{scores[i], ages[i]}
	}

	// sort by score, ascending
	sort.Slice(table, func(i, j int) bool {
		if table[i][0] < table[j][0] {
			return true
		} else if table[i][0] > table[j][0] {
			return false
		}

		// same score, the lower age the earlier
		return table[i][1] < table[j][1]
	})

	// dp[i] means max score from a team build range 0 - i people
	dp := make([]int, size)
	dp[0] = table[0][0]
	maxScore := dp[0]

	for i := 1; i < size; i++ {
		dp[i] = table[i][0]
		for j := 0; j < i; j++ {
			if table[j][0] == table[i][0] || table[j][1] <= table[i][1] {
				dp[i] = max(dp[i], dp[j]+table[i][0])
			}
		}
		maxScore = max(maxScore, dp[i])
	}

	return maxScore
}

func bestTeamScore1(scores []int, ages []int) int {
	size := len(scores)

	table := make([][]int, size)
	for i := range scores {
		table[i] = []int{scores[i], ages[i]}
	}

	// sort by age, ascending
	sort.Slice(table, func(i, j int) bool {
		if table[i][1] < table[j][1] {
			return true
		} else if table[i][1] > table[j][1] {
			return false
		}

		// same age, the lower score the earlier
		return table[i][0] < table[j][0]
	})

	// dp[i] means max score from a team build range 0 - i people
	dp := make([]int, size)
	dp[0] = table[0][0]
	maxScore := dp[0]

	for i := 1; i < size; i++ {
		dp[i] = table[i][0]
		for j := 0; j < i; j++ {
			if table[j][1] == table[i][1] || table[j][0] <= table[i][0] {
				dp[i] = max(dp[i], dp[j]+table[i][0])
			}
		}
		maxScore = max(maxScore, dp[i])
	}

	return maxScore
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	younger people with higher score, two factors make it complicated to do
//		decision, so sort by one factor and choose this other by rule

//		rule: younger people cannot have strictly higher score

//		sort by age ascending, since every next one is older or equal, goal
//		becomes to find longest non-decreasing score sequence

//		sort by score ascending, every next one has lower or equal score, goal
//		becomes to find longest non-decreasing age sequence

//	2.	inspired from https://leetcode.com/problems/best-team-with-no-conflicts/discuss/900342/C%2B%2B-DP-IS-Solution

//		author provides a very good observation, younger cannot have strictly
//		higher score than older people means both age & score are in non
//		decreasing sequence

//		I should have think about this, it's the key to solve this problem

//	3.	inspired from https://leetcode.com/problems/best-team-with-no-conflicts/discuss/899467/O-(n*log-n)-Dynamic-programming-explanation-source-code-and-video-solution

//		author provides a better tc in O(n log(n)) using segment tree, not
//		studying it
