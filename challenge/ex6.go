package main

import (
	"fmt"
	"log"
)

type Square struct {
	Center Point
	Length int
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0")
	}

	s := Square{Center: Point(x, y), Length: length}

	return s, nil
}

func main() {
	s, err := NewSquare(1, 1, 20)
	if err != nil {
		log.Fatalf("Error: Can't create square")
	}
	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Printf(s.Area())
}
