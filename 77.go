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
	ans := make([][]int, 0)

	recursive(n, k, 1, []int{}, &ans)

	return ans
}

func recursive(n, k, start int, cur []int, ans *[][]int) {
	if k == 0 {
		*ans = append(*ans, cur)
		return
	}

	// beware of limit, since there's not enough number left, terminate process
	for i := start; i <= n-k+1; i++ {
		tmp := make([]int, len(cur)+1)
		copy(tmp, cur)
		tmp[len(cur)] = i

		recursive(n, k-1, i+1, tmp, ans)
	}
}

// tc: O(k* combination n pick k), k for string operation
func combine4(n int, k int) [][]int {
	used := make([]bool, n)
	ans := make([][]int, 0)

	dfs(used, []int{}, 0, k, &ans)

	return ans
}

func dfs(used []bool, cur []int, idx, k int, ans *[][]int) {
	size := len(cur)

	if size == k {
		*ans = append(*ans, cur)
		return
	}

	// since dfs is pretty high tc,
	// prune path that is uncessary to accomplish task
	if idx > len(used)-(k-len(cur)) {
		return
	}

	for i := idx; i < len(used); i++ {
		if used[i] {
			continue
		}

		used[i] = true

		tmp := make([]int, len(cur)+1)
		copy(tmp, cur)
		tmp[len(tmp)-1] = i + 1
		dfs(used, tmp, i+1, k, ans)

		used[i] = false
	}
}

func combine3(n int, k int) [][]int {
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

func combine2(n int, k int) [][]int {
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

func combine1(n int, k int) [][]int {
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

//	5.	inspired from https://leetcode.com/problems/combinations/discuss/27002/Backtracking-Solution-Java

//		first comment explains why my recursive loop so slow, because some
//		paths won't get enough numbers, no need to do it
