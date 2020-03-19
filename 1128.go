package main

//Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j] = [c, d] if and only if either (a==c and b==d), or (a==d and b==c) - that is, one domino can be rotated to be equal to another domino.
//
//Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length, and dominoes[i] is equivalent to dominoes[j].
//
//
//
//Example 1:
//
//Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
//Output: 1
//
//
//Constraints:
//
//1 <= dominoes.length <= 40000
//1 <= dominoes[i][j] <= 9

func numEquivDominoPairs(dominoes [][]int) int {
	duplicates := 0
	mapping := make(map[int]int)

	for _, d := range dominoes {
		var num int
		if d[0] <= d[1] {
			num = d[0]*10 + d[1]
		} else {
			num = d[1]*10 + d[0]
		}

		if c, ok := mapping[num]; !ok {
			mapping[num] = 1
		} else {
			duplicates += c
			mapping[num]++
		}
	}

	return duplicates
}

//	problems
//	1.	I thought counting is viewed by overall, but problem says it's
//		paris(i, j) and 0 <= i < j < len(dominoes), so every count is considered
//		for all dominoes before that one
//	2.	when two number is identical, the algorithm will count additional times
//	3.	optimization, it's too slow, I think the reason comes from fmt.Sprintf,
//		since every number is in range of 1-9, so number can be distinguished
//		by simpler way, first number * 10 + second number
//	4.	optimization, no need to separate number is same or different, just add
//		them into same category.
