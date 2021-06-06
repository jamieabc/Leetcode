package problem1888

// You are given a binary string s. You are allowed to perform two types of operations on the string in any sequence:

//     Type-1: Remove the character at the start of the string s and append it to the end of the string.
//     Type-2: Pick any character in s and flip its value, i.e., if its value is '0' it becomes '1' and vice-versa.

// Return the minimum number of type-2 operations you need to perform such that s becomes alternating.

// The string is called alternating if no two adjacent characters are equal.

//     For example, the strings "010" and "1010" are alternating, while the string "0100" is not.



// Example 1:

// Input: s = "111000"
// Output: 2
// Explanation: Use the first operation two times to make s = "100011".
// Then, use the second operation on the third and sixth elements to make s = "101010".

// Example 2:

// Input: s = "010"
// Output: 0
// Explanation: The string is already alternating.

// Example 3:

// Input: s = "1110"
// Output: 1
// Explanation: Use the second operation on the second element to make s = "1010".



// Constraints:

//     1 <= s.length <= 105
//     s[i] is either '0' or '1'.

// tc: O(n)
func minFlips(s string) int {
	size := len(s)
	minCount := math.MaxInt32

	s += s
    var flip1, flip2 int

	for i := 0; i < len(s); i++ {
		// even index w/ 1, odd index w/ 0
		if (i & 1 == 0 && s[i] == '1') || (i & 1 == 1 && s[i] == '0') {
			flip1++
		}

		// even index w/ 0, odd index w/ 1
		if (i & 1 == 0 && s[i] == '0') || (i & 1 == 1 && s[i] == '1') {
			flip2++
		}

		if i >= size-1 {
			// out of sliding window, remove previously add flip
			if i >= size {
				if ((i-size) & 1 == 0 && s[i-size] == '1') || ((i-size) & 1 == 1 && s[i-size] == '0') {
					flip1--
				}

				if ((i-size) & 1 == 0 && s[i-size] == '0') || ((i-size) & 1 == 1 && s[i-size] == '1') {
					flip2--
				}
			}
			minCount = min(minCount, min(flip1, flip2))
		}
	}

	return minCount
}

// tc: O(n^2)
func minFlips1(s string) int {
    size := len(s)

    minCount := math.MaxInt32

    for i := 0; i < size; i++ {
        tmp := s[i:]+s[:i]
        minCount = min(minCount, min(count(tmp, 0), count(tmp, 1)))
    }

    return minCount
}

func count(s string, target int) int {
    var total int

    for i := range s {
        if int(s[i] -'0') != target {
            total++
        }
        target = (target+1)%2
    }

    return total
}

func min(i ,j int) int {
    if i <= j {
        return i
    }
    return j
}

//	Notes
//	1.	didn't solve in contest
//
//		I know alternating should be key point to solve the problem, because alternating property
//		doesn't change by rotation, e.g. 0101 rotate to 1010 still holds the property
//
//		the only part to influence flip count is at start/end position, because start position no
//		need to consider previous char, and end position no need to consider successive char,
//		so flip can be reduced
//
//		e.g. 00111 => rotate once such that string becomes 01110, 00111 needs to consider first 00,
//		but 01110 doesn't need to consider start & end 0 because they are separated
//
//		but I don't know how to leverage this property to effectively solve the problem, got TLE
//
//	2.	inspired from https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/discuss/1253874/C%2B%2B-Solution-sliding-window.-O(N)
//
//		author solves by a really clever way, for alternating string, it's either 010101... or 101010...
//		so just generate two strings with double length starting from 0 or 1 and calculate for each
//		position # of flips to be same as target string
//
//		the reason this algorithm works is because, minimum flips comes from directly to the target
//		string
//
//		e.g. s = "1110", target = "0101"
//		it could be rotate 3 times (0111) and 1 flip
//					rotate 7 times (0111) and 1 flip
//					rotate 11 times (0111) and 1 flip, etc
//
//		it's not efficient to try from beginning of s because there are too many possibilities,
//		but it's efficient to start from end, because once target string is fixed and rotation
//		don't care, flips are easy to count
//
//		I remind the problem of https://leetcode.com/problems/maximum-number-of-visible-points/
//		also uses same technique to solve

//	3.	inspired from https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/discuss/1253886/C++-O(N)-time-O(1)-space-Concise-and-easy-to-understand/963759
//		author gives clear explanation of the solution

// 	4.	inspired from https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/discuss/1254148/Sliding-Window
//
//		voturbac with very elegant code

//		interestingly, using this is quite slow, I guess it's because operation of %
