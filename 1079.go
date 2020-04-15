package main

import (
	"sort"
)

// You have a set of tiles, where each tile has one letter tiles[i] printed on it.  Return the number of possible non-empty sequences of letters you can make.
//
//
//
// Example 1:
//
// Input: "AAB"
// Output: 8
// Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".
//
// Example 2:
//
// Input: "AAABBC"
// Output: 188
//
//
//
// Note:
//
//     1 <= tiles.length <= 7
//     tiles consists of uppercase English letters.

//	thinking process
//	this problem relates to both combination & permutation
//	combination:  5 = 1 + 1 + 1 + 1 + 1
//					= 2 + 1 + 1 + 1
//					= 2 + 2 + 1
//					= 3 + 1 + 1
//					= 4 + 1
//					= 3 + 2
//					= 5
//	I think it as putting ball into boxes, first is 5 box each with 1 ball.
//	4 boxes, one box with 2 balls and other with 1 ball.
//	3 boxes has 2 conditions, one box with 3 and other 2 boxes with 1,
//	two boxes with 2 and one box with 1.

// 	When add numbers are decided, next problem is to find possible chars, so I
// 	need a table that stores occurrence count for every char, and starts
//	to permute chars.
//	permutation: find all possible for AAAAB (4+1), AAABB(3+2)

// 	The way of generating combination is so complicated, I have to generate numbers
//	that sum to specific number (4+1), and by those numbers, find potential chars
//	that meets count. Then designate each group a char which cannot duplicate, then
//	generate permutations. No need to do it and know it's extremely complicated.

//	However, I think of another method, use subtraction to do. Start from full
//	string, then subtract one each time, and permute. Or further more, once string
//	is decided, use math equation to calculate
func numTilePossibilities(tiles string) int {
	length := len(tiles)
	if length <= 1 {
		return length
	}

	bs := []byte(tiles)
	sort.Slice(bs, func(i, j int) bool { return bs[i] <= bs[j] })
	tiles = string(bs)

	var total int
	track := make([]bool, length)
	permutation(tiles, track, &total)
	return total
}

func permutation(str string, track []bool, total *int) {
	for i := 0; i < len(str); {
		// this char is already used in this round
		if track[i] {
			i++
			continue
		}

		track[i] = true
		*total++
		permutation(str, track, total)
		track[i] = false

		// find next different character
		for i += 1; i < len(str); i++ {
			if str[i] != str[i-1] {
				break
			}
		}
	}
}

//	problems
//	1.	duplicates happens in this algorithm, e.g.
//		AAABBC => AABBC ... AAABB ... => AABB ... AABB ...
//	2.	wrong logic, string is calculated should be checked before calculate
//		combination
//	3.	permutation needs 2 variables: what are used in whole loop, what are used
//		in this loop
//	4.	when skipping repeated chars, make sure to consider when char is same
//		to end
