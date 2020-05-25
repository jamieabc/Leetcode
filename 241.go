package main

// Given a string of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. The valid operators are +, - and *.
//
// Example 1:
//
// Input: "2-1-1"
// Output: [0, 2]
// Explanation:
// ((2-1)-1) = 0
// (2-(1-1)) = 2
//
// Example 2:
//
// Input: "2*3-4*5"
// Output: [-34, -14, -10, -10, 10]
// Explanation:
// (2*(3-(4*5))) = -34
// ((2*3)-(4*5)) = -14
// ((2*(3-4))*5) = -10
// (2*((3-4)*5)) = -10
// (((2*3)-4)*5) = 10

func diffWaysToCompute(input string) []int {
	mapping := make(map[string][]int)

	traverse(input, mapping)

	return mapping[input]
}

func traverse(str string, mapping map[string][]int) {
	idx := make([]int, 0)
	for i := range str {
		if isOperator(str[i]) {
			idx = append(idx, i)
		}
	}

	if len(idx) == 0 {
		n, _ := strconv.Atoi(str)
		mapping[str] = append(mapping[str], n)
		return
	}

	for i := range idx {
		if _, ok := mapping[str[idx[i]+1:]]; !ok {
			traverse(str[idx[i]+1:], mapping)
		}

		if _, ok := mapping[str[:idx[i]]]; !ok {
			traverse(str[:idx[i]], mapping)
		}

		num1 := mapping[str[:idx[i]]]
		num2 := mapping[str[idx[i]+1:]]
		for j := range num1 {
			for k := range num2 {
				mapping[str] = append(mapping[str], operate(num1[j], num2[k], str[idx[i]]))
			}
		}
	}
}

func isOperator(b byte) bool {
	return b == '+' || b == '-' || b == '*'
}

func operate(num1, num2 int, op byte) int {
	switch op {
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	default:
		return num1 * num2
	}
}
