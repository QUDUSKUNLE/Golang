package main

import "testing"

func TestSum(t *testing.T) {
	assert_Array_Sum_Function := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	
	t.Run("Testing sum function", func(t *testing.T) {
		number := []int{1, 2, 3, 4, 5}
		want := 15
		assert_Array_Sum_Function(t, Sum(number), want)
	})

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}
		want := 15
		got := Sum(numbers)
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Collection of any size", func(t*testing.T)  {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
