// In this problem, a tree is an undirected graph that is connected and has no cycles.

// You are given a graph that started as a tree with n nodes labeled from 1 to n, with one additional edge added. The added edge has two different vertices chosen from 1 to n, and was not an edge that already existed. The graph is represented as an array edges of length n where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the graph.

// Return an edge that can be removed so that the resulting graph is a tree of n nodes. If there are multiple answers, return the answer that occurs last in the input.



// Example 1:

// Input: edges = [[1,2],[1,3],[2,3]]
// Output: [2,3]

// Example 2:

// Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
// Output: [1,4]



// Constraints:

//     n == edges.length
//     3 <= n <= 1000
//     edges[i].length == 2
//     1 <= ai < bi <= edges.length
//     ai != bi
//     There are no repeated edges.
//     The given graph is connected.

#include <vector>
using namespace std;

class Solution {
private:
    class UnionFind {
        private:
            vector<int> group;
            vector<int> rank;

        public:
            UnionFind(int n) {
                group = vector<int>(n);
                rank = vector<int>(n);

                for (int i = 0; i < group.size(); ++i) {
                    group[i] = i;
                    rank[i] = 1;
                }
            }

            int find(int index) {
                if (group[index] != index) {
                    group[index] = find(group[index]);
                }

                return group[index];
            }

            bool unionNode(int &index1, int &index2) {
                int p1 = find(index1);
                int p2 = find(index2);

                if (p1 == p2) {
                    return false;
                }

                if (rank[p1] >= rank[p2]) {
                    group[p2] = p1;
                    rank[p1] += rank[p2];
                } else {
                    group[p1] = p2;
                    rank[p2] += rank[p1];
                }

                return true;
            }
    };

public:
    vector<int> findRedundantConnection(vector<vector<int>>& edges) {
        int n = edges.size();

        UnionFind uf(n+1);

        for (auto &e : edges) {

            if (!uf.unionNode(e[0], e[1])) {
                return e;
            }
        }

        return vector<int>(2);
    }
};

//  Notes
//  1.  inspired from sample code, use private class UnionFind

//  2.  beware of how to use constructor for variable declaration
//
//  3.  problem wants to find last redundant edge, initially i thought keep union nodes till last
//      node is found, but it's not necessary
//
//      a - b - c, each node with distinct parent
//
//      a - b - c, a & c with same parent, this is redundant
//      |       |
//      ---------
//
//      because there's only one redundant edge, the one found should be the answer

//  4.  inspired from sample code, since input range only from 3 ~ 1000, could create array for faster
//      initialization

//  5.  inspired from https://leetcode.com/problems/redundant-connection/discuss/107990/Python-52ms-DFS-Algorithm-with-Explanations-(compared-with-52ms-union-find)
//
//      dfs can be used to traverse, the concept is that for each edge, check if start from edge[0]
//      can reach edge[1], takes O(n), and there are n edges overall tc: O(n^2)
