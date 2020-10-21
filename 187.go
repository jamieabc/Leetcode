package main

// All DNA is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T', for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.
//
// Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.
//
//
//
// Example 1:
//
// Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
// Output: ["AAAAACCCCC","CCCCCAAAAA"]
// Example 2:
//
// Input: s = "AAAAAAAAAAAAA"
// Output: ["AAAAAAAAAA"]
//
//
// Constraints:
//
// 0 <= s.length <= 105
// s[i] is 'A', 'C', 'G', or 'T'.

func findRepeatedDnaSequences(s string) []string {
	L := 10
	size := len(s)
	if size < L {
		return []string{}
	}

	var num int
	visited := make(map[int]int)
	for i := 0; i < 10; i++ {
		num <<= 2
		num += toNum(s[i])
	}
	visited[num]++

	leftWeight := 1 << ((L - 1) * 2)
	ans := make([]string, 0)
	for i := 10; i < size; i++ {
		num -= leftWeight * toNum(s[i-10])
		num <<= 2
		num += toNum(s[i])

		visited[num]++
		if visited[num] == 2 {
			ans = append(ans, s[i-9:i+1])
		}
	}

	return ans
}

func toNum(b byte) int {
	switch b {
	case 'A':
		return 0
	case 'C':
		return 1
	case 'T':
		return 2
	default:
		return 3
	}
}

func findRepeatedDnaSequences1(s string) []string {
	seq := make(map[string]int)
	size := 10

	for i := 0; i <= len(s)-size; i++ {
		seq[s[i:i+size]]++
	}

	ans := make([]string, 0)

	for str, count := range seq {
		if count >= 2 {
			ans = append(ans, str)
		}
	}

	return ans
}

//	Notes
//	1.	if S & L both very large, it's putting every possible strings into memory.
//		encode L into number to reduce used memory space
