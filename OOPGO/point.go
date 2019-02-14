package main

import "fmt"

// Point is a 2d point
type Point struct {
	X, Y int
}

// Move movs the point
func (p Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func (mp *Point) MovePointer(dx int, dy int) {
	mp.X += dx
	mp.Y += dy
}

func main() {
	p := Point{1, 2}
	p.Move(2, 3)
	fmt.Printf("%+v\n", p) // {X:1, Y:2}

	mp := &Point{2, 3}
	mp.MovePointer(1, 7)
	fmt.Printf("%+v\n", mp) // {X:1, Y:2}
}
