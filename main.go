package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

func (c circle) area() {
	fmt.Printf("circle area: %f\n", c.radius*c.radius*math.Pi)
}

type square struct {
	sideLen float32
}

func (s square) area() {
	fmt.Printf("square area: %f\n", s.sideLen*s.sideLen)
}
func main() {
	c := circle{radius: 5}
	c.area()
	s := square{sideLen: 2}
	s.area()
}
