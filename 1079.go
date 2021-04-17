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

func numTilePossibilities(tiles string) int {
	length := len(tiles)
	if length <= 1 {
		return length
	}

	bs := []byte(tiles)
	sort.Slice(bs, func(i, j int) bool { return bs[i] <= bs[j] })
	tiles = string(bs)

	flags := make([]bool, length)
	return permutation(tiles, flags)
}

func permutation(str string, flags []bool) int {
	var count int

	for i := 0; i < len(str); {
		// this char is already used in this round
		if flags[i] {
			i++
			continue
		}

		flags[i] = true
		count += 1 + permutation(str, flags)
		flags[i] = false

		// find next different character
		for i++; i < len(str); i++ {
			if str[i] != str[i-1] {
				break
			}
		}
	}

	return count
}

//	Notes

//	1.	thinking process
//		this problem relates to both combination & permutation
//		combination:  5 = 1 + 1 + 1 + 1 + 1
//						= 2 + 1 + 1 + 1
//						= 2 + 2 + 1
//						= 3 + 1 + 1
//						= 4 + 1
//						= 3 + 2
//						= 5
//		I think it as putting ball into boxes, first is 5 box each with 1 ball.
//		4 boxes, one box with 2 balls and other with 1 ball.
//		3 boxes has 2 conditions, one box with 3 and other 2 boxes with 1,
//		two boxes with 2 and one box with 1.

// 		When add numbers are decided, next problem is to find possible chars,
// 		so I need a table that stores occurrence count for every char, and
// 		starts to permute chars.
//		permutation: find all possible for AAAAB (4+1), AAABB(3+2)

// 		The way of generating combination is so complicated, I have to generate
// 		numbers that sum to specific number (4+1), and by those numbers, find
// 		potential chars that meets count. Then designate each group a char
// 		which cannot duplicate, then generate permutations. No need to do it
// 		and know it's extremely complicated.

//		However, I think of another method, use subtraction to do. Start from
//		full string, then subtract one each time, and permute. Or further more,
//		once string is decided, use math equation to calculate

//	2.	duplicates happens in this algorithm, e.g.
//		AAABBC => AABBC ... AAABB ... => AABB ... AABB ...

//	3.	wrong logic, string is calculated should be checked before calculate
//		combination

//	4.	permutation needs 2 variables: what are used in whole loop, what are used
//		in this loop

//	5.	when skipping repeated chars, make sure to consider when char is same
//		to end

//	6.	spend one hour and stuck at permutation & combination, I try to first
//		iterate find combination, then use math to calculate permutations

//		but I failed spending for one hour...

//	7.	inspired from https://leetcode.com/problems/letter-tile-possibilities/discuss/308284/Concise-java-solution

//		use counter to store distinct characters, then start to iterate

//		each time, use a character at that position, go to next recursion call
//		and that position will never be same again because it uses distinct
//		character, this is the combination (largest for loop)

//		inside loop, it acts as permutation, because arr[i]--, arr[i]++

//		that's all i can understand for now, need to further review it again
