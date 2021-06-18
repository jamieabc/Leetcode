// You are given an integer array matchsticks where matchsticks[i] is the length of the ith matchstick. You want to use all the matchsticks to make one square. You should not break any stick, but you can link them up, and each matchstick must be used exactly one time.

// Return true if you can make this square and false otherwise.



// Example 1:

// Input: matchsticks = [1,1,2,2,2]
// Output: true
// Explanation: You can form a square with length 2, one side of the square came two sticks with length 1.

// Example 2:

// Input: matchsticks = [3,3,3,3,4]
// Output: false
// Explanation: You cannot find a way to form a square with all the matchsticks.



// Constraints:

//     1 <= matchsticks.length <= 15
//     0 <= matchsticks[i] <= 109

#include <vector>
#include <algorithm>
#include <numeric>
using namespace std;

class Solution {
public:
  // tc: O(4^n)
  bool makesquare(vector<int> &sticks) {
    int n = sticks.size();
    int sum = accumulate(sticks.begin(), sticks.end(), 0);

    if (sum % 4 || !sum) {
      return false;
    }
    int length = sum / 4;

    sort(sticks.begin(), sticks.end(), [](const int &a, const int &b) { return a > b; });

    if (sticks[0] > length) {
      return false;
    }

    int sides[4] = {0};

    return dfs(sticks, sides, length, 0);
  }

  bool dfs(vector<int> &sticks, int sides[], const int &length, const int idx) {
    if (idx == sticks.size()) {
      return sides[0] == sides[1] && sides[1] == sides[2] &&
             sides[2] == sides[3];
    }

    for (int i = 0; i < 4; ++i) {
      if (sides[i] + sticks[idx] <= length) {
        sides[i] += sticks[idx];

        if (dfs(sticks, sides, length, idx + 1)) {
          return true;
        }

        sides[i] -= sticks[idx];
      }
    }

    return false;
  }
};

//  Notes
//  1.  inspired from sample code
//
//      can use int x = accumulate();
//
//      ++i
//
//      passing int array to function is denoted as "int array[]"
