package main

import (
	"testing"
)

// Part 1 tests
func TestFactorial(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		want      int
		wantError bool
	}{
		{name: "factorial of 0", input: 0, want: 1, wantError: false},
		{name: "factorial of 1", input: 1, want: 1, wantError: false},
		{name: "factorial of 6", input: 6, want: 720, wantError: false},
		{name: "factorial of 12", input: 12, want: 479001600, wantError: false},
		{name: "factorial of negative -1", input: -1, want: 0, wantError: true},
		{name: "factorial of negative -53", input: -53, want: 0, wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("Factorial() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test cases for IsPrime
func TestIsPrime(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		want      bool
		wantError bool
	}{
		{name: "prime 2", input: 2, want: true, wantError: false},
		{name: "prime 3", input: 3, want: true, wantError: false},
		{name: "prime 19", input: 19, want: true, wantError: false},
		{name: "prime 23", input: 23, want: true, wantError: false},
		{name: "prime 29", input: 29, want: true, wantError: false},
		{name: "composite 20", input: 20, want: false, wantError: false},
		{name: "composite 25", input: 25, want: false, wantError: false},
		{name: "negative number -7", input: -7, want: false, wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("IsPrime() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test cases for Power
func TestPower(t *testing.T) {
	tests := []struct {
		name      string
		base      int
		exponent  int
		want      int
		wantError bool
	}{
		{name: "base^0", base: 5, exponent: 0, want: 1, wantError: false},
		{name: "0^exponent", base: 0, exponent: 5, want: 0, wantError: false},
		{name: "3^4", base: 3, exponent: 4, want: 81, wantError: false},
		{name: "4^5", base: 4, exponent: 5, want: 1024, wantError: false},
		{name: "2^8", base: 2, exponent: 8, want: 256, wantError: false},
		{name: "negative exponent", base: 2, exponent: -2, want: 0, wantError: true},
		{name: "1^exponent", base: 1, exponent: 100, want: 1, wantError: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantError {
				t.Errorf("Power() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Part 2 Tests
func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name  string
		start int
		calls int
		want  []int
	}{
		{name: "start at 5 three calls", start: 5, calls: 3, want: []int{6, 7, 8}},
		{name: "start at 50 three calls", start: 50, calls: 3, want: []int{51, 52, 53}},
		{name: "start at 0 three calls", start: 0, calls: 3, want: []int{1, 2, 3}},
		{name: "start negative -4", start: -4, calls: 2, want: []int{-3, -2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.start)

			for i := 0; i < tt.calls; i++ {
				got := counter()
				if got != tt.want[i] {
					t.Errorf("call %d = %d, want %d", i+1, got, tt.want[i])
				}
			}
		})
	}
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		input  int
		want   int
	}{
		{name: "double 8", factor: 2, input: 8, want: 16},
		{name: "triple 8", factor: 3, input: 8, want: 24},
		{name: "times five", factor: 5, input: 4, want: 20},
		{name: "zero input", factor: 7, input: 0, want: 0},
		{name: "negative input", factor: 3, input: -2, want: -6},
		{name: "negative factor", factor: -2, input: 3, want: -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MakeMultiplier(tt.factor)
			got := m(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	tests := []struct {
		name    string
		initial int
		adds    []int
		subs    []int
		want    int
	}{
		{name: "start 100, add 25 and 50, sub 15", initial: 100, adds: []int{25, 50}, subs: []int{15}, want: 160},
		{name: "start 50, add 30 and 100, sub 15", initial: 50, adds: []int{30, 100}, subs: []int{15}, want: 165},
		{name: "add only", initial: 10, adds: []int{5, 5}, subs: []int{}, want: 20},
		{name: "subtract only", initial: 10, adds: []int{}, subs: []int{4}, want: 6},
		{name: "mixed ops", initial: 0, adds: []int{10, 2}, subs: []int{5}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, subtract, get := MakeAccumulator(tt.initial)

			for _, v := range tt.adds {
				add(v)
			}
			for _, v := range tt.subs {
				subtract(v)
			}

			got := get()
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

// Part 3 tests
func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		operation func(int) int
		want      []int
	}{
		{name: "square", nums: []int{2, 4, 6, 8}, operation: func(x int) int { return x * x }, want: []int{4, 16, 36, 64}},
		{name: "double", nums: []int{2, 4, 6}, operation: func(x int) int { return x * 2 }, want: []int{4, 8, 12}},
		{name: "add 5", nums: []int{1, 2, 3}, operation: func(x int) int { return x + 5 }, want: []int{6, 7, 8}},
		{name: "negate", nums: []int{1, -2, 3}, operation: func(x int) int { return -x }, want: []int{-1, 2, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.nums, tt.operation)
			if len(got) != len(tt.want) {
				t.Errorf("Apply() length = %d, want %d", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Apply()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{name: "even numbers", nums: []int{1, 2, 3, 4, 5, 6}, predicate: func(x int) bool { return x%2 == 0 }, want: []int{2, 4, 6}},
		{name: "divisible by 4", nums: []int{2, 4, 6, 8, 10, 12}, predicate: func(x int) bool { return x%4 == 0 }, want: []int{4, 8, 12}},
		{name: "greater than 10", nums: []int{5, 11, 20, 7, 15}, predicate: func(x int) bool { return x > 10 }, want: []int{11, 20, 15}},
		{name: "positive numbers", nums: []int{-1, 2, -3, 4, -5}, predicate: func(x int) bool { return x > 0 }, want: []int{2, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.predicate)
			if len(got) != len(tt.want) {
				t.Errorf("Filter() length = %d, want %d", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Filter()[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		initial   int
		operation func(int, int) int
		want      int
	}{
		{name: "sum of even numbers 2-20", nums: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, initial: 0, operation: func(acc, cur int) int { return acc + cur }, want: 110},
		{name: "product", nums: []int{1, 2, 3, 4}, initial: 1, operation: func(acc, cur int) int { return acc * cur }, want: 24},
		{name: "max", nums: []int{1, 5, 3, 4, 20, 15}, initial: -999, operation: func(acc, cur int) int {
			if cur > acc {
				return cur
			}
			return acc
		}, want: 20},
		{name: "min", nums: []int{1, 5, 3, 4, -2, 15}, initial: 999, operation: func(acc, cur int) int {
			if cur < acc {
				return cur
			}
			return acc
		}, want: -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.nums, tt.initial, tt.operation)
			if got != tt.want {
				t.Errorf("Reduce() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	tests := []struct {
		name  string
		f     func(int) int
		g     func(int) int
		input int
		want  int
	}{
		{name: "double then add 5", f: func(x int) int { return x + 5 }, g: func(x int) int { return x * 2 }, input: 7, want: 19},
		{name: "square after double", f: func(x int) int { return x * x }, g: func(x int) int { return x * 2 }, input: 3, want: 36},
		{name: "double after square", f: func(x int) int { return x * 2 }, g: func(x int) int { return x * x }, input: 3, want: 18},
		{name: "add 10 then double", f: func(x int) int { return x * 2 }, g: func(x int) int { return x + 10 }, input: 5, want: 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Compose(tt.f, tt.g)
			got := h(tt.input)
			if got != tt.want {
				t.Errorf("Compose() = %d, want %d", got, tt.want)
			}
		})
	}
}

// Part 5 tests
func TestSwapValues(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		wantA int
		wantB int
	}{
		{"15 and 25", 15, 25, 25, 15},
		{"zeros", 0, 0, 0, 0},
		{"negative + positive", -5, 2, 2, -5},
		{"both negative", -3, -9, -9, -3},
		{"5 and 10", 5, 10, 10, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := SwapValues(tt.a, tt.b)

			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("got (%d,%d) want (%d,%d)",
					gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestSwapPointers(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		wantA int
		wantB int
	}{
		{"15 and 25", 15, 25, 25, 15},
		{"20 and 30", 20, 30, 30, 20},
		{"zeros", 0, 0, 0, 0},
		{"negative + positive", -3, 5, 5, -3},
		{"both negative", -2, -8, -8, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.a
			b := tt.b

			SwapPointers(&a, &b)

			if a != tt.wantA || b != tt.wantB {
				t.Errorf("got (%d,%d) want (%d,%d)",
					a, b, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestDoubleValue(t *testing.T) {
	// Test that DoubleValue doesn't modify original
	val := 15
	original := val
	DoubleValue(val)
	if val != original {
		t.Errorf("DoubleValue modified original value from %d to %d", original, val)
	}
}

func TestDoublePointer(t *testing.T) {
	val := 15
	expected := 30
	DoublePointer(&val)
	if val != expected {
		t.Errorf("DoublePointer = %d, want %d", val, expected)
	}
}
