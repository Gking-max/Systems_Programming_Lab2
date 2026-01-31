package main

import "testing"

func TestFactorial(t *testing.T) {
    tests := []struct {
        name    string
        input   int
        want    int
        wantErr bool
    }{
        {"factorial of 0", 0, 1, false},
        {"factorial of 1", 1, 1, false},
        {"factorial of 5", 5, 120, false},
        {"factorial of 10", 10, 3628800, false},
        {"factorial of negative number", -5, 0, true},
        {"factorial of 3", 3, 6, false},
        {"factorial of 7", 7, 5040, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Factorial(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Factorial() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestIsPrime(t *testing.T) {
    tests := []struct {
        name    string
        input   int
        want    bool
        wantErr bool
    }{
        {"2 is prime", 2, true, false},
        {"3 is prime", 3, true, false},
        {"4 is not prime", 4, false, false},
        {"17 is prime", 17, true, false},
        {"20 is not prime", 20, false, false},
        {"25 is not prime", 25, false, false},
        {"1 is not prime", 1, false, false},
        {"0 is not prime", 0, false, false},
        {"negative number", -5, false, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := IsPrime(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("IsPrime() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPower(t *testing.T) {
    tests := []struct {
        name      string
        base      int
        exponent  int
        want      int
        wantErr   bool
    }{
        {"2^3", 2, 3, 8, false},
        {"5^0", 5, 0, 1, false},
        {"0^5", 0, 5, 0, false},
        {"3^4", 3, 4, 81, false},
        {"negative exponent", 2, -3, 0, true},
        {"10^2", 10, 2, 100, false},
        {"1^100", 1, 100, 1, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Power(tt.base, tt.exponent)
            if (err != nil) != tt.wantErr {
                t.Errorf("Power() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Power() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestMakeCounter(t *testing.T) {
    counter1 := MakeCounter(0)
    

    if got := counter1(); got != 1 {
        t.Errorf("counter1() first call = %v, want 1", got)
    }
    if got := counter1(); got != 2 {
        t.Errorf("counter1() second call = %v, want 2", got)
    }
    
    
    counter2 := MakeCounter(10)
    if got := counter2(); got != 11 {
        t.Errorf("counter2() first call = %v, want 11", got)
    }
    
    if got := counter1(); got != 3 {
        t.Errorf("counter1() third call = %v, want 3", got)
    }
    if got := counter2(); got != 12 {
        t.Errorf("counter2() second call = %v, want 12", got)
    }
}

func TestMakeMultiplier(t *testing.T) {
    doubler := MakeMultiplier(2)
    tripler := MakeMultiplier(3)
    
    tests := []struct {
        name   string
        fn     func(int) int
        input  int
        want   int
    }{
        {"doubler with 5", doubler, 5, 10},
        {"doubler with 0", doubler, 0, 0},
        {"doubler with -3", doubler, -3, -6},
        {"tripler with 4", tripler, 4, 12},
        {"tripler with -2", tripler, -2, -6},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.fn(tt.input)
            if got != tt.want {
                t.Errorf("%s(%d) = %v, want %v", tt.name, tt.input, got, tt.want)
            }
        })
    }
}

func TestMakeAccumulator(t *testing.T) {
    add, subtract, get := MakeAccumulator(100)

    if got := get(); got != 100 {
        t.Errorf("initial get() = %v, want 100", got)
    }
    
    add(50)
    if got := get(); got != 150 {
        t.Errorf("after add(50), get() = %v, want 150", got)
    }
    
    
    subtract(30)
    if got := get(); got != 120 {
        t.Errorf("after subtract(30), get() = %v, want 120", got)
    }
    
    add(100)
    subtract(20)
    add(5)
    if got := get(); got != 205 {
        t.Errorf("after multiple operations, get() = %v, want 205", got)
    }
}