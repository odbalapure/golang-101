package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	length float64
}

func (s square) getArea() float64 {
	return float64(s.length * s.length)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func getAllTypesOfArea(s shape) {
	fmt.Printf("Area of %v is %v\n", reflect.TypeOf(s).Name(), s.getArea())
}

func main() {
	sq := square{
		length: 10.5,
	}
	tr := triangle{
		base:   10.2,
		height: 23.3,
	}

	shapes := []shape{tr, sq}
	for _, sh := range shapes {
		fmt.Println("getArea::Area:", sh.getArea())
	}

	getAllTypesOfArea(sq)
	getAllTypesOfArea(tr)
}
