package v

import "sync"

type Stack[T any] struct {
	v []T
	m sync.Mutex
}

func (s *Stack[T]) Push(e T) {
	s.m.Lock()
	defer s.m.Unlock()
	s.v = append(s.v, e)
}

func (s *Stack[T]) Pop() T {
	if len(s.v) == 0 {
		panic(" stack is empty")
	}
	e := s.v[len(s.v)-1]
	s.v = s.v[:len(s.v)-1]
	return e
}

func (s *Stack[T]) Peek() T {
	l := len(s.v)
	if l == 0 {
		panic(" stack is empty")
	}
	return s.v[l-1]
}
func (s *Stack[T]) IsEmpty() bool {
	return len(s.v) == 0
}
