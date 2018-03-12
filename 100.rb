#  Given two binary trees, write a function to check if they are the same or not.

# Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

# Example 1:

# Input:     1         1
#           / \       / \
#          2   3     2   3

#         [1,2,3],   [1,2,3]

# Output: true

# Example 2:

# Input:     1         1
#           /           \
#          2             2

#         [1,2],     [1,null,2]

# Output: false

# Example 3:

# Input:     1         1
#           / \       / \
#          2   1     1   2

#         [1,2,1],   [1,1,2]

# Output: false

# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} p
# @param {TreeNode} q
# @return {Boolean}
def is_same_tree(p, q)
  # three situations:
  # 1. empty nodes
  # 2. single node
  # 3. multiple nodes
  return true if p&.val.nil? && q&.val.nil?
  same?(p, q)
end


def same?(node1, node2)
  # three situations:
  # 1. val is empty
  # 2. left is empty
  # 3. right is empty
  return true if node1.nil? && node2.nil?
  return false if node1&.val != node2&.val
  same?(node1.left, node2.left) && same?(node1.right, node2.right)
end
