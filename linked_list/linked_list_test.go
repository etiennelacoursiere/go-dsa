package linked_list

import "testing"

func TestAppend(t *testing.T) {
	l := LinkedList{}

	l.Append(7)
	l.Append(11)
	l.Append(92)

	l.Print()

	if l.Head.value != 7 {
		t.Fatalf("Expected 7 got %v", l.Head.value)
	}

	if l.Head.next.value != 11 {
		t.Fatalf("Expected 11 got %v", l.Head.next.value)
	}

	if l.Head.next.next.value != 92 {
		t.Fatalf("Expected 92 got %v", l.Head.next.next.value)
	}

	if l.Size != 3 {
		t.Fatalf("Expected size to be 3 got %v", l.Size)
	}
}

func TestPrepend(t *testing.T) {
	l := LinkedList{}

	l.Append(2)
	l.Append(3)
	l.Prepend(1)

	l.Print()

	if l.Head.value != 1 {
		t.Fatalf("Expected 1 got %v", l.Head.value)
	}

	if l.Size != 3 {
		t.Fatalf("Expected size to be 3 got %v", l.Size)
	}
}

func TestRemoveFirst(t *testing.T) {
	l := LinkedList{}

	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.RemoveFirst()

	l.Print()

	if l.Head.value != 1 {
		t.Fatalf("Expected 1 got %v", l.Head.value)
	}

	if l.Size != 2 {
		t.Fatalf("Expected size to be 2 got %v", l.Size)
	}
}

func TestRemoveLast(t *testing.T) {
	l := LinkedList{}

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.RemoveLast()

	l.Print()

	if l.Head.next.next != nil {
		t.Fatalf("Expected nil got %v", l.Head.next.next)
	}

	if l.Size != 2 {
		t.Fatalf("Expected size to be 2 got %v", l.Size)
	}

}

func TestInsertAt(t *testing.T) {
	l := LinkedList{}

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(5)
	l.InsertAt(3, 4)

	l.Print()

	if l.Head.next.next.next.value != 4 {
		t.Fatalf("Expected 4 got %v", l.Head.next.next.value)
	}
	if l.Size != 5 {
		t.Fatalf("Expected size to be 5 got %v", l.Size)
	}
}

func TestRemoveAt(t *testing.T) {
	l := LinkedList{}

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.RemoveAt(1)

	l.Print()

	if l.Head.next.value != 3 {
		t.Fatalf("Expected 3 got %v", l.Head.next.value)
	}

	if l.Size != 2 {
		t.Fatalf("Expected size to be 2 got %v", l.Size)
	}

	l.RemoveAt(1)
	l.Print()

	if l.Head.next != nil {
		t.Fatalf("Expected nil got %v", l.Head.next.value)
	}

	if l.Size != 1 {
		t.Fatalf("Expected size to be 1 got %v", l.Size)
	}
}
