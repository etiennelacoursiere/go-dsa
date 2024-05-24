package queue

import (
	"reflect"
	"testing"
)

type Test struct {
	name string
	want any
	got  any
}

func TestEnqueue(t *testing.T) {
	q := Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	var tests = []Test{
		{"head should be 1", 1, q.head.value},
		{"tail should be 3", 3, q.tail.value},
		{"size should be 3", 3, q.size},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert(t, tt.got, tt.want)
		})
	}
}

func TestDequeue(t *testing.T) {
	q := Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	shouldBe1, _ := q.Dequeue()
	shouldBe2, _ := q.Dequeue()
	shouldBe3, _ := q.Dequeue()
	_, err := q.Dequeue()

	var tests = []Test{
		{"should be 1", 1, shouldBe1},
		{"should be 3", 2, shouldBe2},
		{"should be 3", 3, shouldBe3},
		{"should be 3", 0, q.size},
		{"Should have an error", queueEmptyError, err},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert(t, tt.got, tt.want)
		})
	}
}

func TestPeek(t *testing.T) {
	q := Queue[int]{}

	_, err := q.Peek()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	peek1, _ := q.Peek()
	peek2, _ := q.Peek()

	var tests = []Test{
		{"Should have an error", queueEmptyError, err},
		{"should be 1", 1, peek1},
		{"should be 3", 1, peek2},
		{"should be 3", 3, q.size},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert(t, tt.got, tt.want)
		})
	}
}

func TestContains(t *testing.T) {
	q := Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	var tests = []Test{
		{"should not contain 0", false, q.Contains(0)},
		{"should contain 1", true, q.Contains(1)},
		{"should contain 2", true, q.Contains(2)},
		{"should contain 3", true, q.Contains(3)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert(t, tt.got, tt.want)
		})
	}
}

func TestFlush(t *testing.T) {
	q := Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	queue := q.Flush()

	var tests = []Test{
		{"Flushed queue should be:", []int{1, 2, 3}, queue},
		{"q should be empty", true, q.IsEmpty()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert(t, tt.got, tt.want)
		})
	}
}

func assert(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, Expected: %v", got, want)
	}
}
