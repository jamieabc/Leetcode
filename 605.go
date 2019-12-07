package main

import "fmt"

//Suppose you have a long flowerbed in which some of the plots are planted and some are not. However, flowers cannot be planted in adjacent plots - they would compete for water and both would die.
//
//Given a flowerbed (represented as an array containing 0 and 1, where 0 means empty and 1 means not empty), and a number n, return if n new flowers can be planted in it without violating the no-adjacent-flowers rule.
//
//Example 1:
//
//Input: flowerbed = [1,0,0,0,1], n = 1
//Output: True
//
//Example 2:
//
//Input: flowerbed = [1,0,0,0,1], n = 2
//Output: False
//
//Note:
//
//    The input array won't violate no-adjacent-flowers rule.
//    The input array size is in the range of [1, 20000].
//    n is a non-negative integer which won't exceed the input array size.

// [0, 0, 0, 0, 0] 3
// [0, 0, 1, 0, 0, 0, 1, 0, 0]
// [1, 0, 0, 0, 0, 1]
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) == 0 && n != 0 {
		return false
	}

	var i int
	length := len(flowerbed)
	leftBoundary := 0
	rightBoundary := length - 1

	leftZeros := 0
	for i = 0; i < length; i++ {
		if flowerbed[i] == 0 {
			leftZeros++
		} else {
			leftBoundary = i
			break
		}
	}

	rightZeros := 0
	for i = length - 1; i >= 0; i-- {
		if flowerbed[i] == 0 {
			rightZeros++
		} else {
			rightBoundary = i
			break
		}
	}

	// check if it's all zero
	if leftZeros == length {
		return (length+1)/2 >= n
	}

	// count in-between numbers
	count := 0
	var j int
	for i = leftBoundary + 1; i <= rightBoundary; {
		if flowerbed[i] == 0 {
			for j = i + 1; j <= rightBoundary; j++ {
				if flowerbed[j] == 1 {
					count += (j - 1 - i + 1 - 1) / 2
					break
				}
			}
			i = j + 1
		} else {
			i++
		}
	}

	return (count + leftZeros/2 + rightZeros/2) >= n
}

func main() {
	fmt.Printf("answer: %t\n", canPlaceFlowers([]int{1, 0, 1, 0, 1, 0, 1}, 1))
}
