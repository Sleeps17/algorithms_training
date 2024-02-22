package main

type Stack[T any] struct {
	arr []T
	len int
	cap int
}

func (s *Stack[T]) Create() {
	s.arr = make([]T, 0, 8)
	s.len = 0
	s.cap = 8
}

func (s *Stack[T]) Push(x T) {
	s.arr = append(s.arr, x)
	s.len++
	s.cap = cap(s.arr)
}

func (s *Stack[T]) Pop() {
	if s.len == 0 {
		return
	}

	s.len--
	s.arr = s.arr[:s.len]
}

func (s *Stack[T]) Top() T {
	if s.len == 0 {
		var zeroValue T
		return zeroValue
	}

	return s.arr[s.len-1]
}

func (s *Stack[T]) Size() int {
	return s.len
}

func (s *Stack[T]) Empty() bool {
	return s.len == 0
}

func (s *Stack[T]) Clear() {
	s.arr = []T{}
	s.len = 0
	s.cap = 0
}
