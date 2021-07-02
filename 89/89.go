package main

// The gray code is a binary numeral system where two successive values differ in only one bit.
//
// Given an integer n representing the total number of bits in the code, return any sequence of gray code.
//
// A gray code sequence must begin with 0.
//
//
//
// Example 1:
//
// Input: n = 2
// Output: [0,1,3,2]
// Explanation:
// 00 - 0
// 01 - 1
// 11 - 3
// 10 - 2
// [0,2,3,1] is also a valid gray code sequence.
// 00 - 0
// 10 - 2
// 11 - 3
// 01 - 1
//
// Example 2:
//
// Input: n = 1
// Output: [0,1]
//
//
//
// Constraints:
//
// 1 <= n <= 16

// tc: O(2^n), bottom-up, need to build all previous n-1 sequences
// sc: O(n)
func grayCode(n int) []int {
	ans := make([]int, 0)
	ans = append(ans, 0, 1)

	for i := 2; i <= n; i++ {
		size := len(ans)
		for j := size - 1; j >= 0; j-- {
			ans = append(ans, ans[j]|(1<<(i-1)))
		}
	}

	return ans
}

// tc: O(2^n), sc: O(1)
func grayCode3(n int) []int {
	visited := make(map[int]bool)
	ans := []int{0}
	visited[0] = true

	for len(ans) < (1 << n) {
		num := ans[len(ans)-1]
		for j := 0; j < n; j++ {
			tmp := num ^ (1 << j)
			if !visited[tmp] {
				ans = append(ans, tmp)
				visited[tmp] = true
				break
			}
		}
	}

	return ans
}

// tc: O(2^n), sc: O(n)
func grayCode2(n int) []int {
	return dfs2(n, 0, []int{0})
}

func dfs2(n, num int, cur []int) []int {
	if n == num {
		return cur
	}

	size := len(cur)
	tmp := append([]int{}, cur...)

	for i := size-1; i >= 0; i-- {
		tmp = append(tmp, cur[i] | (1 << num))
	}

	return dfs2(n, num+1, tmp)
}

// tc: O(n * 2^n)
func grayCode1(n int) []int {
    used := make(map[int]bool)
    used[0] = true
    ans := []int{0}

    dfs1(n, 0, used, &ans)

    return ans
}

func dfs1(n, cur int, used map[int]bool, ans *[]int) bool {
    if len(*ans) == (1 << n) {
        var count int
        for i := 0; i < n; i++ {
            if (*ans)[0] & (1 << i) != (*ans)[len(*ans)-1] & (1 << i) {
                count++
            }
        }

        return count == 1
    }

    var next int
    for i := 0; i < n; i++ {
        // iterate through all bits
        next = cur ^ (1 << i)

        if exist, ok := used[next]; !ok || !exist {
            used[next] = true

            *ans = append(*ans, next)

            if dfs1(n, next, used, ans) {
                return true
            }

            *ans = (*ans)[:len(*ans)-1]

            used[next] = false
        }
    }

    return false
}

//	Notes
//	1.	using visited to store if a number is visited, it's kind of slow because
//		additional checks are needed

//	2.	inspired from https://leetcode.com/problems/gray-code/discuss/29891/Share-my-solution

//		from previous sequence, append 0 from left position, iterate forward
//		then append 1 from left position, iterate backward
//
//		the reason this works is because, from previous sequence, it already valid that each number
//		differs from 1 bit, so append 0 won't change any difference
//
//		for last number, append 0 at left & append 1 at left differs exactly 1 bit, still meets
//		then traverse previuos sequence backward, append 1 at left will still hold condition
//
//		very brilliant solution
//
//		1 bit:
//		0 -> 1
//
//		2 bits:
//		00 -> 01 (from 1 bit seq, append 0 at left) -> 11 -> 10 (from 1 bit seq, append 1 at left, backward)
//
//		3 bits:
//		000 -> 001 -> 011 -> 010 -> 110 -> 111 -> 101 -> 100
//
//		for 1 bit, iterate over previous [0], 2^0
//		for 2 bits, iterate over previous [0, 1], 2^1
//		for 3 bits, iterate over previous [00, 01, 11, 10], 2^2
//
//		for n bits, bttom-up needs to build all previous, iterate from 2^0 + 2^1 +  2^2 + .. + 2^(n-1) = 2^n
//		overall tc O(2^n)

//	3.	after some time, cannot find pattern to generate code, use backtracking which is really slow
//
//	4.	inspired from sample code, the recursion solution can be optimized, because all variables are static
//		no need stack to store
