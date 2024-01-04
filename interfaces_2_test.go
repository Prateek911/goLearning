package main

import (
	"math"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Length: 5.00, Width: 5.00}
	expectedArea := 25.00
	if area := rect.Area(); area != expectedArea {
		t.Errorf("Rectangle area calculation is incorrect, got: %.2f, want: %.2f", area, expectedArea)
	}
}

func TestSquareArea(t *testing.T) {
	sq := Square{Length: 5.00}
	expectedArea := 25.00
	if area := sq.Area(); area != expectedArea {
		t.Errorf("Square area calculation is incorrect, got: %.2f, want: %.2f", area, expectedArea)
	}
}

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 5.00}
	expectedArea := math.Pi * 25.00
	if area := circle.Area(); area != expectedArea {
		t.Errorf("Circle area calculation is incorrect, got: %.2f, want: %.2f", area, expectedArea)
	}
}
