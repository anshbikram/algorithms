package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{1, 7, 34, 89, 2, 5, 5, -50, 28, 12}
	expected := []int{-50, 1, 2, 5, 5, 7, 12, 28, 34, 89}
	actual := MergeSort(nums)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual %v, expected %v", actual, expected)
	}
}
