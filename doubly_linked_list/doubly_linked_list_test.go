package doubly_linked_list

import "testing"

func TestInsertFirst(t *testing.T) {
	l := DoublyLinkedList[int]{}

	l.InsertFirst(2)
	l.InsertFirst(1)

	var tests = []struct {
		name string
		want any
		got  any
	}{
		{"head should be 1", 1, l.head.value},
		{"head.next should be 2", 2, l.head.next.value},
		{"size should be 2", 2, l.size},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("Got %v, Expected: %v", tt.got, tt.want)
			}
		})
	}
}

func TestInsertLast(t *testing.T) {
	l := DoublyLinkedList[int]{}

	l.InsertLast(1)
	l.InsertLast(2)

	var tests = []struct {
		name string
		want any
		got  any
	}{
		{"head should be 1", 1, l.head.value},
		{"head.next should be 2", 2, l.head.next.value},
		{"size should be 2", 2, l.size},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("Got %v, Expected: %v", tt.got, tt.want)
			}
		})
	}
}

func TestRemoveFirst(t *testing.T) {
	l := DoublyLinkedList[int]{}

	l.InsertLast(3)
	l.InsertLast(1)
	l.InsertLast(2)
	removed, _ := l.RemoveFirst()

	var tests = []struct {
		name string
		want any
		got  any
	}{
		{"head should be 1", 1, l.head.value},
		{"head.next should be 2", 2, l.head.next.value},
		{"size should be 2", 2, l.size},
		{"removed should be 3", 3, removed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("Got %v, Expected: %v", tt.got, tt.want)
			}
		})
	}
}

func TestRemoveLast(t *testing.T) {
	l := DoublyLinkedList[int]{}

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)
	removed, _ := l.RemoveLast()

	var tests = []struct {
		name string
		want any
		got  any
	}{
		{"head should be 1", 1, l.head.value},
		{"head.next should be 2", 2, l.head.next.value},
		{"size should be 2", 2, l.size},
		{"removed should be 3", 3, removed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("Got %v, Expected: %v", tt.got, tt.want)
			}
		})
	}
}

func TestRemoveAt(t *testing.T) {
	l := DoublyLinkedList[int]{}

	l.InsertLast(1)
	l.InsertLast(3)
	l.InsertLast(2)
	removed, err := l.RemoveAt(1)

	var tests = []struct {
		name string
		want any
		got  any
	}{
		{"head should be 1", 1, l.head.value},
		{"head.next should be 2", 2, l.head.next.value},
		{"size should be 2", 2, l.size},
		{"removed should be 3", 3, removed},
		{"err should be nil", nil, err},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("Got %v, Expected: %v", tt.got, tt.want)
			}
		})
	}
}
