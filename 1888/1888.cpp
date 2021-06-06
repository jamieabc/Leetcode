// You are given a binary string s. You are allowed to perform two types of operations on the string in any sequence:

//     Type-1: Remove the character at the start of the string s and append it to the end of the string.
//     Type-2: Pick any character in s and flip its value, i.e., if its value is '0' it becomes '1' and vice-versa.

// Return the minimum number of type-2 operations you need to perform such that s becomes alternating.

// The string is called alternating if no two adjacent characters are equal.

//     For example, the strings "010" and "1010" are alternating, while the string "0100" is not.



// Example 1:

// Input: s = "111000"
// Output: 2
// Explanation: Use the first operation two times to make s = "100011".
// Then, use the second operation on the third and sixth elements to make s = "101010".

// Example 2:

// Input: s = "010"
// Output: 0
// Explanation: The string is already alternating.

// Example 3:

// Input: s = "1110"
// Output: 1
// Explanation: Use the second operation on the second element to make s = "1010".



// Constraints:

//     1 <= s.length <= 105
//     s[i] is either '0' or '1'.

#include <string>
#include <limits.h>
using namespace std;

class Solution {
public:
    int minFlips(string s) {
        int size = s.size();
        s += s;

        int flip1 = 0, flip2 = 0;
        int min_count = INT_MAX;

        for (int i = 0; i < s.size(); i++) {
            if (((i & 1) == 0 && s[i] == '1') || ((i & 1) == 1 && s[i] == '0')) {
                ++flip1;
            }

            if (((i & 1) == 0 && s[i] == '0') || ((i & 1) == 1 && s[i] == '1')) {
                ++flip2;
            }

            if (i >= size-1) {
                if (i >= size) {
                  if ((((i - size) & 1) == 0 && s[i - size] == '1') || (((i - size) & 1) == 1 && s[i - size] == '0')) {
                      --flip1;
                    }

                    if ((((i - size) & 1) == 0 && s[i - size] == '0') || (((i - size) & 1) == 1 && s[i - size] == '1')) {
                        --flip2;
                    }
                }
                min_count = min(min_count, min(flip1, flip2));
            }
        }

        return min_count;
    }
};

//  Notes
//  1.  inspired from https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/discuss/1254148/Sliding-Window
//
//      author provides very elegant code

//  2.  inspired from https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/discuss/1254148/Sliding-Window/963721
//
//      author provides another way of code
