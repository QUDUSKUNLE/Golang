package iterations

import "testing"

func TestRepeat(t *testing.T) {
	assert_Functions := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Testing repeat function", func(t *testing.T) {
		assert_Functions(t, Repeat("a"), "aaaaa")
	})
}

// How to run Bench mark repeat.
// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("aaaaaaa")
	}
}
