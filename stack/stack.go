package stack

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	value T
	next  *Node[T]
}

type Stack[T constraints.Ordered] struct {
	top  *Node[T]
	size int
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Push(value T) {
	s.top = &Node[T]{value: value, next: s.top}
	s.size++
}

func (s *Stack[T]) Pop() (T, error) {
	if s.size > 0 {
		node := s.top
		s.top = s.top.next
		s.size--

		return node.value, nil
	}

	return *new(T), errors.New("Stack is empty")
}

func (s *Stack[T]) Peek() (T, error) {
	if s.size > 0 {
		node := s.top

		return node.value, nil
	}

	return *new(T), errors.New("Stack is empty")
}
