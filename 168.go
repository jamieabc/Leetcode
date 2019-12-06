package main

import "strings"

//Given a positive integer, return its corresponding column title as appear in an Excel sheet.
//
//For example:
//
//    1 -> A
//    2 -> B
//    3 -> C
//    ...
//    26 -> Z
//    27 -> AA
//    28 -> AB
//    ...
//Example 1:
//
//Input: 1
//Output: "A"
//Example 2:
//
//Input: 28
//Output: "AB"
//Example 3:
//
//Input: 701
//Output: "ZY"

func convertToTitle(n int) string {
	mapping := make(map[int]string)
	mapping[0] = "Z"
	mapping[1] = "A"
	mapping[2] = "B"
	mapping[3] = "C"
	mapping[4] = "D"
	mapping[5] = "E"
	mapping[6] = "F"
	mapping[7] = "G"
	mapping[8] = "H"
	mapping[9] = "I"
	mapping[10] = "J"
	mapping[11] = "K"
	mapping[12] = "L"
	mapping[13] = "M"
	mapping[14] = "N"
	mapping[15] = "O"
	mapping[16] = "P"
	mapping[17] = "Q"
	mapping[18] = "R"
	mapping[19] = "S"
	mapping[20] = "T"
	mapping[21] = "U"
	mapping[22] = "V"
	mapping[23] = "W"
	mapping[24] = "X"
	mapping[25] = "Y"
	mapping[26] = "Z"

	var tmp strings.Builder
	var remainder int
	for n > 26 {
		remainder = n % 26
		if remainder == 0 {
			n -= 26
			tmp.WriteString(mapping[26])
			n /= 26
		} else {
			n /= 26
			tmp.WriteString(mapping[remainder])
		}
	}
	tmp.WriteString(mapping[n%26])
	str := tmp.String()

	result := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		result[i] = str[len(str)-1-i]
	}
	return string(result)
}
