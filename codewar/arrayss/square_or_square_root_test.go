package arrayss

import (
	"testing"
)

func TestSquareOrSquareRoot(t *testing.T) {
	t.Run("Testing square or square root function method 1", func(t *testing.T) {
		got := SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1})
		want := []int{2, 9, 3, 49, 4, 1}
		for i := 0; i < len(got); i++ {
			if got[i] != want[i] {
				t.Errorf("got %v want %v", got, want)
			}	
		}
	})
	t.Run("Testing square or square root function method 2", func(t *testing.T) {
		assertFunction := func(t *testing.T, got, want []int) {
			t.Helper()
			for i := 0; i < len(got); i++ {
				if got[i] != want[i] {
					t.Errorf("got %v want %v", got, want)
				}
			}
		}
		assertFunction(t, SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1}), []int{2, 9, 3, 49, 4, 1})
	})
}
