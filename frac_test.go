package dectofrac

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRatIterations(t *testing.T) {
	type testCase struct {
		expected   string
		float      float64
		iterations int64
	}
	tests := []testCase{
		{"1/4", 0.25, 4},
		{"1/4", 0.25, 100},
		{"2/3", 2 / 3.0, 100},
		{"2/3", 0.6666, 4},
		{"3331/4997", 0.6666, 5},
		{"2/3", 0.6667, 3},
		{"6667/10000", 0.6667, 4},
		{"1/7", 1 / 7.0, 7},
		{"99/101", 99 / 101.0, 5},
		{"108851651149874/111050674405427", 99 / 101.0, 6},
	}
	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			assert.Equal(t, test.expected, NewRatIterations(test.float, test.iterations).String())
		})
	}
}

func TestNewRatPrecision(t *testing.T) {
	type testCase struct {
		expected  string
		float     float64
		precision float64
	}
	tests := []testCase{
		{"1/4", 0.25, 1e-3},
		{"1/4", 0.25, 1e-3},
		{"2/3", 2 / 3.0, 1e-3},
		{"2/3", 0.6666, 1e-3},
		{"3333/5000", 0.6666, 1e-4},
		{"2/3", 0.6667, 1e-3},
		{"6667/10000", 0.6667, 1e-4},
		{"1/7", 1 / 7.0, 1e-3},
		{"99/101", 99 / 101.0, 1e-2},
		{"99/101", 99 / 101.0, 1e-11},
		{"108851651149874/111050674405427", 99 / 101.0, 1e-14},
	}
	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			assert.Equal(t, test.expected, NewRatPrecision(test.float, test.precision).String())
		})
	}
}

func ExampleNewRatPrecision() {
	fmt.Println(NewRatPrecision(0.6666, 0.01).String())
	fmt.Println(NewRatPrecision(0.981, 0.001).String())
	fmt.Println(NewRatPrecision(0.75, 0.01).String())
	// Output:
	// 2/3
	// 981/1000
	// 3/4
}

func ExampleNewRatIterations() {
	fmt.Println(NewRatIterations(0.6667, 3).String())
	fmt.Println(NewRatIterations(0.6667, 4).String())
	// Output:
	// 2/3
	// 6667/10000
}
