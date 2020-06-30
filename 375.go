package main

import "math"

// We are playing the Guess Game. The game is as follows:
//
// I pick a number from 1 to n. You have to guess which number I picked.
//
// Every time you guess wrong, I'll tell you whether the number I picked is higher or lower.
//
// However, when you guess a particular number x, and you guess wrong, you pay $x. You win the game when you guess the number I picked.
//
// Example:
//
// n = 10, I pick 8.
//
// First round:  You guess 5, I tell you that it's higher. You pay $5.
// Second round: You guess 7, I tell you that it's higher. You pay $7.
// Third round:  You guess 9, I tell you that it's lower. You pay $9.
//
// Game over. 8 is the number I picked.
//
// You end up paying $5 + $7 + $9 = $21.
//
// Given a particular n ≥ 1, find out how much money you need to have to guarantee a win.

func getMoneyAmount(n int) int {
	table := make([][]int, n+1)
	for i := range table {
		table[i] = make([]int, n+1)
	}

	return topDown(table, 1, n)
}

func topDown(table [][]int, start, end int) int {
	if start >= end {
		return 0
	}

	// already exist
	if table[start][end] != 0 {
		return table[start][end]
	}

	minCost := math.MaxInt32
	for i := start; i <= end; i++ {
		lowerChoices := topDown(table, start, i-1)
		higherChoices := topDown(table, i+1, end)

		minCost = min(minCost, i+max(lowerChoices, higherChoices))
	}

	table[start][end] = minCost
	return minCost
}

func getMoneyAmount1(n int) int {
	memoization := make(map[int]map[int]int)
	return dfs(memoization, 1, n)
}

func dfs(memoization map[int]map[int]int, start, end int) int {
	// initialize
	if _, ok := memoization[start]; !ok {
		memoization[start] = make(map[int]int)
	}

	// check table
	if val, ok := memoization[start][end]; ok {
		return val
	}

	// only one choice, no cost
	if start >= end {
		memoization[start][end] = 0
		return 0
	}

	// 2 numbers, choose smaller one
	if end-start == 1 {
		memoization[start][end] = start
		return start
	}

	// 3 numbers, choose middle one
	if end-start == 2 {
		memoization[start][end] = start + 1
		return start + 1
	}

	minCost := math.MaxInt32
	for i := start; i <= end; i++ {
		smallerChoices := dfs(memoization, start, i-1)
		largerChoices := dfs(memoization, i+1, end)

		// worst case, so it's a max in two choices
		minCost = min(minCost, i+max(smallerChoices, largerChoices))
	}

	memoization[start][end] = minCost
	return minCost
}

//	problems
//	1.	don't know how to write this...I thought it was binary search problem

//	2.	inspired from https://leetcode.com/problems/guess-number-higher-or-lower-ii/discuss/84766/Clarification-on-the-problem-description.-Problem-description-need-to-be-updated-!!!

//		I can understand what this problem about. In my words, the problem can
//		be described as: if every time I am smart enough to choose best
//		strategy but my guesses are always wrong (worst case scenario), what is
//		minimum cost to find out correct number?

//		some scenarios:
//		- one number: cost = 0
//		- two numbers: minimum number between two
//		- three numbers: middle number among them
//		- four numbers: e.g. 1 2 3 4
//		  1 + 3 = 4
//		  2 + 3 = 5
//        3 + 1 = 4
//		  4 + 2 = 6
//		  best strategy is firsts number then third number

//		in brief, I can divide the problem into sub-problem as follows:
//		for range start ~ end, for any chosen number start <= i <= end, worst
//		case scenario is i + max(DP(start, i-1), DP(i+1, end))

//	3.	I was having difficulty to find out time complexity, but inspired from
//		https://leetcode.com/problems/guess-number-higher-or-lower-ii/discuss/84764/Simple-DP-solution-with-explanation~~

//		for a range, I need to try for every possible number in n, for every number
//		is chosen, I need to find its sub-intervals, e.g. n = 10, choose 5, then
//		intervals are 2-4,3-4,4,7-10,8-10, ,etc.

//		tc: O(n^3) first n: 1 ~ n, finding possible intervals for a chosen
//		number: O(n^2)

//	4.	add reference https://leetcode.com/problems/guess-number-higher-or-lower-ii/discuss/84794/DP-JAVA-O(n3)-Solution-With-Explanation-15ms-17-lines
//
//		don't know how to write from bottom-up
