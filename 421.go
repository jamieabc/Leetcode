package main

//Given an integer array nums, return the maximum result of nums[i] XOR nums[j], where 0 ≤ i ≤ j < n.
//
//Follow up: Could you do this in O(n) runtime?
//
//
//
//Example 1:
//
//Input: nums = [3,10,5,25,2,8]
//Output: 28
//Explanation: The maximum result is 5 XOR 25 = 28.
//Example 2:
//
//Input: nums = [0]
//Output: 0
//Example 3:
//
//Input: nums = [2,4]
//Output: 6
//Example 4:
//
//Input: nums = [8,10,2]
//Output: 10
//Example 5:
//
//Input: nums = [14,70,53,83,49,91,36,80,92,51,66,70]
//Output: 127
//
//
//Constraints:
//
//1 <= nums.length <= 2 * 104
//0 <= nums[i] <= 231 - 1

func findMaximumXOR(nums []int) int {
	var result, check, mask int

	for i := 31; i >= 0; i-- {
		mask |= 1 << i

		set := make(map[int]bool)
		for _, n := range nums {
			set[n&mask] = true
		}

		check = result | (1 << i)
		for n := range set {
			// a ^ b = c, a ^ c = b
			// assumes ith bit is set, check if two expected pairs does exist
			if set[n^check] {
				result = check
				break
			}
		}
	}

	return result
}

type Trie struct {
	One, Zero *Trie
	Value     int
}

func add(node *Trie, num int) int {
	var msb int

	for i := 31; i >= 0; i-- {
		one := (1<<i)&num > 0

		if one {
			msb = max(msb, i)
			if node.One == nil {
				node.One = &Trie{}
			}
			node = node.One
		} else {
			if node.Zero == nil {
				node.Zero = &Trie{}
			}
			node = node.Zero
		}
	}
	node.Value = num

	return msb
}

func findMaximumXOR1(nums []int) int {
	// build prefix tree
	root := &Trie{}
	var msb int

	for _, num := range nums {
		msb = max(msb, add(root, num))
	}

	// find max XOR
	var maxXOR int
	var node *Trie

	for _, num := range nums {
		// msb not 1, no need to check
		if (1<<msb)&num == 0 {
			continue
		}

		node = root

		for i := 0; i < 32; i++ {
			one := (1<<(31-i))&num > 0

			// try to find invert bit
			if (one && node.Zero != nil) || (!one && node.One == nil) {
				node = node.Zero
			} else {
				node = node.One
			}
		}

		maxXOR = max(maxXOR, num^node.Value)
	}

	return maxXOR
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	inspired from https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/discuss/91046/How-do-you-even-think-of-getting-an-O(n)-solution-for-this

//		O(n^2) has some waste of comparisons, since for any given number, maximum
//		XOR result only comes from number w/ most inverted bits, the left the
//		better.

//		maximum XOR result comes from left most inverted bit:
//		e.g. 01010101
//           1....... (1)
//			 0....... (2)
//		(1) is better than (2) because (1) has larger msb

//		in this way, O(n^2) can be reduced because some number is no need to
//		compare. As former example, (2) is un-necessary to compare

//		next question is: how to quickly find the number? It becomes a search
//		problem:
//		- does a number exist w/ specific bit 1 or 0
//		- when specific bit is chosen, how to search for next bit

//		prefix tree (trie) can help here, because every number is decomposed
//		by bits for quicker search bit one by one

//	2.	when doing bit operation, it should be 1 << offset

//	3.	inspired from https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/discuss/91049/Java-O(n)-solution-using-bit-manipulation-and-HashMap

//		author provides another way to tackle this problem: scan by bits
//		since maximum XOR result is bounded to 32 bits, the problems can also
//		be seen in each bit perspective

//		for XOR operation, element is interchangeable (pair) a ^ b = c, a ^ c = b

//		assumes this bit is set, and check if numbers pass this check, if passes
//		than this bit is set for maximum result; otherwise, this bit is not set
