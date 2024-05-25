package priority_queue

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	isEmptyError  = errors.New("queue is empty")
	notEmptyError = errors.New("queue is not empty")
	notFoundError = errors.New("value not in queue")
)

type PriorityQueue[T constraints.Ordered] struct {
	heap []T
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	if pq.IsEmpty() {
		return *new(T), isEmptyError
	}

	return pq.heap[0], nil
}

func (pq *PriorityQueue[T]) Add(value T) {
	pq.heap = append(pq.heap, value)
	lastElementIndex := pq.Size() - 1
	pq.swim(lastElementIndex)
}

func (pq *PriorityQueue[T]) Poll() (T, error) {
	return pq.removeAt(0)
}

func (pq *PriorityQueue[T]) Remove(value T) (removed bool, err error) {
	if pq.IsEmpty() {
		return false, isEmptyError
	}

	for i, v := range pq.heap {
		if v == value {
			pq.removeAt(i)
			return true, nil
		}
	}

	return false, notFoundError
}

func (pq *PriorityQueue[T]) Fill(values []T) error {
	if !pq.IsEmpty() {
		return notEmptyError
	}

	pq.heap = values

	// Heapify
	for i := max(0, (pq.Size()/2)-1); i >= 0; i-- {
		pq.sink(i)
	}

	return nil
}

func (pq *PriorityQueue[T]) Contains(value T) bool {
	if pq.IsEmpty() {
		return false
	}

	for _, v := range pq.heap {
		if v == value {
			return true
		}
	}

	return false
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *PriorityQueue[T]) Clear() {
	pq.heap = nil
}

func (pq *PriorityQueue[T]) Size() int {
	return len(pq.heap)
}

func (pq *PriorityQueue[T]) removeAt(index int) (T, error) {
	if pq.IsEmpty() {
		return *new(T), isEmptyError
	}

	lastElementIndex := pq.Size() - 1
	removed := pq.heap[index]

	pq.swap(index, lastElementIndex)
	pq.heap = pq.heap[:lastElementIndex]

	if index == lastElementIndex {
		return removed, nil
	}

	value := pq.heap[index]
	pq.sink(index)

	if pq.heap[index] == value {
		pq.swim(index)
	}

	return removed, nil
}

func (pq *PriorityQueue[T]) swim(index int) {
	parentIndex := pq.getParent(index)

	for index > 0 && pq.less(index, parentIndex) {
		pq.swap(parentIndex, index)
		index = parentIndex
		parentIndex = pq.getParent(index)
	}
}

func (pq *PriorityQueue[T]) sink(index int) {
	for true {
		left, right := pq.getChildren(index)
		smallest := left

		if right < pq.Size() && pq.less(right, left) {
			smallest = right
		}

		if left >= pq.Size() || pq.less(index, smallest) {
			break
		}

		pq.swap(smallest, index)
		index = smallest
	}
}

func (pq *PriorityQueue[T]) swap(index1, index2 int) {
	value1 := pq.heap[index1]
	value2 := pq.heap[index2]

	pq.heap[index1] = value2
	pq.heap[index2] = value1
}

func (pq *PriorityQueue[T]) less(index1, index2 int) bool {
	value1 := pq.heap[index1]
	value2 := pq.heap[index2]

	return value1 <= value2
}

func (pq *PriorityQueue[T]) getParent(index int) int {
	return (index - 1) / 2
}

func (pq *PriorityQueue[T]) getChildren(index int) (int, int) {
	left := 2*index + 1
	right := 2*index + 2

	return left, right
}

// This method is just for testing.
func (pq *PriorityQueue[T]) isMinHeap(index int) bool {
	if index >= pq.Size() {
		return true
	}

	left, right := pq.getChildren(index)

	if left < pq.Size() && !pq.less(index, left) {
		return false
	}

	if right < pq.Size() && !pq.less(index, right) {
		return false
	}

	return pq.isMinHeap(left) && pq.isMinHeap(right)
}
