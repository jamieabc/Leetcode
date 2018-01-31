# Given an array of n integers where n > 1, nums, return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

# Solve it without division and in O(n).

# For example, given [1,2,3,4], return [24,12,8,6].

# Follow up:
# Could you solve it with constant space complexity? (Note: The output array does not count as extra space for the purpose of space complexity analysis.)

# @param {Integer[]} nums
# @return {Integer[]}
def product_except_self(nums)
  product = 1
  zero_count = 0
  nums.each do |num|
    if num.zero?
      zero_count += 1
    else
      product = num * product
    end
  end

  if zero_count.zero?
    nums.map { |num| product / num }
  elsif zero_count > 1
    Array.new(nums.length, 0)
  else
    nums.map do |num|
      num.zero? ? product : 0
    end
  end
end

# [1, 2, 3, 4]
# [1, 0]
