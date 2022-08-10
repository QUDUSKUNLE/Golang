package main

import "testing"

func TestHome(t *testing.T) {
	t.Run("Hey, say Hello to Qudus Yekeen", func(t *testing.T) {
		got := Hello("World", "Yoruba")
		want := "Hello, World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Chris", "Yoruba")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "Yoruba")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying Hello to people in a loop", func(t *testing.T) {
		got := Hello("Chris", "Yoruba")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "Yoruba")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("testing adder function", func(t *testing.T) {
		sum := Add(2, 4)
		expected := 6

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

