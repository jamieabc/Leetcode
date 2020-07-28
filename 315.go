package main

// You are given an integer array nums and you have to return a new counts array. The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].
//
// Example:
//
// Input: [5,2,6,1]
// Output: [2,1,1,0]
// Explanation:
// To the right of 5 there are 2 smaller elements (2 and 1).
// To the right of 2 there is only 1 smaller element (1).
// To the right of 6 there is 1 smaller element (1).
// To the right of 1 there is 0 smaller element.
func countSmaller(nums []int) []int {
	size := len(nums)

	if size == 0 {
		return []int{}
	}

	sorting := make([]int, size)
	for i := range sorting {
		sorting[i] = i
	}

	result := make([]int, size)
	mergeSort(nums, sorting, result, 0, size-1)

	return result
}

func mergeSort(nums, sorting, result []int, start, end int) {
	if end == start {
		return
	}

	mid := start + (end-start)/2
	mergeSort(nums, sorting, result, start, mid)
	mergeSort(nums, sorting, result, mid+1, end)
	merge(nums, sorting, result, start, mid, end)
}

func merge(nums, sorting, result []int, start, mid, end int) {
	tmp := make([]int, end-start+1)
	var inversion int

	for idx, i, j := 0, start, mid+1; i <= mid || j <= end; idx++ {
		if i <= mid && j <= end {
			if nums[sorting[i]] <= nums[sorting[j]] {
				tmp[idx] = sorting[i]
				result[sorting[i]] += inversion
				i++
			} else {
				tmp[idx] = sorting[j]
				j++
				inversion++
			}
		} else if i <= mid {
			tmp[idx] = sorting[i]
			result[sorting[i]] += inversion
			i++
		} else {
			tmp[idx] = sorting[j]
			inversion++
			j++
		}
	}

	for i := range tmp {
		sorting[start+i] = tmp[i]
	}
}

func countSmaller1(nums []int) []int {
	result := make([]int, len(nums))

	// count smaller before is just sorted array index
	sorting := make([]int, 0)
	for i := range nums {
		if len(sorting) == 0 {
			sorting = append(sorting, nums[i])
			result[i] = 0
		} else {
			low, high := 0, len(sorting)-1
			for low <= high {
				mid := low + (high-low)/2

				if sorting[mid] == nums[i] {
					low = mid
					break
				}

				if sorting[mid] > nums[i] {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}

			tmp := append([]int{}, sorting[:low]...)
			tmp = append(tmp, nums[i])
			tmp = append(tmp, sorting[low:]...)
			sorting = tmp
			result[i] = low
		}
	}

	// sort every number one by one, find smaller number before (index)
	// with every number sorted, final index - in sort index

	indexes := make(map[int]int)
	for i := range sorting {
		if _, exist := indexes[sorting[i]]; !exist {
			indexes[sorting[i]] = i
		}
	}

	for i := range result {
		if i > 0 && nums[i] == nums[i-1] {
			result[i] = result[i-1]
		} else {
			result[i] = indexes[nums[i]] - result[i]
		}
	}

	return result
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	problems
//	1.	need to careful check insert position

//	2.	when target number equal middle of sorting array, before break loop,
//		need to update index to mid

//	3.	inspired from https://leetcode.com/problems/count-of-smaller-numbers-after-self/discuss/445769/merge-sort-CLEAR-simple-EXPLANATION-with-EXAMPLES-O(n-lg-n)

//		count of number smaller can be found during merge sort, since every time
//		merge left & right subarray, any number chosen from right means a number
//		"crossing", and "crossing" means the number is smaller than number in
// 		right array, which is problem wants

//		another dp technique is to store number that has inversion, since left
//		& right subarray is already sorted, means if any number is from right,
//		all remain left number shares this inversion.
