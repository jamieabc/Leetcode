package main

// Given a non-empty string s and an abbreviation abbr, return whether the string matches with the given abbreviation.
//
// A string such as "word" contains only the following valid abbreviations:
//
// ["word", "1ord", "w1rd", "wo1d", "wor1", "2rd", "w2d", "wo2", "1o1d", "1or1", "w1r1", "1o2", "2r1", "3d", "w3", "4"]
//
// Notice that only the above abbreviations are valid abbreviations of the string "word". Any other string is not a valid abbreviation of "word".
//
// Note:
// Assume s contains only lowercase letters and abbr contains only lowercase letters and digits.
//
// Example 1:
//
// Given s = "internationalization", abbr = "i12iz4n":
//
// Return true.
//
// Example 2:
//
// Given s = "apple", abbr = "a2e":
//
// Return false.

func validWordAbbreviation(word string, abbr string) bool {
	// only number
	// number + char
	// all char

	var count, idx, i int
	var b byte
	for i = 0; i < len(word); {
		b, count, idx = nextCountAndChar(abbr, idx)

		// find nothing but words has not traversed finish
		if b == 0 && count == 0 {
			return false
		}

		// difference byte
		if b != 0 {
			if b != word[i] {
				return false
			} else {
				i++
			}
		}

		// count
		if count > 0 {
			i += count
		}
	}

	return i == len(word) && idx == len(abbr)
}

func nextCountAndChar(abbr string, start int) (byte, int, int) {
	var count int
	idx := len(abbr)
	var b byte

	// already reach end of string
	if start == len(abbr) {
		return b, count, idx
	}

	// char
	if abbr[start]-'0' > 9 {
		b = abbr[start]
		start++
	}

	if start < len(abbr) && abbr[start]-'0' == 0 {
		return 0, 0, idx
	}

	for i := start; i < len(abbr); i++ {
		if abbr[i]-'0' <= 9 {
			count *= 10
			count += int(abbr[i] - '0')
		} else {
			idx = i
			break
		}
	}

	return b, count, idx
}

//	problems
//	1.	number should exactly match word, e.g. word = "ab", abbr = "3" should be wrong
//	2.	should not exist 0 at beginning of number
//	3.	when indexing in string, cannot just reference, need to check first
//	4.	forget to check if abbr is fully traversed
