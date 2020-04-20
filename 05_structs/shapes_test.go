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
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{20, 20}, 400},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := Area(tt.shape)
		if got != tt.want {
			t.Errorf("expected area %g, got %g", tt.want, got)
		}
	}
}
