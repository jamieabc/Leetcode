# Suppose you have n integers labeled 1 through n. A permutation of those n integers perm (1-indexed) is considered a beautiful arrangement if for every i (1 <= i <= n), either of the following is true:
#
#     perm[i] is divisible by i.
#     i is divisible by perm[i].
#
# Given an integer n, return the number of the beautiful arrangements that you can construct.
#
#
#
# Example 1:
#
# Input: n = 2
# Output: 2
# Explanation:
# The first beautiful arrangement is [1,2]:
#     - perm[1] = 1 is divisible by i = 1
#     - perm[2] = 2 is divisible by i = 2
# The second beautiful arrangement is [2,1]:
#     - perm[1] = 2 is divisible by i = 1
#     - i = 2 is divisible by perm[2] = 1
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
#     1 <= n <= 15

# @param {Integer} n
# @return {Integer}
def count_arrangement(n)
  @count = 0
  used = Array.new(n+1, false)

  backtracking(n, n, used)

  @count
end

def backtracking(n, store, used)
  if store.zero?
    @count += 1
    return
  end

  for i in 1..n
    next if used[i]
    used[i] = true

    backtracking(n, store-1, used) if store % i == 0 || i % store == 0

    used[i] = false
  end
end

