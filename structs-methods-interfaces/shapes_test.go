package structsmethodsinterfaces

import "testing"

type Shape interface {
	Area() float64
}

// type Square struct {
// 	width  float64
// 	height float64
// }

// type Circle struct {
// 	radius float64
// }

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0

	if got != want {
		t.Errorf("got %.f, want %.2f", got, want)
	}
}
