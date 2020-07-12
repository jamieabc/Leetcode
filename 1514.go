package main

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	if n == 0 || len(edges) == 0 || start == end {
		return 0
	}

	// create probability table
	probs := make(map[int]map[int]float64)
	for i := 0; i < n; i++ {
		probs[i] = make(map[int]float64)
	}

	for i, edge := range edges {
		probs[edge[0]][edge[1]] = float64(1) - succProb[i]
		probs[edge[1]][edge[0]] = float64(1) - succProb[i]
	}

	// bellman-ford
	// maxSuccess[i] means max success probability from start to i
	maxSuccess := make([]float64, n)
	for i := range maxSuccess {
		maxSuccess[i] = -1
	}
	maxSuccess[end] = 0
	queue := []info{{start, 0}}

	var p float64
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]

		for reachable, prob := range probs[popped.pos] {
			p = (float64(1) - popped.fail) * (1 - prob)

			if maxSuccess[reachable] == -1 || p > maxSuccess[reachable] {
				maxSuccess[reachable] = p
				queue = append(queue, info{reachable, float64(1) - p})
			}
		}
	}

	return maxSuccess[end]
}

type info struct {
	pos  int
	fail float64
}

type infos []info

func (p infos) Len() int {
	return len(p)
}

func (p infos) Less(i, j int) bool {
	return p[i].fail < p[j].fail
}

func (p infos) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *infos) Push(x interface{}) {
	*p = append(*p, x.(info))
}

func (p *infos) Pop() interface{} {
	old := *p
	popped := old[len(old)-1]
	*p = old[:len(old)-1]

	return popped
}

func maxProbability1(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	size := len(edges)
	if n == 0 || size == 0 || start == end {
		return 0
	}

	probs := make(map[int]map[int]float64)
	for i := 0; i < n; i++ {
		probs[i] = make(map[int]float64)
	}

	for i, edge := range edges {
		probs[edge[0]][edge[1]] = succProb[i]
		probs[edge[1]][edge[0]] = succProb[i]
	}

	seen := make(map[int]bool)

	// dijkstra
	p := &infos{{start, 0}}
	heap.Init(p)

	for p.Len() > 0 {
		popped := heap.Pop(p).(info)

		if popped.pos == end {
			return float64(1) - popped.fail
		}

		if seen[popped.pos] {
			continue
		}

		seen[popped.pos] = true

		for reachable, suc := range probs[popped.pos] {
			heap.Push(p, info{
				pos:  reachable,
				fail: float64(1) - (1-popped.fail)*suc,
			})
		}
	}

	return 0
}

//	problems
//	1.	panic, spend 1 hour and cannot solve it in time

//	2.	max success probability = min fail probability => dijkstra

//	3.	for bellman-ford, be careful about initial condition
