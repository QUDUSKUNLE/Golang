package arrayss

import (
	"testing"
)

func TestArrayDiff(t *testing.T) {
	t.Run("Testing array diff function method 1", func(t *testing.T) {
		got := ArrayDiff([]int{1, 2, 3}, []int{1})
		want := []int{2, 3}

		for i := 0; i < len(got); i++ {
			if got[i] != want[i] {
				t.Errorf("got %v want %v", got, want)
			}
		}
	})
	t.Run("Testing array diff function method 2", func(t *testing.T) {
		assertFunction := func(t *testing.T, got, want []int) {
			t.Helper()
			if got[0] != want[0] || got[1] != want[1] {
				t.Errorf("got %v want %v", got, want)
			}
		}
		assertFunction(t, ArrayDiff([]int{1, 2, 3}, []int{1}), []int{2, 3})
	})
}
