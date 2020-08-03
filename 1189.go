package main

// Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

// You can use each character in text at most once. Return the maximum number of instances that can be formed.



// Example 1:

// Input: text = "nlaebolko"
// Output: 1

// Example 2:

// Input: text = "loonbalxballpoon"
// Output: 2

// Example 3:

// Input: text = "leetcode"
// Output: 0



// Constraints:

//     1 <= text.length <= 10^4
//     text consists of lower case English letters only.

func maxNumberOfBalloons(text string) int {
    counter := make([]int, 26)
    for i := range text {
        counter[text[i]-'a']++
    }

    // balloon, b: 1, a: 1, l: 2, o: 2, n: 1
    factor := math.MaxInt32

    factor = min(factor, counter[0]) // a
    factor = min(factor, counter[1]) // b
    factor = min(factor, counter['l'-'a']/2) // l
    factor = min(factor, counter['o'-'a']/2) // o
    factor = min(factor, counter['n'-'a']) // n

    if factor == math.MaxInt32 {
        return 0
    }

    return factor
}

func min(i, j int) int {
    if i <= j {
        return i
    }
    return j
}
