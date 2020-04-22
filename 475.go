package main

import "sort"

//Winter is coming! Your first job during the contest is to design a standard heater with fixed warm radius to warm all the houses.
//
//Now, you are given positions of houses and heaters on a horizontal line, find out minimum radius of heaters so that all houses could be covered by those heaters.
//
//So, your input will be the positions of houses and heaters separately, and your expected output will be the minimum radius standard of heaters.
//
//Note:
//
//Numbers of houses and heaters you are given are non-negative and will not exceed 25000.
//Positions of houses and heaters you are given are non-negative and will not exceed 10^9.
//As long as a house is in the heaters' warm radius range, it can be warmed.
//All the heaters follow your radius standard and the warm radius will the same.
//
//
//Example 1:
//
//Input: [1,2,3],[2]
//Output: 1
//Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.
//
//
//Example 2:
//
//Input: [1,2,3,4],[1,4]
//Output: 1
//Explanation: The two heater was placed in the position 1 and 4. We need to use radius 1 standard, then all the houses can be warmed.

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	var radius, idx int
	length := len(heaters)

	for _, h := range houses {
		// already covered
		if h >= heaters[idx]-radius && h <= heaters[idx]+radius {
			continue
		}

		// check if next heater exist and covers house
		if idx+1 < length && h >= heaters[idx+1]-radius && h <= heaters[idx+1]+radius {
			idx++
			continue
		}

		// house is not covered, increase radius
		// in this case, house is fixed, find closest heater, but new radius cannot below existing radius
		// the tricky part is that if house at 1000, and heaters at 100, 101, 102, ..., 999,
		// need to find closest heater to house, but radius might not update
		var r int
		for r = abs(heaters[idx] - h); idx+1 < length; idx++ {
			tmp := abs(heaters[idx+1] - h)
			if tmp <= r {
				r = tmp
			} else {
				break
			}
		}

		if r > radius {
			radius = r
		}
	}

	return radius
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

// 	problems
//	1.	Too complicate, re-do the problem with simpler way.
//		first sort houses & heaters, then if a house if included by current
//		heater +- radius, continue to next house.
//		If a house is not included in current heater, try to check if next
//		heater +- radius include this house. Since house & heater are mono
//		increasing, so next heater is guaranteed no less than current heater.

//		This method needs to take care of one thing: if a house at 1000, and
//		heaters are 500, 501, 502, ... 999, then in this round of checking
//		need to find minimum radius to include the house, which is 1 (999-1000).

//		But new radius cannot be lower than previous radius to make sure
//		previous houses are included.
