# You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

# Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

# @param {Integer[]} nums
# @return {Integer}
def rob(nums)

  # max sum comes from max(current, previous + this integer)

  # use dynamic programming to solve the problem:
  # variable current means sum that doesn't include this integer,
  # variable previous means sum that includes this integer
  # only 2 situations can appear, choose this integer or not to choose this integer

  return 0 if nums.empty?

  return nums[0] if nums.size == 1

  previous = nums[0]
  current = max(nums[0], nums[1])
  idx = 2
  while idx < nums.size
    temp = current
    current = max(previous + nums[idx], current)
    previous = temp
    idx += 1
  end
  current
end

def max(a, b)
  a >= b ? a : b
end

# p rob([1, 2, 3, 4, 5])
# p rob([1, 3, 1, 1, 9])
# p rob([4, 1, 2, 7, 5, 3, 1])
# p rob([1, 1, 1, 2])
# p rob([6,3,10,8,2,10,3,5,10,5,3])
# p rob([1, 1, 1])
# p rob([2, 3, 2])
p rob([2, 1, 1, 2])