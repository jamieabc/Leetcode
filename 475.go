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

// [1, 2, 3], [2]
// [1, 2, 3, 4], [2, 3]
// [1, 2, 3, 4], [1, 4]
// [1, 4, 7, 10, 16, 22, 30], [2, 15, 30]
// [282475249,622650073,984943658,144108930,470211272,101027544,457850878,458777923]
//[823564440,115438165,784484492,74243042,114807987,137522503,441282327,16531729,823378840,143542612]
func findRadius(houses []int, heaters []int) int {
	lenHouses := len(houses)
	lenHeaters := len(heaters)

	sort.Ints(houses)
	sort.Ints(heaters)

	if lenHouses == 0 {
		return 0
	}

	radius := -1
	var i, diff int

	if lenHouses == 1 {
		radius = 25001
		for i = 0; i < lenHeaters; i++ {
			diff = abs(houses[0] - heaters[i])
			if diff < radius {
				radius = diff
			}
		}
		return radius
	}

	if lenHeaters == 1 {
		for i = 0; i < lenHouses; i++ {
			diff = abs(houses[i] - heaters[0])
			if diff > radius {
				radius = diff
			}
		}
		return radius
	}

	var leftHouseIndex, rightHouseIndex int

	for i = 0; i < lenHouses; i++ {
		if houses[i] > heaters[0] {
			break
		}
		// not necessary, only need to consider index 0
		diff = abs(houses[i] - heaters[0])
		if diff > radius {
			radius = diff
		}
	}

	// find left side radius and range
	if houses[0] < heaters[0] {
		diff := abs(houses[0] - heaters[0])
		if diff > radius {
			diff = radius
		}
	}
	for leftHouseIndex = 0; leftHouseIndex < lenHouses; leftHouseIndex++ {
		if houses[leftHouseIndex] > heaters[0] {
			break
		}
	}

	// find right side radius and range
	if houses[lenHouses-1] > heaters[lenHeaters-1] {
		diff := abs(houses[lenHouses-1] - heaters[lenHeaters-1])
		if diff > radius {
			radius = diff
		}
	}
	for rightHouseIndex = lenHouses - 1; rightHouseIndex >= 0; rightHouseIndex-- {
		if houses[rightHouseIndex] < heaters[lenHeaters-1] {
			break
		}
	}

	i = leftHouseIndex
	for j := 0; j < lenHeaters-1; j++ {
		// find houses within heaters j & j+1
		for ; i <= rightHouseIndex; i++ {
			if houses[i] > heaters[j+1] {
				break
			}
			toLeft := abs(houses[i] - heaters[j])
			toRight := abs(houses[i] - heaters[j+1])

			if toLeft < toRight && toLeft > radius {
				radius = toLeft
			} else if toRight < toLeft && toRight > radius {
				radius = toRight
			}
		}
	}

	if radius == -1 {
		return 0
	}

	return radius
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
