package main

import (
    "fmt"
    "strconv"
)

func main() {
    n := 0
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        word := ""
        fmt.Scan(&word)
        fmt.Println(solution(word))
    }
}

func solution(s string) string {
    length := len(s)
    if length <= 10 {
        return s
    }
    return fmt.Sprintf("%v%v%v", string(s[0]), strconv.Itoa(length - 2), string(s[length - 1]))
}