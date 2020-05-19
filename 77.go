package main

import "crypto/x509"

// Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
//
// Example:
//
// Input: n = 4, k = 2
// Output:
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	stack := make([]int, k)

	for s, num := 0, 1; true; {
		if s < k {
			stack[s] = num
			s, num = s+1, num+1
		} else {
			t := make([]int, k)
			copy(t, stack)
			result = append(result, t)

			if stack[0] == n-k+1 {
				break
			}

			for s--; stack[s] == n-(k-1-s); s-- {
			}
			num = stack[s] + 1
		}
	}

	return result
}

func combine1(n int, k int) [][]int {
	result := make([][]int, 0)
	if n < k {
		return result
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	stack := make([]int, k)
	for s, i := 0, 0; true; {
		if s < k {
			stack[s] = i
			s, i = s+1, i+1
		} else {
			tmp := make([]int, k)
			for j := range tmp {
				tmp[j] = nums[stack[j]]
			}
			result = append(result, tmp)

			if stack[0] == n-k {
				break
			}

			for s--; stack[s] == (n-1)-(k-1-s); s-- {
			}
			i = stack[s] + 1
		}
	}

	return result
}

func combine2(n int, k int) [][]int {
	result := make([][]int, 0)

	if n < k {
		return result
	}

	flags := make([]bool, n)
	tmp := make([]int, k)

	combinations(n, k, 1, 0, flags, tmp, &result)

	return result
}

func combinations(n, k, start, idx int, flags []bool, tmp []int, result *[][]int) {
	if idx == k {
		t := append([]int{}, tmp...)
		*result = append(*result, t)
		return
	}

	for i := start; i <= n; i++ {
		if flags[i-1] {
			continue
		}

		tmp[idx] = i
		idx++
		flags[i-1] = true
		combinations(n, k, i+1, idx, flags, tmp, result)
		flags[i-1] = false
		idx--
	}
}

//	problems
//	1.	too slow, I think it's because too many memory allocation of tmp

//	2.	inspired from https://leetcode.com/problems/combinations/discuss/27111/My-shortest-c%2B%2B-solutionusing-dfs

//		if n < k, no need to do anything

//	3.	add iterative solution, pretty ugly, need improvement

//	4.	add reference https://leetcode.com/problems/combinations/discuss/27090/DP-for-the-problem

//		author use dp to solve problem, but I didn't take time to read it
