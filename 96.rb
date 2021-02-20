# Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.
#
#
#
# Example 1:
#
# Input: n = 3
# Output: 5
#
# Example 2:
#
# Input: n = 1
# Output: 1
#
#
#
# Constraints:
#
#     1 <= n <= 19

# @param {Integer} n
# @return {Integer}
def num_trees(n)
  dp = Array.new(n+1, 0)
  dp[0] = 1
  dp[1] = 1

  (2..n).each do |i|
    j = 1
    count = 0
    while j <= i
      count += dp[j-1] * dp[i-j]
      j += 1
    end

    dp[i] = count
  end

  dp[n]
end
