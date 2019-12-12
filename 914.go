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
	if len(deck) < 2 {
		return false
	}

	mapping := make(map[int]int)

	// count all cards occur times
	for _, n := range deck {
		if _, ok := mapping[n]; !ok {
			mapping[n] = 1
		} else {
			mapping[n]++
		}
	}

	// find out min
	count := make([]int, 0)
	min := 10000
	for _, v := range mapping {
		count = append(count, v)
		if v < min {
			min = v
		}
	}

	// find common divisor
	divisor := make([]int, min+1)
	length := len(count)
	var j int
	for i := 2; i <= min; i++ {
		if divisor[i] == 0 {
			// mark all multiply
			for j = i; j <= min; j += i {
				divisor[j] = 1
			}

			// find out if any divisor that separate groups
			for j = 0; j < length; j++ {
				if count[j]%i != 0 {
					break
				}
			}

			if j == length {
				return true
			}
		}
	}
	return false
}
