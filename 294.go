package main

// You are playing a Flip Game with your friend.
//
// You are given a string currentState that contains only '+' and '-'. You and your friend take turns to flip two consecutive "++" into "--". The game ends when a person can no longer make a move, and therefore the other person will be the winner.
//
// Return true if the starting player can guarantee a win, and false otherwise.
//
//
//
// Example 1:
//
// Input: currentState = "++++"
// Output: true
// Explanation: The starting player can guarantee a win by flipping the middle "++" to become "+--+".
//
// Example 2:
//
// Input: currentState = "+"
// Output: false
//
//
//
// Constraints:
//
// 1 <= currentState.length <= 60
// currentState[i] is either '+' or '-'.

func canWin(currentState string) bool {
	var state uint64
	for i := range currentState {
		if currentState[i] == '+' {
			state |= 1 << (63 - i)
		}
	}

	// memo[i]: for current play with state i wins
	memo := make(map[uint64]bool)
	memo[0] = false

	dfs(state, memo)

	return memo[state]
}

func dfs(state uint64, memo map[uint64]bool) bool {
	if val, ok := memo[state]; ok {
		return val
	}

	var found, win bool
	for i := 0; i < 64; i++ {
		if state&(1<<(63-i)) > 0 && state&(1<<(62-i)) > 0 {
			found = true
			cur := state
			cur ^= 1 << (63 - i)
			cur ^= 1 << (62 - i)
			win = win || !dfs(cur, memo)
		}
	}

	memo[state] = found && win

	return memo[state]
}

//	Notes
//	1.	at first my thought was too complicated, dfs function needs to know
//		which user's turn, but this problem has only 2 outcomes, either win
//		or lose for current situation, so I remove player, focus more on memo,
//		and simplify it to be current player with this state has a change to
//		win or not

//	2.	dfs means all possibilities, so "guarantee win" can be changed to any
//		dfs result with opponent lose
