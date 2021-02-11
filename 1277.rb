# Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.
#
#
#
#   Example 1:
#
#   Input: matrix =
#   [
#     [0,1,1,1],
#     [1,1,1,1],
#     [0,1,1,1]
#   ]
# Output: 15
# Explanation:
#   There are 10 squares of side 1.
#   There are 4 squares of side 2.
#   There is  1 square of side 3.
#   Total number of squares = 10 + 4 + 1 = 15.
#
#   Example 2:
#
#   Input: matrix =
#   [
#     [1,0,1],
#     [1,1,0],
#     [1,1,0]
#   ]
# Output: 7
# Explanation:
#   There are 6 squares of side 1.
#   There is 1 square of side 2.
#   Total number of squares = 6 + 1 = 7.
#
#
#
#   Constraints:
#
#   1 <= arr.length <= 300
# 1 <= arr[0].length <= 300
# 0 <= arr[i][j] <= 1

# @param {Integer[][]} matrix
# @return {Integer}
def count_squares(matrix)
  # dp[i][j]: # of square end at [i, j]
  dp = Array.new(matrix.size) { Array.new(matrix[0].size, 0) }

  count = 0
  matrix.each_with_index do |row, i|
    row.each_with_index do |col, j|
      if matrix[i][j] == 1
        if i > 0 && j > 0
          dp[i][j] = [dp[i-1][j], dp[i][j-1], dp[i-1][j-1]].min + 1
        else
          dp[i][j] = 1
        end

        count += dp[i][j]
      end
    end
  end

  count
end