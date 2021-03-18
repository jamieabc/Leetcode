package main

// A word's generalized abbreviation can be constructed by taking any number of non-overlapping substrings and replacing them with their respective lengths. For example, "abcde" can be abbreviated into "a3e" ("bcd" turned into "3"), "1bcd1" ("a" and "e" both turned into "1"), and "23" ("ab" turned into "2" and "cde" turned into "3").
//
// Given a string word, return a list of all the possible generalized abbreviations of word. Return the answer in any order.
//
//
//
// Example 1:
//
// Input: word = "word"
// Output: ["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]
//
// Example 2:
//
// Input: word = "a"
// Output: ["1","a"]
//
//
//
// Constraints:
//
//     1 <= word.length <= 15
//     word consists of only lowercase English letters.

type Data struct {
	Cur        []byte
	Idx, Count int
}

func generateAbbreviations(word string) []string {
	ans := make([]string, 0)
	stack := []Data{{make([]byte, 0), 0, 0}}
	size := len(word)

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if item.Idx >= size {
			if item.Count > 0 {
				if item.Count >= 10 {
					item.Cur = append(item.Cur, '1')
					item.Count -= 10
				}
				item.Cur = append(item.Cur, byte('0'+item.Count))
			}
			ans = append(ans, string(item.Cur))
			continue
		}

		stack = append(stack, Data{
			Cur:   item.Cur,
			Idx:   item.Idx + 1,
			Count: item.Count + 1,
		})

		if item.Count > 0 {
			if item.Count >= 10 {
				item.Cur = append(item.Cur, '1')
				item.Count -= 10
			}
			item.Cur = append(item.Cur, byte('0'+item.Count))
		}

		item.Cur = append(item.Cur, word[item.Idx])
		item.Idx++
		item.Count = 0
		stack = append(stack, item)
	}

	return ans
}

func generateAbbreviations1(word string) []string {
	ans := make([]string, 0)

	dfs(word, []byte{}, 0, 0, &ans)

	return ans
}

func dfs(word string, cur []byte, idx, count int, ans *[]string) {
	size := len(word)

	if idx >= size {
		if count > 0 {
			if count >= 10 {
				cur = append(cur, '1')
				count -= 10
			}
			cur = append(cur, byte('0'+count))
		}

		*ans = append(*ans, string(cur))
		return
	}

	dfs(word, cur, idx+1, count+1, ans)

	tmp := append([]byte{}, cur...)
	if count > 0 {
		if count >= 10 {
			tmp = append(tmp, '1')
			count -= 10
		}
		tmp = append(tmp, byte('0'+count))
	}

	tmp = append(tmp, word[idx])
	dfs(word, tmp, idx+1, 0, ans)
}

func isChar(b byte) bool {
	return b >= 'a' && b <= 'z'
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/generalized-abbreviation/discuss/77256/Could-anyone-help-to-understand-this-question

//		the reason abcde cannot abbreviated into 23 is because don't know if
//		there are 23 chars or 2+3 chars

//	2.	use iterative, still TLE

//	3.	inspired from https://leetcode.com/problems/generalized-abbreviation/discuss/77193/Simple-Python-Solution-with-Explanation

//		no need to add number in advance, can use a variable to store, clear
//		this value when adding char to it

//		with this technique, time improved from TLE to pass
