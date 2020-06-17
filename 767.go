package main

import (
	"container/heap"
	"sort"
)

// Given a string S, check if the letters can be rearranged so that two characters that are adjacent to each other are not the same.
//
// If possible, output any possible result.  If not possible, return the empty string.
//
// Example 1:
//
// Input: S = "aab"
// Output: "aba"
// Example 2:
//
// Input: S = "aaab"
// Output: ""
// Note:
//
// S will consist of lowercase letters and have length in range [1, 500].

type stat struct {
	b     byte
	count int
}

type stats []stat

func (s stats) Len() int {
	return len(s)
}

func (s stats) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s stats) Less(i, j int) bool {
	return s[i].count > s[j].count
}

func (s *stats) Push(x interface{}) {
	*s = append(*s, x.(stat))
}

func (s *stats) Pop() interface{} {
	old := *s
	popped := old[len(old)-1]
	*s = old[:len(old)-1]
	return popped
}

func reorganizeString(S string) string {
	counts := make([]int, 26)
	for i := range S {
		counts[int(S[i]-'a')]++
	}

	s := stats{}
	heap.Init(&s)

	for i := range counts {
		if counts[i] > 0 {
			if counts[i] > (len(S)+1)/2 {
				return ""
			}

			heap.Push(&s, stat{byte('a' + i), counts[i]})
		}
	}

	bs := make([]byte, len(S))

	for idx := 0; s.Len() > 0; idx += 2 {
		s1 := heap.Pop(&s).(stat)
		if s.Len() == 0 {
			bs[idx] = s1.b
			s1.count--
			break
		}

		s2 := heap.Pop(&s).(stat)

		bs[idx] = s1.b
		bs[idx+1] = s2.b

		s1.count--
		s2.count--

		if s1.count > 0 {
			heap.Push(&s, s1)
		}

		if s2.count > 0 {
			heap.Push(&s, s2)
		}
	}

	return string(bs)
}

func reorganizeString1(S string) string {
	counts := make([]int, 26)
	for i := range S {
		counts[int(S[i]-'a')]++
	}

	type info struct {
		b     byte
		count int
	}

	arr := make([]info, 0)
	for i := range counts {
		if counts[i] > 0 {
			arr = append(arr, info{
				b:     byte('a' + i),
				count: counts[i],
			})
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].count > arr[j].count
	})

	bs := make([]byte, len(S))
	var idx int
	var round bool

	for i := 0; i < len(arr); i++ {
		if arr[i].count > (len(S)+1)/2 {
			return ""
		}

		for j := arr[i].count; j > 0; j-- {
			if idx >= len(S) {
				if !round {
					round = true
					idx = 1
				} else {
					break
				}
			}
			bs[idx] = arr[i].b
			idx += 2
		}
	}

	return string(bs)
}

//	problems
//	1.	my attempt is really complicated, and didn't write it out

//	2.	inspired from solution, sort string into continuous block by occurrence
//		time, then insert into new string in the order of odd index first, then
//		even index., tc: O(n + a log a), n: word length, a: char group length

//		fail criteria, any char occurrence time < (N+1)/2

//	3.	the reason I didn't write it out is because I know the algorithm should
//		write interleaving, but I think in a more complicated way, write a then
//		b then a, so I encounter a problem of choosing next which means I have
//		to track 2 groups of chars at the same time.

//		solution uses a cleaver way, write even index first, then odd index, and
//		write sequence is based on char occurrence time.

//	4.	inspired from solution, use heap to get top 2 frequent chars and write
//		to result, tc: O(n log a), n: word length, a: char group length.

//		every time write result, write most frequent char then second frequent
//		char
