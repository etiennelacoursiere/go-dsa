package linked_list

import (
	"reflect"
	"testing"
)

func TestInsertLast(t *testing.T) {
	l := LinkedList[int]{}

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

func TestInsertFirst(t *testing.T) {
	l := LinkedList[int]{}

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

func TestRemoveFirst(t *testing.T) {
	l := LinkedList[int]{}

	l.InsertLast(3)
	l.InsertLast(1)
	l.InsertLast(2)
	l.RemoveFirst()

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

func TestRemoveLast(t *testing.T) {
	l := LinkedList[int]{}

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)
	l.RemoveLast()

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
				t.Log(reflect.DeepEqual(tt.got, tt.want))
				t.Fatalf("Got %v, Expected: %v", tt.got, tt.want)
			}
		})
	}
}

func TestInsertAt(t *testing.T) {
	l := LinkedList[int]{}

	l.InsertLast(1)
	l.InsertAt(1, 2)

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
				t.Log(reflect.DeepEqual(tt.got, tt.want))
				t.Fatalf("Got %v, Expected: %v", tt.got, tt.want)
			}
		})
	}
}

func TestRemoveAt(t *testing.T) {
	l := LinkedList[int]{}

	l.InsertLast(1)
	l.InsertLast(3)
	l.InsertLast(2)
	l.RemoveAt(1)

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
				t.Log(reflect.DeepEqual(tt.got, tt.want))
				t.Fatalf("Got %v, Expected: %v", tt.got, tt.want)
			}
		})
	}
}
