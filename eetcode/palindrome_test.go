package eetcode

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	got := IsPalindrome(121)
	want := true
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
