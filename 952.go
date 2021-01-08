package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

//Given a non-empty array of unique positive integers A, consider the following graph:
//
//There are A.length nodes, labelled A[0] to A[A.length - 1];
//There is an edge between A[i] and A[j] if and only if A[i] and A[j] share a common factor greater than 1.
//
//Return the size of the largest connected component in the graph.
//
//
//
//Example 1:
//
//Input: [4,6,15,35]
//Output: 4
//
//Example 2:
//
//Input: [20,50,9,63]
//Output: 2
//
//Example 3:
//
//Input: [2,3,6,7,4,12,21,39]
//Output: 8
//
//Note:
//
//1 <= A.length <= 20000
//1 <= A[i] <= 100000

func main() {
	start := time.Now()

	data, err := ioutil.ReadFile("txt.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)

	if len(input) == 0 {
		return
	}

	if input[0] == '[' {
		input = input[1 : len(input)-2]
	}

	strNums := strings.Split(string(input), ",")
	nums := make([]int, 0)
	for i := range strNums {
		num, _ := strconv.Atoi(strNums[i])
		nums = append(nums, num)
	}

	fmt.Println(largestComponentSize(nums))

	elapsed := time.Since(start)
	fmt.Println("elapsed: ", elapsed)
}

func largestComponentSize(A []int) int {
	primeGraph := make(map[int][]int) // 2: [0, 1], primes number 2 can fully divide A[0] & A[1]

	// group factors for every number
	for i := range A {
		groupPrimeFactors(i, A, primeGraph)
	}

	groups := make([]int, len(A))
	for i := range groups {
		groups[i] = i
	}

	// union
	for _, indexes := range primeGraph {
		for i := 1; i < len(indexes); i++ {
			union(groups, indexes[i-1], indexes[i])
		}
	}

	// flatten groups, avoid order dependency
	for i := range groups {
		parent(groups, i)
	}

	// count groups
	counter := make(map[int]int)
	maxCount := math.MinInt32
	for _, f := range groups {
		counter[groups[f]]++
		maxCount = max(maxCount, counter[groups[f]])
	}

	return maxCount
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

// in order to have all groups follow same order, choose to align to smaller number
func union(groups []int, to, from int) {
	groups[parent(groups, from)] = groups[parent(groups, to)]
}

func parent(groups []int, idx int) int {
	if groups[idx] != idx {
		groups[idx] = parent(groups, groups[idx])
	}
	return groups[idx]
}

// O(nk), k: average count of number's primes factors
func groupPrimeFactors(idx int, nums []int, primeGroups map[int][]int) {
	num := nums[idx]

	for i := 2; i*i <= num; i++ {
		if num%i != 0 {
			continue
		}

		primeGroups[i] = append(primeGroups[i], idx)
		for num /= i; num%i == 0; num /= i {
		}
	}

	if num > 1 {
		primeGroups[num] = append(primeGroups[num], idx)
	}
}

func largestComponentSize1(A []int) int {
	factors := make(map[int]int) // smallest factor for a number
	parents := make(map[int]int) // what factors are grouped

	for _, num := range A {
		findFactors(num, factors, parents)
	}

	counter := make(map[int]int)
	for _, num := range A {
		p := find(parents, factors[num])
		counter[p]++
	}

	var maxGroup int
	for _, count := range counter {
		maxGroup = max(maxGroup, count)
	}

	return maxGroup
}

func findFactors(num int, factors map[int]int, parents map[int]int) {
	var smallest int
	for i := 2; i*i <= num; i++ {
		if tmp := num / i; tmp*i == num {
			// use first factor, to avoid using numbers has other
			// factors, e.g. 4
			if _, ok := factors[num]; !ok {
				factors[num] = i
				smallest = i
			}

			p1, p2 := find(parents, i), find(parents, tmp)
			smallest = min(smallest, min(p1, p2))
			parents[p1] = smallest
			parents[p2] = smallest
		}
	}

	// prime factor
	if _, ok := factors[num]; !ok {
		factors[num] = num

		// could already be added before
		if _, ok := parents[num]; !ok {
			parents[num] = num
		}
	}
}

// factors: 2, 5, 4
func find(parents map[int]int, target int) int {
	if _, ok := parents[target]; !ok {
		parents[target] = target
	} else if parents[target] != target {
		parents[target] = find(parents, parents[target])
	}

	return parents[target]
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

//	Notes
//	1.	be careful about common factor limit, not only satisfies n/2, but also
//		need to satisfies >= 2

//	2.	groups need to be merged when any number has root smaller than number
//		with smaller index

//	3.	union happens whenever common factor is found, because it also needs to
//		update root belong to which group

//	4.	to find common factors, every number program starts from 2, so first
//		number will iterate n times to find common factors, which is a huge
//		waste of time.

//	5.	as long as a number's group is same as another number's group, no need
//		to check common factors, avoid un-necessary operations

//	6.	search prime factor up to square root of the number, but need to check
//		every number even if it's not prime, e.g. sqrt(68) = 9, 2*34 = 4*17

//	7.	when comparing group, need to find group parent instead of it's current
//		parent, without checking parent, it might cause a loop, e.g. 2's parent
//		is 1, and 1's parent is 2

//	8.	numbers increase very fast, but primes increase more slower, I would
//		like to use prime numbers as group index, such that operations could be
//		reduced dramatically

//	9.	after merge primes into group by numbers, eventually, need to map number
//		to primes so that count is possible

//	10.	1 is not prime, but since I assume prime starts from 2 with index 0, it
//		will merge 1 into 2, so need to specially process it

//	11.	for every union, it needs to update at root, otherwise, it might be
//		changed in the next find

//	12.	inspired from sample code, it's sufficient to check primes up to sqrt(n),
//		but in a more brilliant way. n = sqrt(n) * sqrt(n), if some prime number
//		is smaller than sqrt(n) means it "might" exist another prime number
//		larger than sqrt(n)

//		e.g. sqrt(35) = 5.xxxx, check up to 5, 35 / 5 = 7

//		so, how do I know if 7 is another prime, because prime numbers is pre-
//		calculated only up to 5, which is sqrt(35)

//		the way is to divide n for every prime factors, and it it turns out that
//		this n after all prime number division still larger than 1, which means
//		remaining n is prime number

//		e.g. n = 100000, sqrt(100000) = 316, primes : 2, 3, 5, 7, 11, ..., 313
//		for a given number 94238,, 94238 / 2 = 47119, 47119 cannot be fully
//		divided by rest of primes => 47119 is another prime number

//		so, with this observation, calculations can be dramatically reduced

//		though it's a brilliant solution to find prime number, I can't think of
//		a way to use this technique, because when union numbers, each number
//		should have a group. And in this technique, not all prime numbers are
//		found in the first time, which means some numbers are missing and is
//		found during processing numbers.

//		this technique can be used by author is because he uses a graph to
//		demonstrate edges among numbers, and after graph is built, he traverses
//		this graph to find maximum groups

//	13.	after thinking, even with union, technique to find primes still can be
//		used. just add new find prime into all data structures, and also put into
//		hash to more quickly found

//	14.	inspired from sample code, author provides a quick way to union: not to
//		to union for every number, instead, create a map to store what prime
// 		numbers has edges, and do union when this graph is done.

//		it's fast because it can find minimum parent when do the union one time,
//		unlike my algorithm, it takes several loops for intermediate (un-necessary)
//		computations

//		since union is done last, so it's okay to use technique to check prime
//		factors up to sqrt(n), and add new prime number when processing

//		because I think finding all prime number at first is actually not that
//		efficient, especially when numbers in the array are not that big

//	15.	inspired from solution, union-find tc: O(N * log M), M: largest number
//		in input, N: size of input, but if union-find is called upon building
//		up, average tc would be O(N)

//		overall tc: O(N * sqrt(K)), K: max number in the list

//	16.	to avoid duplicate number in prime group, author uses a technique to
//		divide number by found factor until that factor no longer exists, this
//		avoids redundant factors such as 2, 4, 8, 16, etc.

//		also, factors are generated in ascending order

//	17.	factors are used to union numbers, so it's not necessary to know what
//		that factors are union, should focus on what numbers are union

//		I still make mistake on focusing factors, but that's not important at
//		all, factors acts an in-direct indicator
