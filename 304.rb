# Given a 2D matrix matrix, find the sum of the elements inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).

# Range Sum Query 2D
# The above rectangle (with the red border) is defined by (row1, col1) = (2, 1) and (row2, col2) = (4, 3), which contains sum = 8.

# Example:

# Given matrix = [
#   [3, 0, 1, 4, 2],
#   [5, 6, 3, 2, 1],
#   [1, 2, 0, 1, 5],
#   [4, 1, 0, 1, 7],
#   [1, 0, 3, 0, 5]
# ]

# sumRegion(2, 1, 4, 3) -> 8
# sumRegion(1, 1, 2, 2) -> 11
# sumRegion(1, 2, 2, 4) -> 12

# Note:

#     You may assume that the matrix does not change.
#     There are many calls to sumRegion function.
#     You may assume that row1 ≤ row2 and col1 ≤ col2.

class NumMatrix
  attr_reader :matrix, :sum

=begin
    :type matrix: Integer[][]
=end
  def initialize(matrix)
    @matrix = matrix
  end


=begin
    :type row1: Integer
    :type col1: Integer
    :type row2: Integer
    :type col2: Integer
    :rtype: Integer
=end
    def sum_region(row1, col1, row2, col2)
      # validate each index number equal or greater than zero
      # find out horizontal & vertical distance
      # sum
      @sum = 0

      x_length = col2 - col1

      (row1..row2).each { |y| sum_row(matrix[y].slice(col1, x_length + 1)) }
      sum
    end

    def sum_row(row)
      @sum = sum + row.reduce(:+)
    end
end

# Your NumMatrix object will be instantiated and called as such:
# obj = NumMatrix.new(matrix)
# param_1 = obj.sum_region(row1, col1, row2, col2)

obj = NumMatrix.new(
  [[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]
)

p obj.sum_region(1,2,2,4)
