package main

import "fmt"

func fibonacci(num int) int {
    a := 0
    b := 1
    // Iterate until num position in sequence.
    for i := 0; i < num; i++ {
        // Used temporary variable for swaping the values.
        temp := a
        a = b
        b = temp + a
    }
    return a
}

func main() {
    // Display first 10 Fibonacci numbers.
    for num := 0; num < 10; num++ {
        result := fibonacci(num)
        fmt.Printf("%v\t", result)
    }
}
