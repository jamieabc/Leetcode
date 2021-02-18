# A sequence of numbers is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.
#
# For example, these are arithmetic sequences:
#
# 1, 3, 5, 7, 9
# 7, 7, 7, 7
# 3, -1, -5, -9
#
# The following sequence is not arithmetic.
#
# 1, 1, 2, 5, 7
#
#
#
# A zero-indexed array A consisting of N numbers is given. A slice of that array is any pair of integers (P, Q) such that 0 <= P < Q < N.
#
# A slice (P, Q) of the array A is called arithmetic if the sequence:
# A[P], A[P + 1], ..., A[Q - 1], A[Q] is arithmetic. In particular, this means that P + 1 < Q.
#
# The function should return the number of arithmetic slices in the array A.
#
#
# Example:
#
# A = [1, 2, 3, 4]
#
# return: 3, for 3 arithmetic slices in A: [1, 2, 3], [2, 3, 4] and [1, 2, 3, 4] itself.

# @param {Integer[]} a
# @return {Integer}
def number_of_arithmetic_slices(a)
  count = 0
  idx = 0

  while idx < a.size-2
    j = idx+2
    tmp = 0
    diff = a[idx+1] - a[idx]

    while j < a.size && a[j] - a[j-1] == diff
      j += 1
      tmp += 1
      count += tmp
    end

    idx = j-1 if j > idx+2

    idx += 1
  end

  count
end

#   Notes
#   1.  inspired from sample code, use (1...3).each