package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	expectedValues := []int{1, 2, 3, 4, 5}

	BubbleSort(values)

	for i, value := range values {
		if value != expectedValues[i] {
			t.Error("BubbleSort() failed. Got", values, "Expected", expectedValues)
		}
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	expectedValues := []int{1, 2, 3, 5, 5}

	BubbleSort(values)

	for i, value := range values {
		if value != expectedValues[i] {
			t.Error("BubbleSort() failed. Got", values, "Expected", expectedValues)
		}
	}
}


func TestBubbleSort3(t *testing.T) {
	values := []int{5}

	BubbleSort(values)

	if values[0] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 5")
	}
}
