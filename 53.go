package main

//Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
//
//Example:
//
//Input: [-2,1,-3,4,-1,2,1,-5,4],
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
//
//Follow up:
//
//If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	//return dp(nums)

	max := -99999
	divide(nums, &max)

	return max
}

func dp(nums []int) int {
	max := nums[0]
	maxToCurrent := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+prev > nums[i] {
			maxToCurrent = nums[i] + prev
		} else {
			maxToCurrent = nums[i]
		}
		prev = maxToCurrent

		if maxToCurrent > max {
			max = maxToCurrent
		}
	}

	return max
}

func divide(nums []int, tmpMax *int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	var l, r int
	if length == 2 {
		l = nums[0]
		r = nums[1]
	} else {
		mid := (length - 1) / 2
		l = divide(nums[:mid], tmpMax)
		r = divide(nums[mid:], tmpMax)
	}

	return conquer(nums, l, r, tmpMax)
}

func conquer(nums []int, l, r int, tmpMax *int) int {
	length := len(nums)
	if length <= 2 {
		tmp := max(l+r, max(l, r))
		if tmp > *tmpMax {
			*tmpMax = tmp
		}
		return tmp
	}

	mid := (length - 1) / 2

	// crossing
	tmpL := 0
	sl := 0
	for i := mid - 1; i >= 0; i-- {
		tmpL += nums[i]
		if tmpL > sl {
			sl = tmpL
		}
	}

	tmpR := 0
	sr := 0
	for i := mid + 1; i < length; i++ {
		tmpR += nums[i]
		if tmpR > sr {
			sr = tmpR
		}
	}

	tmp := max(sl+nums[mid]+sr, max(l, r))
	if tmp > *tmpMax {
		*tmpMax = tmp
	}
	return tmp
}

// problems
// 1. reduce memory usage, arr is not necessary
// 2. when dividing array, there's a situation of infinite loop when length equals 2
// 3. wrong conquer method, only way to keep summing is add every elements
// 4. iteration keeps going only if array length >= 2
// 5. since use length - 1 to determine crossing index, it could be only left-crossing along as comparison
// 6. crossing could be left, right, left+right
// 7. I had a fundamental thinking error, divide and conquer should separate into left, right, and return max(left, right); for crossing, do another calculation
// 8. wrong increment of i
// 9. wrong calculation of crossing, because it should be mid, mid-left, mid-right
// 10. didn't consider the situation when length is 2, mid is 0, mid-1 is negative number which is invalid
// 11. crossing use mid to (length-1)/2, divide use (length+1)/2, they are different and cause error
// 12. process sub-result & merge result should be in conquer method
// 13. what I think about crossing is wrong, not necessary needs to include at least one of left or right
