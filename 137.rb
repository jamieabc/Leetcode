#  Given an array of integers, every element appears three times except for one, which appears exactly once. Find that single one.

# Note:
# Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

# solution:
# 1. for linear runtime complexity, use hash to record if a number has occured twice
# 2. for no extra memeory, scan array for each element to see if another one can be found

# use bit-wise operation to solve the problem
# goal: if a bit adds three times, set it to 0
# 0 + 1 = 0 ^ 1, 1 + 1 = 1 ^ 1
# since there are three states: occur firts/second/third time
# need two bits to denote a bit's occurance status: ab
# a bit occurs first time: ab = 0b01
# a bit occurs second time: ab = 0b10
# a bit occurs third time: ab = 0b11 -> convert to 0b00
#
# for number X = ...1...0...
#            X = ...1...0...
#            X = ...1...0...
#          sum = ...1...0...
#    converted = ...0...0...

# @param {Integer[]} nums
# @return {Integer}
def single_number1(nums)
  hsh = Hash.new(-1)
  nums.each { |n| hsh[n] += 1 }
  hsh.select { |k, v| v.zero? }.keys.first
end

def single_number(nums)
  a = 0
  b = 0
  nums.each do |n|
    a = a ^ (b & n)             # a added only if b == 1 & c == 1
    b = b ^ n                   # b = b + c

    # convert 0b11 -> 0b00
    c = a & b
    a &= ~c                     # a = 0 if c == 1, otherwise, a = a
    b &= ~c                     # b = 0 if c == 1, otherwise, b = b
  end

  a | b
end

p single_number([1])