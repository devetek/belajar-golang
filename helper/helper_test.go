package helper

import (
	"testing"
)

func TestFloatEqualizerIsEqual(t *testing.T) {
	var eqParams float64 = 314.1592653589793

	isEqual := FloatEqualizer(eqParams, eqParams)

	if !isEqual {
		t.Errorf("TestFloatEqualizerIsEqual got %v want %v", isEqual, true)
	}
}

func TestFloatEqualizerIsNotEqual(t *testing.T) {
	var p1 float64 = 314.1592653589793
	var p2 float64 = 314.1592653589791

	isNotEqual := FloatEqualizer(p1, p2)

	if isNotEqual {
		t.Errorf("TestFloatEqualizerIsNotEqual got %v want %v", isNotEqual, true)
	}
}
