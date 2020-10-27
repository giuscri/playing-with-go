package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
    fmt.Println(fib(100))
}

func fib(n int) uint64 {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    var a uint64 = 0
    var b uint64 = 1
    for i := 0; i < n; i++ {
        a, b = b, a+b
    }

    return b
}
