package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Length float64
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (sq *Square) Area() float64 {
	return sq.Length * sq.Length
}

func PrintArea(s Shape) {
	fmt.Println("Area is :", s.Area())
}

func main() {
	rectangle := Rectangle{Length: 5.00, Width: 5.00}
	square := Square{Length: 5.00}
	circle := Circle{Radius: 5.00}

	PrintArea(&rectangle)
	PrintArea(&square)
	PrintArea(&circle)

}
