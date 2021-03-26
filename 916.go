package main

// We are given two arrays A and B of words.  Each word is a string of lowercase letters.
//
// Now, say that word b is a subset of word a if every letter in b occurs in a, including multiplicity.  For example, "wrr" is a subset of "warrior", but is not a subset of "world".
//
// Now say a word a from A is universal if for every b in B, b is a subset of a.
//
// Return a list of all universal words in A.  You can return the words in any order.
//
//
//
// Example 1:
//
// Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","o"]
// Output: ["facebook","google","leetcode"]
//
// Example 2:
//
// Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["l","e"]
// Output: ["apple","google","leetcode"]
//
// Example 3:
//
// Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","oo"]
// Output: ["facebook","google"]
//
// Example 4:
//
// Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["lo","eo"]
// Output: ["google","leetcode"]
//
// Example 5:
//
// Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["ec","oc","ceo"]
// Output: ["facebook","leetcode"]
//
//
//
// Note:
//
// 1 <= A.length, B.length <= 10000
// 1 <= A[i].length, B[i].length <= 10
// A[i] and B[i] consist only of lowercase letters.
// All words in A[i] are unique: there isn't i != j with A[i] == A[j].

func wordSubsets(A []string, B []string) []string {
	var counter [26]int
	for i := range B {
		var tmp [26]int
		for j := range B[i] {
			tmp[B[i][j]-'a']++
		}

		for j := 0; j < 26; j++ {
			counter[j] = max(counter[j], tmp[j])
		}
	}

	ans := make([]string, 0)

	var j int
	for i := range A {
		var tmp [26]int
		for j = range A[i] {
			tmp[A[i][j]-'a']++
		}

		for j = 0; j < len(tmp); j++ {
			if tmp[j] < counter[j] {
				break
			}
		}

		if j == 26 {
			ans = append(ans, A[i])
		}
	}

	return ans
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func wordSubsets2(A []string, B []string) []string {
	sizeA := len(A)
	counterA := make([][26]int, sizeA)

	for i := range A {
		var counter [26]int
		for j := range A[i] {
			counter[A[i][j]-'a']++
		}
		counterA[i] = counter
	}

	table := make(map[[26]int]bool)
	for i := range B {
		var counter [26]int
		for j := range B[i] {
			counter[B[i][j]-'a']++
		}
		table[counter] = true
	}

	ans := make([]string, 0)

	var j, k int
	for i := range counterA {
		j = 0
		for counter := range table {
			for k = 0; k < 26; k++ {
				if counterA[i][k] < counter[k] {
					break
				}
			}

			if k == 26 {
				j++
			} else {
				break
			}
		}

		if j == len(table) {
			ans = append(ans, A[i])
		}
	}

	return ans
}

func wordSubsets1(A []string, B []string) []string {
	sizeA := len(A)
	counterA := make([][26]int, sizeA)

	for i := range A {
		var counter [26]int
		for j := range A[i] {
			counter[A[i][j]-'a']++
		}
		counterA[i] = counter
	}

	sizeB := len(B)
	counterB := make([][26]int, sizeB)
	for i := range B {
		var counter [26]int
		for j := range B[i] {
			counter[B[i][j]-'a']++
		}
		counterB[i] = counter
	}

	ans := make([]string, 0)

	var j, k int
	for i := range counterA {
		for j = 0; j < sizeB; j++ {
			if len(B[j]) > len(A[i]) {
				break
			}

			for k = 0; k < 26; k++ {
				if counterA[i][k] < counterB[j][k] {
					break
				}
			}

			if k != 26 {
				break
			}
		}

		if j == sizeB {
			ans = append(ans, A[i])
		}
	}

	return ans
}

//	Notes
//	1.	got TLE, but tc O(mn), cannot think of any way to reduce complexity...

//	2.	try to dedup B

//	3.	inspired form sample code, select max char count, because that's the
//		check to pass if meets criteria
