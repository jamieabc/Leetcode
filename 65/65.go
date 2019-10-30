package problem_65

import (
	"strings"
	"unicode"
)

//Validate if a given string can be interpreted as a decimal number.
//
//Some examples:
//"0" => true
//" 0.1 " => true
//"abc" => false
//"1 a" => false
//"2e10" => true
//" -90e3   " => true
//" 1e" => false
//"e3" => false
//" 6e-1" => true
//" 99e2.5 " => false
//"53.5e93" => true
//" --6 " => false
//"-+3" => false
//"95a54e53" => false
//
//Note: It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one. However, here is a list of characters that can be in a valid decimal number:
//
//Numbers 0-9
//Exponent - "e"
//Positive/negative sign - "+"/"-"
//Decimal point - "."
//
//Of course, the context of these characters also matters in the input.
//
//Update (2015-02-10):
//The signature of the C++ function had been updated. If you still see your function signature accepts a const char * argument, please click the reload button to reset your code definition.

// skip leading white spaces
// skip +, -, .
// after sign character can only be digit, except following 2 exceptions
// if "." exist, needs to be a digit afterwards
// if "e" exist, needs to be a digit afterwards or end of string
// if space exist, it can only be space to end
func isNumber(s string) bool {
	trimmed := strings.Trim(s, " ")
	if "" == trimmed {
		return false
	}

	for _, s := range trimmed {
		if s == ' ' {
			return false
		}
	}

	var dotExist, exponentExist, digitBeforeExponent, digitAfterExponent, digitExist bool

	runes := []rune(trimmed)
	length := len(runes)
	start := 0

	// skip +, -, .
	if !unicode.IsDigit(runes[start]) {
		if runes[start] != '+' && runes[start] != '-' && runes[start] != '.' {
			return false
		}

		if runes[start] == '.' {
			dotExist = true
		}
		start++
	}

	//	invalid: exponent before dot, multiple e, multiple dot, e without digit, no digits,
	// 			 no digits before exponent, no digits after exponent
	for start < length {
		if unicode.IsDigit(runes[start]) {
			digitExist = true
			if exponentExist && !digitAfterExponent {
				digitAfterExponent = true
			}
			start++
			continue
		}

		if runes[start] == '.' {
			if exponentExist || dotExist {
				return false
			}

			dotExist = true
			start++
			continue
		}

		if runes[start] == 'e' {
			if exponentExist {
				return false
			}

			if digitExist {
				digitBeforeExponent = true
			}

			exponentExist = true

			start++
			continue
		}

		// 2.3e+6
		if runes[start] == '+' || runes[start] == '-' {
			if start >= 1 && runes[start-1] == 'e' {
				start++
				continue
			}
			return false
		}

		if !unicode.IsDigit(runes[start]) {
			return false
		}
	}

	if !digitExist {
		return false
	}

	if exponentExist && !digitAfterExponent {
		return false
	}

	if exponentExist && !digitBeforeExponent {
		return false
	}

	return true
}
