package structure

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func (r Rectangle) AreaSecond() float64 {
	return r.Width * r.Height
}

func (c Circle) AreaSecond() float64 {
	return math.Pi * c.Radius * c.Radius
}
