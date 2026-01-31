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