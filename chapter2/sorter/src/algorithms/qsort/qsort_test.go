package qsort

import "testing"

func TestQuickSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	expectedValues := []int{1, 2, 3, 4, 5}

	QuickSort(values)

	for i, value := range values {
		if value != expectedValues[i] {
			t.Error("TestQuickSort() failed. Got", values, "Expected", expectedValues)
		}
	}
}

func TestQuickSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	expectedValues := []int{1, 2, 3, 5, 5}

	QuickSort(values)

	for i, value := range values {
		if value != expectedValues[i] {
			t.Error("TestQuickSort() failed. Got", values, "Expected", expectedValues)
		}
	}
}


func TestQuickSort3(t *testing.T) {
	values := []int{5}

	QuickSort(values)

	if values[0] != 5 {
		t.Error("TestQuickSort() failed. Got", values, "Expected 5")
	}
}
