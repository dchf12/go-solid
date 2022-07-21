package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	sideLen float32
}

func (s square) area() float32 {
	return s.sideLen * s.sideLen
}

type shape interface {
	area() float32
}
type outPrinter struct{}

func (op outPrinter) toText(s shape) string {
	return fmt.Sprintf("Area: %v\n", s.area())
}

func main() {
	c := circle{radius: 5}
	c.area()
	s := square{sideLen: 2}
	s.area()
}
