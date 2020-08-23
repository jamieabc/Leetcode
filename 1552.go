package main

import (
	"fmt"
	"math"
	"sort"
)

//In universe Earth C-137, Rick discovered a special form of magnetic force between two balls if they are put in his new invented basket. Rick has n empty baskets, the ith basket is at position[i], Morty has m balls and needs to distribute the balls into the baskets such that the minimum magnetic force between any two balls is maximum.
//
//Rick stated that magnetic force between two different balls at positions x and y is |x - y|.
//
//Given the integer array position and the integer m. Return the required force.
//
//
//
//Example 1:
//
//Input: position = [1,2,3,4,7], m = 3
//Output: 3
//Explanation: Distributing the 3 balls into baskets 1, 4 and 7 will make the magnetic force between ball pairs [3, 3, 6]. The minimum magnetic force is 3. We cannot achieve a larger minimum magnetic force than 3.
//
//Example 2:
//
//Input: position = [5,4,3,2,1,1000000000], m = 2
//Output: 999999999
//Explanation: We can use baskets 1 and 1000000000.
//
//
//
//Constraints:
//
//    n == position.length
//    2 <= n <= 10^5
//    1 <= position[i] <= 10^9
//    All integers in position are distinct.
//    2 <= m <= position.length

//	tc: O(n log 10^9)
func maxDistance(position []int, m int) int {
	sort.Ints(position)

	maxDist := (position[len(position)-1] - position[0]) / (m - 1)

	low, high := 0, maxDist

	for low < high {
		mid := low + (high-low+1)/2

		// need to check for m-1, because it might happen that
		// 0, a1, a2, ..., ak meets criteria, but ak to an doesn't
		// meet criteria
		if check(position, mid, m-1) {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return low
}

func check(position []int, target, count int) bool {
	for prev, i := 0, 1; i < len(position); i++ {
		if position[i] - position[prev] >= target {
			prev = i
			count--
			if count == 0 {
				return true
			}
		}
	}

	return count == 0
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	tc: O(m * n^2)
func maxDistance1(position []int, m int) int {
	sort.Ints(position)

	maxDist := math.MinInt32

	// memo[i][j] means maximum min distance start from index i, with remaining j balls
	memo := make([][]int, len(position))
	for i := range memo {
		memo[i] = make([]int, m+1)
	}

	for i := 0; i <= len(position)-m; i++ {
		maxDist = max(maxDist, recursive(position, memo, i, m-1))
	}

	return maxDist
}

func recursive(position []int, memo [][]int, start, remain int) int {
	if remain == 0 {
		return math.MaxInt32
	}

	if memo[start][remain] > 0 {
		return memo[start][remain]
	}

	maxDist := math.MinInt32

	for i := start + 1; i <= len(position)-remain; i++ {
		d := min(position[i]-position[start], recursive(position, memo, i, remain-1))
		maxDist = max(maxDist, d)
	}

	memo[start][remain] = maxDist

	return maxDist
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	first I want to use a memo to store intermediate value, memo[i][j]
//		means start from index i with j balls remain, maximum of minimum
//		distance

//		then I found it's not correct, because for same i, j, might still with
//		different value

//		e.g. position = [1, 2, 3, 4, 7], m = 3
//		select position 1, 2, 7 => memo[4][1] = 1
//		select position 1, 4, 7 => memo[4][1] = 3

//		because minimum distance is determined by previous selected positions,
//		notation of storing start index & remain balls is not in correct form.
//		notation should be appropriate to store result that means exhausive
//		result

//		goal of memo is to reduce computation, technically it's better to store
//		as many middle computation result

//	2.	inspired from https://leetcode.com/problems/magnetic-force-between-two-balls/discuss/794188/Top-Bottom-DP-with-Memoization-Java-TLE

//		when for any position, there's 2 conditions: choose this position or not
//		choose, so I do it as this.

//		but author provides another way of doing this: not choose this position
//		means choose next position with same ball remaining, so there's a
//		duplicate calculation in my implementation

//		the other thing is that when iterate through possible position, could
//		just limit i to length - remain, and reduce i accordingly, this avoids
//		additional checking that start index out of boundary

//	3.	Suddenly realize that this is problem is about separate (group) numbers.
//		for n numbers, group those number into m groups, each group is
//		represented by its largest number. The goal is to find maximum of
//		minimum distance among adjacent groups

//		because maximum of min distance, it means smallest & largest numbers
//		should be selected to get largest possible distance to distribute.

//		theoretically, maximum of min distance is less or equal to
//		(largest - smallest) / (m-1)

//      |    there are m-1 spaces among selected numbers	|
//		_    _     _ _ 	_		_ _ _		_ _		_ _		_
//      ^													^
//   smallest									         largest

//		it's possible to try any given distance ranges from 1 - theory max,
//		if not able to fit into this criteria, reduce target distance, if meets,
//		increase target distance (binary search)

//	4.	inspired from https://leetcode.com/problems/magnetic-force-between-two-balls/discuss/794066/Simple-Explanation

//		after binary search for valid distance, use linear to check if it's
//		valid.

//		also, author summarized binary search criteria, as I copy below:
//		Consider, when we have lo = 1 and hi = 2.
//		If I used, int mi = (lo + hi) / 2; mi = 1
//		If I used, int mi = (lo + hi + 1) / 2; mi = 2
//
//		Now since, lo = mi, we will keep falling back to lo = 1 and hi = 2
//		(Hence TLE), if I used first version.
//
//		General rule of thumb I use:
//		If we are doing lo = mi, and hi = mi-1, use 2nd version.
//		If we are doing hi = mi, lo = mi + 1, use 1st version.
