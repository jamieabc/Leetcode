# Follow up for "Remove Duplicates":
# What if duplicates are allowed at most twice?

# For example,
# Given sorted array nums = [1,1,1,2,2,3],

# Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3. It doesn't matter what you leave beyond the new length.

# @param {Integer[]} nums
# @return {Integer}
def remove_duplicates(nums)
  hsh = Hash.new(0)
  duplicates = []
  nums.each_with_index do |num, idx|
    if hsh[num] != 2
      hsh[num] += 1
    else
      duplicates.push(idx)
    end
  end
  duplicates.reverse.each { |idx| nums.delete_at(idx) }
  nums.length
end
