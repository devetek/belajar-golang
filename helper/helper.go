package helper

import (
	"math/big"
)

func FloatEqualizer(a float64, b float64) bool {
	result := big.NewFloat(a).Cmp(big.NewFloat(b))

	if result == 0 {
		return true
	}

	return false
}
