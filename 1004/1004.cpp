// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.



// Example 1:

// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
// Explanation: [1,1,1,0,0,1,1,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

// Example 2:

// Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
// Output: 10
// Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.



// Constraints:

//     1 <= nums.length <= 105
//     nums[i] is either 0 or 1.
//     0 <= k <= nums.length

#include <vector>
using namespace std;

class Solution {
public:
    int longestOnes(vector<int>& nums, int k) {
        int i = 0, j = 0, longest = 0, n = nums.size();

        for (; j < n; ++j) {
          for (; j < n; ++j) {
            if (nums[j] == 0) {
              --k;

              if (k < 0) {
                break;
              }
            }
          }

          longest = max(longest, j-i);

          // forward i either to j or next 0
          for (; i < j && nums[i] == 1; ++i) {}

          ++k;
          ++i;
        }

        return longest;
    }
};

//  Notes
//  1.  inspired from sample code, use int a = 0, b = 0
