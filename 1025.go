package main

// Alice and Bob take turns playing a game, with Alice starting first.
//
// Initially, there is a number N on the chalkboard.  On each player's turn, that player makes a move consisting of:
//
//     Choosing any x with 0 < x < N and N % x == 0.
//     Replacing the number N on the chalkboard with N - x.
//
// Also, if a player cannot make a move, they lose the game.
//
// Return True if and only if Alice wins the game, assuming both players play optimally.
//
//
//
// Example 1:
//
// Input: 2
// Output: true
// Explanation: Alice chooses 1, and Bob has no more moves.
//
// Example 2:
//
// Input: 3
// Output: false
// Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.
//
//
//
// Note:
//
//     1 <= N <= 1000

func divisorGame(N int) bool {
	// dp[i] means alice plays when N = i, true means win, false means lose
	dp := make([]bool, N+1)
	if N == 1 {
		return false
	}
	dp[2] = true

	for i := 3; i <= N; i++ {
		nums := divisable(i)
		tmp := false
		for _, j := range nums {
			tmp = tmp || !dp[i-j]
			if tmp {
				break
			}
		}
		dp[i] = tmp
	}

	return dp[N]
}

func divisable(n int) []int {
	result := []int{1}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

//	problems
//	1.	boundary case, when 1 then I cannot assign dp[2]

//	2.	inspired from https://leetcode.com/problems/divisor-game/discuss/274590/C%2B%2B-Recursive-DP

//		to check value at most sqrt(n) because N % i == 0, so at most it
//		should be N = i * i

//	3.	add reference https://leetcode.com/problems/divisor-game/discuss/382233/Solution-in-Python-3-(With-Detailed-Proof)

//		it's mathematical/induction proof
