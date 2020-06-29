package main

import (
	"math"
	"strings"
)

// dp: O(n^2)
func fullJustify(words []string, maxWidth int) []string {
	size := len(words)
	costs := buildCosts(words, maxWidth)

	// minCost[i] = minimum line cost from words i to end
	minCost, lineSplitter := make([]int, size), make([]int, size)

	var currentMin int
	for i := size - 1; i >= 0; i-- {
		currentMin = math.MaxInt32
		for j := size - 1; j >= i; j-- {
			if costs[i][j] > -1 {
				if j == size-1 {
					currentMin = costs[i][j]
					lineSplitter[i] = j + 1
				} else {
					if costs[i][j]+minCost[j+1] < currentMin {
						currentMin = costs[i][j] + minCost[j+1]
						lineSplitter[i] = j + 1
					}
				}
			}
		}
		minCost[i] = currentMin
	}

	result := make([]string, 0)
	var builder strings.Builder
	for i := 0; i < size; {
		result = append(result, justify(words, i, lineSplitter[i]-1, builder, maxWidth))
		i = lineSplitter[i]
	}

	return result
}

func buildCosts(words []string, maxWidth int) [][]int {
	size := len(words)

	costs := make([][]int, size)
	for i := range costs {
		costs[i] = make([]int, size)
		for j := range costs {
			costs[i][j] = -1 // -1 means unreachable
		}
	}

	var width int
	for i := range words {
		width = 0
		for j := 0; i+j < size; j++ {
			width += len(words[i+j])
			if width+j <= maxWidth {
				costs[i][i+j] = lineCost(maxWidth, width+j)
			} else {
				// reach limit, go to next word
				break
			}
		}
	}

	return costs
}

func lineCost(maxWidth, width int) int {
	return maxWidth - width
}

func fullJustify2(words []string, maxWidth int) []string {
	size := len(words)
	result := make([]string, 0)
	var builder strings.Builder

	for i := 0; i < size; i++ {
		lastIdx := lastWordIndexInLine(words, i, maxWidth)
		result = append(result, justify(words, i, lastIdx, builder, maxWidth))
		i = lastIdx
	}

	return result
}

func lastWordIndexInLine(words []string, start int, maxWidth int) int {
	var width, idx int
	for width, idx = len(words[start]), 1; start+idx < len(words); idx++ {
		if width+idx+len(words[start+idx]) <= maxWidth {
			width += len(words[start+idx])
		} else {
			break
		}
	}

	return start + idx - 1
}

func justify(words []string, start, end int, builder strings.Builder, maxWidth int) string {
	builder.Reset()

	spaces := spacesBetweenWords(words, start, end, maxWidth)
	for i := 0; i+start <= end; i++ {
		builder.WriteString(words[start+i])

		if i < len(spaces) {
			for j := 0; j < spaces[i]; j++ {
				builder.WriteByte(' ')
			}
		}
	}

	return builder.String()
}

func spacesBetweenWords(words []string, start, end int, maxWidth int) []int {
	var spaces []int

	var totalWordLength int
	for i := 0; start+i <= end; i++ {
		totalWordLength += len(words[start+i])
	}
	remainSpace := maxWidth - totalWordLength

	if end == len(words)-1 {
		// last line, space between words is 1, remaining spaces at end
		spaces = make([]int, end-start+1)

		for i := 0; i < len(spaces)-1; i++ {
			spaces[i] = 1
			remainSpace--
		}
		spaces[len(spaces)-1] = remainSpace
	} else if end == start {
		// only one word, all remaining spaces at end
		spaces = []int{remainSpace}
	} else {
		// spaces evenly distributed in-between
		spaces = make([]int, end-start)
		for i := 0; remainSpace > 0; i, remainSpace = i+1, remainSpace-1 {
			spaces[i%len(spaces)]++
		}
	}

	return spaces
}

func fullJustify1(words []string, maxWidth int) []string {
	size := len(words)
	result := make([]string, 0)
	if size == 0 {
		return result
	}

	var sb strings.Builder
	var i, j, currentWidth, remainSpace, normalSpace, leftJustifiedCount int

	for i = 0; i < size; i++ {
		sb.Reset()

		for currentWidth, j = 0, 0; i+j < size; j++ {
			if currentWidth+len(words[i+j])+j > maxWidth {
				j--
				break
			}
			currentWidth += len(words[i+j])
		}

		if i+j == size {
			j--
			// last line
			leftJustifiedCount = 0
			if j == 0 {
				normalSpace = 0
				remainSpace = maxWidth - currentWidth
			} else {
				normalSpace = 1
				remainSpace = maxWidth - currentWidth - j
			}
		} else {
			// current use 8 spaces with 3 words, each words with (16-8)/3 = 2 spaces, first (16-8)-3*2 = 2 need to be 3 spaces
			remainSpace = maxWidth - currentWidth
			if j == 0 {
				normalSpace = remainSpace
			} else {
				normalSpace = remainSpace / j
			}

			leftJustifiedCount = remainSpace - normalSpace*j
		}

		for k := 0; k < j && i+k < size; k++ {
			sb.WriteString(words[i+k])

			for l := 0; l < normalSpace; l++ {
				sb.WriteByte(' ')
			}

			if leftJustifiedCount > 0 {
				sb.WriteByte(' ')
				leftJustifiedCount--
			}
		}
		sb.WriteString(words[i+j])

		// last line or only word in a line
		if j == 0 || i+j == size-1 {
			for ; remainSpace > 0; remainSpace-- {
				sb.WriteByte(' ')
			}
		}

		result = append(result, sb.String())
		i += j
	}

	return result
}

//	problems
//	1.	inspired from https://leetcode.com/problems/text-justification/discuss/24902/Java-easy-to-understand-broken-into-several-functions

//		autor provides a cleaner way to read

//	2.	inspired from https://www.youtube.com/watch?v=RORuwHiblPc

//		the way of judging is to build a 2D matrix, and iterate through
//		each to decide min cost
