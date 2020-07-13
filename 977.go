package main

func sortedSquares(A []int) []int {
	var next, low, high int
	result := make([]int, len(A))

	for next, low, high = len(A)-1, 0, len(A)-1; low <= high; next-- {
		i, j := abs(A[low]), abs(A[high])

		if i <= j {
			result[next] = j * j
			high--
		} else {
			result[next] = i * i
			low++
		}
	}

	return result[next+1:]
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
