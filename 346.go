package main

type MovingAverage struct {
	nums  []int
	sum   int
	count int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		nums:  make([]int, size),
		sum:   0,
		count: 0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.add(val)
	size := len(this.nums)
	if this.count < len(this.nums) {
		size = this.count
	}

	return float64(this.sum) / float64(size)
}

func (this *MovingAverage) add(val int) {
	idx := this.count % len(this.nums)
	this.sum += -this.nums[idx] + val // default 0, so this line is okay
	this.nums[idx] = val
	this.count++
}

type MovingAverage1 struct {
	sum  int
	nums []int
	size int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage1 {
	return MovingAverage1{
		sum:  0,
		nums: make([]int, 0),
		size: size,
	}
}

func (this *MovingAverage1) Next(val int) float64 {
	if len(this.nums) < this.size {
		this.nums = append(this.nums, val)
		this.sum += val
	} else {
		this.sum -= this.nums[0]
		this.sum += val
		this.nums = this.nums[1:]
		this.nums = append(this.nums, val)
	}

	if this.sum == 0 {
		return 0
	}

	return float64(this.sum) / float64(len(this.nums))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

//	problems
//	1.	this number only access first number, linked-list maybe more suitable
//		than array

//	2.	inspired from sample code, array with circular behavior is better
//		also, author adds another method "add" for abstraction
