package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{8, 1, 4, 3, 6, 5, 2, 9, 7}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := mergeSort(arr)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}
