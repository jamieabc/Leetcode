package main

// You are standing at position 0 on an infinite number line. There is a goal at position target.
//
// On each move, you can either go left or right. During the n-th move (starting from 1), you take n steps.
//
// Return the minimum number of steps required to reach the destination.
//
// Example 1:
// Input: target = 3
// Output: 2
// Explanation:
// On the first move we step from 0 to 1.
// On the second step we step from 1 to 3.
// Example 2:
// Input: target = 2
// Output: 3
// Explanation:
// On the first move we step from 0 to 1.
// On the second move we step  from 1 to -1.
// On the third move we step from -1 to 2.
// Note:
// target will be a non-zero integer in the range [-10^9, 10^9].

func reachNumber(target int) int {
	if target < 0 {
		target *= -1
	}

	var sum, steps int
	for steps = 1; true; steps++ {
		sum += steps
		if sum >= target {
			break
		}
	}

	if (sum-target)&1 == 0 {
		return steps
	}

	for steps++; true; steps++ {
		sum += steps
		if (sum-target)&1 == 0 {
			break
		}
	}

	return steps
}

//	Notes
//	1.	some observations
//		n = 1, sequences: -1, 1
//		n = 2, sequences: -3, -1, 1, 3
//		n = 3, sequences: -6, -4, -2, 2, 4, 6
//		n = 4, sequences: -10, -8, -6, -2, 2, 6, 8, 10

//		so, sequences is symmetric to 0, k & -k both exist, this reduces
//		solution space form [-10^9, 10^9] down to [0, 10^9]

//	2.	the operations is like there are some number k such that
//		1 +- 2 +- 3 +- 4 +- ... +- k

//		the goal is to find + or - on each number, the time complexity is 2^n

//	3.	inspired from https://leetcode.com/problems/reach-a-number/discuss/112968/Short-JAVA-Solution-with-Explanation

//		there's a brilliant way to find it: find 1 + 2 + ... + k >= target

//		let s = 1 + 2 + ... + k, if s - target is even number, so we can always
//		flip + to - to form this even number.

//		because switch +1 to -1 means a -2 difference, switch +2 to -2 is a -4
//		difference, etc.

//		if difference is odd, then keep increasing k till difference is even
