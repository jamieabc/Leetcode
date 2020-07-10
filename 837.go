package main

func new21Game(N int, K int, W int) float64 {
	// not possible to win
	if K > N {
		return 0
	}

	// definitely win
	if K == 0 || N >= K+W {
		return float64(1)
	}

	prob := make([]float64, K+W)
	prevProb := float64(1)
	var result float64

	for i := 1; i <= N; i++ {
		prob[i] = prevProb / float64(W)

		if i < K {
			prevProb += prob[i]
		} else {
			result += prob[i]
		}

		// remove previous probability sum when number > W
		if i >= W {
			prevProb -= prob[i-W]
		}
	}

	return result
}

//  problems
//  1.  time limit exceed

//	2.	count # of possible combinations when points = K, K+1, K+2, ..., N, N+1,
//		N+2, ..., K-1+N

//	3.	inspired from https://leetcode.com/problems/new-21-game/discuss/132394/Can-someone-explain-why-counting-number-of-ways-to-get-to-X-approach-doesn't-give-correct-answer

//		I am trying to count all possible combinations to specific point, then
//		sum(N-K, N-K+1, ..., N) / sum(N-K, N-K+1, ..., N, N+1, N+2, ..., K-1+N)
//		this is wrong because depends on different path, possibility is different
//		even if they sum to same count

//		e.g. points 5 = 5    (1 card)
//			 points 5 = 1+4  (2 cards)
//		these two situation are same when counting points, but different
//		possibility

//	4.	inspired from https://leetcode.com/problems/new-21-game/discuss/132334/One-Pass-DP-O(N)

//		really hard to think of...
