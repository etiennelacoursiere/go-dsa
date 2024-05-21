package linked_list

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	value T
	next  *Node[T]
}

type LinkedList[T constraints.Ordered] struct {
	Head *Node[T]
	Size int
}

func (l *LinkedList[T]) Append(value T) {
	new_node := &Node[T]{value: value}

	if l.Head == nil {
		l.Head = new_node
	} else {
		current := l.Head
		for current.next != nil {
			current = current.next
		}

		current.next = new_node
	}

	l.Size++
}

func (l *LinkedList[T]) Prepend(value T) {
	new_node := &Node[T]{value: value}
	new_node.next = l.Head
	l.Head = new_node
	l.Size++
}

func (l *LinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}

	l.Head = l.Head.next
	l.Size--
}

func (l *LinkedList[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}

	current := l.traverseTo(l.Size - 2) // traverse to the next to last
	fmt.Println(l.Size, current)
	current.next = nil
	l.Size--
}

func (l *LinkedList[T]) InsertAt(index int, value T) error {
	err := l.checkIndex(index)

	if err != nil {
		return err
	}

	if index == 0 {
		l.Prepend(value)
	} else if index == l.Size {
		l.Append(value)
	} else {
		new_node := &Node[T]{value: value}
		current := l.traverseTo(index - 1)
		new_node.next = current.next
		current.next = new_node
		l.Size++
	}

	return nil
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
		l.Size--
	}

	return nil
}

func (l *LinkedList[T]) traverseTo(index int) *Node[T] {
	current := l.Head

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current
}

func (l *LinkedList[T]) checkIndex(index int) error {
	if index < 0 || index > l.Size {
		return errors.New("Index out of range")
	}

	return nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.Size == 0
}

func (l *LinkedList[T]) Print() {
	current := l.Head

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
