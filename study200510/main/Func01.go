package main

import "fmt"

// 다른 언어와 다르게 타입이 뒤에 온다! 주의!
func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}
