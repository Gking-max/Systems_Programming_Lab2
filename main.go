package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

// Part 1 Functions
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Factorial is not defined for negative numbers")
	}

	if n == 0 {
		return 1, nil
	}

	result, err := Factorial(n - 1)
	if err != nil {
		return 0, err
	}

	return n * result, nil
}

// A prime number can only be divided (without a remainder) by itself and 1
func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}

	if n == 2 {
		return true, nil
	}

	if n%2 == 0 {
		return false, nil
	}

	//For loop starting from 2 and does not go beyond √n
	for i := 3; float64(i) <= math.Sqrt(float64(n)); i = i + 2 {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	//any number raised to zero exponent is equal 1
	if exponent == 0 {
		return 1, nil
	}

	if base == 0 {
		return 0, nil
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result = result * base
	}
	return result, nil
}

// Part 2 Functions
func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(num int) int {
		return factor * num
	}
}

func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	total := initial
	add = func(num int) {
		total += num
	}

	subtract = func(num int) {
		total -= num
	}

	get = func() int {
		return total
	}

	return add, subtract, get
}

// Part 3 functions
func Apply(nums []int, operation func(int) int) []int {
	newSlice := []int{}

	//For loop to loop through the slice and apply the function to each of the elements
	for _, v := range nums {
		newSlice = append(newSlice, operation(v))
	}

	return newSlice
}

func Filter(nums []int, predicate func(int) bool) []int {
	newSlice := []int{}

	for _, v := range nums {
		if predicate(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	accumulator := initial

	for i := 0; i < len(nums); i++ {
		current := nums[i]
		accumulator = operation(accumulator, current)
	}
	return accumulator
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		//returns a composite function that gets the answer to g(x) first then f() is run and given the result of g(x)
		return f(g(x))
	}
}

// Part 4 function
func ExploreProcess() {
	fmt.Println("=== Process Information ===")
	pid := os.Getpid()
	fmt.Printf("Current Process ID: %d\n", pid)
	ppid := os.Getppid()
	fmt.Printf("Parent Process ID: %d\n", ppid)

	data := []int{10, 20, 30, 40, 50}
	fmt.Printf("Memory address of slice: %p\n", &data)
	fmt.Printf("Memory address of first element: %p\n", &data[0])

	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("- Process ID (PID) is a unique identifier assigned to each running process")
	fmt.Println("- Parent Process ID (PPID) identifies the process that created this process")
	fmt.Println("- Process isolation ensures each process has its own memory space")
	fmt.Println("- The slice header address is where Go stores slice metadata (length, capacity)")
	fmt.Println("- The element address points to the actual data in memory")
	fmt.Println("- Other processes cannot access these addresses due to memory protection")
}

// Part 5 functions:
func DoubleValue(x int) {
	x = x * 2
	// This modifies the local copy, not the original variable
}

func DoublePointer(x *int) {
	*x = *x * 2
	// This modifies the original variable because we're working with its address
}

func CreateOnStack() int {
	localVar := 42
	return localVar
	// This variable stays on the stack
}

func CreateOnHeap() *int {
	localVar := 42
	var p *int = &localVar
	return p
	// This variable escapes to the heap
}

func SwapValues(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

func SwapPointers(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func AnalyzeEscape() {
	fmt.Println("\n=== Escape Analysis ===")

	stackVal := CreateOnStack()
	fmt.Printf("Stack value: %d (address cannot be taken)\n", stackVal)

	heapPtr := CreateOnHeap()
	fmt.Printf("Heap value via pointer: %d (address: %p)\n", *heapPtr, heapPtr)

	fmt.Println()
	fmt.Println("Escape Analysis Explanation:")
	fmt.Println("1. Variables in CreateOnStack() stay on the stack because:")
	fmt.Println("   - We return only the value, not its address")
	fmt.Println("   - The compiler can determine its lifetime is limited to the function")
	fmt.Println()
	fmt.Println("2. Variables in CreateOnHeap() escape to the heap because:")
	fmt.Println("   - We return a pointer to the variable")
	fmt.Println("   - The variable's lifetime extends beyond the function")
	fmt.Println("   - Go's garbage collector will manage this memory")
	fmt.Println()
	fmt.Println("3. What 'escapes to heap' means:")
	fmt.Println("   - Variable is allocated in heap memory instead of stack")
	fmt.Println("   - Allows variable to outlive the function call")
	fmt.Println("   - Managed by garbage collector (automatic memory management)")
	fmt.Println("   - Slightly slower access than stack variables")
}

func main() {
	fmt.Println("=== Go Advanced Lab Demo ===")
	fmt.Println()

	ExploreProcess()

	fmt.Println("\n====== Math Operations ======")

	for _, n := range []int{0, 6, 12} {
		if v, err := Factorial(n); err != nil {
			fmt.Println("Factorial error:", err)
		} else {
			fmt.Printf("Factorial(%d) = %d\n", n, v)
		}
	}

	for _, n := range []int{19, 23, 29} {
		if v, err := IsPrime(n); err != nil {
			fmt.Printf("IsPrime(%d) error: %v\n", n, err)
		} else {
			fmt.Printf("IsPrime(%d) = %v\n", n, v)
		}
	}

	if v, err := Power(3, 4); err == nil {
		fmt.Println("Power(3, 4) =", v)
	} else {
		fmt.Println(err)
	}

	if v, err := Power(4, 5); err == nil {
		fmt.Println("Power(4, 5) =", v)
	} else {
		fmt.Println(err)
	}

	fmt.Println("\n====== Closure Demonstration ======")

	counter1 := MakeCounter(5)
	counter2 := MakeCounter(50)

	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter2:", counter2())
	fmt.Println("Counter2:", counter2())

	doubler := MakeMultiplier(2)
	tripler := MakeMultiplier(3)

	fmt.Println("Doubler(8):", doubler(8))
	fmt.Println("Tripler(8):", tripler(8))

	// Accumulator demo
	add, subtract, get := MakeAccumulator(100)
	add(25)
	fmt.Printf("After add(25): %d\n", get())
	subtract(15)
	fmt.Printf("After subtract(15): %d\n", get())
	add(50)
	fmt.Printf("After add(50): %d\n", get())

	fmt.Println("\n====== Higher-Order Functions ======")

	nums := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Println("Original:", nums)

	squared := Apply(nums, func(x int) int { return x * x })
	fmt.Println("Squared:", squared)

	divisibleBy4 := Filter(nums, func(x int) bool { return x%4 == 0 })
	fmt.Println("Divisible by 4:", divisibleBy4)

	sum := Reduce(nums, 0, func(a, c int) int { return a + c })
	fmt.Println("Sum:", sum)

	doubleThenAdd5 := Compose(
		func(x int) int { return x + 5 },
		func(x int) int { return x * 2 },
	)

	fmt.Println("Compose(double then +5)(7):", doubleThenAdd5(7))

	//Pointer demonstration
	fmt.Println("\n====== Pointer Demonstration ======")

	a, b := 15, 25
	fmt.Printf("Before SwapValues: a=%d b=%d\n", a, b)
	x, y := SwapValues(a, b)
	fmt.Printf("After SwapValues: a=%d b=%d (unchanged originals)\n", a, b)
	fmt.Printf("Returned swapped values: %d %d\n", x, y)

	p, q := 15, 25
	fmt.Printf("Before SwapPointers: p=%d q=%d\n", p, q)
	SwapPointers(&p, &q)
	fmt.Printf("After SwapPointers: p=%d q=%d (originals changed)\n", p, q)

	// DoublePointer demo
	val := 15
	fmt.Printf("\nBefore DoublePointer: val=%d\n", val)
	DoublePointer(&val)
	fmt.Printf("After DoublePointer: val=%d\n", val)

	// DoubleValue demo
	val2 := 15
	fmt.Printf("\nBefore DoubleValue: val2=%d\n", val2)
	DoubleValue(val2)
	fmt.Printf("After DoubleValue: val2=%d (unchanged - passed by value)\n", val2)

	AnalyzeEscape()

	fmt.Println("\n=== Demo Complete ===")
}
