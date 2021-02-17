# Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.
#
# Return a list of all possible strings we could create. You can return the output in any order.
#
#
#
# Example 1:
#
# Input: S = "a1b2"
# Output: ["a1b2","a1B2","A1b2","A1B2"]
#
# Example 2:
#
# Input: S = "3z4"
# Output: ["3z4","3Z4"]
#
# Example 3:
#
# Input: S = "12345"
# Output: ["12345"]
#
# Example 4:
#
# Input: S = "0"
# Output: ["0"]
#
#
#
# Constraints:
#
#     S will be a string with length between 1 and 12.
#     S will consist only of letters or digits.

# @param {String} s
# @return {String[]}
def letter_case_permutation(s)
  @ans = []

  dfs(s, '', 0)

  @ans
end

def dfs(s, cur, idx)
  return @ans.push(cur) if idx == s.size

  dfs(s, cur+s[idx], idx+1)

  if s[idx] >= 'a' && s[idx] <= 'z'
    dfs(s, cur + s[idx].upcase, idx+1)
  elsif s[idx] >= 'A' && s[idx] <= 'Z'
    dfs(s, cur + s[idx].downcase, idx+1)
  end
end