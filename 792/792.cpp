// Given a string s and an array of strings words, return the number of words[i] that is a subsequence of s.

// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

//     For example, "ace" is a subsequence of "abcde".



// Example 1:

// Input: s = "abcde", words = ["a","bb","acd","ace"]
// Output: 3
// Explanation: There are three strings in words that are a subsequence of s: "a", "acd", "ace".

// Example 2:

// Input: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
// Output: 2



// Constraints:

//     1 <= s.length <= 5 * 104
//     1 <= words.length <= 5000
//     1 <= words[i].length <= 50
//     s and words[i] consist of only lowercase English letters.

#include <vector>
#include <string>
using namespace std;

class Solution {
public:

    // tc: O(n+k), k: sum of each word length
    int numMatchingSubseq(string s, vector<string>& words) {
        vector<const char*> table[128];     // a ~ z = 97 ~ 122

        // put each word into table
        for (auto &w : words) {
            table[w[0]].push_back(w.c_str());
        }

        for (char &c : s) {
            auto array = table[c];
            table[c].clear();

            for (auto &it : array) {
                // forward one character
                table[*++it].push_back(it);
            }
        }

        return table[0].size();
    }

    // tc: O(nm), sc: O(m)
    int numMatchingSubseq2(string s, vector<string>& words) {
        int n = words.size();
        int count = 0;
        vector<int> indexes(n, 0);

        for (auto& c : s) {
            for (int i = 0; i < words.size(); ++i) {
                if (indexes[i] < words[i].size() && c == words[i][indexes[i]]) {
                    ++indexes[i];

                    if (indexes[i] == words[i].size()) {
                        ++count;
                    }
                }
            }
        }

        return count;
    }

    // tc: O(26*n+m), n: s.size(), m: average size of word
    // sc: O(26*n)
    int numMatchingSubseq1(string s, vector<string>& words) {
        int n = s.size();
        vector<vector<int>> table(26, vector<int>(n+1, -1));

        int last;
        for (int i = 0; i < 26; ++i) {
            last = -1;

            for (int j = n-1; j >= 0; --j) {
                if (int(s[j]-'a') == i) {
                    last = j;
                }
                table[i][j] = last;
            }
        }

        int sub_sequence = 0;
        int index;

        for (int i = 0; i < words.size(); ++i) {
            last = index = 0;

            for (; index < words[i].size(); ++index, ++last) {
                last = table[int(words[i][index]-'a')][last];

                if (last == -1) {
                    break;
                }
            }

            if (index == words[i].size()) {
                ++sub_sequence;
            }
        }

        return sub_sequence;
    }
};

//  Notes
//  1.  prebuilt table takes O(26*n) to generate table that knows what's the next index of
//      specific character
//
//      e.g. "abcde"
//      a: 0, -1, -1, -1, -1
//      b: 1,  1, -1, -1, -1
//      c: 2,  2,  2, -1, -1
//      d: 3,  3,  3,  3, -1
//      e: 4,  4,  4,  4,  4

//      traverse table for each test word to find if it's a substring of original word
//      tc: O(m), m: average size of word
//
//      overall tc: O(26*n+m), sc: O(26*n)

//  2.  inspired from https://leetcode.com/problems/number-of-matching-subsequences/discuss/117634/Efficient-and-simple-go-through-words-in-parallel-with-explanation
//
//      no need to build whole table, actually, traverse from s once and update each word to find word that
//      is not sub-string
//
//      e.g. s = "abcde", words = ["a", "bb", "acd", "ace"]
//
//      s[0] = a, words = ["", "bb", "cd", "ce"], because a exist, so update all words that starts with a
//      s[1] = b, words = ["", "b", "cd", "ce"]
//      s[2] = c, words = ["", "b", "d", "e"]
//      s[3] = d, words = ["", "b", "", "e"]
//      s[4] = e, words = ["", "b", "", ""]
//
//      there are 3 empty strings, which means 3 sub-strings
//
//      this is very clever...
//
//      borrowing this idea, use index array to store next char position
//      tc: O(nm), n: s.size(), m: words.size()
//
//      it's very slow, author uses hash table to avoid whole words scan
