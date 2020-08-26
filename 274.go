package main

//Given an array of citations (each citation is a non-negative integer) of a researcher, write a function to compute the researcher's h-index.
//
//According to the definition of h-index on Wikipedia: "A scientist has index h if h of his/her N papers have at least h citations each, and the other N − h papers have no more than h citations each."
//
//Example:
//
//Input: citations = [3,0,6,1,5]
//Output: 3
//Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had
//received 3, 0, 6, 1, 5 citations respectively.
//Since the researcher has 3 papers with at least 3 citations each and the remaining
//two with no more than 3 citations each, her h-index is 3.
//Note: If there are several possible values for h, the maximum one is taken as the h-index.

func hIndex(citations []int) int {
	counter := make([]int, len(citations)+1)
	for i := range citations {
		counter[min(len(counter)-1, citations[i])]++
	}

	for sum, i := 0, len(counter)-1; i >= 0; i-- {
		sum += counter[i]
		if sum < i {
			continue
		}

		return i
	}

	return 0
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

//	Notes
//	1.	0 <= h <= N(paper count)

//	2.	inspired from solution, sorting graph is really elegant and expressive
