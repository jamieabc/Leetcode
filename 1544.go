package main

//Given a string s of lower and upper case English letters.
//
//A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:
//
//    0 <= i <= s.length - 2
//    s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice-versa.
//
//To make the string good, you can choose two adjacent characters that make the string bad and remove them. You can keep doing this until the string becomes good.
//
//Return the string after making it good. The answer is guaranteed to be unique under the given constraints.
//
//Notice that an empty string is also good.
//
//
//
//Example 1:
//
//Input: s = "leEeetcode"
//Output: "leetcode"
//Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".
//
//Example 2:
//
//Input: s = "abBAcC"
//Output: ""
//Explanation: We have many possible scenarios, and all lead to the same answer. For example:
//"abBAcC" --> "aAcC" --> "cC" --> ""
//"abBAcC" --> "abBA" --> "aA" --> ""
//
//Example 3:
//
//Input: s = "s"
//Output: "s"
//
//
//
//Constraints:
//
//    1 <= s.length <= 100
//    s contains only lower and upper case English letters.

func makeGood(s string) string {
	if s == "" {
		return s
	}

	stack := []byte{s[0]}

	for i := 1; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			if isBad(stack[len(stack)-1], s[i]) {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		}
	}

	return string(stack)
}

func makeGood1(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if isBad(s[i], s[i+1]) {
			s = s[:i] + s[i+2:]
			i = -1
		}
	}

	return s
}

func isBad(s1, s2 byte) bool {
	if s1 >= 'a' && s1 <= 'z' {
		return s2 == s1-'a'+'A'
	}
	return s2 == s1-'A'+'a'
}

//	problems
//	1.	using recursive to check from start is not efficient

//		tc: O(n^2)

//	2.	in fact, for char already checked, no need to do again, just
//		check most recent one => stack
