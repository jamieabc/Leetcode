package main

import "sort"

// N cars are going to the same destination along a one lane road.  The destination is target miles away.
//
// Each car i has a constant speed speed[i] (in miles per hour), and initial position position[i] miles towards the target along the road.
//
// A car can never pass another car ahead of it, but it can catch up to it, and drive bumper to bumper at the same speed.
//
// The distance between these two cars is ignored - they are assumed to have the same position.
//
// A car fleet is some non-empty set of cars driving at the same position and same speed.  Note that a single car is also a car fleet.
//
// If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.
//
//
// How many car fleets will arrive at the destination?
//
//
//
// Example 1:
//
// Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
// Output: 3
// Explanation:
// The cars starting at 10 and 8 become a fleet, meeting each other at 12.
// The car starting at 0 doesn't catch up to any other car, so it is a fleet by itself.
// The cars starting at 5 and 3 become a fleet, meeting each other at 6.
// Note that no other cars meet these fleets before the destination, so the answer is 3.
//
//
// Note:
//
// 0 <= N <= 10 ^ 4
// 0 < target <= 10 ^ 6
// 0 < speed[i] <= 10 ^ 6
// 0 <= position[i] < target
// All initial positions are different.

func carFleet(target int, position []int, speed []int) int {
	size := len(position)
	idx := make([]int, size)
	for i := range idx {
		idx[i] = i
	}

	sort.Slice(idx, func(i, j int) bool {
		return position[idx[i]] > position[idx[j]]
	})

	var fleet int

	for i, j := 0, 0; i < size; {
		fleet++
		timeToDestination := float64(target-position[idx[i]]) / float64(speed[idx[i]])

		for j = i + 1; j < size; j++ {
			if timeToDestination < float64(target-position[idx[j]])/float64(speed[idx[j]]) {
				break
			}
		}

		i = j
	}

	return fleet
}

//	Notes
//	1.	slower car blocks faster car, so it only matters the time car reaches destination
//		if smaller position cars reaches destination time <= larger position cars reach destination time,
//		all of them are considered as one group of cars
