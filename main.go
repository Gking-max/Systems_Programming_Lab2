package main

import "errors"

func Factorial(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("factorial is not defined for negative numbers")
    }
    
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result, nil
}

func IsPrime(n int) (bool, error) {
    if n < 2 {
        return false, nil
    }
    
    for i := 2; i*i <= n; i++ {
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
    
    result := 1
    for i := 0; i < exponent; i++ {
        result *= base
    }
    return result, nil
}


func MakeCounter(start int) func() int {
    count := start
    return func() int {
        count++
        return count
    }
}

func MakeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func MakeAccumulator(initial int) (func(int), func(int), func() int) {
    accumulator := initial
    
    add := func(x int) {
        accumulator += x
    }
    
    subtract := func(x int) {
        accumulator -= x
    }
    
    get := func() int {
        return accumulator
    }
    
    return add, subtract, get
}

func Apply(nums []int, operation func(int) int) []int {
    result := make([]int, len(nums))
    for i, v := range nums {
        result[i] = operation(v)
    }
    return result
}

func Filter(nums []int, predicate func(int) bool) []int {
    result := []int{}
    for _, v := range nums {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
    accumulator := initial
    for _, v := range nums {
        accumulator = operation(accumulator, v)
    }
    return accumulator
}

func Compose(f func(int) int, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}

func ExploreProcess() {
    fmt.Println("=== Process Information ===")
    
    pid := os.Getpid()
    fmt.Printf("Current Process ID: %d\n", pid)
    
    ppid := os.Getppid()
    fmt.Printf("Parent Process ID: %d\n", ppid)
    
    data := []int{1, 2, 3, 4, 5}
    
    fmt.Printf("Memory address of slice: %p\n", &data)
    fmt.Printf("Memory address of first element: %p\n", &data[0])
    
    fmt.Println("\nExplanation:")
    fmt.Println("- Process ID (PID) is a unique identifier assigned to each running process")
    fmt.Println("- Parent Process ID (PPID) identifies the process that created this process")
    fmt.Println("- Process isolation ensures each process has its own memory space")
    fmt.Println("- The slice header address is where Go stores slice metadata (length, capacity)")
    fmt.Println("- The element address points to the actual data in memory")
    fmt.Println("- Other processes cannot access these addresses due to memory protection")
}


func DoubleValue(x int) {
    x = x * 2
}

func DoublePointer(x *int) {
    *x = *x * 2
}

func CreateOnStack() int {
    x := 42
    return x
}

func CreateOnHeap() *int {
    x := 42
    return &x
}

func SwapValues(a, b int) (int, int) {
    return b, a
}

func SwapPointers(a, b *int) {
    *a, *b = *b, *a
}

func AnalyzeEscape() {
    fmt.Println("\n=== Escape Analysis ===")
    
    stackVal := CreateOnStack()
    fmt.Printf("Stack value: %d (address cannot be taken)\n", stackVal)
    
    heapPtr := CreateOnHeap()
    fmt.Printf("Heap value via pointer: %d (address: %p)\n", *heapPtr, heapPtr)
    
    fmt.Println("\nEscape Analysis Explanation:")
    fmt.Println("1. Variables in CreateOnStack() stay on the stack because:")
    fmt.Println("   - We return only the value, not its address")
    fmt.Println("   - The compiler can determine its lifetime is limited to the function")
    fmt.Println("")
    fmt.Println("2. Variables in CreateOnHeap() escape to the heap because:")
    fmt.Println("   - We return a pointer to the variable")
    fmt.Println("   - The variable's lifetime extends beyond the function")
    fmt.Println("   - Go's garbage collector will manage this memory")
    fmt.Println("")
    fmt.Println("3. What 'escapes to heap' means:")
    fmt.Println("   - Variable is allocated in heap memory instead of stack")
    fmt.Println("   - Allows variable to outlive the function call")
    fmt.Println("   - Managed by garbage collector (automatic memory management)")
    fmt.Println("   - Slightly slower access than stack variables")
}