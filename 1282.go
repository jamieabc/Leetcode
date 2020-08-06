package main

//There are n people whose IDs go from 0 to n - 1 and each person belongs exactly to one group. Given the array groupSizes of length n telling the group size each person belongs to, return the groups there are and the people's IDs each group includes.
//
//You can return any solution in any order and the same applies for IDs. Also, it is guaranteed that there exists at least one solution.
//
//
//
//Example 1:
//
//Input: groupSizes = [3,3,3,3,3,1,3]
//Output: [[5],[0,1,2],[3,4,6]]
//Explanation:
//Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
//
//Example 2:
//
//Input: groupSizes = [2,1,3,3,3,2]
//Output: [[1],[0,5],[2,3,4]]
//
//
//
//Constraints:
//
//    groupSizes.length == n
//    1 <= n <= 500
//    1 <= groupSizes[i] <= n

// tc: O(n), sp: O(n)
func groupThePeople(groupSizes []int) [][]int {
	result := make([][]int, 0)
	counter := make(map[int][]int)

	for i := range groupSizes {
		if _, ok := counter[groupSizes[i]]; !ok {
			counter[groupSizes[i]] = make([]int, 0)
		}
		counter[groupSizes[i]] = append(counter[groupSizes[i]], i)

		if len(counter[groupSizes[i]]) == groupSizes[i] {
			result = append(result, counter[groupSizes[i]])
			counter[groupSizes[i]] = make([]int, 0)
		}
	}

	// put remaining groups
	for _, value := range counter {
		if len(value) > 0 {
			result = append(result, value)
		}
	}

	return result
}
