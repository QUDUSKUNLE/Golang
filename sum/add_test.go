package integers

import "testing"

func TestAdd(t *testing.T) {
	assert_Functions := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Testing add function", func(t*testing.T) {
		assert_Functions(t, Add(3, 4), 7)
	})
}

func TestSum(t *testing.T) {
	assert_Functions := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Testing sum function", func(t *testing.T) {
		assert_Functions(t, Sum([]int{1,2,3,4,5,6}), 21)
	})
}
