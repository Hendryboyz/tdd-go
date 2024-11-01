package generic2

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	last := len(s.values) - 1
	element := s.values[last]
	s.values = s.values[:last]
	return element, true
}

type StackOfInts = Stack[int]

type StackOfStrings = Stack[string]
