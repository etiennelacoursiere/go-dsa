package vector2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	a := Vector2{x: 16.0, y: 24.0}
	b := Vector2{x: 32.0, y: 32.0}

	c := a.Add(b)

	fmt.Printf("%#v\n", c)

	expected := &Vector2{x: 48.0, y: 56.0}
	got := c

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected: %v, Got: %v", expected, got)
	}
}
