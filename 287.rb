# Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
#
#   There is only one repeated number in nums, return this repeated number.
#
#
#
#   Example 1:
#
#   Input: nums = [1,3,4,2,2]
# Output: 2
#
# Example 2:
#
#   Input: nums = [3,1,3,4,2]
# Output: 3
#
# Example 3:
#
#   Input: nums = [1,1]
# Output: 1
#
# Example 4:
#
#   Input: nums = [1,1,2]
# Output: 1
#
#
#
# Constraints:
#
#   2 <= n <= 3 * 104
# nums.length == n + 1
# 1 <= nums[i] <= n
# All the integers in nums appear only once except for precisely one integer which appears two or more times.
#
#
#
#   Follow up:
#
#            How can we prove that at least one duplicate number must exist in nums?
# Can you solve the problem without modifying the array nums?
# Can you solve the problem using only constant, O(1) extra space?
# Can you solve the problem with runtime complexity less than O(n2)?

# @param {Integer[]} nums
# @return {Integer}

# slow/fast pointer
def find_duplicate(nums)
  slow = nums[0]
  fast = nums[nums[0]]

  while slow != fast
    slow = nums[slow]
    fast = nums[nums[fast]]
  end

  fast = 0

  while slow != fast
    slow = nums[slow]
    fast = nums[fast]
  end

  slow
end

# cyclic sort
def find_duplicate2(nums)
  nums.each_with_index do |_, idx|
    # index 0 ~ n-1, number 1 ~ n
    while nums[idx] - 1 != idx
      return nums[idx] if nums[idx] == nums[nums[idx]-1]

      tmp = nums[nums[idx]-1]
      nums[nums[idx]-1] = nums[idx]
      nums[idx] = tmp
    end
    p nums

  end
end

def find_duplicate1(nums)
  nums.sort!

  # could exist more times, just check duplicates
  (1...nums.size).each do |i|
    if nums[i] == nums[i-1]
      return nums[i]
    end
  end
end

#   Notes
#   1.  becareful about problem, repeated numb appears more than 1 time
#
#   2.  becareful, ruby's array element swap cannot use index as reference
#
#     e.g. nums[nums[idx]-1], nums[idx] = nums[idx], nums[nums[idx]-1] is not expected