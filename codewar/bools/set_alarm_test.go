package bools

import (
	"testing"
)

func TestSetAlarm(t *testing.T) {
	t.Run("Testing set alarm function method 1", func(t *testing.T) {
		got :=  SetAlarm(true, true)
		want := false
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Testing set alarm function method 2", func(t *testing.T) {
		assertFunction := func(t *testing.T, got, want bool) {
			t.Helper()
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		}
		assertFunction(t, SetAlarm(false, true), false)
	})
}
