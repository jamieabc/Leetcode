# A gene string can be represented by an 8-character long string, with choices from "A", "C", "G", "T".

# Suppose we need to investigate about a mutation (mutation from "start" to "end"), where ONE mutation is defined as ONE single character changed in the gene string.

# For example, "AACCGGTT" -> "AACCGGTA" is 1 mutation.

# Also, there is a given gene "bank", which records all the valid gene mutations. A gene must be in the bank to make it a valid gene string.

# Now, given 3 things - start, end, bank, your task is to determine what is the minimum number of mutations needed to mutate from "start" to "end". If there is no such a mutation, return -1.

# Note:

# Starting point is assumed to be valid, so it might not be included in the bank.
# If multiple mutations are needed, all mutations during in the sequence must be valid.
# You may assume start and end string is not the same.
# Example 1:

# start: "AACCGGTT"
# end:   "AACCGGTA"
# bank: ["AACCGGTA"]

# return: 1
# Example 2:

# start: "AACCGGTT"
# end:   "AAACGGTA"
# bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]

# return: 2
# Example 3:

# start: "AAAAACCC"
# end:   "AACCCCCC"
# bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]

# return: 3

# @param {String} start
# @param {String} end
# @param {String[]} bank
# @return {Integer}
$min = -1

def min_mutation(start, dest, bank)
  $min = -1
  bank.unshift(start)
  arr = []
  calculate_distance(arr, bank)

  # p arr

  bfs(bank, dest, arr, [0])
  $min
end

def bfs(bank, dest, arr, path)

  # search all distance to 1
  # iterate all possible routes
  # terminate if destination is arrived

  if bank[path.last] == dest
    # p path
    if $min == -1
      $min = path.size - 1
    else
      $min = path.size - 1 < $min ? path.size - 1 : $min
    end
    return
  end

  current_idx = path.last
  possible = []
  arr[current_idx].each_index do |idx|
    possible.push(idx) if arr[current_idx][idx] == 1 && !path.include?(idx)
  end
  return -1 if possible.empty?

  # p "path #{path}, possible #{possible}"
  possible.each do |i|
    new_arr = path.dup
    bfs(bank, dest, arr, new_arr.push(i))
  end
end

# compare number of digits is different
def dist(str1, str2)
  result = 0
  for idx in 0..str1.length - 1
    result += 1 if str1[idx] != str2[idx]
  end
  result
end

# generate a mapping of gene
#   A B C D E
# A 0 1 2 3 4
# B 1 0 1 2 3
# C 2 1 0 1 2
# D 3 2 1 0 1
# E 4 3 2 1 0
def calculate_distance(arr, bank)
  for idx in 0..bank.size - 1
    result = []
    for j in 0..bank.size - 1
      result.push(dist(bank[idx], bank[j]))
    end
    arr.push(result)
  end
end

p min_mutation("AACCGGTT",
"AAACGGTA",
["AACCGGTA","AACCGCTA","AAACGGTA"])