# Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most one element.
#
#   We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).
#
#
#
#   Example 1:
#
#   Input: nums = [4,2,3]
# Output: true
# Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
#   Example 2:
#
#   Input: nums = [4,2,1]
# Output: false
# Explanation: You can't get a non-decreasing array by modify at most one element.
#
#
# Constraints:
#
# n == nums.length
# 1 <= n <= 104
# -105 <= nums[i] <= 105

# @param {Integer[]} nums
# @return {Boolean}
def check_possibility(nums)
  idx = -1
  size = nums.size

  for i in 1...size
    if nums[i] < nums[i-1]
      if idx == -1
        idx = i
      else
        return false
      end
    end
  end

  idx == -1 || idx == 1 || idx == size-1 || nums[idx-2] <= nums[idx] || nums[idx-1] <= nums[idx+1]
end