package No1

type Rectangle struct {
	Width  float32
	Height float32
}

func NewRectangle(width, height float32) *Rectangle {
	return &Rectangle{width, height}
}

func (r *Rectangle) Area() float32 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float32 {
	return 2 * (r.Width + r.Height)
}

type Square struct {
	*Rectangle
}

func NewSquare(sideLen float32) *Square {
	return &Square{
		Rectangle: NewRectangle(sideLen, sideLen),
	}
}
