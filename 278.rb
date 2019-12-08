# You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.
#
# Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
#
# You are given an API bool isBadVersion(version) which will return whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.
#
# Example:
#
# Given n = 5, and version = 4 is the first bad version.
#
# call isBadVersion(3) -> false
# call isBadVersion(5) -> true
# call isBadVersion(4) -> true
#
# Then 4 is the first bad version.


# The is_bad_version API is already defined for you.
# @param {Integer} version
# @return {boolean} whether the version is bad
# def is_bad_version(version):

# @param {Integer} n
# @return {Integer}
def first_bad_version(n)
  hsh = {}
  hsh[n.to_s.to_sym] = true
  recursive(1, n, n, hsh)
end

def recursive(small, large, max, hsh)
  middle = (small + large) / 2
  key = middle.to_s.to_sym

  next_middle = middle + 1
  next_key = next_middle.to_s.to_sym

  prev_middle = middle - 1
  prev_key = prev_middle.to_s.to_sym

  unless hsh.key?(key)
    hsh[key] = is_bad_version(middle)
  end

  if hsh.key?(next_key)
    return next_middle if !hsh[key] && hsh[next_key]
  end

  if hsh.key?(prev_key)
    return middle if !hsh[prev_middle] && hsh[middle]
  end

  if hsh[key]
    recursive(small, middle - 1, max, hsh)
  else
    recursive(middle+1, large, max, hsh)
  end
end

