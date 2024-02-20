package structsmethodsinterfaces

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

func (r Rectangle) Area() (area float64) {
	return
}

func (c Circle) Area() (area float64) {
	return
}

func Perimeter(r Rectangle) (perimeter float64) {
	perimeter = 2 * (r.width + r.height)
	return
}
