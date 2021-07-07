// Given an n x n matrix where each of the rows and columns are sorted in ascending order, return the kth smallest element in the matrix.

// Note that it is the kth smallest element in the sorted order, not the kth distinct element.



// Example 1:

// Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
// Output: 13
// Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13

// Example 2:

// Input: matrix = [[-5]], k = 1
// Output: -5



// Constraints:

//     n == matrix.length
//     n == matrix[i].length
//     1 <= n <= 300
//     -109 <= matrix[i][j] <= 109
//     All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
//     1 <= k <= n2

#include <vector>
#include <queue>
using namespace std;

struct compare1 {
public:
    bool operator()(vector<int> &a, vector<int> &b) {
        return a[0] > b[0];
    }
};

#define pair2 pair<int, pair<int, int>>
class Solution {
public:

    // tc: O(mn log(k))
    int kthSmallest(vector<vector<int>>& matrix, int k) {
        priority_queue<int> pq;

        for (auto &i : matrix) {
            for (auto &j : i) {
                pq.push(j);
                if (pq.size() > k) {
                    pq.pop();
                }
            }
        }

        return pq.top();
    }

    // tc: O(k log(n)), n: height of matrix
    int kthSmallest2(vector<vector<int>>& matrix, int k) {
        priority_queue<pair2, vector<pair2>, greater<pair2>> pq;
        int h = matrix.size();
        int w = matrix[0].size();

        for (int i = 0; i < h; ++i) {
            pq.push({matrix[i][0], {i, 0}});
        }

        int ans = 0;
        while (!pq.empty() && k-- > 0) {
            auto cur = pq.top();
            pq.pop();

            ans = cur.first;
            int row = cur.second.first;
            int col = cur.second.second;

            if (cur.second.second < w-1) {
                pq.push({matrix[row][col+1], {row, col+1}});
            }
        }

        return ans;
    }

    int kthSmallest1(vector<vector<int>>& matrix, int k) {
        priority_queue<vector<int>, vector<vector<int>>, compare1> pq;
        int w = matrix[0].size();
        int h = matrix.size();

        // push into priority queue
        for (int i = 0; i < h; ++i) {
            pq.push(vector<int>{matrix[i][0], i, 0});
        }

        int ans = 0;

        for (int i = 0; i < k; ++i) {
            auto vec = pq.top();
            pq.pop();

            ans = vec[0];
            if (vec[2] < w-1) {
                ++vec[2];
                vec[0] = matrix[vec[1]][vec[2]];
                pq.push(vec);
            }
        }

        return ans;
    }
};

//  Notes
//  1.  first version of implementation is very slow

//  2.  inspired from sample code, use different type of priority_queue<int, pair<int, int>>
//      also, pair can be initialized by {}
//
//      it's faster due to avoid using vector<vector<int>>

//  3.  inspired form sample code, could use priority queue (top is largest) to remove number not in
//      kth range
