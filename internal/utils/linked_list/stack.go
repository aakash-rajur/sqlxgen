package linked_list

type Stack[T any] struct {
	top    *Node[T]
	length int
}

func (s *Stack[T]) Length() int {
	return s.length
}

func (s *Stack[T]) Push(value *T) *Stack[T] {
	node := &Node[T]{
		Value: value,
		Next:  s.top,
	}

	s.top = node

	s.length += 1

	return s
}

func (s *Stack[T]) Pop() (*T, bool) {
	if s.top == nil {
		return nil, false
	}

	value := s.top.Value

	s.top = s.top.Next

	s.length -= 1

	return value, true
}

func (s *Stack[T]) Peek() (*T, bool) {
	if s.top == nil {
		return nil, false
	}

	return s.top.Value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack[T]) Clear() {
	s.top = nil

	s.length = 0
}

func (s *Stack[T]) ToSlice() []T {
	slice := make([]T, s.length)

	node := s.top

	for i := 0; i < s.length; i++ {
		slice[i] = *node.Value

		node = node.Next
	}

	return slice
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		top:    nil,
		length: 0,
	}
}

func FromSlice[T any](slice []T) *Stack[T] {
	stack := NewStack[T]()

	for _, item := range slice {
		value := item

		stack.Push(&value)
	}

	return stack
}
