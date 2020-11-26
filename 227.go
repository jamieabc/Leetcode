package main

// Implement a basic calculator to evaluate a simple expression string.
//
// The expression string contains only non-negative integers, +, -, *, / operators and empty spaces . The integer division should truncate toward zero.
//
// Example 1:
//
// Input: "3+2*2"
// Output: 7
// Example 2:
//
// Input: " 3/2 "
// Output: 1
// Example 3:
//
// Input: " 3+5 / 2 "
// Output: 5
// Note:
//
// You may assume that the given expression is always valid.
// Do not use the eval built-in library function.

func calculate(s string) int {
	var ans, prev, op, num int
	size := len(s)

	for i := 0; i < size; i++ {
		if s[i] == ' ' {
			continue
		}

		if isOperator(s, i) {
			op = i
			continue
		}

		num = 0
		for ; i < size && s[i] >= '0' && s[i] <= '9'; i++ {
			num = num*10 + int(s[i]-'0')
		}
		i-- // it's on operator or space, if it's on operator, cannot skip

		if op == 0 || s[op] == '+' {
			// initial status or operator found, because first character cannot
			// be operator
			ans += num
			prev = num
		} else if s[op] == '-' {
			num *= -1
			ans += num
			prev = num
		} else if s[op] == '*' {
			ans = ans - prev + prev*num
			prev *= num
		} else {
			ans = ans - prev + prev/num
			prev /= num
		}
	}

	return ans
}

func isOperator(s string, idx int) bool {
	return s[idx] == '+' || s[idx] == '-' || s[idx] == '*' || s[idx] == '/'
}

func calculate2(s string) int {
	stack := make([]int, 0)
	size := len(s)
	var num, op int

	for i := 0; i < size; i++ {
		if s[i] == ' ' {
			continue
		}

		if isOperator(s, i) {
			op = i
			continue
		}

		num = 0
		for ; i < size && s[i] >= '0' && s[i] <= '9'; i++ {
			num = num*10 + int(s[i]-'0')
		}
		i--

		if op == 0 || s[op] == '+' {
			stack = append(stack, num)
		} else if s[op] == '*' {
			num *= stack[len(stack)-1]
			stack[len(stack)-1] = num
		} else if s[op] == '/' {
			num = stack[len(stack)-1] / num
			stack[len(stack)-1] = num
		} else {
			stack = append(stack, -num)
		}
	}

	var ans int
	for i := range stack {
		ans += stack[i]
	}

	return ans
}

func isOperator(s string, idx int) bool {
	return s[idx] == '+' || s[idx] == '-' || s[idx] == '*' || s[idx] == '/'
}

func calculate1(s string) int {
	var ans, prev int
	size := len(s)

	for i := 0; i < size; i++ {
		num, next := parse(s, i)

		if prev == 0 || s[prev] == '+' {
			ans += num
		} else {
			ans -= num
		}
		prev = next
		i = next
	}

	return ans
}

// deal with * / operators
func parse(s string, idx int) (int, int) {
	var num, prev, i int
	size := len(s)

	for prev, i = idx, idx; i < size; i++ {
		val, next := nextNum(s, i)

		if i == prev {
			num += val
		} else {
			if s[prev] == '*' {
				num *= val
			} else {
				num /= val
			}
		}

		prev = next
		i = next

		if i == size || s[i] == '+' || s[i] == '-' {
			break
		}
	}

	return num, i
}

// find next number and locate to next valid position
func nextNum(s string, idx int) (int, int) {
	var num, i int
	size := len(s)

	for i = idx; i < size && s[i] == ' '; i++ {
	}

	for ; i < size && s[i] >= '0' && s[i] <= '9'; i++ {
		num = num*10 + int(s[i]-'0')
	}

	for ; i < size && s[i] == ' '; i++ {
	}

	return num, i
}

//	Notes
//	1.	I was confused about finding right order to process numbers, because
//		if * / is involved, need to know previous + - operator

//		I try to separate calculation process into two: one is for + -, the other
//		is * /.

//	2.	inspired from sample code, operator - means invert number
//		and it could be a recursion process to always add next number, and if
//		* / is encountered, reduce previous number

//		3 + 2 * 4 * 5

//		ans = 0, num = 3, ans = 0+3 = 3
//   	ans = 3, num = 2, ans = 3+2 = 5, prev = 2
//	 	ans = 5, num = 4, prev = 2, ans = 5-2+2*4, prev = 2*4
//		ans = 11, num = 5, prev = 8, ans = 11 - 8 + 8 * 5, prev = 8 * 5

// 		the way to see this problem is treat every number as add operation, and
//		deduct previous value if * / is encountered

//		if - is encountered, invert number and keep adding

//	3.	solution provides a very good thinking process

//		4+3*5 => 4+15
//		4+3-5 => 7-5
//		4*3/5 => 12/5
//		4*3-5 => 12-5

//		for operator + & -, next number is evaluated by next operator,
//		e.g. 4+3*5, cannot evaluate 4+3 immediately, because next operator is *
//		and * has higher precedence

//		for operator * & /, can evaluate ignore next operator because those two
//		operator has higher precedence

//		because + & - needs to check later operator, so a stack can be used to
//		solve the problem

//		this only relates to observation, and it's how problem is tackled, I
//		should learn how to think in this way
