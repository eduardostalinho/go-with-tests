package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		got := Sum(numbers)
		want := 28

		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6, 7})
	want := []int{6, 22}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %d, got %d", want, got)
	}

}
