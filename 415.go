package main

//Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.
//
//Note:
//
//The length of both num1 and num2 is < 5100.
//Both num1 and num2 contains only digits 0-9.
//Both num1 and num2 does not contain any leading zero.
//You must not use any built-in BigInteger library or convert the inputs to integer directly.

func addStrings(num1 string, num2 string) string {
	var carry uint
	len1 := len(num1)
	len2 := len(num2)

	base := num1
	add := num2
	if len2 > len1 {
		base = num2
		add = num1
	}

	length := len(base) + 1 // in case it has carry
	result := make([]byte, length)

	var i int
	for i = 0; i < len(add); i++ {
		n1 := str2uint(base[len(base)-1-i])
		n2 := str2uint(add[len(add)-1-i])
		sum := n1 + n2 + carry

		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}

		result[length-1-i] = int2byte(sum)
	}

	for j := 0; j < len(base)-len(add); j++ {
		n1 := str2uint(base[len(base)-1-i-j])
		sum := n1 + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		result[length-1-i-j] = int2byte(sum)
	}

	// in case final carry exist
	if carry == 1 {
		result[0] = '1'
		return string(result)
	}

	return string(result[1:])
}

func str2uint(b byte) uint {
	return uint(b - '0')
}

func int2byte(i uint) byte {
	return byte(i + '0')
}

// problems
// 1. wrong start condition, because it needs to start from least significant, which means different start location
// 2. optimize for memory
// 3. wrong index of base
// 4. second loop write destination is wrong (missing -1), this is stupid, takes me 1 hour and printf to find, I need to be aware of this, wrong position might end up wrong data at all
