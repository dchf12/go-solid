package main

import (
	"fmt"
	"math"
)

type transport interface {
	getName() string
}

type vehicle struct {
	name string
}

func (c vehicle) getName() string {
	return c.name
}

type car struct {
	vehicle
	wheel int
	gates int
}

type motocycle struct {
	vehicle
	wheel int
}

type printer struct{}

func (p printer) printTransportName(t transport) {
	fmt.Println("name: ", t.getName())
}

type shape interface {
	area() float32
}

type object interface {
	shape
	volume() float32
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

type cube struct {
	sideLen float32
}

func (c cube) area() float32 {
	return c.sideLen * c.sideLen
}

func (c cube) volume() float32 {
	return c.sideLen * c.sideLen * c.sideLen
}

func areaSum(shapes ...shape) float32 {
	var sum float32
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func areaVolumeSum(shapes ...object) float32 {
	var sum float32
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
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

	v := vehicle{name: "Ford"}
	car := car{
		vehicle: vehicle{name: "car-name"},
		wheel:   4,
		gates:   2,
	}
	m := motocycle{
		vehicle: vehicle{name: "motocycle-name"},
		wheel:   2,
	}
	p := printer{}
	p.printTransportName(v)
	p.printTransportName(car)
	p.printTransportName(m)

	cube := cube{sideLen: 5}
	fmt.Printf("cube: %v, %v\n", cube.area(), cube.volume())
	fmt.Printf("%v\n", areaVolumeSum(cube))
	fmt.Printf("%v\n", areaSum(cube, s))
}
