package main

// There is a fence with n posts, each post can be painted with one of the k colors.
//
// You have to paint all the posts such that no more than two adjacent fence posts have the same color.
//
// Return the total number of ways you can paint the fence.
//
// Note:
// n and k are non-negative integers.
//
// Example:
//
// Input: n = 3, k = 2
// Output: 6
// Explanation: Take c1 as color 1, c2 as color 2. All possible ways are:
//
//             post1  post2  post3
//  -----      -----  -----  -----
//    1         c1     c1     c2
//    2         c1     c2     c1
//    3         c1     c2     c2
//    4         c2     c1     c1
//    5         c2     c1     c2
//    6         c2     c2     c1

func numWays(n int, k int) int {
	dp := make([]int, n)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return k
	}

	if n == 2 {
		return k * k
	}

	dp[0] = k
	dp[1] = k * k

	for i := 2; i < n; i++ {
		dp[i] = dp[i-2]*(k-1) + dp[i-1]*(k-1)
	}

	return dp[n-1]
}

//	problems
//	1.	having hard time to solve it...I know this is tree problem, and
//		by no more than two adjacent are same color, I know it's a filter
//		criteria, but I don't know how to write it down into recursive
//		format

//	2.	add reference https://leetcode.com/problems/paint-fence/discuss/178010/The-only-solution-you-need-to-read

//		I have same idea, but I miss something that for bot same & different
//		color, selections are k-1

//	3.	forget about boundary conditions, when n =0, 1, and so on
