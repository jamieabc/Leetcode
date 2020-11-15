package main

import "math"

// There are 1000 buckets, one and only one of them is poisonous, while the rest are filled with water. They all look identical. If a pig drinks the poison it will die within 15 minutes. What is the minimum amount of pigs you need to figure out which bucket is poisonous within one hour?
//
// Answer this question, and write an algorithm for the general case.
//
//
//
// General case:
//
// If there are n buckets and a pig drinking poison will die within m minutes, how many pigs (x) you need to figure out the poisonous bucket within p minutes? There is exactly one bucket with poison.
//
//
//
// Note:
//
// A pig can be allowed to drink simultaneously on as many buckets as one would like, and the feeding takes no time.
// After a pig has instantly finished drinking buckets, there has to be a cool down time of m minutes. During this time, only observation is allowed and no feedings at all.
// Any given bucket can be sampled an infinite number of times (by an unlimited number of pigs).

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets <= 1 {
		return 0
	}

	turn := minutesToTest / minutesToDie
	if turn*minutesToDie != minutesToTest {
		turn++
	}

	for p := 1; true; p++ {
		if math.Pow(float64(turn+1), float64(p)) >= float64(buckets) {
			return p
		}
	}

	return 0
}

//	Notes
//	1.	for example of 1000 buckets, 4 rounds to check, I was thinking 1000^0.25
//		= 6, each time check will reduce size down to 1/6. But his is not most
//		efficient.

//	2.	inspired from https://leetcode.com/problems/poor-pigs/discuss/94266/Another-explanation-and-solution

//		if there are 2 pigs & 4 turns, then up to 25 buckets can be determined

//		1	2	3	4	5
//		6	7	8	9	10
//		11	12	13	14	15
//		16	17	18	19	20
//		21	22	23	24	25

//		one pig tests for row (first round drinks 1+2+3+4+5, second round drinks
//		6+7+8+9+10, etc.)

//		the other pig tests for column (first round drinks 1+6+11+16+21, second
//		round drinks 2+7+12+17+22, etc.)

//		so it's clear that (turn+1)^pig >= buckets

//	3.	becareful about boundary condition: 0 & 1 pigs, no need to test

//	4.	inspired from solution, very interesting concept of quantum bits, provides
//		different dimension of information
