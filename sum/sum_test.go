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
	assert_Sum_Function := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Testing sum function", func(t *testing.T) {
		assert_Sum_Function(t, Sum([]int{1,2,3,4,5,6}), 21)
	})
}

func TestExampleAdd(t *testing.T) {
	assert_Example_Add_Function := func(t *testing.T, sum, got int ) {
		t.Helper()
		if sum != got {
			t.Errorf("sum '%d' got '%d'", sum, got)
		}
	}

	t.Run("Testing ExampleAdd function", func(t *testing.T) {
		assert_Example_Add_Function(t, ExampleAdd(2, 10), 12)
	})
}
