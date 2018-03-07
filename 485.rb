# Given a binary array, find the maximum number of consecutive 1s in this array.

# Example 1:

# Input: [1,1,0,1,1,1]
# Output: 3
# Explanation: The first two digits or the last three digits are consecutive 1s.
#     The maximum number of consecutive 1s is 3.

# Note:

#     The input array will only contain 0 and 1.
#     The length of input array is a positive integer and will not exceed 10,000

# @param {Integer[]} nums
# @return {Integer}
def find_max_consecutive_ones(nums)
  max_cnt = 0
  tmp_cnt = 0
  nums.each do |n|
    if n.zero?
      max_cnt = tmp_cnt if max_cnt < tmp_cnt
      tmp_cnt = 0
    else
      tmp_cnt += 1
    end
  end
  max_cnt = tmp_cnt if max_cnt < tmp_cnt
  max_cnt
end