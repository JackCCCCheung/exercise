package v

type Stack[T any] struct {
	v []T
}

func (s *Stack[T]) Push(e T) {
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
