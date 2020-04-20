package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rect := Rectangle{20, 20}
		got := Perimeter(rect)
		want := 80.

		if got != want {
			t.Errorf("got %.2f but wanted %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rect := Rectangle{20, 20}
		got := Area(rect)
		want := 400.

		if got != want {
			t.Errorf("got %.2f but wanted %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := Area(circle)
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g but wanted %g", got, want)
		}
	})
}
