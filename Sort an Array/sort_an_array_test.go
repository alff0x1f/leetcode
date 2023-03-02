package Sort_an_Array

import (
	"fmt"
	"testing"
)

func TestSortArray(t *testing.T) {
	tests := []struct {
		array  []int
		sorted []int
	}{
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
		{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, test := range tests {
		test.array = sortArray(test.array)
		if len(test.array) != len(test.sorted) {
			t.Errorf("Sorted array must contain %d elements", len(test.sorted))
			continue
		}
		for i := 0; i < len(test.sorted); i++ {
			if test.array[i] != test.sorted[i] {
				arrStr := fmt.Sprint(test.array)
				t.Errorf("array %s Wrong sorted on position %d", arrStr, i)
			}
		}
	}
}
