package stack

import "testing"

func TestPush(t *testing.T) {
	stack := &Stack[int]{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size() != 3 {
		t.Fatalf("Expected stack size to be 3. Got: %d", stack.Size())
	}
}

func TestPop(t *testing.T) {
	stack := &Stack[int]{}

	stack.Push(1)
	stack.Push(2)

	value, err := stack.Pop()

	if err != nil {
		t.Fatal(err)
	}

	if value != 2 {
		t.Fatalf("Expected returning value to be 2. Got: %d", value)
	}

	if stack.Size() != 1 {
		t.Fatalf("Expected stack size to be 1. Got: %d", stack.Size())
	}

	stack.Pop()
	_, err = stack.Pop()

	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
}

func TestPeek(t *testing.T) {
	stack := &Stack[int]{}

	stack.Push(1)
	stack.Push(2)

	value, err := stack.Peek()

	if err != nil {
		t.Fatal(err)
	}

	if value != 2 {
		t.Fatalf("Expected returning value to be 2. Got: %d", value)
	}

	if stack.Size() != 2 {
		t.Fatalf("Expected stack size to be 2. Got: %d", stack.Size())
	}
}
