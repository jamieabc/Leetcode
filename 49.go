package main

// Given an array of strings, group anagrams together.
//
// Example:
//
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
//
// Note:
//
//     All inputs will be in lowercase.
//     The order of your output does not matter.

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	signatures := make(map[string][]string)

	var j int
	tmp := make([]int, 26)
	for i := range strs {
		reset(&tmp)

		for j = range strs[i] {
			tmp[strs[i][j]-'a']++
		}

		key := mapKey(tmp)
		if _, ok := signatures[key]; ok {
			signatures[key] = append(signatures[key], strs[i])
		} else {
			signatures[key] = []string{strs[i]}
		}
	}

	for _, v := range signatures {
		result = append(result, v)
	}

	return result
}

func reset(tmp *[]int) {
	for i := range *tmp {
		(*tmp)[i] = 0
	}
}

// tea -> a1e1t1
func mapKey(data []int) string {
	var b []byte

	for i := range data {
		if data[i] > 0 {
			b = append(b, byte('a'+i))
			b = append(b, byte(data[i])-byte(0))
		}
	}

	return string(b)
}

//	problems
//	1.	too slow, remove some memory allocation

//	2.	still too slow, I think problems comes from too many groups to search,
//		I can use map for faster find, with custom key strucutre tea -> a1e1t1

//	3.	when grouping, it's a hidden criteria that dedup is needed

//	4.	use map for faster finding, but be aware that to convert char # into
//		byte, it needs to deduct byte(0), otherwise it might cause duplicates

//	5.	from reference https://leetcode.com/problems/group-anagrams/discuss/19176/Share-my-short-JAVA-solution

//		it's also possible to convert those char bucket into string
