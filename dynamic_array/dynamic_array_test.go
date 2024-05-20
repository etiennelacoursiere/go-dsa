package dynamic_array

import (
	"reflect"
	"testing"
)

func TestAddInt(t *testing.T) {
	da := DynamicArray[int]{}

	da.Add(0)
	da.Add(1)
	da.Add(2)
	da.Add(3)

	expected := []int{0, 1, 2, 3}
	got := da.GetData()

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected: %v, Got: %v", expected, got)
	}
}

func TestAddString(t *testing.T) {
	da := DynamicArray[string]{}

	da.Add("salut")
	da.Add("bonjour")

	expected := []string{"salut", "bonjour"}
	got := da.GetData()

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected: %v, Got: %v", expected, got)
	}
}

func TestRemove(t *testing.T) {
	da := DynamicArray[int]{}

	da.Add(0)
	da.Add(1)
	da.Add(2)
	da.Add(3)
	da.Remove(2)

	expected := []int{0, 1, 3}
	got := da.GetData()

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected: %v, Got: %v", expected, got)
	}
}

func TestGet(t *testing.T) {
	da := DynamicArray[int]{}

	da.Add(0)
	da.Add(1)
	da.Add(2)
	da.Add(3)

	expected := 3
	err, got := da.Get(3)

	if err != nil || !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected: %v, Got: %v", expected, got)
	}
}

func TestSet(t *testing.T) {
	da := DynamicArray[int]{}

	da.Add(0)
	da.Add(1)
	da.Add(2)
	da.Add(3)
	da.Set(2, 10)

	expected := []int{0, 1, 10, 3}
	got := da.GetData()

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected: %v, Got: %v", expected, got)
	}
}
