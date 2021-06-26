// Given the edges of a directed graph where edges[i] = [ai, bi] indicates there is an edge between nodes ai and bi, and two nodes source and destination of this graph, determine whether or not all paths starting from source eventually, end at destination, that is:

//     At least one path exists from the source node to the destination node
//     If a path exists from the source node to a node with no outgoing edges, then that node is equal to destination.
//     The number of possible paths from source to destination is a finite number.

// Return true if and only if all roads from source lead to destination.



// Example 1:

// Input: n = 3, edges = [[0,1],[0,2]], source = 0, destination = 2
// Output: false
// Explanation: It is possible to reach and get stuck on both node 1 and node 2.

// Example 2:

// Input: n = 4, edges = [[0,1],[0,3],[1,2],[2,1]], source = 0, destination = 3
// Output: false
// Explanation: We have two possibilities: to end at node 3, or to loop over node 1 and node 2 indefinitely.

// Example 3:

// Input: n = 4, edges = [[0,1],[0,2],[1,3],[2,3]], source = 0, destination = 3
// Output: true

// Example 4:

// Input: n = 3, edges = [[0,1],[1,1],[1,2]], source = 0, destination = 2
// Output: false
// Explanation: All paths from the source node end at the destination node, but there are an infinite number of paths, such as 0-1-2, 0-1-1-2, 0-1-1-1-2, 0-1-1-1-1-2, and so on.

// Example 5:

// Input: n = 2, edges = [[0,1],[1,1]], source = 0, destination = 1
// Output: false
// Explanation: There is infinite self-loop at destination node.



// Constraints:

//     1 <= n <= 104
//     0 <= edges.length <= 104
//     edges.length == 2
//     0 <= ai, bi <= n - 1
//     0 <= source <= n - 1
//     0 <= destination <= n - 1
//     The given graph may have self-loops and parallel edges.

#include <vector>
#include <queue>
#include <unordered_set>
using namespace std;

class Solution {
public:

    bool leadsToDestination(int n, vector<vector<int>>& edges, int source, int destination) {
        // 4 states: 0 not visited, 1 visiting, 2 can reach destination, 3 cannot reach destination
        vector<int> table(n, 0);

        vector<vector<int>> adj_list(n);
        for (const auto &e : edges) {
            adj_list[e[0]].emplace_back(e[1]);
        }

        // destination node can reach other node, has outgoing
        if (!adj_list[destination].empty()) {
            return false;
        }

        return dfs(adj_list, table, source, destination);
    }

    bool dfs(vector<vector<int>> &adj_list, vector<int> &table, int cur, int dest) {
        if (cur == dest || table[cur] == 2) {
            return true;
        }

        // if a node cannot go further and not dest, condition not meet
        if (adj_list[cur].empty() || table[cur] == 1 || table[cur] == 3) {
            return false;
        }

        // mark current node as visiting
        table[cur] = 1;

        // keep traversing
        for (const auto &to : adj_list[cur]) {
            if (!dfs(adj_list, table, to, dest)) {
                    table[to] = 3;
                    return false;
            }
        }

        table[cur] = 2;
        return true;
    }

    bool leadsToDestination1(int n, vector<vector<int>>& edges, int source, int destination) {
        vector<bool> dp(n, false);
        unordered_set<int> set;
        vector<vector<int>> adj_list(n);

        for (const auto &e : edges) {
            adj_list[e[0]].push_back(e[1]);
        }

        if (adj_list[destination].size() != 0) {
            return false;
        }

        return dfs1(adj_list, dp, set, source, destination);
    }

private:
    bool dfs1(vector<vector<int>> &adj_list, vector<bool> &dp, unordered_set<int> &set, int from, int dest) {
        if (from == dest) {
            return true;
        }

        if (adj_list[from].size() == 0) {
            return false;
        }

        for (const auto &to : adj_list[from]) {
            // to is already checked okay
            if (dp[to]) {
                continue;
            }

            // cycle
            if (set.find(to) != set.end()) {
                return false;
            }
            set.insert(to);

            dp[to] = dfs(adj_list, dp, set, to, dest);

            if (!dp[to]) {
                return false;
            }

            set.erase(to);
        }

        return true;
    }
};

//  Notes
//  1.  check for cycle but not for visited, because visited node could still be a valid one, but go to
//      a node on current path must be invalid

//  2.  inspired from sample code, because n is already known, could put all states into vector
//
//      0: not visited
//      1: visiting
//      2: can go to destination

//      return false immediately if cannot reach destination
//
//      use vector.emplace_back() for faster inserting integer
//      use vector.empty() to check empty

//  3.  this problem is about checking cycle exists
