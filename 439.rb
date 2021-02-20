# Given a string representing arbitrarily nested ternary expressions, calculate the result of the expression. You can always assume that the given expression is valid and only consists of digits 0-9, ?, :, T and F (T and F represent True and False respectively).
#
#   Note:
#
#   The length of the given string is ≤ 10000.
#     Each number will contain only one digit.
#       The conditional expressions group right-to-left (as usual in most languages).
#   The condition will always be either T or F. That is, the condition will never be a digit.
#   The result of the expression will always evaluate to either a digit 0-9, T or F.
#
#   Example 1:
#
#   Input: "T?2:3"
#
# Output: "2"
#
# Explanation: If true, then result is 2; otherwise result is 3.
#
#   Example 2:
#
#   Input: "F?1:T?4:5"
#
# Output: "4"
#
# Explanation: The conditional expressions group right-to-left. Using parenthesis, it is read/evaluated as:
#
#   "(F ? 1 : (T ? 4 : 5))"                   "(F ? 1 : (T ? 4 : 5))"
# -> "(F ? 1 : 4)"                 or       -> "(T ? 4 : 5)"
# -> "4"                                    -> "4"
#
# Example 3:
#
#   Input: "T?T?F:5:3"
#
# Output: "F"
#
# Explanation: The conditional expressions group right-to-left. Using parenthesis, it is read/evaluated as:
#
#   "(T ? (T ? F : 5) : 3)"                   "(T ? (T ? F : 5) : 3)"
# -> "(T ? F : 3)"                 or       -> "(T ? F : 5)"
# -> "F"                                    -> "F"

# @param {String} expression
# @return {String}


def parse_ternary(exp)
  @idx = 0

  recursive(exp)
end

def recursive(exp)
  char = exp[@idx]

  if @idx == exp.size-1 || exp[@idx+1] == ':'
    @idx += 2
    return char
  end

  @idx += 2

  first, second = recursive(exp), recursive(exp)

  if char == 'T'
    return first
  end

  second
end

def parse_ternary2(exp)
  stack = [].push(exp[-1])

  i = exp.size - 2

  while i >= 0
    if exp[i] == ':'
      stack.push(exp[i-1])
    else
      # ?, choose one
      if exp[i-1] == 'T'
        tmp = stack.last
      else
        tmp = stack[-2]
      end
      stack.pop(2)
      stack.push(tmp)
    end
    i -= 2
  end

  return stack.first
end

# tc: O(n^2)
# find last ?, and decompose from it
def parse_ternary1(exp)
  if exp.size > 1
    i = exp.rindex('?')
    if exp[i-1] == 'T'
      return parse_ternary1(exp[0...i-1] + exp[i+1] + exp[i+4..-1])
    else
      return parse_ternary1(exp[0...i-1] + exp[i+3..-1])
    end
  end

  exp
end