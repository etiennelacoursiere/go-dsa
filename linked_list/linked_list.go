package linked_list

import (
	"errors"
	"fmt"
)

type Node struct {
	value any
	next  *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func (l *LinkedList) Append(value any) {
	new_node := &Node{value: value}

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

func (l *LinkedList) Prepend(value any) {
	new_node := &Node{value: value}
	new_node.next = l.Head
	l.Head = new_node
	l.Size++
}

func (l *LinkedList) RemoveFirst() {
	if l.IsEmpty() {
		return
	}

	l.Head = l.Head.next
	l.Size--
}

func (l *LinkedList) RemoveLast() {
	if l.IsEmpty() {
		return
	}

	current := l.traverseTo(l.Size - 2) // traverse to the next to last
	fmt.Println(l.Size, current)
	current.next = nil
	l.Size--
}

func (l *LinkedList) InsertAt(index int, value any) error {
	err := l.checkIndex(index)

	if err != nil {
		return err
	}

	if index == 0 {
		l.Prepend(value)
	} else if index == l.Size {
		l.Append(value)
	} else {
		new_node := &Node{value: value}
		current := l.traverseTo(index - 1)
		new_node.next = current.next
		current.next = new_node
		l.Size++
	}

	return nil
}

func (l *LinkedList) RemoveAt(index int) error {
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

func (l *LinkedList) traverseTo(index int) *Node {
	current := l.Head

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current
}

func (l *LinkedList) checkIndex(index int) error {
	if index < 0 || index > l.Size {
		return errors.New("Index out of range")
	}

	return nil
}

func (l *LinkedList) IsEmpty() bool {
	return l.Size == 0
}

func (l *LinkedList) Print() {
	current := l.Head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}
