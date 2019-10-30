package main

import (
	"fmt"
	"unicode"
)

//Implement atoi which converts a string to an integer.
//
//The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.
//
//The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.
//
//If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.
//
//If no valid conversion could be performed, a zero value is returned.
//
//Note:
//
//Only the space character ' ' is considered as whitespace character.
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
//
//Example 1:
//
//Input: "42"
//Output: 42
//
//Example 2:
//
//Input: "   -42"
//Output: -42
//Explanation: The first non-whitespace character is '-', which is the minus sign.
//Then take as many numerical digits as possible, which gets 42.
//
//Example 3:
//
//Input: "4193 with words"
//Output: 4193
//Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
//
//Example 4:
//
//Input: "words and 987"
//Output: 0
//Explanation: The first non-whitespace character is 'w', which is not a numerical
//digit or a +/- sign. Therefore no valid conversion could be performed.
//
//Example 5:
//
//Input: "-91283472332"
//Output: -2147483648
//Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
//Thefore INT_MIN (−231) is returned.

// remove leading white spaces then check positive or negative
// parse every character afterwards, terminates when any non-digit character exist
// check for maximum/minimum

const (
	maxInt    = 2147483647
	minInt    = -2147483648
	zeroAscii = 48
)

// remove all leading white spaces
// check if first non-digit character is + or -
// remove leading 0 digits
// get number when encounter fir non-digit character
// convert to number
func myAtoi(str string) int {
	start := 0
	length := len(str)

	if 0 == length {
		return 0
	}

	runes := []rune(str)

	// skip leading white spaces
	for i, s := range runes {
		if s == ' ' {
			continue
		}
		start = i
		break
	}

	minus := false

	// leading character not - or digit
	if !unicode.IsDigit(runes[start]) {
		if runes[start] != '-' && runes[start] != '+' {
			return 0
		}

		if runes[start] == '-' {
			minus = true
		}
		start++
	}

	// skip leading zeros
	for start < length && runes[start] == '0' {
		start++
	}

	i := start
	end := length - 1

loop:
	for i < length {
		if unicode.IsDigit(runes[i]) {
			i++
		} else {
			break loop
		}
	}

	end = i - 1

	var result int64
	multiply := int64(1)

	for i = end; i >= start; i-- {
		result += (int64(runes[i]) - zeroAscii) * multiply
		newMultiply := multiply * 10
		if newMultiply < multiply || result > maxInt {
			if minus {
				return minInt
			} else {
				return maxInt
			}
		}
		multiply = newMultiply

	}

	// check limit
	if minus {
		return int(-result)
	} else {
		return int(result)
	}
}

func main() {
	fmt.Printf("100: %d\n", myAtoi("    03455"))
}
