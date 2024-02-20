package structsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func (r Rectangle) Area() (area float64) {
	area = r.height * r.width
	return
}

func (c Circle) Area() (area float64) {
	area = math.Pi * c.radius * c.radius
	return
}

func (t Triangle) Area() (area float64) {
	area = (t.base * t.height) / 2
	return
}

func Perimeter(r Rectangle) (perimeter float64) {
	perimeter = 2 * (r.width + r.height)
	return
}
