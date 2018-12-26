# go-decimal-to-rational
Go library to convert decimal (float64) to rational fraction with required precision

Relies on [Continued Fraction](http://mathworld.wolfram.com/ContinuedFraction.html) algorythm.

It's sometimes more appropriate than default big.Rat SetString, because
you can get `2/3` from `0.6666` by specifiing required precision. In big.Rat SetString
you can only get `3333/50000`, and have no way to manipulate than (as of go 1.11).

# Example
```go
TODO
```