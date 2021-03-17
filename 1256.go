package main

// Given a non-negative integer num, Return its encoding string.
//
// The encoding is done by converting the integer to a string using a secret function that you should deduce from the following table:
//
//
//
// Example 1:
//
// Input: num = 23
// Output: "1000"
//
// Example 2:
//
// Input: num = 107
// Output: "101100"
//
//
//
// Constraints:
//
//     0 <= num <= 10^9

func encode(num int) string {
	var offset, diff int

	for ; offset < 32; offset++ {
		if tmp := (1 << offset) - 1; tmp >= num {
			if tmp != num {
				offset--
			}
			diff = num - (1 << offset) + 1
			break
		}
	}

	ans := make([]byte, offset)

	for i := 0; i < offset; i++ {
		if 1<<(offset-1-i)&diff > 0 {
			ans[i] = '1'
		} else {
			ans[i] = '0'
		}
	}

	return string(ans)
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/encode-number/discuss/430488/JavaC%2B%2BPython-Binary-of-n-%2B-1

//		lee always provides great solution

//	2.	inspired from https://leetcode.com/problems/encode-number/discuss/430502/Intuitive-Explanation-With-Logic-And-Picture

//		author draw a more clear picture, it's same as lee's first idea
