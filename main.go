package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}

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

type triangle struct {
	height float32
	base   float32
}

func (t triangle) area() float32 {
	return t.height * t.base / 2
}

type calculator struct {
	total float32
}

func (c calculator) sumAreas(shapes ...shape) float32 {
	var sum float32
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{sideLen: 2}
	t := triangle{height: 10, base: 5}

	calc := calculator{}

	fmt.Println(calc.sumAreas(c, s, t))
}
