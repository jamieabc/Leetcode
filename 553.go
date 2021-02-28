package main

import (
	"strconv"
	"strings"
)

// Given a list of positive integers, the adjacent integers will perform the float division. For example, [2,3,4] -> 2 / 3 / 4.
//
// However, you can add any number of parenthesis at any position to change the priority of operations. You should find out how to add parenthesis to get the maximum result, and return the corresponding expression in string format. Your expression should NOT contain redundant parenthesis.
//
// Example:
//
// Input: [1000,100,10,2]
// Output: "1000/(100/10/2)"
// Explanation:
// 1000/(100/10/2) = 1000/((100/10)/2) = 200
// However, the bold parenthesis in "1000/((100/10)/2)" are redundant,
// since they don't influence the operation priority. So you should return "1000/(100/10/2)".
//
// Other cases:
// 1000/(100/10)/2 = 50
// 1000/(100/(10/2)) = 50
// 1000/100/10/2 = 0.5
// 1000/100/(10/2) = 2
//
// Note:
//
//     The length of the input array is [1, 10].
//     Elements in the given array will be in range [2, 1000].
//     There is only one optimal division for each test case.

func optimalDivision(nums []int) string {
	ans := make([]byte, 0)

	tmp := make([]byte, 0)
	for n := nums[0]; n > 0; n /= 10 {
		tmp = append(tmp, byte((n%10)+'0'))
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		ans = append(ans, tmp[i])
	}

	size := len(nums)
	if size == 1 {
		return string(ans)
	}

	ans = append(ans, byte('/'))
	if size > 2 {
		ans = append(ans, byte('('))
	}

	for i := 1; i < size; i++ {
		if i > 1 {
			ans = append(ans, byte('/'))
		}

		tmp := make([]byte, 0)
		for j := nums[i]; j > 0; j /= 10 {
			tmp = append(tmp, byte((j%10)+'0'))
		}
		for j := len(tmp) - 1; j >= 0; j-- {
			ans = append(ans, tmp[j])
		}
	}

	if size > 2 {
		ans = append(ans, byte(')'))
	}

	return string(ans)
}

func optimalDivision1(nums []int) string {
	length := len(nums)
	if length == 1 {
		return strconv.Itoa(nums[0])
	}

	var sb strings.Builder
	sb.WriteString(strconv.Itoa(nums[0]))

	if length == 2 {
		sb.WriteString("/")
		sb.WriteString(strconv.Itoa(nums[1]))
		return sb.String()
	} else {
		sb.WriteString("/(")
		sb.WriteString(strconv.Itoa(nums[1]))
	}

	for i := 2; i < length; i++ {
		sb.WriteString("/")
		sb.WriteString(strconv.Itoa(nums[i]))
	}

	sb.WriteString(")")
	return sb.String()
}

//  Notes
//  1.  add reference https://leetcode.com/problems/optimal-division/discuss/101687/Easy-to-understand-simple-O(n)-solution-with-explanation

//      this problem has many dislike, so I just refer to others for answer
