package main

//  Given a list of sorted characters letters containing only lowercase letters, and given a target letter target, find the smallest element in the list that is larger than the given target.
//
// Letters also wrap around. For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.
//
// Examples:
//
// Input:
// letters = ["c", "f", "j"]
// target = "a"
// Output: "c"
//
// Input:
// letters = ["c", "f", "j"]
// target = "c"
// Output: "f"
//
// Input:
// letters = ["c", "f", "j"]
// target = "d"
// Output: "f"
//
// Input:
// letters = ["c", "f", "j"]
// target = "g"
// Output: "j"
//
// Input:
// letters = ["c", "f", "j"]
// target = "j"
// Output: "c"
//
// Input:
// letters = ["c", "f", "j"]
// target = "k"
// Output: "c"
//
// Note:
//
//     letters has a length in range [2, 10000].
//     letters consists of lowercase letters, and contains at least 2 unique letters.
//     target is a lowercase letter.

func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)-1

	if greaterThan(letters[0], target) || letters[high] == target || greaterThan(target, letters[high]) {
		return letters[0]
	}

	for low <= high {
		mid := low + (high-low)/2

		if greaterThan(letters[mid], target) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return letters[low]
}

func greaterThan(from, to byte) bool {
	return from > to
}

//	problems
//	1.	could exist duplicates
