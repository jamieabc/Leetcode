package main

// Given a string, sort it in decreasing order based on the frequency of characters.
//
// Example 1:
//
// Input:
// "tree"
//
// Output:
// "eert"
//
// Explanation:
// 'e' appears twice while 'r' and 't' both appear once.
// So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
//
// Example 2:
//
// Input:
// "cccaaa"
//
// Output:
// "cccaaa"
//
// Explanation:
// Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
// Note that "cacaca" is incorrect, as the same characters must be together.
//
// Example 3:
//
// Input:
// "Aabb"
//
// Output:
// "bbAa"
//
// Explanation:
// "bbaA" is also a valid answer, but "Aabb" is incorrect.
// Note that 'A' and 'a' are treated as two different characters.

type info struct {
	str   byte
	count int
}

func frequencySort(s string) string {
	freq := make([]info, 256)

	// calculate char frequency
	for i := range s {
		freq[int(s[i])].count++

		if freq[int(s[i])].str == 0 {
			freq[int(s[i])].str = s[i]
		}
	}

	// bucket sort
	bucket := make([][]byte, len(s)+1)
	for _, i := range freq {
		if i.count > 0 {
			if len(bucket[i.count]) == 0 {
				bucket[i.count] = []byte{i.str}
			} else {
				bucket[i.count] = append(bucket[i.count], i.str)
			}
		}
	}

	var result []byte

	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) > 0 {
			for j := range bucket[i] {
				for k := 0; k < i; k++ {
					result = append(result, bucket[i][j])
				}
			}
		}
	}
	return string(result)
}

//	problems
//	1.	a & A are both valid

//	2.	with numbers and spaces

//	3. from sample code, use byte array

//	4.	from reference https://leetcode.com/problems/sort-characters-by-frequency/discuss/93519/Python-O(N)-solution-using-Hash-Map.

//		first hash map is char -> freq
//		second hash map is freq -> char

//		similar to this, I can create a array w/ size equal to len(s), put
//		character c into index f if frequency equals (bucket sort)

//	5.	from reference https://leetcode.com/problems/sort-characters-by-frequency/discuss/93404/C%2B%2B-O(n)-solution-without-sort()

//		when string is long, a lot of memory will used. one way is to use
//		map freq -> char, and sort map keys, then iterate through slice

//		go's map is not guaranteed to have same order every time, so need
//		to sort map keys self (https://stackoverflow.com/a/23332089)
