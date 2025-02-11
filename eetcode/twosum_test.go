package eetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	// One method of writing test
	assert_Functions := func(t *testing.T, got, want []int) {
		t.Helper()
		if got[0] != want[0] || got[1] != want[1] {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Testing two sum function 1", func(t *testing.T) {
		assert_Functions(t, TwoSum([]int{1, 2, 3}, 3), []int{0, 1})
	})

	// Second method of writing test
	t.Run("Testing two sum function 2", func(t *testing.T) {
		got := TwoSum([]int{1, 2, 3}, 4)
		want := []int{0, 2}
		if got[0] != want[0] || got[1] != want[1] {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
