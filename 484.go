package main

//  By now, you are given a secret signature consisting of character 'D' and 'I'. 'D' represents a decreasing relationship between two numbers, 'I' represents an increasing relationship between two numbers. And our secret signature was constructed by a special integer array, which contains uniquely all the different number from 1 to n (n is the length of the secret signature plus 1). For example, the secret signature "DI" can be constructed by array [2,1,3] or [3,1,2], but won't be constructed by array [3,2,4] or [2,1,3,4], which are both illegal constructing special string that can't represent the "DI" secret signature.
//
// On the other hand, now your job is to find the lexicographically smallest permutation of [1, 2, ... n] could refer to the given secret signature in the input.
//
// Example 1:
//
// Input: "I"
// Output: [1,2]
// Explanation: [1,2] is the only legal initial spectial string can construct secret signature "I", where the number 1 and 2 construct an increasing relationship.
//
// Example 2:
//
// Input: "DI"
// Output: [2,1,3]
// Explanation: Both [2,1,3] and [3,1,2] can construct the secret signature "DI",
// but since we want to find the one with the smallest lexicographical permutation, you need to output [2,1,3]
//
// Note:
// The input string will only contain the character 'D' and 'I'.
// The length of input string is a positive integer and will not exceed 10,000

func findPermutation(s string) []int {
	result := make([]int, 0)

	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == 'I' {
			for j, size := i+1, len(result); j > size; j-- {
				result = append(result, j)
			}
		}
	}

	return result
}

func findPermutation2(s string) []int {
	size := len(s) + 1

	// initialize ans, assume all chars are 'I'
	ans := make([]int, size)
	for i := range ans {
		ans[i] = i + 1
	}

	for start, end := 0, 0; start < size; {
		if end == size-1 || s[end] == 'I' {
			if start != end {
				// between start ~ end are D
				for i, j := start, min(size-1, end); i < j; i, j = i+1, j-1 {
					ans[i], ans[j] = ans[j], ans[i]
				}
				start, end = end+1, end+1
			} else {
				start, end = start+1, end+1
			}
		} else {
			// D, keep moving till I or end encountered
			end++
		}
	}

	return ans
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func findPermutation1(s string) []int {
	ans := make([]int, 0)
	stack := []int{1}
	size := len(s)

	for i, num := 0, 2; i < size; i, num = i+1, num+1 {
		if s[i] == 'I' {
			if i > 0 && s[i-1] == 'D' {
				for len(stack) > 0 {
					ans = append(ans, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
			} else {
				ans = append(ans, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, num)
	}

	for len(stack) > 0 {
		ans = append(ans, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return ans
}

//	problems
//	1.	brute force, should use better way to find solution

//	2.	use ugly way to find solution

//	3.	from reference https://leetcode.com/problems/find-permutation/discuss/96613/Java-O(n)-clean-solution-easy-to-understand

//		During thinking, I have occurred an idea that D is reverse increase,
//		then I focus on how to do reverse increase.

//		But author does it in another way, reverse increase is normal increase
//		with swap. This is a clever solution, I have chance to come out but
//		miss it. The insight to this problem is still not enough.

//		Also, the other amazing part is swapping, 3 operations of XOR.

//	  	 a <-> b:
//		  a ^= b
//		  b ^= a
//		  a ^= b

//	4.	from reference https://leetcode.com/problems/find-permutation/discuss/96663/Greedy-O(n)-JAVA-solution-with-explanation

//		author has a brilliant solution to add '.' to end of s, makes coding
//		easier.

//		also, he uses != 'D' to do operation, which includes added '.'

//	5.	from reference https://leetcode.com/problems/find-permutation/discuss/96644/C%2B%2B-simple-solution-in-72ms-and-9-lines

//		I view D as limitation, so I try to find D, then I's value is
//		determined by D.

//		author uses array size to calculate D. difference between I's index &
//		array size is D amount. very impressive....

//	6.	stack solution from answer, when I is reached, push into stack, then
//		pop out all numbers in stack, when D is reached, push into stack.
