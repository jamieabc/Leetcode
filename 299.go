package main

import "fmt"

//You are playing the following Bulls and Cows game with your friend: You write down a number and ask your friend to guess what the number is. Each time your friend makes a guess, you provide a hint that indicates how many digits in said guess match your secret number exactly in both digit and position (called "bulls") and how many digits match the secret number but locate in the wrong position (called "cows"). Your friend will use successive guesses and hints to eventually derive the secret number.
//
//Write a function to return a hint according to the secret number and friend's guess, use A to indicate the bulls and B to indicate the cows.
//
//Please note that both secret number and friend's guess may contain duplicate digits.
//
//Example 1:
//
//Input: secret = "1807", guess = "7810"
//
//Output: "1A3B"
//
//Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.
//
//Example 2:
//
//Input: secret = "1123", guess = "0111"
//
//Output: "1A1B"
//
//Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a cow.
//
//Note: You may assume that the secret number and your friend's guess only contain digits, and their lengths are always equal.

func getHint(secret string, guess string) string {
	if secret == guess {
		length := len(secret)
		return fmt.Sprintf("%dA0B", length)
	}

	var countA, countB int
	var s, g uint8
	count := make([]int, 10)

	for i := range secret {
		if secret[i] == guess[i] {
			countA++
			continue
		}

		s = secret[i] - '0'
		g = guess[i] - '0'

		if count[s] < 0 {
			countB++
		}

		if count[g] > 0 {
			countB++
		}

		count[s]++
		count[g]--
	}

	return fmt.Sprintf("%dA%dB", countA, countB)
}

// problems
// 1. length is not always 4
// 2. remove delete, try to speed up
// 3. remove some procedure to see if it's faster
// 4. wrong scope
// 5. not wrong scope, it actually wrong logic, each comparision can only add by certain condition
// 6. discard that complicate compare both side and try to keep map in sync
// 7. forget to add element to remain slice
// 8. wrong position of putting remain slice
// 9. wrong condition to add B
// 10. using single loop to check for existence, from secret +1, from guess -1
// 11. further improvement, use array instead of map
