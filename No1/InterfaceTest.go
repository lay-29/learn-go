package No1

import "fmt"

func InterfaceTest() {
	rect := NewRectangle(10, 5)
	circle := NewCrile(5)
	square := NewSquare(5)
	PrintShapeInfo(rect)
	PrintShapeInfo(circle)
	PrintShapeInfo(square)
}

func PrintShapeInfo(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter", s.Perimeter())
}
