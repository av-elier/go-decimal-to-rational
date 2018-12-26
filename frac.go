package dectofrac

import "math/big"

// NewRat returns rational from decimal using iterations number of iterations in Continued Fraction algorythm
func NewRat(val *big.Float, iterations int) *big.Rat {
	res, _ := val.Rat(nil)
	return res
}
