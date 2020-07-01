package main

// Alex and Lee play a game with piles of stones.  There are an even number of piles arranged in a row, and each pile has a positive integer number of stones piles[i].
//
// The objective of the game is to end with the most stones.  The total number of stones is odd, so there are no ties.
//
// Alex and Lee take turns, with Alex starting first.  Each turn, a player takes the entire pile of stones from either the beginning or the end of the row.  This continues until there are no more piles left, at which point the person with the most stones wins.
//
// Assuming Alex and Lee play optimally, return True if and only if Alex wins the game.
//
//
//
// Example 1:
//
// Input: [5,3,4,5]
// Output: true
// Explanation:
// Alex starts first, and can only take the first 5 or the last 5.
// Say he takes the first 5, so that the row becomes [3, 4, 5].
// If Lee takes 3, then the board is [4, 5], and Alex takes 5 to win with 10 points.
// If Lee takes the last 5, then the board is [3, 4], and Alex takes 4 to win with 9 points.
// This demonstrated that taking the first 5 was a winning move for Alex, so we return true.
//
//
//
// Note:
//
//     2 <= piles.length <= 500
//     piles.length is even.
//     1 <= piles[i] <= 500
//     sum(piles) is odd.

func stoneGame(piles []int) bool {
	size := len(piles)
	// memo[i][j] means max stones could get if remaining stones from i to j
	memo := make([][]int, size)
	for i := range memo {
		memo[i] = make([]int, size)
		memo[i][i] = piles[i]
	}
	dfs(piles, 0, size-1, memo)

	// alex always have full choices, bob has either 1~size-1 or 0~size-2
	return memo[0][size-1] > max(memo[1][size-1], memo[0][size-2])
}

func dfs(stones []int, start, end int, memo [][]int) int {
	// terminate
	if start > end {
		return 0
	}

	// already visited
	if memo[start][end] > 0 {
		return memo[start][end]
	}

	// either choose start or end
	var chooseStart, chooseEnd int
	if memo[start+1][end] > 0 {
		chooseStart = memo[start+1][end]
	} else {
		chooseStart = dfs(stones, start+1, end, memo)
		memo[start+1][end] = chooseStart
	}

	if memo[start][end-1] > 0 {
		chooseEnd = memo[start][end-1]
	} else {
		chooseEnd = dfs(stones, start, end-1, memo)
		memo[start][end-1] = chooseEnd
	}

	memo[start][end] = max(stones[start]+chooseStart, stones[end]+chooseEnd)
	return memo[start][end]
}

func stoneGame3(piles []int) bool {
	length := len(piles)
	dp := make([]int, length)
	for i := range dp {
		dp[i] = piles[i]
	}

	for d := 1; d < length; d++ {
		for i := 0; i < length-d; i++ {
			dp[i] = max(piles[i]-dp[i], piles[i+d]-dp[i+d-1])
		}
	}

	return dp[length-1] > 0
}

func stoneGame2(piles []int) bool {
	length := len(piles)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
		dp[i][i] = piles[i]
	}

	for d := 1; d < length; d++ {
		for i := 0; i < length-d; i++ {
			dp[i][i+d] = max(piles[i]-dp[i+1][i+d], piles[i+d]-dp[i][i+d-1])
		}
	}

	return dp[0][length-1] > 0
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func stoneGame1(piles []int) bool {
	return play(piles, 0, 0, 0, 0, len(piles)-1)
}

func play(piles []int, alex, lee, turn, start, end int) bool {
	if start > end {
		return alex > lee
	}

	nextTurn := turn ^ 1

	if turn == 0 {
		return play(piles, alex+piles[start], lee, start+1, end, nextTurn) || play(piles, alex+piles[end], lee, start, end-1, nextTurn)
	} else {
		return play(piles, alex, lee+piles[start], start+1, end, nextTurn) || play(piles, alex, lee+piles[end], start, end-1, nextTurn)
	}
}

//	problems
//	1.	too slow, try to trim some un-necessary branch, time complexity is
//		O(2^n), cause at every length, people can choose either first or last

//	2.	try not to modify slice but updating index, see if it's faster

//	3.	inspired from https://leetcode.com/problems/stone-game/discuss/154610/DP-or-Just-return-true

//		first thing is that when piles length is even, alex can choose all
//		even or odd pile, and since total stones are odd which means alex
//		can always win

//		this is a clever insight, I didn't even take time to observe
//		problem and find out this simple conclusion. I guess the problem is
//		about finding rules for alex:
//		- choose even piles: pick first element, then depending on lee's
//		  choice, if lee picks first then alex pick first, if lee picks last
//		  then alex picks last
//		- choose odd piles: pick last element, then depending on lee's
//		  choice, if lee picks first then alex picks first, if lee picks
//		  last then alex picks last

//		this kind of conclusion comes with 2 presumptions: total stones are
//		odd, and pile length is even

//		when using 2D dp, it's very clever. goal is alex score > lee score.
//		but dp can pursuit one variable, which become diff = alex - lee.
//		adn score relationship can be decided by choose first or choose
//		last, choose whatever max of it:
//		dp(i, j) = max(p[i] - dp(i+1, j), p[j] - dp(i, j-1))

//		author also uses beautiful dp, starts from dp[i][i] then expands
//		to dp[i][j]

//		at last also reduces 2D dp to 1D dp, kind of observing dp formula,
//		dp[i][j] relates to dp(i+1, j) & dp(i, j-1), and dp(i+1, j) is
//		1D dp(i), dp(i, j-1) is dp(i-1)

//	4.	add another reference https://leetcode.com/problems/stone-game/discuss/154660/Java-This-is-minimax-%2B-dp-(fully-detailed-explanation-%2B-generalization-%2B-easy-understand-code)

//		author explains how dp come out, and the time complexity

//	5.	inspired from https://leetcode.com/problems/stone-game/discuss/261718/Step-by-Step-Recursive-TopDown-BottomUp-and-BottomUp-using-O(n)-space-in-Java

//		start from recursive w/ memo

//	6.	inspired from https://leetcode.com/problems/stone-game/discuss/154660/Java-This-is-minimax-%2B-dp-(fully-detailed-explanation-%2B-generalization-%2B-easy-understand-code)

//		I didn't aware this is a minimax problem.

//		score can be defined as score(alex) - score(lee), then alex wants to maximize score, and lee
//		wants to minimize score.
