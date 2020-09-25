package main

import (
	"sort"
)

// Suppose you have a random list of people standing in a queue. Each person is described by a pair of integers (h, k), where h is the height of the person and k is the number of people in front of this person who have a height greater than or equal to h. Write an algorithm to reconstruct the queue.
//
// Note:
// The number of people is less than 1,100.
//
//
// Example
//
// Input:
// [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//
// Output:
// [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		return people[i][0] > people[j][0]
	})

	result := make([][]int, len(people))

	for i := 0; i < len(people); i++ {
		copy(result[people[i][1]+1:], result[people[i][1]:])
		result[people[i][1]] = people[i]
	}

	return result
}

// tc: O(n^2), because when placing people into result, time depends on # of
// people in result, e.g. 1 + 2 + 3 + ... + n-1
func reconstructQueue2(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		return people[i][0] < people[j][0]
	})

	result := make([][]int, len(people))
	var j, count int
	for i := range people {
		for j, count = 0, 0; count < people[i][1] || len(result[j]) != 0; j++ {
			if len(result[j]) == 0 || result[j][0] >= people[i][0] {
				count++
			}
		}
		result[j] = people[i]
	}

	return result
}

//	problem
//	1.	when moving by difference, it should be finding # of people higher/lower
//		than me, and put myself there

//		e.g. if higher than me is 4, but now is 2, means I need to move myself
//			 after 2 people higher than me

//		e.g. if higher than me is 4, but now is 6, means too many people higher
//			 than me, I have to move myself before 2 people higher me

//	2.	when deciding next position, it could be possible that to end of array,
//		which means dst could beyond range

//	3.	can't think of any way to improve, so I check reference
//		https://leetcode.com/problems/queue-reconstruction-by-height/discuss/89345/Easy-concept-with-PythonC%2B%2BJava-Solution
//		this is really clever solution, it founds the rule of sorting, looks like
//		divide & conquer. The point here is highest people order is always fixed.

//		My thinking goes from minimum to maximum, but the problem is that
//		minimum position will be changed when larger number moves.

//		What can be better when I think of the problem? I should focus more on
//		nature of problem: the higher count means lower one can be inserted into
//		higher ones. Which means higher ones position are still fixed, the change
//		comes from lower one.

//	4.	the other reference: https://leetcode.com/problems/queue-reconstruction-by-height/discuss/89359/Explanation-of-the-neat-Sort%2BInsert-solution
//		this one starts from lowest people, since lowest people has nothing to
//		do with higher people, the higher count of lowest people is actually
//		the index value of still free spaces.

//	5.	After seeing 2 beautiful solutions, what did I miss during thinking? I
//		focus on surface and fail to notice actual rule that makes it work.
//		If I notice that smaller people moving won't affect higher people,
//		then this could be a clue to solve the problem. It's hard, but not that
//		hard, I should learn how to think through clues.

//	6.	when start from minimum, count is increased in 2 conditions:
//		- empty slot
//		- people higher than or equal to me

//	7.	refactor

//	8.	optimize, too slow when starts from highest...the reason is many memory
//		allocation/dump, it can be in-place insert
