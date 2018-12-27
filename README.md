# go-decimal-to-rational

[![Build Status](https://travis-ci.org/av-elier/go-decimal-to-rational.svg?branch=master)](https://travis-ci.org/av-elier/go-decimal-to-rational)

Go library to convert decimal (float64) to rational fraction with required precision

Relies on [Continued Fraction](http://mathworld.wolfram.com/ContinuedFraction.html) algorythm.

It's sometimes more appropriate than default big.Rat SetString, because
you can get `2/3` from `0.6666` by specifiing required precision. In big.Rat SetString
you can only get `3333/50000`, and have no way to manipulate than (as of go 1.11).

# Example
```go
func ExampleNewRatPrecision() {
	fmt.Println(NewRatPrecision(0.6666, 0.01).String())
	fmt.Println(NewRatPrecision(0.981, 0.001).String())
	fmt.Println(NewRatPrecision(0.75, 0.01).String())
	// Output:
	// 2/3
	// 981/1000
	// 3/4
}
```
```go
func ExampleNewRatIterations() {
	fmt.Println(NewRatIterations(0.6667, 3).String())
	fmt.Println(NewRatIterations(0.6667, 4).String())
	// Output:
	// 2/3
	// 6667/10000
}
```

# Docs
```
import dectofrac "github.com/av-elier/go-decimal-to-rational"
```

#### func NewRatIterations

```go
func NewRatIterations(val float64, iterations int64) *big.Rat
```
NewRatIterations returns rational from decimal using `iterations` number of
iterations in Continued Fraction algorythm

#### func NewRatPrecision

```go
func NewRatPrecision(val float64, precision float64) *big.Rat
```
NewRatPrecision returns rational from decimal by going as mush iterations, until
next fraction is less than precision
