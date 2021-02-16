# Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
#
#
#
# Example 1:
#
# Input: root = [1,2,3,null,5,null,4]
# Output: [1,3,4]
#
# Example 2:
#
# Input: root = [1,null,3]
# Output: [1,3]
#
# Example 3:
#
# Input: root = []
# Output: []
#
#
#
# Constraints:
#
#     The number of nodes in the tree is in the range [0, 100].
#     -100 <= Node.val <= 100

# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val = 0, left = nil, right = nil)
#         @val = val
#         @left = left
#         @right = right
#     end
# end
# @param {TreeNode} root
# @return {Integer[]}

def right_side_view(root)
  ans = []
  return ans if root.nil?

  queue = Queue.new
  queue.push(root)

  until queue.empty?
    idx = queue.size

    idx.times do |i|
      n = queue.pop

      queue.push(n.left) unless n.left.nil?
      queue.push(n.right) unless n.right.nil?

      # store right most node's value
      ans.push(n.val) if i == idx-1
    end
  end

  ans
end

def right_side_view1(root)
  @ans = []

  dfs(root, 0)

  @ans
end

def dfs(node, level)
  return if node.nil?

  diff = level + 1 -@ans.size
  diff.times { |t| @ans.push(0) } if diff > 0

  @ans[level] = node.val

  dfs(node.left, level+1)
  dfs(node.right, level+1)
end

#   Notes
#   1.  inspired form sample code, since right side view is wanted, traverse
#       first right then left and use variable prev_level to detect level change