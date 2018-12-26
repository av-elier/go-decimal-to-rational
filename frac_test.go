package dectofrac

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRat(t *testing.T) {
	type testCase struct {
		expected   string
		float      float64
		iterations int
	}
	tests := []testCase{
		{"1/4", 0.25, 4},
		{"1/4", 0.25, 100},
		{"2/3", 2 / 3.0, 6},
		{"1/7", 1 / 7.0, 7},
	}
	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			assert.Equal(t, test.expected, NewRat(big.NewFloat(test.float), test.iterations).String())
		})
	}
}
