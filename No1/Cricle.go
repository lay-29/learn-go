package No1

import "math"

type Circle struct {
	Radius float32
}

func NewCrile(radius float32) *Circle {
	return &Circle{radius}
}

func (c *Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}
