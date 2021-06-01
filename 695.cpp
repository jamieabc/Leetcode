class Solution {
public:
    int maxAreaOfIsland(vector<vector<int>>& grid) {
        int w = grid[0].size();
        int h = grid.size();
        vector<vector<bool>> visited(h, vector<bool>(w, false));

        int maxArea = 0;

        for (int i = 0; i < h; i++) {
            for (int j = 0; j < w; j++) {
                if (grid[i][j] == 1 && !visited[i][j]) {
                    visited[i][j] = true;
                    maxArea = max(maxArea, 1+dfs(grid, visited, i, j));
                }
            }
        }

        return maxArea;
    }

    vector<vector<int>> dirs = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};

    int dfs(vector<vector<int>>& grid, vector<vector<bool>>& visited, int& y, int& x) {
        int cur = 0;
        int w = grid[0].size();
        int h = grid.size();

        for (int i = 0; i < dirs.size(); i++) {
            int newY = dirs[i][0]+y;
            int newX = dirs[i][1]+x;

            if (newX >= 0 && newY >= 0 && newX < w && newY < h && grid[newY][newX] == 1 && !visited[newY][newX]) {
                visited[newY][newX] = true;
                cur += 1+dfs(grid, visited, newY, newX);
            }
        }

        return cur;
    }
};