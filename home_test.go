package main

import "testing"

func TestHome(t *testing.T) {

	assert_Hello_func := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assert_Add_func := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	assert_Subtract_func := func(t *testing.T, x, y int) {
		t.Helper()
		if x != y {
			t.Errorf("got %d want %d", x, y)
		}
	}

	t.Run("Hey, say Hello to Qudus Yekeen", func(t *testing.T) {
		got := Hello("World", "Yoruba")
		want := "Hello, World"

		assert_Hello_func(t, got, want)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Chris", "Yoruba")
		want := "Hello, Chris"

		assert_Hello_func(t, got, want)
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "Yoruba")
		want := "Hello, World"

		assert_Hello_func(t, got, want)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying Hello to people in a loop", func(t *testing.T) {
		got := Hello("Chris", "Yoruba")
		want := "Hello, Chris"

		assert_Hello_func(t, got, want)

	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "Yoruba")
		want := "Hello, World"
		assert_Hello_func(t, got, want)
	})

	t.Run("testing adder function", func(t *testing.T) {
		sum := Add(2, 4)
		expected := 6
		assert_Add_func(t, sum, expected)
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("testing adder function", func(t*testing.T) {
		summ := Add(2, 9)
		expected := 11
		assert_Add_func(t, summ, expected)
		if summ != expected {
			t.Fail()
			t.Errorf("expected '%d' but got '%d'", expected, summ)
		}
	})

	t.Run("testing subtract function", func(t*testing.T) {
		sub := Subtract(10, 5)
		expected := 5
		assert_Subtract_func(t, sub, expected)
		if sub != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sub)
		}
	})
}

