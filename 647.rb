# Given a string, your task is to count how many palindromic substrings in this string.
#
#   The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.
#
#   Example 1:
#
#   Input: "abc"
# Output: 3
# Explanation: Three palindromic strings: "a", "b", "c".
#
#
#
#   Example 2:
#
#   Input: "aaa"
# Output: 6
# Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
#
#
#
#   Note:
#
#   The input string length won't exceed 1000.

# @param {String} s
# @return {Integer}
def count_substrings(s)
  count = 0

  for i in 0..(s.size-1) do
    count += 1

    # center at i
    j = 1
    while i-j >= 0 && i+j < s.size do
      if s[i-j] == s[i+j]
        count += 1
      else
        break
      end

      j += 1
    end

    # center at i & i+1
    if i+1 < s.size && s[i] == s[i+1]
      count += 1

      j = 2
      while i-j+1 >= 0 && i+j < s.size do
        if s[i-j+1] == s[i+j]
          count += 1
        else
          break
        end

        j += 1
      end
    end
  end

  count
end