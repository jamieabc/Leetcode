# Given an array of integers nums, sort the array in ascending order.
#
#
#
#   Example 1:
#
#   Input: nums = [5,2,3,1]
# Output: [1,2,3,5]
# Example 2:
#
#   Input: nums = [5,1,1,2,0,0]
# Output: [0,0,1,1,2,5]
#
#
# Constraints:
#
#   1 <= nums.length <= 50000
# -50000 <= nums[i] <= 50000

# @param {Integer[]} nums
# @return {Integer[]}

def sort_array(nums)
  quick_sort(nums, 0, nums.size-1)
  nums
end

def quick_sort(nums, from, to)
  if from >= to
    return
  end

  idx = partition(nums, from, to)

  if idx != -1
    quick_sort(nums, from, idx-1)
    quick_sort(nums, idx+1, to)
  end
end

def partition(nums, from, to)
  store, pivot = from, nums[from]
  nums[from], nums[to] = nums[to], nums[from]

  i = from
  while i < to
    if nums[i] < pivot
      nums[i], nums[store] = nums[store], nums[i]
      store += 1
    end

    i += 1
  end

  nums[store], nums[to] = nums[to], nums[store]

  store
end

def sort_array1(nums)
  if nums.size <= 1
    return nums
  end

  mid = nums.size >> 1
  left, right = sort_array1(nums[0..mid-1]), sort_array1(nums[mid..-1])

  merge(left, right)
end

def merge(arr1, arr2)
  ans = Array.new(arr1.size + arr2.size)

  idx, p1, p2 = 0, 0, 0

  while idx < ans.size
    if p2 == arr2.size || (p1 < arr1.size && arr1[p1] <= arr2[p2])
      ans[idx] = arr1[p1]
      p1 += 1
    else
      ans[idx] = arr2[p2]
      p2 += 1
    end

    idx += 1
  end

  ans
end

#   Notes
#   1.  inspired from sample code
#       w/o index first/last, can use array to check
#       use left/right to denote partitioned array
#       use array[-1] to denote last number
#
#       becareful about array range, it's both index are inclusion