package main

func backspaceCompare(S string, T string) bool {
	var p1, p2 int

	for p1, p2 = len(S)-1, len(T)-1; p1 >= 0 || p2 >= 0; {
		p1 = nextChar(S, p1)
		p2 = nextChar(T, p2)

		if (p1 < 0 && p2 >= 0) || (p1 >= 0 && p2 < 0) || (p1 >= 0 && p2 >= 0 && S[p1] != T[p2]) {
			break
		}

		p1--
		p2--
	}

	return p1 < 0 && p2 < 0
}

func nextChar(str string, idx int) int {
	var delCount int
	for idx >= 0 {
		if str[idx] == '#' {
			idx--
			delCount++
		} else if delCount > 0 {
			idx--
			delCount--
		} else {
			return idx
		}
	}

	return idx
}

func backspaceCompare3(S string, T string) bool {
	backS, backT := 0, 0

	// traverse backward, since backspace can only affect char before that position
	var i, j int
	for i, j = len(S)-1, len(T)-1; i >= 0 || j >= 0; {

		// find next not deleted char
		for i >= 0 {
			if S[i] == '#' {
				backS++
				i--
			} else {
				if backS > 0 {
					i--
					backS--
				} else {
					break
				}
			}
		}

		for j >= 0 {
			if T[j] == '#' {
				backT++
				j--
			} else {
				if backT > 0 {
					j--
					backT--
				} else {
					break
				}
			}
		}

		if i < 0 && j < 0 {
			return true
		}

		if (i < 0 && j >= 0) || (i >= 0 && j < 0) {
			return false
		}

		if S[i] != T[j] {
			return false
		}

		i--
		j--
	}

	return true
}

func backspaceCompare2(S string, T string) bool {
	s1 := stack{
		data: make([]byte, 0),
	}
	s2 := stack{
		data: make([]byte, 0),
	}

	for i := 0; i < len(S); i++ {
		if S[i] == '#' {
			_ = s1.pop()
		} else {
			s1.push(S[i])
		}
	}

	for i := 0; i < len(T); i++ {
		if T[i] == '#' {
			_ = s1.pop()
		} else {
			s1.push(T[i])
		}
	}

	for !s1.empty() && !s2.empty() {
		s := s1.pop()
		t := s2.pop()
		if s != t {
			return false
		}
	}

	return s1.empty() && s2.empty()
}

type stack struct {
	data []byte
}

func (s *stack) push(i byte) {
	s.data = append(s.data, i)
}

func (s *stack) pop() byte {
	length := len(s.data)
	if length == 0 {
		return '0'
	}
	popped := s.data[length-1]
	s.data = s.data[:length-1]
	return popped
}

func (s *stack) empty() bool {
	return len(s.data) == 0
}

// problems
// 1. use wrong variable, when processing T, forget to update variable
// 2. wrong condition of comparing, if should be neither of stack is empty
// 3. still wrong condition of point 2, because it should be neither, so it's !empty
// 4. typo
// 5. wrong situation if all characters are deleted, e.g. i will be 0 & j will be 0 which is not true
// 6. if i & j both < 0, break loop

//	7.	inspired from https://leetcode.com/problems/backspace-string-compare/discuss/145786/Python-tm

//		use function to get next char

//	8.	the point of this problem is to identify performing delete process,
//		still need to check if next char is another delete
