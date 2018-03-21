#  There are 1000 buckets, one and only one of them contains poison, the rest are filled with water. They all look the same. If a pig drinks that poison it will die within 15 minutes. What is the minimum amount of pigs you need to figure out which bucket contains the poison within one hour.

# Answer this question, and write an algorithm for the follow-up general case.

# Follow-up:

# If there are n buckets and a pig drinking poison will die within m minutes, how many pigs (x) you need to figure out the "poison" bucket within p minutes? There is exact one bucket with poison.

# Each pig can try at most 4 times (60/15)
# but to remember additional part no need to test
# (if all pigs safe, then parts left should include poison )

# @param {Integer} buckets
# @param {Integer} minutes_to_die
# @param {Integer} minutes_to_test
# @return {Integer}
def poor_pigs(buckets, minutes_to_die, minutes_to_test)
  times_to_try = minutes_to_test / minutes_to_die
  return 0 if buckets == 1
  Math.log(buckets,  1 + times_to_try).ceil
end

p poor_pigs(1000, 15, 60)