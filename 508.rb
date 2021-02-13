# Given the root of a tree, you are asked to find the most frequent subtree sum. The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself). So what is the most frequent subtree sum value? If there is a tie, return all the values with the highest frequency in any order.
#
#   Examples 1
# Input:
#
#   5
# /  \
# 2   -3
# return [2, -3, 4], since all the values happen only once, return all of them in any order.
# Examples 2
# Input:
#
#   5
#  /  \
# 2   -5
# return [2], since 2 happens twice, however -5 only occur once.
#   Note: You may assume the sum of values in any subtree is in the range of 32-bit signed integer.


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

def find_frequent_tree_sum(root)
  @maxOccurrence = 0

  counter = Hash.new(0)
  recursive(root, counter)

  counter.keys.select do |key|
    counter[key] == @maxOccurrence
  end
end

def recursive(node, counter)
  if node.nil?
    return 0
  end

  cur = node.val + recursive(node.left, counter) + recursive(node.right, counter)

  counter[cur] += 1

  if counter[cur] > @maxOccurrence
    @maxOccurrence = counter[cur]
  end

  cur
end

#   Notes
#   1.  inspired from sample code
#       use select to get answer