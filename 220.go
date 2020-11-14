package main

import (
	"fmt"
	"sort"
)

//Given an array of integers, find out whether there are two distinct indices i and j in the array such that the absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.
//
//
//
//Example 1:
//
//Input: nums = [1,2,3,1], k = 3, t = 0
//Output: true
//
//Example 2:
//
//Input: nums = [1,0,1,1], k = 1, t = 2
//Output: true
//
//Example 3:
//
//Input: nums = [1,5,9,1,5,9], k = 2, t = 3
//Output: false
//
//
//
//Constraints:
//
//    0 <= nums.length <= 2 * 104
//    -231 <= nums[i] <= 231 - 1
//    0 <= k <= 104
//    0 <= t <= 231 - 1

func main() {
	nums := []int{1, 5, 9, 1, 5, 9}
	fmt.Println(containsNearbyAlmostDuplicate(nums, 1, 2))
}

type Sorted []int

func (s *Sorted) Insert(val int) int {
	idx := sort.SearchInts(*s, val)
	*s = append(*s, 0)
	copy((*s)[idx+1:], (*s)[idx:])
	(*s)[idx] = val

	return idx
}

func (s *Sorted) Remove(val int) {
	idx := sort.SearchInts(*s, val)
	copy((*s)[idx:], (*s)[idx+1:])
	*s = (*s)[:len(*s)-1]
}

// tc: O(n log(n))
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k == 0 {
		return false
	}

	s := &Sorted{}

	for i := range nums {
		idx := s.Insert(nums[i])

		if (idx > 0 && abs((*s)[idx-1]-nums[i]) <= t) || (idx < len(*s)-1 &&
			abs((*s)[idx+1]-nums[i]) <= t) {
			return true
		}

		if i >= k {
			s.Remove(nums[i-k])
		}
	}

	return false
}

// tc: O(n)
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	buckets := make(map[int]int)

	for i, num := range nums {
		// remove number from bucket, only if it's expected (because there
		// might be duplicate numbers, so only number with exactly same)
		if i > k {
			// if previous number exist numbers within same bucket, program
			// already returns, so it means no numbers in same bucket, which
			// is safe to remove this bucket
			delete(buckets, bucket(nums[i-k-1], t))
		}

		bkt := bucket(nums[i], t)

		// already number in same bucket, this number difference must <= t
		if _, ok := buckets[bkt]; ok {
			return true
		}

		// check if adjacent buckets exist
		if val, ok := buckets[bkt-1]; ok && num-val <= t {
			return true
		}

		if val, ok := buckets[bkt+1]; ok && val-num <= t {
			return true
		}

		// update bucket to latest index, because if there exist multiple
		// numbers in same bucket, it should already return true in previous
		// conditions
		buckets[bkt] = num
	}

	return false
}

// -1 / t = 0, 1 / t = 0, but these 2 number should in different buckets
func bucket(num, t int) int {
	if t == 0 || num == 0 {
		return num
	}

	if num > 0 {
		return num / (t + 1)
	}
	return num/t - 1
}

// tc: O(n log n + nkm), m: # of same number in array
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	// if k == 0, only one number exist
	if k == 0 {
		return false
	}

	pos := make(map[int][]int)
	for i := range nums {
		pos[nums[i]] = append(pos[nums[i]], i)
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			for i++; i < len(nums) && nums[i] == nums[i-1]; i++ {
			}
			i--
			continue
		}

		for j := i + 1; j < len(nums) && nums[j]-nums[i] <= t; j++ {
			if nums[i] == nums[j] {
				for m := 1; m < len(pos[nums[j]]); m++ {
					if pos[nums[j]][m]-pos[nums[j]][m-1] <= k {
						return true
					}
				}

				for j++; j < len(nums) && nums[j] == nums[i]; j++ {
				}
				j--
			} else {
				// both indexes are sorted, so it's linear operation to find
				// closest distance
				for m, n := 0, 0; m < len(pos[nums[i]]) && n < len(pos[nums[j]]); {
					if abs(pos[nums[i]][m]-pos[nums[j]][n]) <= k {
						return true
					} else if pos[nums[i]][m] < pos[nums[j]][n] {
						m++
					} else {
						n++
					}
				}
			}
		}
	}

	return false
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

//	Notes
//	1.	compare k next numbers results in time complexity O(nk), if k == n,
//		then it ends up to O(n^2)

//		to reduce computation, there are 2 sub-problems to solved:
//		- does number within abs difference exist?
//		- if this number exist, where is it?

//		first problem can be determined by sort the array, check if sorted
//		array next number difference within t

//		second problem can be determined by a hashmap to store same number
//		occurrence index

//		overall tc becomes O(n log n + n + nmp), m: # of same number indexes,
//		p: # of numbers abs difference <= t

//		this method is faster because it reduces some un-necessary
//		computations, after sorted a number doesn't have close number should
//		not be considered, but in naive solution, every number will be
//		checked

//	2.	when finding index of specific number, there are 2 conditions:
//		- numbers are same
//		- numbers are different

//		conditions are separate because in same number situation, also need
//		to check if indexes are same because iteration on same array

//	3.	inspired from solution, instead to store all indexes of every number,
//		another way is to group numbers as buckets, and check if adjacent
//		bucket with same number

//		the idea of bucket sort is to group number with same properties,
//		suck as group lower case letters into array, etc. In this problem,
//		groups can be separated by t
//		e.g.   0 ~  t  => bucket 0
//			 t+1 ~ 2t  => bucket 1
//			2t+1 ~ 3t  => bucket 2

//		also, be careful that 0 <= t means t could be 0, number cannot
//		divided by 0

//		so for any number, as long as its bucket is decided, is pretty fast
//		to check if adjacent buckets (previous & next buckets) contains
//		proper numbers

//		the other thing is, bucket can store latest number's index, because
//		if there exists same number in a bucket, program should already
//		return true

//		value is removed when number's index > k, e.g. when index = k+1,
//		remove index 0 number from bucket

//		since number could be pretty big, use array to store buckets may
//		needs large space, use hashmap to store buckets is more appropriate

//	4.	be careful about overflow, because use t+1 to divide bucket, and
//		0 <= t <= 2^31 - 1, so t+1 might cause overflow

//		also, -1 / t = 0, 1 / t = 0, but these 2 should in different buckets

//	5.	an observation: for a number x, find if any number value difference
//		no more than t, x-t ~ x+t, one part is about smaller, the other part
//		is about larger. this clue implies binary search, which separate
//		numbers into 2 parts: smaller & larger

//	6.	for BST, in order to make sure distance t value exist, needs to find
//		next larger successor & previous smaller successor, and check if these
//		values with in distant t

//	7.	BST and array are basically same, inspired from sampe code, use array
//		to store sorted numbers
