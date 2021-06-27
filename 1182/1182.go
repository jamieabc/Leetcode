package main

// You are given an array colors, in which there are three colors: 1, 2 and 3.

// You are also given some queries. Each query consists of two integers i and c, return the shortest distance between the given index i and the target color c. If there is no solution return -1.



// Example 1:

// Input: colors = [1,1,2,1,3,2,2,3,3], queries = [[1,3],[2,2],[6,1]]
// Output: [3,0,3]
// Explanation:
// The nearest 3 from index 1 is at index 4 (3 steps away).
// The nearest 2 from index 2 is at index 2 itself (0 steps away).
// The nearest 1 from index 6 is at index 3 (3 steps away).

// Example 2:

// Input: colors = [1,2], queries = [[0,3]]
// Output: [-1]
// Explanation: There is no 3 in the array.



// Constraints:

//     1 <= colors.length <= 5*10^4
//     1 <= colors[i] <= 3
//     1 <= queries.length <= 5*10^4
//     queries[i].length == 2
//     0 <= queries[i][0] < colors.length
//     1 <= queries[i][1] <= 3

func shortestDistanceColor(colors []int, queries [][]int) []int {
    var prev int
    n := len(colors)
    dp := make([][]int, 3)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    for i := 1; i <= 3; i++ {
        prev = -1
        for j := range colors {
            if colors[j] == i {
                dp[i-1][j] = 0

                if prev != -1 {
                    for k, l, dist := prev+1, j-1, 1; k <= l; k, l, dist = k+1, l-1, dist+1 {
                        dp[i-1][k] = dist
                        dp[i-1][l] = dist
                    }
                } else {
                    for k, dist := j-1, 1; k >= 0; k, dist = k-1, dist+1 {
                        dp[i-1][k] = dist
                    }
                }

                prev = j
            }
        }

        // update remaining if color i exist
        if prev != -1 {
            for j, dist := prev+1, 1; j < n; j, dist = j+1, dist+1 {
                dp[i-1][j] = dist
            }
        }
    }

    ans := make([]int, len(queries))

    for i, q := range queries {
        ans[i] = dp[q[1]-1][q[0]]
    }

    return ans
}
