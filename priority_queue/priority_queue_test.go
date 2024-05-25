package priority_queue

import (
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	pq := PriorityQueue[int]{}
	values := []int{5, 4, 1, 2, 10}
	pq.Fill(values)
	assert(t, pq.Size(), 5)
	assert(t, pq.heap[0], 1)
}

func TestAdd(t *testing.T) {
	pq := PriorityQueue[int]{}
	pq.Add(7)
	pq.Add(5)
	pq.Add(99)

	assert(t, pq.Size(), 3)
	assert(t, pq.heap[0], 5)

	pq.Add(0)
	assert(t, 0, pq.heap[0])
}

func TestPeek(t *testing.T) {
	pq := PriorityQueue[int]{}
	pq.Add(100)
	pq.Add(50)
	peekedValue, err := pq.Peek()
	assert(t, peekedValue, 50)
	assert(t, peekedValue, 50)
	assert(t, pq.Size(), 2)
	assert(t, err, nil)

	pq.Clear()
	_, err = pq.Peek()
	assert(t, err, isEmptyError)
}

func TestRemove(t *testing.T) {
	pq := PriorityQueue[int]{}
	values := []int{5, 4, 1, 2, 10}
	pq.Fill(values)
	isRemoved, err := pq.Remove(1)

	assert(t, isRemoved, true)
	assert(t, pq.heap[0], 2)
	assert(t, pq.Size(), 4)

	isRemoved, _ = pq.Remove(4)
	assert(t, isRemoved, true)
	assert(t, pq.heap[0], 2)
	assert(t, pq.Size(), 3)

	_, err = pq.Remove(4)
	assert(t, err, notFoundError)

	pq.Clear()
	_, err = pq.Remove(0)
	assert(t, err, isEmptyError)
}

func TestPoll(t *testing.T) {
	pq := PriorityQueue[int]{}
	pq.Fill([]int{5, 4, 1})

	polledValues := []int{}
	for !pq.IsEmpty() {
		polledValue, _ := pq.Poll()
		polledValues = append(polledValues, polledValue)
	}

	assert(t, polledValues, []int{1, 4, 5})

	pq.Clear()
	_, err := pq.Poll()
	assert(t, err, isEmptyError)
}

func TestContains(t *testing.T) {
	pq := PriorityQueue[int]{}
	pq.Fill([]int{5, 4, 1, 100, 8})

	assert(t, pq.Contains(1), true)
	assert(t, pq.Contains(4), true)
	assert(t, pq.Contains(9), false)
	assert(t, pq.Size(), 5)
}

func TestClear(t *testing.T) {
	pq := PriorityQueue[int]{}
	pq.Fill([]int{5, 4, 1, 100, 8})

	assert(t, pq.Size(), 5)

	pq.Clear()
	assert(t, pq.Size(), 0)
}

func assert(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got %v, Expected: %v", got, want)
	}
}
