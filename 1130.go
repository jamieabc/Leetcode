package main

import "math"

// Given an array arr of positive integers, consider all binary trees such that:
//
//     Each node has either 0 or 2 children;
//     The values of arr correspond to the values of each leaf in an in-order traversal of the tree.  (Recall that a node is a leaf if and only if it has 0 children.)
//     The value of each non-leaf node is equal to the product of the largest leaf value in its left and right subtree respectively.
//
// Among all possible binary trees considered, return the smallest possible sum of the values of each non-leaf node.  It is guaranteed this sum fits into a 32-bit integer.
//
//
//
// Example 1:
//
// Input: arr = [6,2,4]
// Output: 32
// Explanation:
// There are two possible trees.  The first has non-leaf node sum 36, and the second has non-leaf node sum 32.
//
//     24            24
//    /  \          /  \
//   12   4        6    8
//  /  \               / \
// 6    2             2   4
//
//
//
// Constraints:
//
//     2 <= arr.length <= 40
//     1 <= arr[i] <= 15
//     It is guaranteed that the answer fits into a 32-bit signed integer (ie. it is less than 2^31).

func mctFromLeafValues(arr []int) int {
	length := len(arr)
	var sum int

	if length == 0 {
		return sum
	}
	stack := []int{math.MaxInt32}

	// generate descending order sequence
	for _, i := range arr {
		for len(stack) > 1 {
			s := stack[len(stack)-1]
			if i <= s {
				break
			}

			stack = stack[:len(stack)-1]
			sum += s * min(stack[len(stack)-1], i)
		}

		stack = append(stack, i)
	}

	// process descending order sequence
	for len(stack) > 2 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += s * stack[len(stack)-1]
	}

	return sum
}

// tc: O(n^2)
func mctFromLeafValues4(arr []int) int {
	var ans int

	for len(arr) > 1 {
		smallest := math.MaxInt32
		var toRemove int

		for i := 0; i < len(arr)-1; i++ {
			tmp := arr[i] * arr[i+1]
			if tmp < smallest {
				smallest = tmp
				toRemove = i
			}
		}

		ans += smallest

		// smaller leaf will never be used anymore
		if arr[toRemove] > arr[toRemove+1] {
			toRemove++
		}

		tmp := make([]int, len(arr)-1)
		copy(tmp, arr[:toRemove])
		copy(tmp[toRemove:], arr[toRemove+1:])
		arr = tmp
	}

	return ans
}

// tc: O(n^2)
func mctFromLeafValues3(arr []int) int {
	nums := make([]int, 16)
	for i := range arr {
		nums[arr[i]]++
	}

	var sum, tmp int
	for len(arr) > 1 {
		// find minimum
		for i := range nums {
			if nums[i] > 0 {
				tmp = i
				nums[i]--
				break
			}
		}

		// select from minimum neighbor
		for i := range arr {
			if arr[i] == tmp {
				if i == 0 {
					sum += tmp * arr[i+1]
				} else if i == len(arr)-1 {
					sum += tmp * arr[i-1]
				} else {
					sum += tmp * min(arr[i-1], arr[i+1])
				}
				t := append([]int{}, arr[:i]...)
				arr = append(t, arr[i+1:]...)
				break
			}
		}
	}

	return sum
}

func mctFromLeafValues2(arr []int) int {
	size := len(arr)
	memo := make([][]int, size)
	for i := range memo {
		memo[i] = make([]int, size)
	}

	return dfs(arr, memo, 0, len(arr)-1)
}

// dfs returns sum of non-leaves from start ~ end
// tc: O(n^3), separate into n sub-problems, each sub-problem scan takes n
// 1 + 2^2 + 3^2 + 4^2 + ... + n^2 = O(n^3)
func dfs(arr []int, memo [][]int, start, end int) int {
	// leaf
	if start == end {
		return 0
	}

	if memo[start][end] != 0 {
		return memo[start][end]
	}

	smallest := math.MaxInt32
	for i := start; i < end; i++ {
		left, right := arr[start], arr[i+1]

		for j := start + 1; j <= i; j++ {
			left = max(left, arr[j])
		}

		for j := i + 2; j <= end; j++ {
			right = max(right, arr[j])
		}

		smallest = min(smallest, dfs(arr, memo, start, i)+dfs(arr, memo, i+1, end)+left*right)
	}

	memo[start][end] = smallest

	return smallest
}

// tc: O(n^2)
func mctFromLeafValues1(arr []int) int {
	length := len(arr)
	maxi := make([][]int, length)
	for i := range maxi {
		maxi[i] = make([]int, length)
	}

	// table to find max number in specific range
	for i := range maxi {
		maxi[i][i] = arr[i]
		for j := i + 1; j < length; j++ {
			maxi[i][j] = max(maxi[i][j-1], arr[j])
		}
	}

	memo := make([][]int, length)
	for i := range memo {
		memo[i] = make([]int, length)
	}

	return dp(0, length-1, memo, maxi)
}

func dp(left, right int, memo, maxi [][]int) int {
	// leaf node
	if left == right {
		return 0
	}

	// already store value
	if memo[left][right] != 0 {
		return memo[left][right]
	}

	tmp := math.MaxInt32
	for i := left; i < right; i++ {
		tmp = min(tmp,
			maxi[left][i]*maxi[i+1][right]+
				dp(left, i, memo, maxi)+
				dp(i+1, right, memo, maxi))
	}
	memo[left][right] = tmp

	return tmp
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
//	1.	don't know how to do it

//	2.	inspired from https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/340033/C%2B%2B-with-comments

//		dp(left, right) = min(max(arr[left, left+1, ..., i]) * max(arr[i+1,
//		i+2, ..., right] + dp(left, i) + dp(i+1, right)

//		honestly, even with this formula, I still don't know how to do it.

//		time complexity is also hard...it's O(n^3). what I understand is for
//		dp, for each left, it's right value can be range from left +1,
//		left + 2, ..., length-1, so complexity for dp is O(n^2), inside dp
//		it also calls n

//	3.	reference from https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/340027/Java-DP-easy-to-understand

//		the structure shows more clearly that time complexity is O(n^3)

//	3.	inspired from https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/478708/RZ-Summary-of-all-the-solutions-I-have-learned-from-Discuss-in-Python

//		for dfs/dp, when every two numbers combined, it always exist left side
//		and right side

//		e.g. [6, 2, 4]
// 			     8
//			   /  \
//			  12   4
//			 /  \
//			6    2

//		left side | right side
//		leaf numbers are divided by final root value, which is 8
//		left [6, 2], right [8]
//		dfs keeps separate array into left & right, from left side boundary
//		ranges from

//		if I didn't observe the way to separate array, I might not think of a
//		way to solve this

//		greedy comes from observation, that minimum sum is to put large
//		number close to root, and smaller number away from root. the
//		procedure is start from minimum of array, choose adjacent values
//		that is smaller, build tree bottom up

//		this approach finds minimum value in the list, so it doesn't matter
//		how many minimum value exit, cause those numbers will eventually
//		be chosen, so just focus on finding minimum item in the list

//		and it's a point to remove whenever minimum number is used, cause
//		it will not be chosen anyway

//	4.	add reference https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/339959/One-Pass-O(N)-Time-and-Space

//		lee transforms problem into another one:
//		Given an array A, choose two neighbors in the array a and b,
//		we can remove the smaller one min(a,b) and the cost is a * b.
//		What is the minimum cost to remove the whole array until only one
//		left?

//		lee's solution is really beautiful, I think he sees nature of
//		problems differently, just like he transforms problem into finding
//		next greater number in array, which can be solved by stack.

//		what I understand is that first he processes array to keep only
//		decreasing sequence, if it's a sequence that sort descending,
//		then finding next greater item is the previous one of current
//		item

//		the question for me is: can I see through nature of problem as he
//		does, or how can I have this kind of ability?

//		this reminds of listening podcasts, learning. I though if I can
//		think more question, my thinking skill can be improved somehow.

//	5.	after 8 months later, don't know how to solve thi sproblem
//
//		inspired from https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/951938/Don't-overthink-about-trees.-It's-a-DPGreedy-problem.

//		with this hint, I kind of know the rule, but still not clear how
//		to solve this problem

//	6.	inspired form https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/349098/From-O(N2)-to-O(n).-Greedy

//		I kind of know the start of thinking, in-order tree leaf is an
//		illusion, the truth is smaller number removed during grouping

//		however, with this, I still not able to solve this

//		this version is easier for me to understand greedy version

//	7.	inspired from https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/474188/I-think-I-able-to-explain-it-to-myself-and-to-you...(Java-DP)-.-Complexity-is-in-the-question/494578

//		for array of size n, separated into n sub-problems, each sub problem
//		scan for it length to find max values, total sequence is
//		1 + 2^2 + 3^2 + 4^2 + ... + N^2 = O(N^3)

//	8.	inspired from https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/349098/From-O(N2)-to-O(n).-Greedy

//		author explains why using mono-decreasing stack guarantees optimal
//		result, very brilliant
