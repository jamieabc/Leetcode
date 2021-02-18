# Given strings A and B of the same length, we say A[i] and B[i] are equivalent characters. For example, if A = "abc" and B = "cde", then we have 'a' == 'c', 'b' == 'd', 'c' == 'e'.
#
# Equivalent characters follow the usual rules of any equivalence relation:
#
#     Reflexivity: 'a' == 'a'
#     Symmetry: 'a' == 'b' implies 'b' == 'a'
#     Transitivity: 'a' == 'b' and 'b' == 'c' implies 'a' == 'c'
#
# For example, given the equivalency information from A and B above, S = "eed", "acd", and "aab" are equivalent strings, and "aab" is the lexicographically smallest equivalent string of S.
#
# Return the lexicographically smallest equivalent string of S by using the equivalency information from A and B.
#
#
#
# Example 1:
#
# Input: A = "parker", B = "morris", S = "parser"
# Output: "makkek"
# Explanation: Based on the equivalency information in A and B, we can group their characters as [m,p], [a,o], [k,r,s], [e,i]. The characters in each group are equivalent and sorted in lexicographical order. So the answer is "makkek".
#
# Example 2:
#
# Input: A = "hello", B = "world", S = "hold"
# Output: "hdld"
# Explanation:  Based on the equivalency information in A and B, we can group their characters as [h,w], [d,e,o], [l,r]. So only the second letter 'o' in S is changed to 'd', the answer is "hdld".
#
# Example 3:
#
# Input: A = "leetcode", B = "programs", S = "sourcecode"
# Output: "aauaaaaada"
# Explanation:  We group the equivalent characters in A and B as [a,o,e,r,s,c], [l,p], [g,t] and [d,m], thus all letters in S except 'u' and 'd' are transformed to 'a', the answer is "aauaaaaada".
#
#
#
# Note:
#
#     String A, B and S consist of only lowercase English letters from 'a' - 'z'.
#     The lengths of string A, B and S are between 1 and 1000.
#     String A and B are of the same length.

# @param {String} a
# @param {String} b
# @param {String} s
# @return {String}
def smallest_equivalent_string(a, b, s)
  parents = Array.new(26) { |i| i }
  ranks = Array.new(26, 1)

  a.each_char.with_index do |_, i|
    p1, p2 = find(parents, offset(a[i])), find(parents, offset(b[i]))

    if ranks[p1] >= ranks[p2]
      ranks[p1] += 1
      parents[p2] = p1
    else
      ranks[p2] += 1
      parents[p1] = p2
    end
  end

  ans = []

  s.each_char.with_index do |char, idx|
    group = find(parents, offset(char))

    parents.each_with_index do |val, i|
      if val == group
        ans[idx] = ('a'.ord + i).chr
        break
      end
    end
  end

  ans.join('')
end

def offset(char)
  char.ord - 'a'.ord
end

def find(parents, i)
  parents[i] = find(parents, parents[i]) if parents[i] != i
  parents[i]
end