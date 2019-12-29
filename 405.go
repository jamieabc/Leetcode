package main

import (
	"strings"
)

// Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, two’s complement method is used.
//
//Note:
//
//    All letters in hexadecimal (a-f) must be in lowercase.
//    The hexadecimal string must not contain extra leading 0s. If the number is zero, it is represented by a single zero character '0'; otherwise, the first character in the hexadecimal string will not be the zero character.
//    The given number is guaranteed to fit within the range of a 32-bit signed integer.
//    You must not use any method provided by the library which converts/formats the number to hex directly.
//
//Example 1:
//
//Input:
//26
//
//Output:
//"1a"
//
//Example 2:
//
//Input:
//-1
//
//Output:
//"ffffffff"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	var sb strings.Builder
	minus := false
	if num < 0 {
		minus = true
		num = -num
	}

	// to hex
	iteration := 0
	for num != 0 {
		remainder := num % 16
		sb.WriteString(digitToHex(remainder))
		num /= 16
		iteration++
	}

	if !minus {
		return reverseString(sb.String())
	}

	// restore string to 8 digits, zero is crucial for 2's compliment
	for iteration < 8 {
		sb.WriteString("0")
		iteration++
	}

	var sb2 strings.Builder
	for _, s := range sb.String() {
		sb2.WriteString(hexCompliment(string(s)))
	}

	// compliment from reference https://electronics.stackexchange.com/a/84820
	tmp := reverseString(sb2.String())
	compliment := hexToDigit(string(tmp[0]))
	for i := 1; i < len(tmp); i++ {
		compliment <<= 4
		compliment += hexToDigit(string(tmp[i]))
	}
	compliment++

	sb.Reset()
	for compliment != 0 {
		remainder := compliment % 16
		sb.WriteString(digitToHex(remainder))
		compliment /= 16
	}

	result := reverseString(sb.String())
	var i int
	for i = 0; i < len(result); i++ {
		if result[i] != '0' {
			break
		}
	}

	return result[i:]
}

func hexToDigit(str string) int {
	switch str {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "a":
		return 10
	case "b":
		return 11
	case "c":
		return 12
	case "d":
		return 13
	case "e":
		return 14
	case "f":
		return 15
	default:
		return 0
	}
}

func digitToHex(num int) string {
	switch num {
	case 0:
		return "0"
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	case 10:
		return "a"
	case 11:
		return "b"
	case 12:
		return "c"
	case 13:
		return "d"
	case 14:
		return "e"
	case 15:
		return "f"
	default:
		return "f"
	}
}

func hexCompliment(str string) string {
	switch str {
	case "0":
		return "f"
	case "1":
		return "e"
	case "2":
		return "d"
	case "3":
		return "c"
	case "4":
		return "b"
	case "5":
		return "a"
	case "6":
		return "9"
	case "7":
		return "8"
	case "8":
		return "7"
	case "9":
		return "6"
	case "a":
		return "5"
	case "b":
		return "4"
	case "c":
		return "3"
	case "d":
		return "2"
	case "e":
		return "1"
	case "f":
		return "0"
	default:
		return "0"
	}
}

func reverseString(str string) string {
	var sb strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte(str[i])
	}
	return sb.String()
}

// problems
// 1. forget about string order is reversed
// 2. when doing minus operation, last operation of converting to hex use wrong number
// 3. when minus, need zero for compliment
// 4. when convert hex to digit, first digit should not always multiply by 16
// 5. leading zero should not returned
// 6. when truncate leading zero, variable not exported
// 7 .digitToHex missing one condition of 0 digit
// 8. missing one condition of hexCompliment
