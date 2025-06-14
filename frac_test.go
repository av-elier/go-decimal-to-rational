package dectofrac

import (
	"fmt"
	"testing"
)

func TestNewRatI(t *testing.T) {
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
			actual := NewRatI(test.float, test.iterations).String()
			if actual != test.expected {
				t.Errorf("expected %s, got %s", test.expected, actual)
			}
		})
	}
}

func TestNewRatP(t *testing.T) {
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
			actual := NewRatP(test.float, test.precision).String()
			if actual != test.expected {
				t.Errorf("expected %s, got %s", test.expected, actual)
			}
		})
	}
}

func ExampleNewRatP() {
	fmt.Println(NewRatP(0.6666, 0.01).String())
	fmt.Println(NewRatP(0.981, 0.001).String())
	fmt.Println(NewRatP(0.75, 0.01).String())
	// Output:
	// 2/3
	// 981/1000
	// 3/4
}

func ExampleNewRatI() {
	fmt.Println(NewRatI(0.6667, 3).String())
	fmt.Println(NewRatI(0.6667, 4).String())
	// Output:
	// 2/3
	// 6667/10000
}
