package linked_list

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	value T
	next  *node[T]
}

type LinkedList[T constraints.Ordered] struct {
	head *node[T]
	size int
}

func (l *LinkedList[T]) InsertLast(value T) {
	new_node := &node[T]{value: value}

	if l.head == nil {
		l.head = new_node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}

		current.next = new_node
	}

	l.size++
}

func (l *LinkedList[T]) InsertFirst(value T) {
	l.head = &node[T]{value: value, next: l.head}
	l.size++
}

func (l *LinkedList[T]) InsertAt(index int, value T) error {
	err := l.checkIndex(index)

	if err != nil {
		return err
	}

	if index == 0 {
		l.InsertFirst(value)
	} else if index == l.size {
		l.InsertLast(value)
	} else {
		new_node := &node[T]{value: value}
		current := l.traverseTo(index - 1)
		new_node.next = current.next
		current.next = new_node
		l.size++
	}

	return nil
}

func (l *LinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}

	l.head = l.head.next
	l.size--
}

func (l *LinkedList[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}

	current := l.traverseTo(l.size - 2) // traverse to the next to last
	current.next = nil
	l.size--
}

func (l *LinkedList[T]) RemoveAt(index int) error {
	err := l.checkIndex(index)

	if err != nil {
		return err
	}

	if index == 0 {
		l.RemoveFirst()
	} else {
		current := l.traverseTo(index - 1)
		current.next = current.next.next
		l.size--
	}

	return nil
}

func (l *LinkedList[T]) traverseTo(index int) *node[T] {
	current := l.head

	for i := 0; i != index; i++ {
		current = current.next
	}

	return current
}

func (l *LinkedList[T]) checkIndex(index int) error {
	if index < 0 || index > l.size {
		return errors.New("Index out of range")
	}

	return nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Print() {
	current := l.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%v ", current.value)
		current = current.next
	}
	fmt.Println()
}
