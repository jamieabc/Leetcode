package main

import "sort"

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

// tc: O(n*2^n)
func makesquare(sticks []int) bool {
    n := len(sticks)

    // find length
    var sum int
    for _, i := range sticks {
        sum += i
    }

    if sum % 4 != 0 {
        return false
    }

    length := sum / 4

    var largest int
    table := make(map[int]int)
    for i := 0; i < (1 << n); i++ {
        sum = 0

        for j := 0; j < n; j++ {
            if (1 << (n-j-1)) & i > 0 {
                sum += sticks[j]
            }
        }

        if sum == length {
            largest = 1

            for k, v := range table {
                if k & i == 0 {
                    newVal := k | i

                    if val, ok := table[newVal]; ok {
                        table[newVal] = max(val, v+1)
                    } else {
                        table[newVal] = v+1
                    }

                    largest = max(largest, table[newVal])
                }
            }

            table[i] = 1
        }

        // since sum % 4 == 0, and 3 sides are found, then it's okay, no need to keep finding
        if largest == 3 {
            return true
        }
    }

    return false
}

func max(i, j int) int {
    if i >= j {
        return i
    }
    return j
}

// tc: O(n * 2^n), there are total n combinations (n bits), each time a group
// of nubmers selected, takes O(n) to check if group sum equals target length
func makesquare4(sticks []int) bool {
    n := len(sticks)

    // find length
    var sum int
    for _, i := range sticks {
        sum += i
    }

    if sum % 4 != 0 {
        return false
    }

    length := sum / 4
    table := make(map[int]bool)
    all := (1 << n) - 1
    for i := 0; i <= all; i++ {
        sum = 0

        for j := 0; j < n; j++ {
            if (1 << (n-j-1)) & i > 0 {
                sum += sticks[j]
            }
        }

        // found one side
        if sum == length {
            for k := range table {
                if k & i == 0 {
                    nextVal := k | i
                    table[nextVal] = true

                    if table[all ^ nextVal] {
                        return true
                    }
                }
            }
            table[i] = true
        }
    }

    return false
}

// tc: O(4^n), with some pruning
func makesquare3(sticks []int) bool {
    var sum int
    for _, i := range sticks {
        sum += i
    }

    // cannot equally divided to 4 sides
    if sum % 4 != 0 {
        return false
    }

    // start from longer sticks, reduce possible paths
    sort.Slice(sticks, func(i, j int) bool {
        return sticks[i] > sticks[j]
    })

    sides := make([]int, 4)

    return dfs3(sticks, sides, 0, sum/4)
}

func dfs3(sticks, groups []int, idx, length int) bool {
    if idx == len(sticks) {
        return groups[0] == groups[1] && groups[1] == groups[2] && groups[2] == groups[3]
    }

    for i := 0; i < 4; i++ {
        if groups[i]+sticks[idx] <= length {
            groups[i] += sticks[idx]

            if dfs3(sticks, groups, idx+1, length) {
                return true
            }

            groups[i] -= sticks[idx]
        }
    }

    return false
}

// tc: O(4^n)
func makesquare2(matchsticks []int) bool {
    var sum int
    for _, i := range matchsticks {
        sum += i
    }

    // cannot equally divided to 4 sides
    if sum % 4 != 0 {
        return false
    }

    n := len(matchsticks)
    groups := make([]int, n)
    for i := range groups {
        groups[i] = -1
    }

    return dfs2(matchsticks, groups, 0, sum/4)
}

func dfs2(sticks, groups []int, idx, length int) bool {
    if idx == len(sticks) {
        // check each group
        var sum int
        for i := 0; i < 4; i++ {
            sum = 0
            for j := range groups {
                if groups[j] == i {
                    sum += sticks[j]
                }
            }

            if sum != length {
                return false
            }
        }

        return true
    }

    for j := 0; j < 4; j++ {
        groups[idx] = j

        if dfs2(sticks, groups, idx+1, length) {
            return true
        }
    }

    return false
}

// tc: O(n!), n: sticks length
func makesquare1(matchsticks []int) bool {
    var sum int
    for _, i := range matchsticks {
        sum += i
    }

    // not able to be divided by 4, must not form a square
    if sum % 4 != 0 {
        return false
    }

    // start from largest sticks, reduce possible combinations
    sort.Slice(matchsticks, func(i, j int) bool {
        return matchsticks[i] > matchsticks[j]
    })


    n := len(matchsticks)
    used := make([]bool, n)
    return dfs1(matchsticks, sum/4, 0, 4, used, sum)
}

func dfs1(sticks []int, target, current, remain int, used []bool, sum int) bool {
    if remain == 0 {
        return true
    }

    if remain == 1 {
        return sum == target
    }

    var result bool

    for i := range used {
        if used[i] {
            continue
        }

        if current + sticks[i] <= target {
            next := current + sticks[i]
            used[i] = true

            if next == target {
                result = result || dfs1(sticks, target, 0, remain-1, used, sum - sticks[i])
            } else {
                result = result || dfs1(sticks, target, next, remain, used, sum - sticks[i])
            }

            used[i] = false
        }
    }

    return result
}

//  Notes
//  1.  use backtracing, TLE

//  2.  inspired from solution, use backtracking with tc O(n!), but it's really back
//
//      view it from grouping, 4 groups with same sum, then each number with 4 choices, tc O(4^n)

//      it's possible to track only 4 sides of square, consider on the total length

//  3.  inspired form solution, if a number is selected, what are other numbers that sum of group of
//      numbers equals expected length?
//
//      the problem doesn't care about order of groups, e.g. it doesn't matter groups in order of
//      [1, 2, 3, 4] or [4, 1, 3, 2], only thing matters is what group of numbers can split original
//      array into 4 groups
//
//      the recurring problem is that if a number is selected, what are other group of numbers that
//      meets expected sum?
//
//      since it's group of numbers, and n <= 15, it's okay to use single int32 to represent group of
//      selected numbers
//
//      e.g. [1, 2, 3, 4]
//      int   0  1  0  1 => 2 & 4 are selected
//
//      the benefit of using int32 is to quickly check group of numbers with overlap by & operator
//
//      e.g.  [0 1 0 1] & [1 0 1 0] = 0
//
//      the other technique is that since total sum can be fully divided by 4, only need to find 3 groups
//      that sum to target length, then it's done

//  4.  inspired from https://leetcode.com/problems/matchsticks-to-square/discuss/95762/Two-different-solutions
//
//      author provides a brilliant way to optimize dp
//
//      similar to two-sum, if all sticks are used, use binary format to represent 111...1
//
//      separate sticks into two half, if all of them are disjoint set a and b, then a & b = 0, which
//      is easy to check
//
//      all = 111...1, all ^ (a | b) to find if other half of group is found, this is similar to two-sum
