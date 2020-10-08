package shape

import (
	"fmt"
	"testing"

	"github.com/devetek/Go-Interface/helper"
)

var r = Rectangle{5, 5}

func TestRectangleArea(t *testing.T) {
	var testNumb float64 = 25.000000

	aRect := r.Area()

	fmt.Println(aRect)

	if !helper.FloatEqualizer(aRect, testNumb) {
		t.Errorf("TestRectangleArea got %f want %f", aRect, testNumb)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	var testNumb float64 = 20.000000

	pRect := r.Perimeter()

	if !helper.FloatEqualizer(pRect, testNumb) {
		t.Errorf("TestRectanglePerimeter got %f want %f", pRect, testNumb)
	}
}
