package dynamic_array

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var defaultCapacity = 16

type DataType interface {
	constraints.Ordered
}

type DynamicArray[T DataType] struct {
	array    []T
	size     int
	capacity int
}

func (da *DynamicArray[T]) Get(index int) (error, any) {
	err := da.checkIndex(index)
	if err != nil {
		return err, nil
	}

	return nil, da.array[index]
}

func (da *DynamicArray[T]) Set(index int, element T) error {
	err := da.checkIndex(index)
	if err != nil {
		return err
	}

	da.array[index] = element

	return nil
}

func (da *DynamicArray[T]) Add(element T) {
	if da.size == da.capacity {
		da.growCapacity()
	}

	da.array[da.size] = element
	da.size++
}

func (da *DynamicArray[T]) Remove(index int) error {
	err := da.checkIndex(index)

	if err != nil {
		return err
	}

	da.array = append(da.array[:index], da.array[index+1:]...)
	da.size--

	return nil
}

func (da *DynamicArray[T]) GetData() []T {
	return da.array[:da.size]
}

func (da *DynamicArray[T]) checkIndex(index int) error {
	if index >= da.size || index < 0 {
		return errors.New("Index out of range")
	}
	return nil
}

func (da *DynamicArray[T]) growCapacity() {
	if da.capacity == 0 {
		da.capacity = defaultCapacity
	} else {
		da.capacity = da.capacity * 2
	}

	newArray := make([]T, da.capacity)
	copy(newArray, da.array)
	da.array = newArray
}
