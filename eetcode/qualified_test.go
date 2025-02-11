package eetcode

import (
	"testing"
	"fmt"
)

func TestReturnDays(t *testing.T) {
	assert_Functions := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Testing return days function 1", func(t *testing.T) {
		assert_Functions(t, fmt.Sprintf("%v", ReturnDays([]map[string]string{})), fmt.Sprintf("%v", []map[string]string{}))
	})
}
