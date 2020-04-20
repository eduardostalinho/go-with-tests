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
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{20, 20}, 400},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Area(tt.shape)
			if got != tt.want {
				t.Errorf("%#v expected area %g, got %g", tt.shape, tt.want, got)
			}
		})
	}
}
