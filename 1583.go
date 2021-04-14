package main

//You are given a list of preferences for n friends, where n is always even.
//
//For each person i, preferences[i] contains a list of friends sorted in the order of preference. In other words, a friend earlier in the list is more preferred than a friend later in the list. Friends in each list are denoted by integers from 0 to n-1.
//
//All the friends are divided into pairs. The pairings are given in a list pairs, where pairs[i] = [xi, yi] denotes xi is paired with yi and yi is paired with xi.
//
//However, this pairing may cause some of the friends to be unhappy. A friend x is unhappy if x is paired with y and there exists a friend u who is paired with v but:
//
//    x prefers u over y, and
//    u prefers x over v.
//
//Return the number of unhappy friends.
//
//
//
//Example 1:
//
//Input: n = 4, preferences = [[1, 2, 3], [3, 2, 0], [3, 1, 0], [1, 2, 0]], pairs = [[0, 1], [2, 3]]
//Output: 2
//Explanation:
//Friend 1 is unhappy because:
//- 1 is paired with 0 but prefers 3 over 0, and
//- 3 prefers 1 over 2.
//Friend 3 is unhappy because:
//- 3 is paired with 2 but prefers 1 over 2, and
//- 1 prefers 3 over 0.
//Friends 0 and 2 are happy.
//
//Example 2:
//
//Input: n = 2, preferences = [[1], [0]], pairs = [[1, 0]]
//Output: 0
//Explanation: Both friends 0 and 1 are happy.
//
//Example 3:
//
//Input: n = 4, preferences = [[1, 3, 2], [2, 3, 0], [1, 3, 0], [0, 2, 1]], pairs = [[1, 3], [0, 2]]
//Output: 4
//
//
//
//Constraints:
//
//    2 <= n <= 500
//    n is even.
//    preferences.length == n
//    preferences[i].length == n - 1
//    0 <= preferences[i][j] <= n - 1
//    preferences[i] does not contain i.
//    All values in preferences[i] are unique.
//    pairs.length == n/2
//    pairs[i].length == 2
//    xi != yi
//    0 <= xi, yi <= n - 1
//    Each person is contained in exactly one pair.

// tc: O(n^2)
func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
	}

	for i, pref := range preferences {
		for j := range pref {
			graph[i][pref[j]] = j
		}
	}

	paired := make([]int, n)
	for _, pair := range pairs {
		paired[pair[0]] = pair[1]
		paired[pair[1]] = pair[0]
	}

	var unhappy int

	for i := 0; i < n; i++ {
		order := graph[i][paired[i]]

		for j := 0; j < order; j++ {
			preferred := preferences[i][j]

			if graph[preferred][i] < graph[preferred][paired[preferred]] {
				unhappy++
				break
			}
		}
	}

	return unhappy
}

// tc: O(n^2)
func unhappyFriends1(n int, preferences [][]int, pairs [][]int) int {
	table := make(map[int]int)
	for _, p := range pairs {
		table[p[0]] = p[1]
		table[p[1]] = p[0]
	}

	// friend distance
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		dist[i][i] = -1
	}

	for i := range preferences {
		for j, k := range preferences[i] {
			dist[i][k] = j
		}
	}

	var unhappy int

	for _, p := range pairs {
		for _, ppl := range p {
			if !happy(table, dist, preferences, ppl) {
				unhappy++
			}
		}
	}

	return unhappy
}

func happy(table map[int]int, dist [][]int, preferences [][]int, target int) bool {
	paired := table[target]
	stop := dist[target][paired]

	for i := 0; i < stop; i++ {
		liker := preferences[target][i]
		likerPair := table[liker]

		if dist[liker][likerPair] > dist[liker][target] {
			return false
		}
	}

	return true
}

//	Notes
//	1.	it's really hard for me to understand the problem, with some long
//		description. But after wrong submission, suddenly understand this is
//		a matching problem. everyone is paired, check if I like other people
//		better than current paired, and the other people also likes me better
//		than paired.

//		In other words: not top pick. This is actually an interesting problem.

//	2.	inspired from https://leetcode.com/problems/count-unhappy-friends/discuss/843963/C%2B%2B-Friend-Distance

//		author uses better naming: friend distance to check, which reminds me
//		to build a 2D array is enough
