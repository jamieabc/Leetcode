package main

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

func twoSum(nums []int, target int) []int {
	mapping := generateMap(nums)

	for _, num := range nums {
		remain := target - num
		if remain == num {
			if 2 <= len(mapping[remain]) {
				return []int{mapping[remain][0], mapping[remain][1]}
			}
			continue
		} else {
			if _, ok := mapping[remain]; ok {
				return []int{mapping[num][0], mapping[remain][0]}
			}
			continue
		}
	}

	return []int{}
}

func generateMap(nums []int) map[int][]int {
	result := make(map[int][]int, 0)

	for i, num := range nums {
		if _, ok := result[num]; !ok {
			result[num] = []int{i}
		} else {
			result[num] = append(result[num], i)
		}
	}
	return result
}
