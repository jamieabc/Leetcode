package main

//Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.
//
//Example 1:
//
//Input: 2
//Output: [0,1,1]
//Example 2:
//
//Input: 5
//Output: [0,1,1,2,1,2]
//Follow up:
//
//It is very easy to come up with a solution with run time O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a single pass?
//Space complexity should be O(n).
//Can you do it like a boss? Do it without using any builtin function like __builtin_popcount in c++ or in any other language.

// https://www.youtube.com/watch?v=QjEyO1137cM
func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}

	result := make([]int, num+1)
	power := 1

	for power <= num {
		i := 0
		for i+power <= num && i < power {
			result[i+power] = result[i] + 1
			i++
		}
		power = power << 1
	}
	return result
}

func countBits1(num int) []int {
	if num == 0 {
		return []int{0}
	}

	result := make([]int, num+1)
	result[0] = 0
	for i := 1; i <= num; i++ {
		result[i] = countOne(i)
	}
	return result
}

func countOne(num int) int {
	i := 0
	for num != 0 {
		if num&1 == 1 {
			i++
		}
		num = num >> 1
	}
	return i
}
