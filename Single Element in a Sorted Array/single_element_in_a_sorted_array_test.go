package Single_Element_in_a_Sorted_Array

import (
	"fmt"
	"testing"
)

func TestSingleNonDuplicate(t *testing.T) {
	tests := []struct {
		arr           []int
		singleElement int
	}{
		{arr: []int{1, 1, 2, 3, 3, 4, 4, 8, 8}, singleElement: 2},
		{arr: []int{3, 3, 7, 7, 10, 11, 11}, singleElement: 10},
		{arr: []int{2, 1, 1, 3, 3, 4, 4, 8, 8}, singleElement: 2},
		{arr: []int{1, 1, 3, 3, 4, 4, 8, 8, 2}, singleElement: 2},
		{arr: []int{1, 1, 3}, singleElement: 3},
		{arr: []int{3, 1, 1}, singleElement: 3},
		{arr: []int{3}, singleElement: 3},
	}
	for _, tst := range tests {
		answer := singleNonDuplicate(tst.arr)
		if answer != tst.singleElement {
			arrStr := fmt.Sprint(tst.arr)
			t.Errorf(
				"%d not single in %s, single - %d",
				answer, arrStr, tst.singleElement,
			)
		}
	}
}
