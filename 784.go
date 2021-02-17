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
	ans := make([]string, 0)

	dfs(S, "", 0, &ans)

	return ans
}

func dfs(S, current string, idx int, ans *[]string) {
	if idx == len(S) {
		*ans = append(*ans, current)
		return
	}

	dfs(S, current+string(S[idx]), idx+1, ans)

	if S[idx] >= 'a' && S[idx] <= 'z' {
		newChar := string('A' + S[idx] - 'a')
		dfs(S, current+newChar, idx+1, ans)
	} else if S[idx] >= 'A' && S[idx] <= 'Z' {
		newChar := string('a' + S[idx] - 'A')
		dfs(S, current+newChar, idx+1, ans)
	}
}

type data struct {
	str []byte
	idx int
}

func letterCasePermutation3(S string) []string {
	queue := []data{{str: []byte(S), idx: 0}}

	ans := make([]string, 0)

	for len(queue) > 0 {
		d := queue[0]
		queue = queue[1:]

		if d.idx == len(S) {
			ans = append(ans, string(d.str))
		} else {
			if c := d.str[d.idx]; c >= 'a' && c <= 'z' {
				tmp := make([]byte, len(S))
				copy(tmp, d.str)

				newData := data{
					str: tmp,
					idx: d.idx + 1,
				}

				newData.str[d.idx] = d.str[d.idx] - 'a' + 'A'
				queue = append(queue, newData)
			} else if c >= 'A' && c <= 'Z' {
				tmp := make([]byte, len(S))
				copy(tmp, d.str)

				newData := data{
					str: tmp,
					idx: d.idx + 1,
				}

				newData.str[d.idx] = d.str[d.idx] - 'A' + 'a'
				queue = append(queue, newData)
			}

			d.idx++
			queue = append(queue, d)
		}
	}

	return ans
}

func letterCasePermutation2(S string) []string {
	ans := make([]string, 0)

	traverse(S, []byte{}, 0, &ans)

	return ans
}

func traverse(S string, cur []byte, idx int, ans *[]string) {
	if len(S) == len(cur) {
		*ans = append(*ans, string(cur))
		return
	}

	for i := idx; i < len(S); i++ {
		if S[i] >= 'a' && S[i] <= 'z' {
			tmp := make([]byte, len(cur)+1)
			copy(tmp, cur)
			tmp[len(tmp)-1] = S[i] - 'a' + 'A'

			traverse(S, tmp, i+1, ans)

			tmp[len(tmp)-1] = S[i]
			traverse(S, tmp, i+1, ans)
		} else if S[i] >= 'A' && S[i] <= 'Z' {
			tmp := make([]byte, len(cur)+1)
			copy(tmp, cur)
			tmp[len(tmp)-1] = S[i] - 'A' + 'a'

			traverse(S, tmp, i+1, ans)

			tmp[len(tmp)-1] = S[i]
			traverse(S, tmp, i+1, ans)
		} else {
			cur = append(cur, S[i])
		}
	}

	// in case last char is digit
	if len(cur) == len(S) {
		*ans = append(*ans, string(cur))
	}
}

func letterCasePermutation1(S string) []string {
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

//	Notes
//	1.	inspired from sample code, instead of copying all string, use "" + ""
//		to append string

//	2.	by the way, it's not backtracking
