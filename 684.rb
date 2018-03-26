# In this problem, a tree is an undirected graph that is connected and has no cycles.

# The given input is a graph that started as a tree with N nodes (with distinct values 1, 2, ..., N), with one additional edge added. The added edge has two different vertices chosen from 1 to N, and was not an edge that already existed.

# The resulting graph is given as a 2D-array of edges. Each element of edges is a pair [u, v] with u < v, that represents an undirected edge connecting nodes u and v.

# Return an edge that can be removed so that the resulting graph is a tree of N nodes. If there are multiple answers, return the answer that occurs last in the given 2D-array. The answer edge [u, v] should be in the same format, with u < v.

# Example 1:

# Input: [[1,2], [1,3], [2,3]]
# Output: [2,3]
# Explanation: The given undirected graph will be like this:
#   1
#  / \
# 2 - 3

# Example 2:

# Input: [[1,2], [2,3], [3,4], [1,4], [1,5]]
# Output: [1,4]
# Explanation: The given undirected graph will be like this:
# 5 - 1 - 2
#     |   |
#     4 - 3

# Note:
# The size of the input 2D-array will be between 3 and 1000.
# Every integer represented in the 2D-array will be between 1 and N, where N is the size of the input array.

# find a way to decide loop - start and end with same node
# decide one node to remove
# for an edge [u, v], generate a hash with it's key is u and [v]
# loop: search all possible nodes to find if any path has same start node
# the edge cannot be removed: only contains one node

# @param {Integer[][]} edges
# @return {Integer[]}
$loops = []
def find_redundant_connection1(edges)
  $loops = []
  hsh = generate_hash(edges)
  find_loop(hsh, hsh.keys.first)
  find_possible_edge(hsh, edges)
end

# edges = [[1, 2], [3, 4]]
# hash = {
#   1: [2],
#   3: [4]
# }
def generate_hash(edges)
  hsh = Hash.new { |hash, key| hash[key] = [] }
  edges.each do |key, value|
    hsh[key] << value
    hsh[value] << key
  end
  hsh
end

# traverse each node to find a loop
def find_loop(hash, node, path = [])

  # cannot traverse to previous path
  if path[-2] && path[-2] == node
    return
  end

  # if index already exist, found one loop
  # trim the shorted loop
  if path.index(node)
    start = path.index node
    path << node
    new_path = path[start, path.size]
    $loops << new_path
    return
  end

  # no loop found
  return if node.nil?

  # continue to next next node
  hash[node].each do |next_node|
    find_loop(hash, next_node, path + [node])
  end
end

def find_possible_edge(hsh, edges)
  # hsh.select { |key, value| $loops.include?(key) && value.size != 1 }
  #   .keys.map do |key|
  #   idx = $loops.index key
  #   next if idx + 1 == $loops.size
  #   [key, $loops[idx + 1]]
  # end
  possible_edges = []
  $loops.each do |loop|
    (1..(loop.size - 1)).each  do |i|
      first, second = loop[i, 2]
      if hsh[first].size > 1 && hsh[second].size > 1
        valid_edge = first < second ? [first, second] : [second, first]
        possible_edges << valid_edge
      end
    end
  end
  edges[possible_edges.map { |e| edges.index(e) }.max]
end

# use disjoint set to decide if an element is already included in a set
# two operations are needed: find_root and union
# find_root: find the root of a set
# union: combine 2 sets
# paretns is an Arary to store each set's parent, index is element, value is set
#
# for example, [0, 1, 1, 2, 2, 3, 3, 3] means:
# 0 in set 0, index 1 in set 1, index 2 in set 1, index 3 in set 2, etc.
#      1   3   5
#    /   /   /  \
#  2   4   6     7
#
#
def find_redundant_connection(edges)
  # initially, each set is viewed as disjoin set
  parents = Array.new(edges.size + 1) { |i| i }

  # initially, each set has same rank (tree depth)
  ranks = Array.new(edges.size + 1, 0)

  possible_edges = edges.map { |x, y| union(x, y, parents, ranks) }

  last_edge = possible_edges.map.with_index do |value, index|
    value == false ? index : 0
  end.max

  edges[last_edge]
end

# find the root of a given set
# also update all element among the path
#       1                 1
#     /                 /   \
#    2       =>       2      3
#   /
# 3

def find_root(current, parents)
  if parents[current] != current
    parents[current] = find_root(parents[current], parents)
    return parents[current]
  end
  current
end

# union two sets if they are not in the same set
# if two sets already in connected (in the same set, e.g. have same parent),
# then this edge can be removed
#     1     3                 1
#   /     /                 /   \
# 2     4       =>        2      3
#                                 \
#                                  4
#
def union(set_a, set_b, parents, ranks)
  root_a = find_root(set_a, parents)
  root_b = find_root(set_b, parents)
  return false if root_a == root_b
  optimize_merge(root_a, root_b, parents, ranks)
  true
end

# during union, it's better not to increase the tree's depth
# merge depth shorter into depth longer will not increase tree's depth
#       1      3
#     /      /
#    2     4
#        /
#      5
#
# for 3 merge into 1 will increase tree depth
# for 1 merge into 3 will not increase tree depth
#
#           1                    3
#         /  \                 /  \
#       2     3              4     1
#              \            /       \
#               4         5          2
#                \
#                 5
#
def optimize_merge(root_a, root_b, parents, ranks)
  rank_a = ranks[root_a]
  rank_b = ranks[root_b]

  if rank_a == rank_b || rank_a > rank_b
    parents[root_b] = root_a
    ranks[root_b] += 1
  else
    parents[root_a] = root_b
    ranks[root_a] += 1
  end
end

# p find_redundant_connection([[1,2], [1,3], [2,3]])
# p find_redundant_connection([[1,2], [2,3], [3,4], [1,4], [1,5]])
# p find_redundant_connection([[30,44],[34,47],[22,32],[35,44],[26,36],[2,15],[38,41],[28,35],[24,37],[14,49],[44,45],[11,50],[20,39],[7,39],[19,22],[3,17],[15,25],[1,39],[26,40],[5,14],[6,23],[5,6],[31,48],[13,22],[41,44],[10,19],[12,41],[1,12],[3,14],[40,50],[19,37],[16,26],[7,25],[22,33],[21,27],[9,50],[24,42],[43,46],[21,47],[29,40],[31,34],[9,31],[14,31],[5,48],[3,18],[4,19],[8,17],[38,46],[35,37],[17,43]])
# p find_redundant_connection([[1,2],[2,3],[3,4],[1,4],[1,5]])

# dot file
# dot test.dot -T png -o test.png
# graph demo1 {
#   "30" -- "44";
#   "34" -- "47";
#   "22" -- "32";
#   "35" -- "44";
#   "26" -- "36";
#   "2" -- "15";
#   "38" -- "41";
#   "28" -- "35";
#   "24" -- "37";
#   "14" -- "49";
#   "44" -- "45";
#   "11" -- "50";
#   "20" -- "39";
#   "7" -- "39";
#   "19" -- "22";
#   "3" -- "17";
#   "15" -- "25";
#   "1" -- "39";
#   "26" -- "40";
#   "5" -- "14";
#   "6" -- "23";
#   "5" -- "6";
#   "31" -- "48";
#   "13" -- "22";
#   "41" -- "44";
#   "10" -- "19";
#   "12" -- "41";
#   "1" -- "12";
#   "3" -- "14";
#   "40" -- "50";
#   "19" -- "37";
#   "16" -- "26";
#   "7" -- "25";
#   "22" -- "33";
#   "21" -- "27";
#   "9" -- "50";
#   "24" -- "42";
#   "43" -- "46";
#   "21" -- "47";
#   "29" -- "40";
#   "31" -- "34";
#   "9" -- "31";
#   "14" -- "31";
#   "5" -- "48";
#   "3" -- "18";
#   "4" -- "19";
#   "8" -- "17";
#   "38" -- "46";
#   "35" -- "37";
#   "17" -- "43";
# }