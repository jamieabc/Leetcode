package main

import "sort"

//Three stones are on a number line at positions a, b, and c.
//
//Each turn, you pick up a stone at an endpoint (ie., either the lowest or highest position stone), and move it to an unoccupied position between those endpoints.  Formally, let's say the stones are currently at positions x, y, z with x < y < z.  You pick up the stone at either position x or position z, and move that stone to an integer position k, with x < k < z and k != y.
//
//The game ends when you cannot make any more moves, ie. the stones are in consecutive positions.
//
//When the game ends, what is the minimum and maximum number of moves that you could have made?  Return the answer as an length 2 array: answer = [minimum_moves, maximum_moves]
//
//
//
//Example 1:
//
//Input: a = 1, b = 2, c = 5
//Output: [1,2]
//Explanation: Move the stone from 5 to 3, or move the stone from 5 to 4 to 3.
//
//Example 2:
//
//Input: a = 4, b = 3, c = 2
//Output: [0,0]
//Explanation: We cannot make any moves.
//
//Example 3:
//
//Input: a = 3, b = 5, c = 1
//Output: [1,2]
//Explanation: Move the stone from 1 to 4; or move the stone from 1 to 2 to 4.
//
//
//
//Note:
//
//    1 <= a <= 100
//    1 <= b <= 100
//    1 <= c <= 100
//    a != b, b != c, c != a

func numMovesStones(a int, b int, c int) []int {
	array := []int{a, b, c}
	sort.Ints(array)

	// 1, 2, 3
	if array[1]-array[0] == 1 && array[2]-array[1] == 1 {
		return []int{0, 0}
	}

	max := (array[1] - array[0] - 1) + (array[2] - array[1] - 1)

	// 1, 3, 5 => 1, 2, 3
	if array[1]-array[0] == 2 || array[2]-array[1] == 2 {
		return []int{1, max}
	}

	// 1, 5, 8
	min := 2

	// 1, 2, 5
	if array[1]-1 == array[0] {
		min--
	}

	// 1, 3, 4
	if array[2]-1 == array[1] {
		min--
	}

	return []int{min, max}
}
