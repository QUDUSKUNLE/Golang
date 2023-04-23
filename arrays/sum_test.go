package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assert_Array_Sum_Function := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	
	t.Run("testing sum function", func(t *testing.T) {
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

func TestSumAll(t *testing.T) {
	t.Run("run SumAll test", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
	
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSum(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}
