package main

// Правильная скобочная последовательность
func PSP(arr []rune) bool {

	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var s Stack[rune]
	s.Create()
	// ()[]
	// 0123
	for _, elem := range arr {
		if elem == '(' || elem == '[' || elem == '{' {
			s.Push(elem)
		} else {
			if s.Size() == 0 {
				return false
			}
			if s.Top() != m[elem] {
				return false
			}
			s.Pop()
		}
	}

	if s.Size() != 0 {
		return false
	}

	return true
}

// Префиксная запись выражения
func CalculateExper(expr string) int {
	calc := func(operator rune, a, b int) int {

		switch operator {
		case '-':
			return b - a
		case '+':
			return a + b
		case '*':
			return a * b
		case '/':
			return b / a
		default:
			return 0
		}
	}

	is_operand := func(a rune) bool {

		if a >= '0' && a <= '9' {
			return true
		}

		return false
	}

	var s Stack[int]
	s.Create()

	for _, elem := range expr {

		if is_operand(elem) {
			s.Push(int(elem - '0'))
		} else {
			a := s.Top()
			s.Pop()
			b := s.Top()
			s.Pop()
			s.Push(calc(elem, a, b))
		}
	}

	return s.Top()
}

func ToPolishNotation(expr string) string {

	is_operator := func(a rune) bool {
		switch a {
		case '-':
			return true
		case '+':
			return true
		case '*':
			return true
		case '/':
			return true
		case '(':
			return true
		case ')':
			return true
		default:
			return false
		}
	}
	priority := func(a rune) int {

		switch a {
		case '-':
			return 1
		case '+':
			return 1
		case '/':
			return 2
		case '*':
			return 2
		default:
			return 0
		}
	}

	var s Stack[rune]
	s.Create()
	ans := make([]rune, 0, len(expr))

	for _, char := range expr {

		if is_operator(char) {

			if s.Size() == 0 {
				s.Push(char)
			} else if char == '(' {
				s.Push(char)
			} else if char == ')' {
				for s.Size() != 0 && s.Top() != '(' {
					ans = append(ans, s.Top())
					s.Pop()
				}
				s.Pop()
			} else {
				for s.Size() != 0 && priority(char) <= priority(s.Top()) {
					ans = append(ans, s.Top())
					s.Pop()
				}
				s.Push(char)
			}
		} else {
			ans = append(ans, char)
		}
	}

	for s.Size() != 0 {
		ans = append(ans, s.Top())
		s.Pop()
	}

	return string(ans)
}

// Поиск ближайшего меньшего справа(слева)
func find_least_right(arr []int) []int {
	ans := make([]int, len(arr))

	var s Stack[interface{}]
	s.Create()

	for ind, val := range arr {
		if s.Empty() {
			s.Push([]int{val, ind})
		} else {
			top := s.Top().([]int)

			if top[0] <= val {
				s.Push([]int{val, ind})
			} else {
				for s.Size() != 0 && top[0] > val {
					ans[top[1]] = ind
					s.Pop()
					if s.Size() != 0 {
						top = s.Top().([]int)
					}
				}
				s.Push([]int{val, ind})
			}
		}
	}

	for s.Size() != 0 {
		top := s.Top().([]int)
		s.Pop()
		ans[top[1]] = len(arr)
	}

	return ans
}
