package main

//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".
//
//Example 1:
//
//Input: ["flower","flow","flight"]
//Output: "fl"
//
//Example 2:
//
//Input: ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.
//
//Note:
//
//All given inputs are in lowercase letters a-z.

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	if length == 1 {
		return strs[0]
	}

	common := strs[0]
	for i := 1; i < length; i++ {
		common = cmp(strs[i], common)
		if common == "" {
			return ""
		}
	}
	return common
}

func cmp(s1, s2 string) string {
	len1 := len(s1)
	len2 := len(s2)
	var j, min int

	if len1 < len2 {
		min = len1
	} else {
		min = len2
	}

	for j = 0; j < min; j++ {
		if s1[j] != s2[j] {
			break
		}
	}

	if j == 0 {
		return ""
	}

	return s1[:j]
}
