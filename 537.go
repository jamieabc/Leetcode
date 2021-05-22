package main

import (
	"fmt"
	"strconv"
)

// Given two strings representing two complex numbers.
//
// You need to return a string representing their multiplication. Note i2 = -1 according to the definition.
//
// Example 1:
//
// Input: "1+1i", "1+1i"
// Output: "0+2i"
// Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
//
// Example 2:
//
// Input: "1+-1i", "1+-1i"
// Output: "0+-2i"
// Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
//
// Note:
//
//    The input strings will not have extra blank.
//    The input strings will be given in the form of a+bi, where the integer a and b will both belong to the range of [-100, 100]. And the output should be also in this form.

func complexNumberMultiply(num1 string, num2 string) string {
	r1, i1 := convertComplex(num1)
	r2, i2 := convertComplex(num2)

	return fmt.Sprintf("%d+%di", r1*r2-i1*i2, r1*i2+i1*r2)
}

func convertComplex(str string) (int, int) {
	var n1, n2, idx int

	for i := range str {
		if str[i] == '+' {
			n1, _ = strconv.Atoi(str[:i])
			idx = i + 1
		} else if str[i] == 'i' {
			n2, _ = strconv.Atoi(str[idx:i])
		}
	}

	return n1, n2
}

type complex struct {
	real    int
	virtual int
}

func parse(str string) complex {
	var i int
	for i = 1; i < len(str); i++ {
		if str[i] == '+' {
			break
		}
	}
	real, _ := strconv.Atoi(str[:i])
	virtual, _ := strconv.Atoi(str[i+1 : len(str)-1])

	return complex{
		real:    real,
		virtual: virtual,
	}
}

func complexNumberMultiply1(a string, b string) string {
	c1 := parse(a)
	c2 := parse(b)

	real := c1.real*c2.real - c1.virtual*c2.virtual
	virtual := c1.real*c2.virtual + c1.virtual*c2.real

	return fmt.Sprintf("%d+%di", real, virtual)
}
