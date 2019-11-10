package main

import (
	"fmt"
)

//Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
//
//Example 1:
//
//Input: 121
//Output: true
//Example 2:
//
//Input: -121
//Output: false
//Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//Example 3:
//
//Input: 10
//Output: false
//Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//Follow up:
//
//Could you solve it without converting the integer to a string?

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// single digit
	if x <= 9 {
		return true
	}

	reversed := reverse(x)
	return reversed == x
}

func reverse(x int) int {
	var reversed, digit int

	for {
		digit = x % 10
		reversed += digit
		if x >= 10 {
			reversed *= 10
			x /= 10
		} else {
			return reversed
		}
	}
}

func main() {
	//x := 121
	x := 9999
	fmt.Printf("%d is palindrone: %t\n", x, isPalindrome(x))
}
