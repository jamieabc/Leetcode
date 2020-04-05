package main

import (
	"sort"
)

//In a deck of cards, every card has a unique integer.  You can order the deck in any order you want.
//
//Initially, all the cards start face down (unrevealed) in one deck.
//
//Now, you do the following steps repeatedly, until all cards are revealed:
//
//Take the top card of the deck, reveal it, and take it out of the deck.
//If there are still cards in the deck, put the next top card of the deck at the bottom of the deck.
//If there are still unrevealed cards, go back to step 1.  Otherwise, stop.
//Return an ordering of the deck that would reveal the cards in increasing order.
//
//The first entry in the answer is considered to be the top of the deck.
//
//
//
//Example 1:
//
//Input: [17,13,11,2,3,5,7]
//Output: [2,13,3,11,5,17,7]
//Explanation:
//We get the deck in the order [17,13,11,2,3,5,7] (this order doesn't matter), and reorder it.
//After reordering, the deck starts as [2,13,3,11,5,17,7], where 2 is the top of the deck.
//We reveal 2, and move 13 to the bottom.  The deck is now [3,11,5,17,7,13].
//We reveal 3, and move 11 to the bottom.  The deck is now [5,17,7,13,11].
//We reveal 5, and move 17 to the bottom.  The deck is now [7,13,11,17].
//We reveal 7, and move 13 to the bottom.  The deck is now [11,17,13].
//We reveal 11, and move 17 to the bottom.  The deck is now [13,17].
//We reveal 13, and move 17 to the bottom.  The deck is now [17].
//We reveal 17.
//Since all the cards revealed are in increasing order, the answer is correct.
//
//
//Note:
//
//1 <= A.length <= 1000
//1 <= A[i] <= 10^6
//A[i] != A[j] for all i != j

func deckRevealedIncreasing(deck []int) []int {
	length := len(deck)
	sort.Ints(deck)

	queue := make([]int, length)
	for i := range deck {
		queue[i] = i
	}

	result := make([]int, length)

	for idx := 0; len(queue) != 0; idx++ {
		result[queue[0]] = deck[idx]
		queue = queue[1:]

		if len(queue) > 1 {
			tmp := queue[0]
			queue = queue[1:]
			queue = append(queue, tmp)
		}
	}

	return result
}

//	problems
//	1.	last shuffle is done only if it's odd number & length enough
//	2.	last shuffle is needed only when initial cards number is odd
//	3. 	the way choosing cards, it acts like a queue. And I don't need another
//		mapping, just use sorted number to store final result.

//		Simplest way is to use queue, pop first item, then put next top item
//		to end of queue, repeat until queue is empty. With this way, no extra
//		memory allocation is needed
//	4.	extra memory to store index order is not needed, whenever queue pop
//		first item, that's the index order.

//	What I learn from this? The behavior of always choosing first card is
//	actually a queue, and second behavior is putting top item back queue.

//	I didn't aware of this, that's I should improve next time. The nature of
//	queue is always choosing first item, keep that in mind.

//	Second problem I have, is complicate the problem. What I need is the order
//	of index, so that I can generate answer by that order, so I have a quite
//	straight forward method, finding index, store it, then remap it back to
//	answer. But the thing is, output from queue is actually the order, so I
//	can directly have it.

//	It takes me so long to realize that... how can thinking faster?
