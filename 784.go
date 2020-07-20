package main

// Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.  Return a list of all possible strings we could create.
//
// Examples:
// Input: S = "a1b2"
// Output: ["a1b2", "a1B2", "A1b2", "A1B2"]
//
// Input: S = "3z4"
// Output: ["3z4", "3Z4"]
//
// Input: S = "12345"
// Output: ["12345"]
//
// Note:
//
//     S will be a string with length between 1 and 12.
//     S will consist only of letters or digits.

func letterCasePermutation(S string) []string {
	result := make([]string, 0)
	flags := make([]bool, len(S))
	for i := range S {
		if S[i] >= '0' && S[i] <= '9' {
			flags[i] = true
		}
	}

	recursive(S, flags, 0, &result)

	return result
}

func recursive(s string, flags []bool, start int, result *[]string) {
	*result = append(*result, s)

	for i := start; i < len(s); i++ {
		if flags[i] {
			continue
		}
		recursive(flip(s, i), flags, i+1, result)
	}
}

func flip(s string, idx int) string {
	newStr := make([]byte, len(s))
	for i := range s {
		if i != idx {
			newStr[i] = s[i]
		} else {
			if s[idx] >= 'a' && s[idx] <= 'z' {
				newStr[i] = byte(int('A') + int(s[idx]-'a'))
			} else {
				newStr[i] = byte(int('a') + int(s[idx]-'A'))
			}
		}
	}

	return string(newStr)
}
