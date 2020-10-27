package main

import (
    "testing"
    "fmt"
)

var fibtests = []struct {
    in int
    out uint64
}{
    {0, 0},
    {1, 1},
    {2, 2},
    {3, 3},
    {4, 5},
    {100, 1298777728820984005},
}

func TestFib(t *testing.T) {
    for _, tt := range(fibtests) {
        t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
            got := fib(tt.in)
            if got != tt.out {
                t.Error("failed")
            }
        })
    }
}
