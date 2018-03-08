# Given a word, you need to judge whether the usage of capitals in it is right or not.

# We define the usage of capitals in a word to be right when one of the following cases holds:

#     All letters in this word are capitals, like "USA".
#     All letters in this word are not capitals, like "leetcode".
#     Only the first letter in this word is capital if it has more than one letter, like "Google".

# Otherwise, we define that this word doesn't use capitals in a right way.

# Example 1:

# Input: "USA"
# Output: True

# Example 2:

# Input: "FlaG"
# Output: False

# Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.

# @param {String} word
# @return {Boolean}
def detect_capital_use(word)
  ords = word.split('').map(&:ord)
  upper_range = 'A'.ord .. 'Z'.ord
  lower_range = 'a'.ord .. 'z'.ord
  word_range = upper_range.to_a + lower_range.to_a

  return true if word.size == 1 && word_range.include?(ords.first)

  # judge all characters are words
  # if first word upper-case, no matter what rest words are, but needs to be in same form
  # if first word lower-case, rest words should be lower-case

  capital = upper_range.include? ords[0]
  if capital
    range = (upper_range.include? word[1].ord) ? upper_range : lower_range
    ords[1..-1].reject { |c| range.include? c }.empty?
  else
    ords[1..-1].reject { |c| lower_range.include? c }.empty?
  end
end

p detect_capital_use('G')