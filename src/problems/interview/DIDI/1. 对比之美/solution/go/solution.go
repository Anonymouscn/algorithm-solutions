package main

import (
    "fmt"
)

func main() {
    T := 0
    fmt.Scan(&T)
    for i := 0; i < T; i++ {
        n, m := 0, 0
        fmt.Scan(&n, &m)
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(best(n , m))
    }
}

func best(n int, m int) int {
    if n == 1 {
        return 0
    } else if n == 2 {
        return m
    }
    return 2*m
}