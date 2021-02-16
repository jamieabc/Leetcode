# There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.
#
# A province is a group of directly or indirectly connected cities and no other cities outside of the group.
#
# You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.
#
# Return the total number of provinces.
#
#
#
# Example 1:
#
# Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
# Output: 2
#
# Example 2:
#
# Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
# Output: 3
#
#
#
# Constraints:
#
#     1 <= n <= 200
#     n == isConnected.length
#     n == isConnected[i].length
#     isConnected[i][j] is 1 or 0.
#     isConnected[i][i] == 1
#     isConnected[i][j] == isConnected[j][i]

# @param {Integer[][]} is_connected
# @return {Integer}


# @param {Integer[][]} is_connected
# @return {Integer}

def find_circle_num(is_connected)
  n = is_connected.size
  parents = Array.new(n) { |i| i }
  ranks = Array.new(n, 1)

  disconnected = n

  is_connected.each_with_index do |arr, i|
    arr.each_with_index do |val, j|
      next if val == 0 || i == j

      p1, p2 = find(parents, i), find(parents, j)

      if p1 != p2
        disconnected -= 1

        if ranks[p1] >= ranks[p2]
          ranks[p1] += 1
          parents[p2] = p1
        else
          ranks[p2] += 1
          parents[p1] = p2
        end
      end
    end
  end

  disconnected
end

def find(parents, idx)
  parents[idx] = find(parents, parents[idx]) if idx != parents[idx]

  parents[idx]
end

def find_circle_num1(is_connected)
  n = is_connected.size
  visited = Array.new(n, false)

  count = 0

  visited.each_with_index do |val, idx|
    unless val
      count += 1
      bfs(is_connected, visited, idx)
    end
  end

  count
end

def bfs(is_connected, visited, idx)
  queue = Queue.new
  queue.push(idx)

  until queue.empty?
    n = queue.pop
    next if visited[n]
    visited[n] = true

    is_connected[n].each_with_index do |val, i|
      queue.push(i) if val == 1 && n != i && !visited[i]
    end
  end
end

#   Notes
#   1.  inspired from sample code, use set to check if key exist