package main

import (
	"fmt"
	"strings"
)

//The count-and-say sequence is the sequence of integers with the first five terms as following:
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//
//1 is read off as "one 1" or 11.
//11 is read off as "two 1s" or 21.
//21 is read off as "one 2, then one 1" or 1211.
//
//Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence. You can do so recursively, in other words from the previous member read off the digits, counting the number of digits in groups of the same digit.
//
//Note: Each term of the sequence of integers will be represented as a string.
//
//
//
//Example 1:
//
//Input: 1
//Output: "1"
//Explanation: This is the base case.
//
//Example 2:
//
//Input: 4
//Output: "1211"
//Explanation: For n = 3 the term was "21" in which we have two groups "2" and "1", "2" can be read as "12" which means frequency = 1 and value = 2, the same way "1" is read as "11", so the answer is the concatenation of "12" and "11" which is "1211".
//

func countAndSay(n int) string {
	return recursive("1", n)
}

func recursive(str string, number int) string {
	if number <= 1 {
		return str
	}

	var sb strings.Builder
	count := 0
	target := uint8(0)

	for i := range str {
		if target == 0 {
			// initial state
			target = str[i]
			count = 1
		} else if str[i] == target {
			// still counting
			count++
		} else {
			// encounter new string
			sb.WriteString(fmt.Sprintf("%d%s", count, string(target)))
			count = 1
			target = str[i]
		}
	}

	// last string is not write
	sb.WriteString(fmt.Sprintf("%d%s", count, string(target)))

	return recursive(sb.String(), number-1)
}

// problems
// 1. forget to update variable name
