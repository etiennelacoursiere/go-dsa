package queue

import (
	"errors"
)

var queueEmptyError = errors.New("Index out of range")

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type Queue[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (q *Queue[T]) Enqueue(value T) {
	new_node := &Node[T]{value: value}

	if q.IsEmpty() {
		q.head = new_node
		q.tail = new_node
	} else if q.size == 1 {
		q.tail = new_node
		q.head.next = q.tail
	} else {
		q.tail.next = new_node
		q.tail = new_node
	}

	q.size++
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return *new(T), queueEmptyError
	}

	value := q.head.value
	q.head = q.head.next
	q.size--
	return value, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return *new(T), queueEmptyError
	}

	return q.head.value, nil
}

func (q *Queue[T]) Contains(value T) bool {
	if q.IsEmpty() {
		return false
	}

	if q.head.value == value || q.tail.value == value {
		return true
	}

	current := q.head
	for current.next != nil {
		if current.value == value {
			return true
		}

		current = current.next
	}

	return false
}

func (q *Queue[T]) Flush() []T {
	queue := make([]T, 0)

	if q.IsEmpty() {
		return queue
	}

	current := q.head
	for q.size != 0 {
		queue = append(queue, current.value)
		current = current.next
		q.size--
	}

	return queue
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
