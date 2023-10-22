package linkedList_test

import (
	_ "fmt"
	"golangCourse/linkedList"
	"testing"
)

func TestCreation(t *testing.T) {
	mockCapacity := 10
	list := linkedList.New(mockCapacity)

	pointer := list.Head

	for pointer.Next != nil {
		// Проверяем значения в нодах
		if pointer.Val != 0 {
			t.Errorf("created without default values")
		}
		pointer = pointer.Next
	}
	// Проверяем количество нод
	if list.Size() != mockCapacity {
		t.Errorf("created with wrong capacity")
	}

}

func TestAddition(t *testing.T) {
	mockValues := []int{1, 2, 3, 4, 5}
	list := linkedList.NewFromSlice(mockValues)

	// Пропускаем первый элемент с 0
	pointer := list.Head
	idx := 0
	// Проверяем значения
	for pointer.Next != nil {
		if (pointer.Val != mockValues[idx]) {
			t.Errorf("added wrong value")
		}
		pointer = pointer.Next
		idx++
	}
}

func TestPop(t *testing.T) {
	mockValues := []int{1, 2, 3, 4, 5}
	popedSize := len(mockValues) - 1
	list := linkedList.NewFromSlice(mockValues)
	// Удаляем последний элемент
	list.Pop()
	// Проверяем количество
	if (list.Size() != popedSize) {
		t.Errorf("wrong amount of nodes")
	}

	// Проверяем значения
	pointer := list.Head
	idx := 0
	for pointer.Next != nil {
		if (pointer.Val != mockValues[idx]) {
			t.Errorf("wrong values")
		}
		pointer = pointer.Next
		idx++
	}
}

func TestAt(t *testing.T) {
	mockValues := []int{1, 2, 3, 4, 5}
	list := linkedList.NewFromSlice(mockValues)
	// Проверяем значения
	for idx, value := range mockValues {
		res, err := list.At(idx)
		if (res != value || err != nil) {
			t.Errorf("wrong value or index")
		}
	}
}

func TestDeleteFrom(t *testing.T) {
	mockValues := []int{1, 2, 3, 4, 5}
	list := linkedList.NewFromSlice(mockValues)
	
	list.DeleteFrom(3)
	res, _ := list.At(3)
	if (res != 5) {
		t.Errorf("Deleted wrong value")
	}
	list.DeleteFrom(0)
	res, _ = list.At(0)
	if (res != 2) {
		t.Errorf("Deleted wrong value")
		}	
	}
	
	func TestUpdateAt(t *testing.T) {
	mockValues := []int{1, 2, 3, 4, 5}
	updatedValues := []int{10, 20, 30, 40, 50}
	list := linkedList.NewFromSlice(mockValues)
	
	for idx, value := range updatedValues {
		list.UpdateAt(idx, value)
	}

	idx := 0
	pointer := list.Head
	for pointer.Next != nil {
		res, _ := list.At(idx)
		if (res != updatedValues[idx]) {
			t.Errorf("Wrong values after update")
		}
		pointer = pointer.Next
		idx++
	}
}