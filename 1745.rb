# Given a string s, return true if it is possible to split the string s into three non-empty palindromic substrings. Otherwise, return false.​​​​​
#
# A string is said to be palindrome if it the same string when reversed.
#
#
#
#   Example 1:
#
#   Input: s = "abcbdd"
# Output: true
# Explanation: "abcbdd" = "a" + "bcb" + "dd", and all three substrings are palindromes.
#   Example 2:
#
#   Input: s = "bcbddxy"
# Output: false
# Explanation: s cannot be split into 3 palindromes.
#
#
#   Constraints:
#
#   3 <= s.length <= 2000
#   s consists only of lowercase English letters.

# @param {String} s
# @return {Boolean}

def check_partitioning(s)
  s = s.chars.map(&:ord)
  n = s.size
  dp = Array.new(n) { Array.new(n, false) }

  for i in 0...n
    dp[i][i] = true
    for j in 0...i
      if s[i] == s[j] && (j == i-1 || dp[j+1][i-1])
        dp[j][i] = true
      end
    end
  end

  for i in 0..n-2
    next unless dp[0][i]

    for j in i+1..n-2
      return true if dp[i+1][j] && dp[j+1][n-1]
    end
  end

  false
end

def check_partitioning2(s)
  size = s.size

  dp = Array.new(size) { Array.new(size, false) }

  for i in 0..size-1
    dp[i][i] = true
    dp[i][i+1] = true if i+1 < size && s[i] == s[i+1]
  end

  for i in (size-1..0).step(-1)
    for j in i+2..size-1
      dp[i][j] = dp[i+1][j-1] if s[i] == s[j]
    end
  end

  for i in 0..size-3
    next unless dp[0][i]

    for j in i+1..size-2
      return true if dp[i+1][j] && dp[j+1][size-1]
    end
  end

  false
end

def check_partitioning1(s)
  size = s.size

  for i in 0..size-3
    for j in i+1..size-2
      return true if palindrome?(s, 0, i) && palindrome?(s, i+1, j) && palindrome?(s, j+1, size-1)
    end
  end

  false
end

def palindrome?(s, i, j)
  return true if i >= j
  s[i] == s[j] && palindrome?(s, i+1, j-1)
end

#   Note
#   1.  inspired from sample code
#       for i in 0..n
#       s.chars.map(&:ord)

#   2.  continue loop as fast as possible

#   3.  inspired from https://youtu.be/IOLi6XbolQQ?t=1082

#       alex uses a beautiful recursion form
#       however, ruby gets timeout

#       as by dp, becareful about range, it's i: size-1 -> 0, j: i+2 -> size-1

#   4.  inspired from sample coe

#       for i in 0..2 => 0, 1, 2
#       for i in 0...2 => 0, 1
#
#       convert string into int array makes execution much faster
#
#       also, it's really beautiful to have bottom-up building