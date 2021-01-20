package main

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//    Open brackets must be closed by the same type of brackets.
//    Open brackets must be closed in the correct order.
//
//Note that an empty string is also considered valid.
//
//Example 1:
//
//Input: "()"
//Output: true
//
//Example 2:
//
//Input: "()[]{}"
//Output: true
//
//Example 3:
//
//Input: "(]"
//Output: false
//
//Example 4:
//
//Input: "([)]"
//Output: false
//
//Example 5:
//
//Input: "{[]}"
//Output: true

func isValid(s string) bool {
	stack := make([]byte, 0)

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, ')')
		} else if s[i] == '[' {
			stack = append(stack, ']')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	st := &stack{data: make([]rune, 0)}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			st.push(c)
		} else if c == ')' || c == ']' || c == '}' {
			if st.pop() != c {
				return false
			}
		} else {
			return false
		}
	}
	return st.length == 0
}

type stack struct {
	data   []rune
	length int
}

func (s *stack) push(r rune) {
	s.data = append(s.data, r)
	s.length++
}

func (s *stack) pop() rune {
	if s.length == 0 {
		return ' '
	}

	popped := s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length--

	switch popped {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	default:
		return ' '
	}
}

//	Notes
//	1.	becareful about boundary conditions such as emtpy stack

//	2.	inspired from https://leetcode.com/problems/valid-parentheses/discuss/9178/Short-java-solution

//		no need to pop stack and check each type matching, can just push valid order into stack
