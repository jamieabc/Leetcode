// You are given a 0-indexed integer array nums and an integer k.

// You are initially standing at index 0. In one move, you can jump at most k steps forward without going outside the boundaries of the array. That is, you can jump from index i to any index in the range [i + 1, min(n - 1, i + k)] inclusive.

// You want to reach the last index of the array (index n - 1). Your score is the sum of all nums[j] for each index j you visited in the array.

// Return the maximum score you can get.



// Example 1:

// Input: nums = [1,-1,-2,4,-7,3], k = 2
// Output: 7
// Explanation: You can choose your jumps forming the subsequence [1,-1,4,3] (underlined above). The sum is 7.

// Example 2:

// Input: nums = [10,-5,-2,4,0,3], k = 3
// Output: 17
// Explanation: You can choose your jumps forming the subsequence [10,4,3] (underlined above). The sum is 17.

// Example 3:

// Input: nums = [1,-5,-20,4,-1,3,-6,-3], k = 2
// Output: 0



// Constraints:

//      1 <= nums.length, k <= 105
//     -104 <= nums[i] <= 104

#include <vector>
#include <deque>
using namespace std;

class Solution {
public:
  int maxResult(vector<int> &nums, int k) {
    int n = nums.size();
    vector<int> dp(n);
    dp[0] = nums[0];
    deque<int> deque;
    deque.push_back(0);

    for (int i = 1; i < nums.size(); i++) {
      if (deque.front() < i - k) {
        deque.pop_front();
      }

      dp[i] = nums[i] + dp[deque.front()];

      while (!deque.empty() && dp[deque.back()] < dp[i]) {
        deque.pop_back();
      }

      deque.push_back(i);
    }

    return dp[n - 1];
  }

  int maxResult1(vector<int> &nums, int k) {
    deque<pair<int, int>> deque;
    deque.push_back(make_pair(nums[0], 0));

    for (int i = 1; i < nums.size(); i++) {
      // remove out of range number
      if (deque.front().second < i - k) {
        deque.pop_front();
      }

      pair<int, int> next(deque.front().first + nums[i], i);

      // keep deque in descending order
      while (deque.size() > 0 && deque.back().first < next.first) {
        deque.pop_back();
      }

      deque.push_back(next);
    }

    return deque.back().first;
    }
};

//  Notes
//  1.  inspired from sample code, the fast way is to use deque & dp
//      i guess it's because it's not necessary to allocate/remove pair
