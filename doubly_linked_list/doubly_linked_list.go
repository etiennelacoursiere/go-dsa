package doubly_linked_list

import (
	"errors"
	"fmt"
)

var (
	EmptyListError = errors.New("List is empty")
)

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type DoublyLinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func (dll *DoublyLinkedList[T]) InsertFirst(value T) {
	if dll.IsEmpty() {
		new_node := &node[T]{value: value}
		dll.head = new_node
		dll.tail = new_node
	} else {
		dll.head.prev = &node[T]{value: value, next: dll.head}
		dll.head = dll.head.prev
	}
	dll.size++
}

func (dll *DoublyLinkedList[T]) InsertLast(value T) {
	if dll.IsEmpty() {
		new_node := &node[T]{value: value}
		dll.head = new_node
		dll.tail = new_node
	} else {
		dll.tail.next = &node[T]{value: value, prev: dll.tail}
		dll.tail = dll.tail.next
	}

	dll.size++
}

func (dll *DoublyLinkedList[T]) GetFirst() (T, error) {
	if dll.IsEmpty() {
		return *new(T), EmptyListError
	}

	return dll.head.value, nil
}

func (dll *DoublyLinkedList[T]) Size() int {
	return dll.size
}

func (dll *DoublyLinkedList[T]) GetLast() (T, error) {
	if dll.IsEmpty() {
		return *new(T), EmptyListError
	}

	return dll.tail.value, nil
}

func (dll *DoublyLinkedList[T]) RemoveFirst() (T, error) {
	if dll.IsEmpty() {
		return *new(T), EmptyListError
	}

	value := dll.head.value
	dll.head = dll.head.next
	dll.size--

	if dll.IsEmpty() {
		dll.tail = nil
	} else {
		dll.head.prev = nil
	}

	return value, nil
}

func (dll *DoublyLinkedList[T]) RemoveLast() (T, error) {
	if dll.IsEmpty() {
		return *new(T), EmptyListError
	}

	value := dll.tail.value
	dll.tail = dll.tail.prev
	dll.size--

	if dll.IsEmpty() {
		dll.head = nil
	} else {
		dll.tail.next = nil
	}

	return value, nil
}

func (dll *DoublyLinkedList[T]) RemoveAt(index int) (T, error) {
	err := dll.checkIndex(index)
	if err != nil {
		return *new(T), err
	}

	node := dll.traverseTo(index)
	return dll.remove(node)

}

func (dll *DoublyLinkedList[T]) remove(node *node[T]) (T, error) {
	if node.prev == nil {
		return dll.RemoveFirst()
	}

	if node.next == nil {
		return dll.RemoveLast()
	}

	node.next.prev = node.prev
	node.prev.next = node.next

	dll.size--
	return node.value, nil
}

func (dll *DoublyLinkedList[T]) traverseTo(index int) *node[T] {
	if index < dll.size/2 {
		current := dll.head

		for i := 0; i != index; i++ {
			current = current.next
		}

		return current
	} else {
		current := dll.tail

		for i := dll.size - 1; i != index; i-- {
			current = current.prev
		}

		return current
	}
}

func (dll *DoublyLinkedList[T]) checkIndex(index int) error {
	if index < 0 || index > dll.size {
		return errors.New("Index out of range")
	}

	return nil
}

func (dll *DoublyLinkedList[T]) IsEmpty() bool {
	return dll.size == 0
}

func (dll *DoublyLinkedList[T]) Print() error {
	current := dll.head

	if current == nil {
		return EmptyListError
	}

	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%v ", current.value)
		current = current.next
	}
	fmt.Println()

	return nil
}
