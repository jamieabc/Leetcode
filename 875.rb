# Koko loves to eat bananas.  There are N piles of bananas, the i-th pile has piles[i] bananas.  The guards have gone and will come back in H hours.
#
# Koko can decide her bananas-per-hour eating speed of K.  Each hour, she chooses some pile of bananas, and eats K bananas from that pile.  If the pile has less than K bananas, she eats all of them instead, and won't eat any more bananas during this hour.
#
# Koko likes to eat slowly, but still wants to finish eating all the bananas before the guards come back.
#
# Return the minimum integer K such that she can eat all the bananas within H hours.
#
#
#
# Example 1:
#
# Input: piles = [3,6,7,11], H = 8
# Output: 4
#
# Example 2:
#
# Input: piles = [30,11,23,4,20], H = 5
# Output: 30
#
# Example 3:
#
# Input: piles = [30,11,23,4,20], H = 6
# Output: 23
#
#
#
# Constraints:
#
#     1 <= piles.length <= 10^4
#     piles.length <= H <= 10^9
#     1 <= piles[i] <= 10^9

# @param {Integer[]} piles
# @param {Integer} h
# @return {Integer}
def min_eating_speed(piles, h)
  total, high = 0, 0
  piles.each do |p|
    total += p
    high = max(high, p)
  end

  low = (total.to_f / h).ceil
  ans = 0

  while low <= high
    mid = low + (high-low)/2

    if finish_in_time?(piles, mid, h)
      ans = mid
      high = mid-1
    else
      low = mid+1
    end
  end

  ans
end

def finish_in_time?(piles, speed, h)
  hour = 0

  piles.each { |p| hour += (p.to_f/speed).ceil }

  return hour <= h
end

def max(i, j )
  return i if i >= j
  j
end

#   Notes
#   1.  inspired from sampe code, to find ceiling of division, another way is
#       to use ((num - 1) / k) + 1