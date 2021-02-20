# Given a string s of '(' , ')' and lowercase English characters.
#
# Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting parentheses string is valid and return any valid string.
#
# Formally, a parentheses string is valid if and only if:
#
#     It is the empty string, contains only lowercase characters, or
#     It can be written as AB (A concatenated with B), where A and B are valid strings, or
#     It can be written as (A), where A is a valid string.
#
#
#
# Example 1:
#
# Input: s = "lee(t(c)o)de)"
# Output: "lee(t(c)o)de"
# Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
#
# Example 2:
#
# Input: s = "a)b(c)d"
# Output: "ab(c)d"
#
# Example 3:
#
# Input: s = "))(("
# Output: ""
# Explanation: An empty string is also valid.
#
# Example 4:
#
# Input: s = "(a(b(c)d)"
# Output: "a(b(c)d)"
#
#
#
# Constraints:
#
#     1 <= s.length <= 10^5
#     s[i] is one of  '(' , ')' and lowercase English letters.

# @param {String} s
# @return {String}
#

def min_remove_to_make_valid(s)
  table = Array.new(s.size, true)
  open = 0

  s.each_char.with_index do |ch, i|
    if ch == '('
      open += 1
    elsif ch == ')'
      if open.positive?
        open -= 1
      else
        table[i] = false
      end
    end
  end

  i = s.size-1
  while open.positive?
    if s[i] == '('
      open -= 1
      table[i] = false
    end
    i -= 1
  end

  table.each_with_index.map { |exist, i| exist ? s[i] : '' }.join
end

def min_remove_to_make_valid2(s)
  open = []
  remove = {}

  s.each_char.with_index do |ch, i|
    if ch == '('
      open.push(i)
    elsif ch == ')'
      if open.any?
        open.pop
      else
        remove[i] = true
      end
    end
  end

  open.each { |i| remove[i] = true }

  s.each_char.with_index.map { |ch, i| remove.has_key?(i) ? '' : s[i] }.join
end

def min_remove_to_make_valid1(s)
  left, right = [], []

  s.each_char.with_index do |_, idx|
    if s[idx] == '('
      left.push(idx)
    elsif s[idx] == ')'
      if left.any?
        left.pop
      else
        right.push(idx)
      end
    end
  end

  ans = []

  s.each_char.with_index do |char, idx|
    if left.any? && idx == left.first
      left.shift
    elsif right.any? && idx == right.first
      right.shift
    else
      ans.push(s[idx])
    end
  end

  ans.join
end
