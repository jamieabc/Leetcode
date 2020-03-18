package main

//Given a sequence of words, check whether it forms a valid word square.
//
//A sequence of words forms a valid word square if the kth row and column read the exact same string, where 0 ≤ k < max(numRows, numColumns).
//
//Note:
//
//    The number of words given is at least 1 and does not exceed 500.
//    Word length will be at least 1 and does not exceed 500.
//    Each word contains only lowercase English alphabet a-z.
//
//Example 1:
//
//Input:
//[
//  "abcd",
//  "bnrt",
//  "crmy",
//  "dtye"
//]
//
//Output:
//true
//
//Explanation:
//The first row and first column both read "abcd".
//The second row and second column both read "bnrt".
//The third row and third column both read "crmy".
//The fourth row and fourth column both read "dtye".
//
//Therefore, it is a valid word square.
//
//Example 2:
//
//Input:
//[
//  "abcd",
//  "bnrt",
//  "crm",
//  "dt"
//]
//
//Output:
//true
//
//Explanation:
//The first row and first column both read "abcd".
//The second row and second column both read "bnrt".
//The third row and third column both read "crm".
//The fourth row and fourth column both read "dt".
//
//Therefore, it is a valid word square.
//
//Example 3:
//
//Input:
//[
//  "ball",
//  "area",
//  "read",
//  "lady"
//]
//
//Output:
//false
//
//Explanation:
//The third row reads "read" while the third column reads "lead".
//
//Therefore, it is NOT a valid word square.

func validWordSquare(words []string) bool {
	length := len(words)
	if length == 0 {
		return true
	}

	for row := range words {
		for column := range words[row] {
			if column >= length {
				return false
			}

			if row >= len(words[column]) {
				return false
			}

			if words[row][column] != words[column][row] {
				return false
			}
		}
	}
	return true
}

// problems
//	1.	wrong about checking length of a line
//	2.	I complicate the situation, since program starts from row -> column,
//		then words[row][column] is always valid, however, on the other half
//		of equation is not, so only thing I need to check is to make sure
//		- column not exceeds total words length
//		- row not exceeds words[column] length
//		previously, I was thinking in terms of graph and brain got stuck,
//		I thought next line of current line should not be longer,
//		and struggle with the criteria. But if I check carefully from
//		code, it already tells what to check.
