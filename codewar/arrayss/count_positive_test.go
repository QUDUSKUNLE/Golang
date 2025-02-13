package arrayss

import (
	"testing"
)

func TestCountPositiveSumNegatives(t *testing.T) {
	got := CountPositiveSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, -8, -9, -10})
	want := []int{7, -27}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
