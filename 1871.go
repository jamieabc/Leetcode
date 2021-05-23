package main

// You are given a 0-indexed binary string s and two integers minJump and maxJump. In the beginning, you are standing at index 0, which is equal to '0'. You can move from index i to index j if the following conditions are fulfilled:
//
// i + minJump <= j <= min(i + maxJump, s.length - 1), and
// s[j] == '0'.
//
// Return true if you can reach index s.length - 1 in s, or false otherwise.
//
//
//
// Example 1:
//
// Input: s = "011010", minJump = 2, maxJump = 3
// Output: true
// Explanation:
// In the first step, move from index 0 to index 3.
// In the second step, move from index 3 to index 5.
//
// Example 2:
//
// Input: s = "01101110", minJump = 2, maxJump = 3
// Output: false
//
//
//
// Constraints:
//
// 2 <= s.length <= 105
// s[i] is either '0' or '1'.
// s[0] == '0'
// 1 <= minJump <= maxJump < s.length

// tc: O(n)
func canReach(s string, minJump int, maxJump int) bool {
	deque := []int{0}
	size := len(s)

	if s[size-1] == '1' {
		return false

	}

	var visited int

	for len(deque) > 0 && deque[len(deque)-1] != size-1 {
		top := deque[0]
		deque = deque[1:]

		for i := max(visited+1, top+minJump); i <= min(top+maxJump, size-1); i++ {
			if s[i] == '0' {
				deque = append(deque, i)
			}
		}

		visited = max(visited, min(top+maxJump, size-1))
	}

	return len(deque) > 0 && deque[len(deque)-1] == size-1
}

// tc: O(n)
func canReach1(s string, minJump int, maxJump int) bool {
	size := len(s)

	if s[size-1] == '1' {
		return false
	}

	dp := make([]bool, size)
	dp[0] = true
	var to int

	for i := 0; i <= to; i++ {
		if dp[i] {
			for j := max(to, i+minJump); j <= min(size-1, i+maxJump); j++ {
				if s[j] == '0' {
					dp[j] = true
				}
			}
			to = min(size-1, i+maxJump)
		}
	}

	return dp[size-1]
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

//	Notes
//	1.	takes me 1 hour 30 minutes to finish this problem, w/ 14 wrong submissions

//		the problem I have is not to check reachable positions, i greedily check
//		only farthest position, but this is wrong

//		0 positions: 0, 5, 8, 12, minJump = 4, maxJump = 5
//		my thinking: 0 jump to 5, 5 can jump up to 10 (by maxJump), since 8 in
//		range of 5~10, 8 is also reachable

//		but 8 is actually not reachable, because minJump = 4, 5 can only jump to
//		9 or 10, not including 8

//		this is to problem of variable 'to', even number with range of 'to', still
//		might not be reachable

//		for i := 0; i < size-minJump; i++ {
//			if s[i] == '0' && i <= to {
//				...
//			}
//		}

//	2.	finally, i came up a solution to use dp[i] to store reachable points,
//		variable 'to' means already checking position, each time from index i can only
//		reach to max(to, i+minJump), and do checking up to i+maxJump

//		although there's 2 for loops, but variable 'to', each position will be
//		visited at most twice

//	3.	inspired from https://youtu.be/7eBeMGxX88k?t=1015

//		alex provides a simpler way: reachable points are used only once because
//		as i increasing, reachable range i+minJump ~ i+maxJump also increase,
//		so it can be modeled by a deque

//		the reason deque can be used is because i is always increasing, such that
//		previous reachable positions are also increasing, as long as one point is
//		invalid, all previous points are also invalid, which fits the behavior
//		of deque

//	4.	key point for this problem: if one position is check, don not check again,
//		reduces tc from O(n^2) to O(n)

//	5.	inspired from https://leetcode.com/problems/jump-game-vii/discuss/1224804/JavaC%2B%2BPython-One-Pass-DP

//		lee uses a variable 'pre' to know if current position can be visited, pre
//		is increased when new number can be reached from previous number, pre is
//		decreased when new number outside of reachable range
