# In an N by N square grid, each cell is either empty (0) or blocked (1).
#
#   A clear path from top-left to bottom-right has length k if and only if it is composed of cells C_1, C_2, ..., C_k such that:
#
#                                                                                                                            Adjacent cells C_i and C_{i+1} are connected 8-directionally (ie., they are different and share an edge or corner)
# C_1 is at location (0, 0) (ie. has value grid[0][0])
# C_k is at location (N-1, N-1) (ie. has value grid[N-1][N-1])
# If C_i is located at (r, c), then grid[r][c] is empty (ie. grid[r][c] == 0).
#   Return the length of the shortest such clear path from top-left to bottom-right.  If such a path does not exist, return -1.
#
#
#
#   Example 1:
#
#   Input: [[0,1],[1,0]]
#
#
# Output: 2
#
# Example 2:
#
#   Input: [[0,0,0],[1,1,0],[1,1,0]]
#
#
# Output: 4
#
#
#
# Note:
#
#   1 <= grid.length == grid[0].length <= 100
# grid[r][c] is 0 or 1

# @param {Integer[][]} grid
# @return {Integer}
def shortest_path_binary_matrix(grid)
  w, h = grid[0].size, grid.size

  if grid[0][0] == 1 || grid[h-1][w-1] == 1
    return -1
  end

  # boundary condition, need to check because steps starts from 2
  if w == 1 && h == 1
    return 1
  end

  visited = Array.new(h) { Array.new(w, false) }
  visited[0][0] = true

  queue = []
  queue << [0, 0]
  steps = 1

  # bfs
  while queue.any?
    steps += 1
    i, size = 0, queue.size

    while i < size
      p = queue.shift

      [-1, 0, 1].each do |j|
        [-1, 0, 1].each do |k|
          newY, newX = p[0]+j, p[1]+k

          if newY == h-1 && newX == w-1
            return steps
          end

          next if newX.negative? || newY.negative? || newX == w || newY == h || grid[newY][newX] == 1 || visited[newY][newX]

          visited[newY][newX] = true
          queue << [newY, newX]
        end
      end

      i += 1
    end
  end

  -1
end

#   Notes
#   1.  inspired from sample code, could use [-1][-1] to check destination
#       use array.empty? (array.any?) to check
#       next if some condition