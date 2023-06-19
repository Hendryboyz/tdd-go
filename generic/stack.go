package generic

type StackOfInt struct {
	values []int32
}

func (s *StackOfInt) Push(val int32) {
	s.values = append(s.values, val)
}

func (s *StackOfInt) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfInt) Pop() (int32, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	lastIndex := len(s.values) - 1
	lastValue := s.values[lastIndex]
	s.values = s.values[:lastIndex]
	return lastValue, true
}
