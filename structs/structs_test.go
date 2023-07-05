package structs

import "testing"

func TestPerimeter(t *testing.T) {
	assert_Functions := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("testing a rectangle perimeter function", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		assert_Functions(t, rectangle.Perimeter(), 40.0)
	})
	t.Run("testing a rectangle area function", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		assert_Functions(t, rectangle.Area(), 72.0)
	})
	t.Run("testing a circle area function", func(t *testing.T) {
		circle := Circle{10}
		assert_Functions(t, circle.Area(), 314.1592653589793)
	})

	t.Run("testing a circle perimeter function", func(t *testing.T) {
		circle := Circle{10.0}
		assert_Functions(t, circle.Perimeter(), 62.83185307179586) 
	})

	t.Run("testing a triangle area function", func(t *testing.T) {
		triangle := Triangle{10, 10}  
		assert_Functions(t, triangle.Area(), 50)
	})

	t.Run("testing a triangle perimeter function", func(t *testing.T) {
		triangle := Triangle{10.0, 10}
		assert_Functions(t, triangle.Perimeter(), 20.0)
	})
}
