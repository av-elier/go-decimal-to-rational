package dectofrac

import (
	"math"
	"math/big"
)

// NewRatIterations returns rational from decimal
// using `iterations` number of iterations in Continued Fraction algorythm
func NewRatIterations(val float64, iterations int64) *big.Rat {
	return newRat(val, iterations, 0)
}

// NewRatPrecision returns rational from decimal
// by going as mush iterations, until next fraction is less than precision
func NewRatPrecision(val float64, precision float64) *big.Rat {
	return newRat(val, math.MaxInt64, precision)
}

func newRat(val float64, iterations int64, precision float64) *big.Rat {
	a0 := int64(math.Floor(val))
	x0 := val - float64(a0)
	rat := cf(x0, 1, iterations, precision)
	return rat.Add(rat, new(big.Rat).SetInt64(a0))
}

func cf(xi float64, i int64, limit int64, precision float64) *big.Rat {
	if i >= limit || xi <= precision {
		return big.NewRat(0, 1)
	}

	inverted := 1 / xi
	aj := int64(math.Floor(inverted))
	xj := inverted - float64(aj)
	ratAJ := new(big.Rat).SetInt64(aj)
	ratNext := cf(xj, i+1, limit, precision)
	res := ratAJ.Add(ratAJ, ratNext)
	res = res.Inv(res)

	return res
}
