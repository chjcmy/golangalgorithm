package _interface

import (
	"fmt"
	"math"
	"testing"
)

type Model interface {
	area() float32
	perimeter() float32
}

func (r *round) area() float32 {
	return (math.Pi*r.half*r.half)*2 + (2*math.Pi*r.half)*r.high
}

func (r *round) perimeter() float32 {
	return (math.Pi * r.half * r.half) * r.high
}

func (s *rightsix) area() float32 {
	return 2*s.a*s.b + 2*s.a*s.c + 2*s.b*s.c
}

func (s *rightsix) perimeter() float32 {
	return s.a * s.b * s.c
}

type round struct {
	half float32
	high float32
}

type rightsix struct {
	a float32
	b float32
	c float32
}

func TestInterface(t *testing.T) {
	round1 := round{10, 10}
	round2 := round{4.2, 15.6}
	rightsix1 := rightsix{10.5, 20.2, 20}
	rightsix2 := rightsix{4, 10, 23}

	printMeasure(&round1, &round2, &rightsix1, &rightsix2)
}

func printMeasure(m ...Model) {
	for _, val := range m {
		fmt.Printf("%0.2f, %0.2f\n", val.area(), val.perimeter())
	}
}
