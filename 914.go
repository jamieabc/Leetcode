package main

//In a deck of cards, each card has an integer written on it.
//
//Return true if and only if you can choose X >= 2 such that it is possible to split the entire deck into 1 or more groups of cards, where:
//
//    Each group has exactly X cards.
//    All the cards in each group have the same integer.
//
//
//
//Example 1:
//
//Input: [1,2,3,4,4,3,2,1]
//Output: true
//Explanation: Possible partition [1,1],[2,2],[3,3],[4,4]
//
//Example 2:
//
//Input: [1,1,1,2,2,2,3,3]
//Output: false
//Explanation: No possible partition.
//
//Example 3:
//
//Input: [1]
//Output: false
//Explanation: No possible partition.
//
//Example 4:
//
//Input: [1,1]
//Output: true
//Explanation: Possible partition [1,1]
//
//Example 5:
//
//Input: [1,1,2,2,2,2]
//Output: true
//Explanation: Possible partition [1,1],[2,2],[2,2]
//
//
//Note:
//
//    1 <= deck.length <= 10000
//    0 <= deck[i] < 10000

func hasGroupsSizeX(deck []int) bool {
	mapping := make(map[int]int)

	// count all cards occur times
	for _, n := range deck {
		if _, ok := mapping[n]; !ok {
			mapping[n] = 1
		} else {
			mapping[n]++
		}
	}

	divisor := mapping[deck[0]]

	// find common divisor
	for _, v := range mapping {
		divisor = gcd(divisor, v)
	}

	return divisor > 1
}

func gcd(i, j int) int {
	if j == 0 {
		return i
	}

	return gcd(j, i%j)
}

//	problems
//	1.	inspired from https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/discuss/469539/Go-GCD-with-explanation

//		gcd can be found by iteratively repeat until b == 0:
//		tmp = a % b, a = b, b = tmp
