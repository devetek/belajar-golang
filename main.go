package main

import (
	"fmt"

	"github.com/devetek/Go-Interface/shape"
)

func main() {
	s := shape.Rectangle{5, 5}

	c := shape.Circle{10}

	pRectangle := s.Perimeter()
	aRectangle := s.Area()
	pCircle := c.Perimeter()
	aCircle := c.Area()

	fmt.Printf("Keliling persegi %0.2f\n", pRectangle)
	fmt.Printf("Keliling lingkaran %0.2f\n", pCircle)
	fmt.Printf("Area lingkaran %f\n", aCircle)
	fmt.Printf("Area persegi %f\n", aRectangle)
}
