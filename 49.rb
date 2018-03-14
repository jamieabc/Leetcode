# Given an array of strings, group anagrams together.

# For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
# Return:

# [
#   ["ate", "eat","tea"],
#   ["nat","tan"],
#   ["bat"]
# ]

# Note: All inputs will be in lower-case.

# @param {String[]} strs
# @return {String[][]}
def group_anagrams1(strs)
  anagrams = []
  hashed_arr = strs.map { |str| calculate_char_occur_times str }

  hashed_arr.each_with_index do |hsh, idx|
    grouped = false
    anagrams.each do |group|
n      if hashed_arr[group.first] == hsh
        group << idx
        grouped = true
      end

    end
    anagrams << [idx] unless grouped
  end
  anagrams.map { |group_of_idx| group_of_idx.map { |idx| strs[idx] } }
end

# decide if two words are anagrams
# iterate through all characters, store each chracter occur
# times into a hash.
# compare hash to decide if it's anagram

def calculate_char_occur_times(str)
  hsh = Hash.new { |h, k| h[k] = 0 }
  str.each_char { |c| hsh[c] += 1 }
  hsh
end

def group_anagrams2(strs)
  anagrams = []
  size = strs.size
  while !size.zero?
    target = strs.first
    group = [target]
    strs[1, strs.size - 1].each do |str|
      group << str if anagram?(target, str)
    end
    anagrams << group
    strs -= group
    size = strs.size
  end
  anagrams
end

def anagram?(str1, str2)
  return false if str1.size != str2.size
  str1 == str2
end

# both methods of group_anagrams1 and group_anagrams2 exceeds time limit
# the problem here is that every word needs a calculation of anagram?
# but this is not necessary, just convert word into sorted form and put into array

def group_anagrams(strs)
  anagrams = Hash.new([])
  strs.each_with_index do |str, idx|
    anagrams[str.chars.sort.join] += [idx]
  end
  anagrams.reduce([]) { |memo, obj| memo << obj[1].map { |idx| strs[idx] } }
end

p group_anagrams(["eat","tea","tan","ate","nat","bat"])
