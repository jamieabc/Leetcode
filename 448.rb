# Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

# Find all the elements of [1, n] inclusive that do not appear in this array.

# Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

# Example:

# Input:
# [4,3,2,7,8,2,3,1]

# Output:
# [5,6]
# @param {Integer[]} nums
# @return {Integer[]}
def find_disappeared_numbers(nums)
  output = Array.new(nums.size, 0)
  nums.each { |n| output[n - 1] = 1 }
  output.map.with_index { |x, i| i + 1 if x.zero? }.compact
end

p find_disappeared_numbers([4, 3, 2, 7, 8, 2, 3, 1])
