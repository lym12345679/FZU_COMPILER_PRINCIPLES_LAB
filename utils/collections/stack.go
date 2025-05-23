// Simple stack implementation

package collections

import (
	"fmt"
)

type Stack[T any] struct {
	data []T
}

// NewStack creates a new stack.
func NewStack[T any]() Stack[T] {
	return Stack[T]{data: []T{}}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top element of the stack.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, true
}

func (s *Stack[T]) TrimTopN(n int) {
	if n <= 0 || n > len(s.data) {
		return
	}
	s.data = s.data[:len(s.data)-n]
}

func (s *Stack[T]) PopTopN(n int) []T {
	if n <= 0 || n > len(s.data) {
		return nil
	}
	values := make([]T, n)
	copy(values, s.data[len(s.data)-n:])
	s.data = s.data[:len(s.data)-n]
	return values
}

// Peek returns the top element of the stack without removing it.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return s.data[len(s.data)-1], true
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	s.data = []T{}
}

// String returns a string representation of the stack.
func (s *Stack[T]) String() string {
	// Convert the stack to a string representation
	result := "["
	for i, v := range s.data {
		result += fmt.Sprintf("%v", v)
		if i < len(s.data)-1 {
			result += ", "
		}
	}
	result += "]"
	return result
}

func (s *Stack[T]) Foreach(f func(T)) {
	for _, v := range s.data {
		f(v)
	}
}

func (s *Stack[T]) PeekAtK(k int) (T, bool) {
	if k < 0 || k >= len(s.data) {
		var zeroValue T
		return zeroValue, false
	}
	return s.data[len(s.data)-1-k], true
}
