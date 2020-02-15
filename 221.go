package main

//Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
//
//Example:
//
//Input:
//
//1 0 1 0 0
//1 0 1 1 1
//1 1 1 1 1
//1 0 0 1 0
//
//Output: 4

func maximalSquare(matrix [][]byte) int {
	yLength := len(matrix)
	if yLength == 0 {
		return 0
	}

	xLength := len(matrix[0])
	dp := make([]int, xLength)

	max := 0
	var tmp, prev int
	for j := 0; j < yLength; j++ {
		for i := 0; i < xLength; i++ {
			if i == 0 || j == 0 {
				// for top and left-most lines
				prev = dp[i]
				if matrix[j][i] == '1' {
					dp[i] = 1
					if max == 0 {
						max = 1
					}
				} else {
					dp[i] = 0
				}
				continue
			}

			tmp = dp[i]
			if matrix[j][i] == '1' {
				m := min(dp[i-1], min(dp[i], prev)) + 1
				if m > max {
					max = m
				}
				dp[i] = m
			} else {
				dp[i] = 0
			}
			prev = tmp
		}
	}
	return max * max
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

// problems
// 1. wrong range of check i, j
// 2. during check square (check), when matrix[j][i] == '0', break only
//    breaks one level of loop, which cause j always to end of matrix
// 3. when length is 1, need to check if 1 exist
// 4. it's not just minPositive, it should be minPositive positive
// 5. cannot just put x into i, there could exist some situation that x+1
//    is valid
// 6. what original thinking was a fix x, find y that fits to this x, but
//    this is wrong
//    e.g.
//    1 1 1 1 1 1 0
//    1 1 1 1 1 1 0
//    1 1 1 1 0 0 0
//    1 1 1 1 0 0 0
//    by original thinking, x is 6, find vertical line that fits length to 6
//    is 2, but then it misses the length of 4
// 7. when shrinking, choose the restriction that is larger because y is not
//    iterate through
// 8. optimize, use another way to check square, for every point, check
//    length increment 1 by 1
//    e.g.    . =>  . .  => . . .
//                  . .     . . .
//                          . . .
// 9. optimize, use dp, another slice that counts for maximum for a node
//    e.g.  0 0 1 1 1     0 0 1 1 1
//          0 0 0 1 1  => 0 0 0 1 2
//          0 0 0 0 0     0 0 0 0 0
//    the rule of dp increment is that left, up, left-up are all same
//    number, so that point can be number+1
// 10. when using dp, if any node is 1, update max to 1
// 11. when using dp, the condition to increment is  min of surrounded + 1
// 12. optimize, dp[[j][i] = min(dp[j-1, i], dp[j, i-1], dp[j-1, i-1]) + 1
