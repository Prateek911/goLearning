package main

import "fmt"

type Rectangle struct {
	width  float64
	length float64
}

func (r *Rectangle) findArea() float64 {
	return r.width * r.length
}

func main() {
	rect := &Rectangle{width: 5.00, length: 5.00}
	fmt.Println("Area is: ", rect.findArea())
}
