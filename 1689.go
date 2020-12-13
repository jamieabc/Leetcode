package main

// A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros. For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.
//
// Given a string n that represents a positive decimal integer, return the minimum number of positive deci-binary numbers needed so that they sum up to n.
//
//
//
// Example 1:
//
// Input: n = "32"
// Output: 3
// Explanation: 10 + 11 + 11 = 32
//
// Example 2:
//
// Input: n = "82734"
// Output: 8
//
// Example 3:
//
// Input: n = "27346209830709182346"
// Output: 9
//
//
//
// Constraints:
//
//     1 <= n.length <= 105
//     n consists of only digits.
//     n does not contain any leading zeros and represents a positive integer.

func minPartitions(n string) int {
	var count int

	for i := range n {
		count = max(count, int(n[i]-'0'))
	}

	return count
}

func minPartitions1(n string) int {
	var count int

	size := len(n)
	for i := 0; i < size; i++ {
		num := int(n[i] - '0')
		if num > count {
			count += num - count
		}
	}

	return count
}

//	Notes
//	1.	for a better observation, it's actually maximum among all digits
