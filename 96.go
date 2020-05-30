package main

// Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
//
// Example:
//
// Input: 3
// Output: 5
// Explanation:
// Given n = 3, there are a total of 5 unique BST's:
//
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// BST(5) = F(1, 5) + F(2, 5) + F(3, 5) + F(4, 5) + F(5, 5)
			//        = G(0) * G(4) + G(1) * G(3) + G(2) * G(2) + G(3) * G(1) +
			//          G(4) * G(0)
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}

//	problems

//	1.	inspired from https://leetcode.com/problems/unique-binary-search-trees/discuss/31666/DP-Solution-in-6-lines-with-explanation.-F(i-n)-G(i-1)-*-G(n-i)

//		I didn't realize it's a dp problem at all, I am trying to use
//		recursion brute force solving problem but fail.

//		when constructing BST, for a chosen root, left part is
//		count(1 ~ root-1), right part is count(root+1 ~ n)

//		the tricky part is count(root+1 ~ n) = count(1 ~ n-root), because
//		number range doesn't matter, it only matters count, so number can be
//		replaced.

//		I can understand by math, but I don't know how this solution come
//		up.

//		BST(0) = 1
//		BST(1) = 1

//		BST(n) = F(1, n) + F(2, n) + F(3, n) + ... + F(n, n)
//				 F(m, n) BST count when m is root, sequence from 1 ~ n

//		further observation can be found that F(m, n) mean left part
//		is count(1 ~ m-1) * count(m+1 ~ n)

//		count(1 ~ m-1) = BST(m-1)
//		count(m+1 ~ n) = BST(n - (m+1) + 1) = BST(n-m) cause it's all about
//		total count, is same count for 5 ~ 8 and 2 ~ 5

///		thanks for author's explanation...

//	2.	inspiref from https://leetcode.com/problems/unique-binary-search-trees/discuss/31706/Dp-problem.-10%2B-lines-with-comments

//		this is more reasonable for me
