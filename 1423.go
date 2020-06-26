package main

// There are several cards arranged in a row, and each card has an associated number of points The points are given in the integer array cardPoints.
//
// In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.
//
// Your score is the sum of the points of the cards you have taken.
//
// Given the integer array cardPoints and the integer k, return the maximum score you can obtain.
//
//
//
// Example 1:
//
// Input: cardPoints = [1,2,3,4,5,6,1], k = 3
// Output: 12
// Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.
//
// Example 2:
//
// Input: cardPoints = [2,2,2], k = 2
// Output: 4
// Explanation: Regardless of which two cards you take, your score will always be 4.
//
// Example 3:
//
// Input: cardPoints = [9,7,7,9,7,7,9], k = 7
// Output: 55
// Explanation: You have to take all the cards. Your score is the sum of points of all cards.
//
// Example 4:
//
// Input: cardPoints = [1,1000,1], k = 1
// Output: 1
// Explanation: You cannot take the card in the middle. Your best score is 1.
//
// Example 5:
//
// Input: cardPoints = [1,79,80,1,1,1,200,1], k = 3
// Output: 202
//
//
//
// Constraints:
//
//     1 <= cardPoints.length <= 10^5
//     1 <= cardPoints[i] <= 10^4
//     1 <= k <= cardPoints.length

func maxScore(cardPoints []int, k int) int {
	size := len(cardPoints)
	if size == 0 {
		return 0
	}

	var total int
	for i := 0; i < k; i++ {
		total += cardPoints[i]
	}

	maxTotal := total
	for i := k - 1; i >= 0; i-- {
		total = total - cardPoints[i] + cardPoints[size-(k-i)]
		maxTotal = max(maxTotal, total)
	}

	return maxTotal
}

func maxScore2(cardPoints []int, k int) int {
	size := len(cardPoints)
	if size == 0 || k == 0 {
		return 0
	}

	// sums[i] means total points from 0 to i, including i
	sums := make([]int, size)
	sums[0] = cardPoints[0]
	for i := 1; i < size; i++ {
		sums[i] = sums[i-1] + cardPoints[i]
	}

	if k == size {
		return sums[size-1]
	}

	var maxScore int
	for i := 0; i <= k; i++ {
		// i means number of cards are chosen from left
		if i == 0 {
			// all from right
			maxScore = max(maxScore, sums[size-1]-sums[size-1-k])
		} else if i == k {
			// all from left
			maxScore = max(maxScore, sums[k-1])
		} else {
			// some from left, some from right
			maxScore = max(maxScore, sums[i-1]+sums[size-1]-sums[size-1-k+i])
		}
	}

	return maxScore
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func maxScore1(cardPoints []int, k int) int {
	if k == 0 {
		return 0
	}

	var score int

	combinations(cardPoints, 0, 0, len(cardPoints)-1, k, &score)

	return score
}

func combinations(cards []int, currentScore int, left, right, target int, score *int) {
	if target == 0 {
		*score = max(*score, currentScore)
		return

	}

	if left > right {
		return
	}

	combinations(cards, currentScore+cards[left], left+1, right, target-1, score)
	combinations(cards, currentScore+cards[right], left, right-1, target-1, score)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	too slow

//	2.	wrong condition that left number should be range from 0 to k

//	3.	this problem is dp, but not the traditional dp that can get result
//		from i-1, j-1, etc. originally I think it as dp[i][j] means max
//		points from remaining cards i to j, but don't know how to proceed

//		then I suddenly think it has to choose k cards, either all in left,
//		all in right, all some from left and some from right. this separates
//		original into sub-problems

//		however, there's a corner case when k == size, which will cause
//		i == 0 index out of range, that's a test case need to be tested

//	4.	inspired from https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/discuss/597933/JAVA-simple-O(n)-solution

//		author has a brilliant solution, if think whole array as a closed
//		loop, then all possible answers come from start = k-1 & end = 0 to
//		start = -1 & end = -(k-1)

//	5.	inspired from https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/discuss/597763/Python3-Easy-Sliding-Window-O(n)%3A-Find-minimum-subarray

//		max score means sum of cards not chosen are minimum
