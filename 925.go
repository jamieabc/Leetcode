package main

//Your friend is typing his name into a keyboard.  Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.
//
//You examine the typed characters of the keyboard.  Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.
//
//
//
//Example 1:
//
//Input: name = "alex", typed = "aaleex"
//Output: true
//Explanation: 'a' and 'e' in 'alex' were long pressed.
//Example 2:
//
//Input: name = "saeed", typed = "ssaaedd"
//Output: false
//Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.
//Example 3:
//
//Input: name = "leelee", typed = "lleeelee"
//Output: true
//Example 4:
//
//Input: name = "laiden", typed = "laiden"
//Output: true
//Explanation: It's not necessary to long press any character.
//
//
//Note:
//
//name.length <= 1000
//typed.length <= 1000
//The characters of name and typed are lowercase letters.

func isLongPressedName(name string, typed string) bool {
	len1 := len(name)
	len2 := len(typed)

	if len1 == 0 && len2 == 0 {
		return true
	}

	// typed should never less than original, first character should always be same
	if len2 < len1 || name[0] != typed[0] {
		return false
	}

	var i, j int
	for i, j = 0, 0; i < len1 && j < len2; {
		// same character
		if name[i] == typed[j] {
			i++
			j++
			continue
		}

		// different character, check if typed equals to previous one
		if name[i-1] == typed[j] {
			j++
			continue
		}

		return false
	}

	// in case last character is duplicated
	if i == len1 && j < len2 {
		for j < len2 {
			if name[i-1] == typed[j] {
				j++
			} else {
				return false
			}
		}
	}

	return i == len1 && j == len2
}

// problems
//	1.	checking criteria should include last character duplicate
