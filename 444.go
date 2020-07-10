package main

import "fmt"

func sequenceReconstruction(org []int, seqs [][]int) bool {
	indexes := make([]int, len(org))
	for i := range org {
		indexes[org[i]-1] = i
	}
	orders := make(map[string]bool)

	for _, seq := range seqs {
		for i := range seq {
			if seq[i] > len(org) {
				return false
			}

			if len(seq) == 1 {
				orders[fmt.Sprintf("%d", seq[0])] = true
				continue
			}

			if i < len(seq)-1 && (seq[i+1] > len(org) || indexes[seq[i]-1] >= indexes[seq[i+1]-1]) {
				return false
			}

			if i < len(seq)-1 {
				orders[keyrize(seq[i], seq[i+1])] = true
			}
		}
	}

	if len(orders) == 0 || len(orders) < len(org)-1 {
		return false
	}

	// make sure consecutive relationships are constructed
	for i := 0; i < len(org)-1; i++ {
		if !orders[keyrize(org[i], org[i+1])] {
			return false
		}
	}

	return true
}

func keyrize(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

//	problems
//	1.	don't know how to do it, it seems like a graph problem that shows
//		connections between nodes

//		org can be validated by seqs, but unique means seqs should also
//		guarantee one solution, I don't know how to check that

//	2.	inspired from https://leetcode.com/problems/sequence-reconstruction/discuss/92574/Very-short-solution-with-explanation

//		to guarantee unique, two conditions are met:
//		- seqs should be part of org
//		- every consecutive number in org should be relations in seqs

//		second point is really important, with this insight, this this problem
//		can be solved

//	3.	seqs should not be same number

//	4.	check seqs in org, and consecutive numbers in org has correct order in
//		seqs

//	5.	from problem description, it already says org is permutation of 1-N,
//		guarantee uniqueness, so it's better to build up relationships in
//		org and check correctness in seq

//	6.	so many corner cases, finally TLE, not keep doing it

//	7. 	didn't write pass solution
