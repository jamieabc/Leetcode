package main

import "math"

// Koko loves to eat bananas.  There are N piles of bananas, the i-th pile has piles[i] bananas.  The guards have gone and will come back in H hours.
//
// Koko can decide her bananas-per-hour eating speed of K.  Each hour, she chooses some pile of bananas, and eats K bananas from that pile.  If the pile has less than K bananas, she eats all of them instead, and won't eat any more bananas during this hour.
//
// Koko likes to eat slowly, but still wants to finish eating all the bananas before the guards come back.
//
// Return the minimum integer K such that she can eat all the bananas within H hours.
//
//
//
// Example 1:
//
// Input: piles = [3,6,7,11], H = 8
// Output: 4
//
// Example 2:
//
// Input: piles = [30,11,23,4,20], H = 5
// Output: 30
//
// Example 3:
//
// Input: piles = [30,11,23,4,20], H = 6
// Output: 23
//
//
//
// Constraints:
//
//     1 <= piles.length <= 10^4
//     piles.length <= H <= 10^9
//     1 <= piles[i] <= 10^9

func minEatingSpeed(piles []int, H int) int {
	var low, high, total int
	for _, i := range piles {
		high = max(high, i)
		total += i
	}

	var ans int
	for low = int(math.Ceil(float64(total) / float64(H))); low <= high; {
		mid := low + (high-low)>>1

		if finishInTime(piles, mid, H) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}

func finishInTime(piles []int, target, H int) bool {
	var hour int

	for i := range piles {
		tmp := piles[i] / target
		if tmp*target < piles[i] {
			tmp++
		}
		hour += tmp
	}

	return hour <= H
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from sample code, if every pile is exactly integer multiple of
//		H, then it would be lower bound, but this need to be same logic as
//		normal counting (math.ceil)

//		since math.ceil won't plus when return is 0, need to convert each number
//		into float such that divide result won't be 0

//	2.	inspired from https://leetcode.com/problems/koko-eating-bananas/discuss/769702/Python-Clear-explanation-Powerful-Ultimate-Binary-Search-Template.-Solved-many-problems.

//		to use binary search, the sequence needs to have a form of
//		t t t t f f f f, thus next direction can decided
